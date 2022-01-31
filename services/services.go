package services

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"testProgect/cash"
	"testProgect/db"
	"testProgect/models"
	"time"
)

func RandomNumber() int {
	return rand.Intn(2555555)
}

func RandomOrder() models.Order {
	var m models.Order
	m.Order_uid = RandomString()
	m.Track_number = RandomString()
	m.Entry = RandomString()
	m.Delivery = RandomDelivery()
	m.Payment = RandomPayment()
	m.Items = RandomItems()
	m.Locale = RandomString()
	m.Internal_signature = RandomString()
	m.Customer_id = RandomString()
	m.Delivery_service = RandomString()
	m.Shardkey = RandomString()
	m.Sm_id = RandomNumber()
	m.Date_created = time.Now().Format("2021-11-26T06:22:19Z")
	m.Oof_shard = RandomString()
	return m
}
func RandomDelivery() models.Delivery {
	var m models.Delivery
	m.Name = RandomString()
	m.Phone = RandomString()
	m.Zip = RandomString()
	m.City = RandomString()
	m.Address = RandomString()
	m.Region = RandomString()
	m.Email = RandomString()
	return m
}
func RandomPayment() models.Payment {
	var m models.Payment
	m.Transaction = RandomString()
	m.Request_id = RandomString()
	m.Currency = RandomString()
	m.Provider = RandomString()
	m.Amount = RandomNumber()
	m.Payment_dt = RandomNumber()
	m.Bank = RandomString()
	m.Delivery_cost = RandomNumber()
	m.Goods_total = RandomNumber()
	m.Custom_fee = RandomNumber()
	return m
}
func RandomItem() models.Item {
	var m models.Item
	m.Chrt_id = RandomNumber()
	m.Track_number = RandomString()
	m.Price = RandomNumber()
	m.Rid = RandomString()
	m.Name = RandomString()
	m.Sale = RandomNumber()
	m.Size = RandomString()
	m.Total_price = RandomNumber()
	m.Nm_id = RandomNumber()
	m.Brand = RandomString()
	m.Status = RandomNumber()
	return m
}
func RandomItems() []models.Item {
	var m []models.Item
	r := rand.Intn(5)
	for i := 0; i <= r; i++ {
		m = append(m, RandomItem())
	}
	return m
}

func RandomString() string {
	n := 1 + rand.Intn(11)
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func MyUnmarshal(jsonText []byte, obj interface{}) error {
	dec := json.NewDecoder(bytes.NewReader(jsonText))
	dec.DisallowUnknownFields()
	err := dec.Decode(&obj)
	return err
}
func SaveOrderInMemoryAndDb(db *db.DBConnect, c *cash.Cash, o models.Order) {
	db.CreateOrder(o)
	c.SaveOrderInCash(o)
}
