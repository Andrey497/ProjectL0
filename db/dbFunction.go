package db

import (
	"fmt"
	"log"
	"testProgect/models"
)

func (db *DBConnect) GetOrCreateDelivery(m *models.Delivery) int {
	command := fmt.Sprintf("Select get_or_create_delivery('%s','%s','%s','%s','%s','%s','%s')",
		m.Name, m.Phone, m.Zip, m.City, m.Address, m.Region, m.Email)
	var id int
	err := db.Db.QueryRow(command).Scan(&id)
	if err != nil {
		log.Printf("Warning:%v\n", err.Error())
		panic(err)
	}
	return id
}

func (db *DBConnect) GetOrCreatePayment(m *models.Payment) int {
	command := fmt.Sprintf("Select get_or_create_payment('%s','%s','%s','%s',%d,%d,'%s',%d,%d,%d)",
		m.Transaction, m.Request_id, m.Currency, m.Provider, m.Amount, m.Payment_dt, m.Bank, m.Delivery_cost, m.Goods_total, m.Custom_fee)
	var id int
	err := db.Db.QueryRow(command).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func (db *DBConnect) GetOrCreateItem(m *models.Item) int {
	command := fmt.Sprintf("Select get_or_create_item(%d,'%s',%d,'%s','%s',%d,'%s',%d,%d,'%s',%d)",
		m.Chrt_id, m.Track_number, m.Price, m.Rid, m.Name, m.Sale, m.Size, m.Total_price, m.Nm_id, m.Brand, m.Status)
	var id int
	err := db.Db.QueryRow(command).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}
func (db *DBConnect) AppendItemInOrder(itemId int, orderId string) {
	command := fmt.Sprintf("Call add_itemInDelivery(%d,'%s')",
		itemId, orderId)
	_, err := db.Db.Exec(command)
	if err != nil {
		panic(err)
	}
}

func (db *DBConnect) CreateOrder(m models.Order) {
	deliveryId := db.GetOrCreateDelivery(&m.Delivery)
	paymentId := db.GetOrCreatePayment(&m.Payment)
	command := fmt.Sprintf("Call create_order('%s','%s','%s',%d,%d,'%s','%s','%s','%s','%s',%d,'%s','%s')",
		m.Order_uid, m.Track_number, m.Entry, deliveryId, paymentId, m.Locale, m.Internal_signature, m.Customer_id, m.Delivery_service, m.Shardkey, m.Sm_id, m.Date_created, m.Oof_shard)
	_, err := db.Db.Exec(command)
	if err != nil {
		panic(err)
	}

	var idItem int
	for _, value := range m.Items {
		idItem = db.GetOrCreateItem(&value)
		db.AppendItemInOrder(idItem, m.Order_uid)
	}
}

func (db *DBConnect) GetDeliveryById(id int) models.Delivery {
	command := fmt.Sprintf("Select* from get_delivery_byId(%d)", id)
	var idDelivery *int
	var delivery models.Delivery
	err := db.Db.QueryRow(command).Scan(&idDelivery, &delivery.Name, &delivery.Phone, &delivery.Zip, &delivery.City,
		&delivery.Address, &delivery.Region, &delivery.Email)
	if err != nil {
		panic(err)
	}
	return delivery
}

func (db *DBConnect) GetPaymentById(id int) models.Payment {
	command := fmt.Sprintf("Select* from  get_payment_byId(%d)", id)
	var payment models.Payment
	var idPayment *int
	err := db.Db.QueryRow(command).Scan(&idPayment, &payment.Transaction, &payment.Request_id, &payment.Currency,
		&payment.Provider, &payment.Amount, &payment.Payment_dt,
		&payment.Bank, &payment.Delivery_cost, &payment.Goods_total, &payment.Custom_fee)
	if err != nil {
		panic(err)
	}
	return payment
}
func (db *DBConnect) GetItemsByOrderId(id string) []models.Item {
	command := fmt.Sprintf("Select* from  get_items_byorderId('%s')", id)
	items := []models.Item{}
	rows, err := db.Db.Query(command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		i := models.Item{}
		err := rows.Scan(&i.Chrt_id, &i.Track_number, &i.Price, &i.Rid, &i.Name, &i.Sale, &i.Size, &i.Total_price, &i.Nm_id, &i.Brand, &i.Status)
		if err != nil {
			fmt.Println(err)
			continue
		}
		items = append(items, i)
	}
	return items
}

func (db *DBConnect) GetAllOrder() []models.Order {
	command := "select * from View_orders_ALL"
	orders := []models.Order{}
	rows, err := db.Db.Query(command)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		order := models.Order{}
		var paymen_id int
		var delivery_id int

		err := rows.Scan(&order.Order_uid, &order.Track_number, &order.Entry, &delivery_id,
			&paymen_id, &order.Locale, &order.Internal_signature, &order.Customer_id,
			&order.Delivery_service, &order.Shardkey, &order.Sm_id, &order.Date_created, &order.Oof_shard)
		if err != nil {
			fmt.Println(err)
			continue
		}
		order.Delivery = db.GetDeliveryById(delivery_id)
		order.Payment = db.GetPaymentById(paymen_id)
		order.Items = db.GetItemsByOrderId(order.Order_uid)

		orders = append(orders, order)
	}
	return orders
}
