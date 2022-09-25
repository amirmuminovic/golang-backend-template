package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type TableProperties struct {
	tableName       string
	tableDefinition string
}

type InsertionEntry struct {
	keys   []string
	values []string
}

type GetQuery struct {
	keys       []string
	conditions []string
}

type UpdateEntry struct {
	newValues  []string
	conditions []string
}

type DeleteQuery struct {
	conditions []string
}

type CountQuery struct {
	conditions []string
}

func createTableIfNotExists(db *sql.DB, tableProperties TableProperties) {
	_, err := db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableProperties.tableName, tableProperties.tableDefinition))

	if err != nil {
		log.Fatal(err)
	}
}

func insertOne(db *sql.DB, tableProperties TableProperties, databaseEntry InsertionEntry) {
	keys := strings.Join(databaseEntry.keys, ",")
	values := strings.Join(databaseEntry.values, ",")

	_, err := db.Query(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableProperties.tableName, keys, values))

	if err != nil {
		log.Fatal(err)
	}
}

func get(db *sql.DB, tableProperties TableProperties, databaseQuery GetQuery) *sql.Rows {
	keys := strings.Join(databaseQuery.keys, ",")
	conditions := strings.Join(databaseQuery.conditions, " AND ")

	rows, err := db.Query(fmt.Sprintf("SELECT %s FROM %s WHERE %s;", keys, tableProperties.tableName, conditions))

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func update(db *sql.DB, tableProperties TableProperties, updateEntry UpdateEntry) {
	_, err := db.Query(fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableProperties.tableName, updateEntry.newValues, updateEntry.conditions))

	if err != nil {
		log.Fatal(err)
	}
}

func delete(db *sql.DB, tableProperties TableProperties, deleteQuery DeleteQuery) {
	_, err := db.Query(fmt.Sprintf("DELETE FROM %s WHERE %s", tableProperties.tableName, deleteQuery.conditions))

	if err != nil {
		log.Fatal(err)
	}
}

func count(db *sql.DB, tableProperties TableProperties, countQuery CountQuery) int64 {
	rows, err := db.Query(fmt.Sprintf("SELECT COUNT * FROM %s WHERE %s", tableProperties.tableName, countQuery.conditions))

	if err != nil {
		log.Fatal(err)
	}

	var count int64
	for rows.Next() {
		rows.Scan(&count)
	}

	return count
}
