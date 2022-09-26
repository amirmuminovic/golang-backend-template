package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type SQLRepository struct {
	db *sql.DB
	tp TableProperties
}

type Repository interface {
	createTableIfNotExists()
	insertOne(i InsertionEntry)
	get(g GetQuery) *sql.Rows
	update(u UpdateEntry)
	delete(d DeleteQuery)
	count(c CountQuery) int64
}

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

func (r SQLRepository) createTableIfNotExists() {
	_, err := r.db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", r.tp.tableName, r.tp.tableDefinition))

	if err != nil {
		log.Fatal(err)
	}
}

func (r SQLRepository) insertOne(databaseEntry InsertionEntry) {
	keys := strings.Join(databaseEntry.keys, ",")
	values := strings.Join(databaseEntry.values, ",")

	_, err := r.db.Query(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", r.tp.tableName, keys, values))

	if err != nil {
		log.Fatal(err)
	}
}

func (r SQLRepository) get(databaseQuery GetQuery) *sql.Rows {
	keys := strings.Join(databaseQuery.keys, ",")
	query := fmt.Sprintf("SELECT %s FROM %s", keys, r.tp.tableName)
	appendQueryWithCondition(query, databaseQuery.conditions)

	rows, err := r.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	return rows
}

func (r SQLRepository) update(updateEntry UpdateEntry) {
	query := fmt.Sprintf("UPDATE %s SET %s", r.tp.tableName, updateEntry.newValues)
	appendQueryWithCondition(query, updateEntry.conditions)

	_, err := r.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
}

func (r SQLRepository) delete(deleteQuery DeleteQuery) {
	query := fmt.Sprintf("DELETE FROM %s", r.tp.tableName)
	appendQueryWithCondition(query, deleteQuery.conditions)
	_, err := r.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}
}

func (r SQLRepository) count(countQuery CountQuery) int64 {
	query := fmt.Sprintf("SELECT COUNT (*) FROM %s", r.tp.tableName)
	appendQueryWithCondition(query, countQuery.conditions)

	rows, err := r.db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	var count int64
	for rows.Next() {
		rows.Scan(&count)
	}

	return count
}

func appendQueryWithCondition(query string, queryCondition []string) string {
	if len(queryCondition) > 0 {
		conditions := strings.Join(queryCondition, " AND ")
		query += fmt.Sprintf(" WHERE %s;", conditions)
	}

	return query
}
