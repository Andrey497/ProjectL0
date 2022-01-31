package cash

import (
	"testProgect/db"
	"testProgect/models"
)

type Cash struct {
	count     int
	orderCash map[int]models.Order
}

func (c *Cash) SaveOrderInCash(order models.Order) {
	c.orderCash[c.count] = order
	c.count++
}

func (c *Cash) SaveListOrderInCash(orders *[]models.Order) {
	for _, value := range *orders {
		c.SaveOrderInCash(value)
	}

}
func (c *Cash) RecoverCash(db *db.DBConnect) {
	orders := db.GetAllOrder()
	c.SaveListOrderInCash(&orders)
}
func (c *Cash) GetItemCashById(id int) (models.Order, bool) {
	if val, ok := c.orderCash[id]; ok {
		return val, true
	}
	return models.Order{}, false
}
func CreateNewCash() Cash {
	return Cash{0, make(map[int]models.Order)}
}
