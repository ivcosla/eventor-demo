package main

import (
	"encoding/json"
	"eventor"
	"eventor/store"
	"fmt"
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
	listener := store.NewListener(brokerList)
	consumer := eventor.NewConsumer("entity1", listener)
	entity := &entity1{}

	fmt.Println("listening")
	consumer.Register("set id", func(b []byte) {
		event := &SetIdEvent{}
		err := json.Unmarshal(b, event)
		if err != nil {
			fmt.Println("set Id event. Error: ", err)
		}

		entity.id = event.Id

		fmt.Printf("set Id event. entity: %+v.\n", entity)
	})

	consumer.Register("increment", func(b []byte) {
		event := &IncrementEvent{}
		err := json.Unmarshal(b, event)
		if err != nil {
			fmt.Println("increment event. Error: ", err)
		}

		entity.value += event.Increment

		fmt.Printf("increment: 3. entity: %+v.\n", entity)
	})

	consumer.Listen()
	for {
	}
}
