package main

import (
	"eventor"
	"eventor/producer"
	"eventor/store"
	"fmt"
	"time"
)

type SetIdEvent struct {
	Id string
}

type IncrementEvent struct {
	Increment int
}

type entity1 struct {
	id    string
	value int
}

func main() {
	brokerList := []string{"localhost:9092"}
	emmitter := store.NewEmmitter(brokerList)
	producer := eventor.NewProducer("entity1", emmitter)
	entity := entity1{
		id: "myId",
	}

	producer.Emmit("set id", SetIdEvent{
		Id: "myId",
	})
	fmt.Printf("set id: myId. entity: %+v\n", entity)

	ticker := time.NewTicker(10 * time.Second)
	for {
		<-ticker.C
		emmitIncrementEvent(&entity, producer)
	}

}

func emmitIncrementEvent(entity *entity1, producer producer.Producer) {
	entity.value += 3

	producer.Emmit("increment", IncrementEvent{
		Increment: 3,
	})

	fmt.Printf("increment: 3. entity: %+v\n", entity)
}
