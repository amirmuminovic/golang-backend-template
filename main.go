package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQConfig struct {
	exchangeName       string
	exchangeType       string
	exchangeDurable    bool
	exchangeAutoDelete bool
	exchangeInternal   bool
	exchangeNoWait     bool
	exchangeArguments  amqp.Table
	routingKey         string
	messageMandatory   bool
	messageImmediate   bool
	messageContentType string
}

func main() {
	ac := getConfig()
	db := connectToDb(ac)
	rmq := connectToRMQ(ac)
	channel := createChannel(rmq)

	toDo := ToDo{
		Title:       "Buy Groceries",
		Description: "Get some food for the week",
	}

	rmqConfig := RabbitMQConfig{
		exchangeName:       "todoapps",
		exchangeType:       amqp.ExchangeTopic,
		exchangeDurable:    true,
		exchangeAutoDelete: false,
		exchangeInternal:   false,
		exchangeNoWait:     false,
		exchangeArguments:  nil,
		routingKey:         "to-do-apps",
		messageMandatory:   false, // TODO: I'd like to revise what this actually means
		messageImmediate:   false, // TODO: I'd like to revise what this actually means
		messageContentType: "application/json",
	}

	createExchange(channel, rmqConfig)

	publishMessage(channel, rmqConfig, toDo.SerializeToJson())

	startServer(ac)

	todoTableProperties := TableProperties{
		tableName:       "foobara",
		tableDefinition: "id serial primary key, title varchar(64), description varchar(255)",
	}

	fmt.Println("waananan: |" + toDo.Title + "|")

	createTableIfNotExists(db, todoTableProperties)
	// insertOne(db, todoTableProperties, toDo.SerializeForInsert())

	// rows := get(db, todoTableProperties, toDo.SerializeForQueryWithAllFields())
	// rows := get(db, todoTableProperties, toDo.SerializeForQueryAllDataWithSelectFields([]string{"title"}))
	count := count(db, todoTableProperties, toDo.SerializeForCountWithFilter())

	fmt.Println("Count", count)

	// for rows.Next() {
	// 	// var id string
	// 	var title string
	// 	// var description string
	// 	// err := rows.Scan(&id, &title, &description)
	// 	err := rows.Scan(&title)
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 		break
	// 	}
	// 	// fmt.Println("id", id)
	// 	fmt.Println("title", title)
	// 	// fmt.Println("description", description)
	// 	// names = append(names, name)
	// }
	// rows.Close()

}
