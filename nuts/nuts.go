package nuts

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"testProgect/cash"
	"testProgect/db"
	"testProgect/models"
	"testProgect/services"
)

const (
	clusterId          = "test-cluster"
	clientId_sender    = "client1"
	clientId_recipient = "client2"
	channel            = "channel1"
)

type Subscriber struct {
	ClientId string
	Sc       stan.Conn
	Sub      stan.Subscription
	Err      error
}

func CreateSubscriber(clienId string) Subscriber {
	var subscriber Subscriber
	subscriber.ClientId = clienId
	return subscriber
}

func Publish(message []byte) {
	sc, err := stan.Connect(clusterId, clientId_sender)
	if err != nil {
		log.Printf("Warning:%v\n", err.Error())
	}
	sc.Publish(channel, message)
	sc.Close()
}

func (s *Subscriber) OpenSubscriber(db *db.DBConnect, c *cash.Cash) {
	s.Sc, s.Err = stan.Connect(clusterId, s.ClientId)
	if s.Err != nil {
		log.Printf("Warning:%v\n", s.Err.Error())
	} else {
		s.Sub, _ = s.Sc.Subscribe(channel, func(m *stan.Msg) {
			var obj models.Order
			services.MyUnmarshal(m.Data, &obj)
			services.SaveOrderInMemoryAndDb(db, c, obj)
		}, stan.DeliverAllAvailable())
	}
}

func (s *Subscriber) CloseSubscriber() {
	s.Sub.Unsubscribe()
	s.Sc.Close()
}
func PublishRandomOrderInNuts() {
	order := services.RandomOrder()
	res, _ := json.Marshal(order)
	Publish(res)
}
