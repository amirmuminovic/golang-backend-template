package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func startServer() {
	ac := getConfig()
	db := connectToDb(ac)

	todoTableProperties := TableProperties{
		tableName:       "foobara",
		tableDefinition: "id serial primary key, title varchar(64), description varchar(255)",
	}

	toDo := ToDo{
		Title: "Buy Groceries",
		// Description: "Get some food for the week",
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

	http.HandleFunc("/health", handleHealthCheck)
	log.Fatal(http.ListenAndServe(":"+ac.appPort, nil))
}
