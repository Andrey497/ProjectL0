package main

import (
	"testProgect/cash"
	db2 "testProgect/db"
	"testProgect/nuts"
	"testProgect/server"
	"time"
)

func Start(db *db2.DBConnect, c *cash.Cash, s *nuts.Subscriber) {
	db.OpenConnect()
	c.RecoverCash(db)
	s.OpenSubscriber(db, c)
	go publishMessage()
	server.RunServer(c)
}

func End(db *db2.DBConnect, s *nuts.Subscriber) {
	s.CloseSubscriber()
	db.CloseConnect()
}
func publishMessage() {
	for true {
		nuts.PublishRandomOrderInNuts() // does not return until an ack has been received from NATS Streaming
		time.Sleep(10 * time.Second)
	}
}

func main() {
	db := db2.DBConnect{}
	c := cash.CreateNewCash()
	subscriber1 := nuts.CreateSubscriber("client2")

	Start(&db, &c, &subscriber1)

	defer End(&db, &subscriber1)

}
