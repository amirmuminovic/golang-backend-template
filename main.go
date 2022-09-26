package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func countSomeStuff(r Repository, t ToDo) {
	count := r.count(t.SerializeForCountAll())
	fmt.Printf("The count is %d", count)
}

func main() {
	ac := getConfig()
	db := connectToDb(ac)
	rmqService := RabbitMQService{
		ac: ac,
	}

	rmqService.connect()
	rmqService.createChannel()
	toDo := ToDo{
		Title:       "Buy Groceries",
		Description: "Get some food for the week",
	}

	ec := ExchangeSettings{
		exchangeName:       "todoapps",
		exchangeType:       amqp.ExchangeTopic,
		exchangeDurable:    true,
		exchangeAutoDelete: false,
		exchangeInternal:   false,
		exchangeNoWait:     false,
		exchangeArguments:  nil,
	}

	rmqService.createTopic(ec)

	rmqMessage := RMQMessageConfig{
		exchangeName:       "todoapps",
		routingKey:         "to-do-apps",
		messageMandatory:   false,
		messageImmediate:   false,
		messageContentType: "application/json",
	}
	rmqService.publishMessage(rmqMessage, toDo.SerializeToJson())

	todoTableProperties := TableProperties{
		tableName:       "foobara",
		tableDefinition: "id serial primary key, title varchar(64), description varchar(255)",
	}

	toDoRepo := SQLRepository{
		db: db,
		tp: todoTableProperties,
	}

	toDoRepo.createTableIfNotExists()
	// count := toDoRepo.count(toDo.SerializeForCountAll())
	countSomeStuff(toDoRepo, toDo)

	// createTableIfNotExists(db, todoTableProperties)
	// insertOne(db, todoTableProperties, toDo.SerializeForInsert())

	// rows := get(db, todoTableProperties, toDo.SerializeForQueryWithAllFields())
	// rows := get(db, todoTableProperties, toDo.SerializeForQueryAllDataWithSelectFields([]string{"title"}))
	// count := count(db, todoTableProperties, toDo.SerializeForCountWithFilter())

	// fmt.Println("Count", count)

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
	startServer(ac)
}
