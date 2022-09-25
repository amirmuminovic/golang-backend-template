package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type ToDo struct {
	Title       string
	Description string
}

func (td ToDo) SerializeForInsert() InsertionEntry {
	e := reflect.ValueOf(&td).Elem()

	keys := []string{}
	values := []string{}

	for i := 0; i < e.NumField(); i++ {
		keys = append(keys, e.Type().Field(i).Name)

		serializedValue := ""
		if e.Type().Field(i).Type.Name() == "string" {
			serializedValue = fmt.Sprintf("'%s'", (e.Field(i).Interface()))
		} else {
			serializedValue = fmt.Sprintf("%s", (e.Field(i).Interface()))
		}

		values = append(values, serializedValue)
	}

	de := InsertionEntry{
		keys:   keys,
		values: values,
	}

	return de
}

func (td ToDo) GetConditions() []string {
	e := reflect.ValueOf(&td).Elem()

	conditions := []string{}

	for i := 0; i < e.NumField(); i++ {
		condition := ""
		if !e.Field(i).IsZero() {
			if e.Type().Field(i).Type.Name() == "string" {
				condition = fmt.Sprintf("%s='%s'", e.Type().Field(i).Name, (e.Field(i).Interface()))
			} else {
				condition = fmt.Sprintf("%s=%s", e.Type().Field(i).Name, (e.Field(i).Interface()))
			}

			conditions = append(conditions, condition)
		}
	}

	return conditions
}

func (td ToDo) SerializeForQueryFilteredDataWithAllFields() GetQuery {
	conditions := td.GetConditions()

	de := GetQuery{
		keys:       []string{"*"},
		conditions: conditions,
	}

	return de
}

func (td ToDo) SerializeForQueryAllDataWithAllFields() GetQuery {
	de := GetQuery{
		keys:       []string{"*"},
		conditions: []string{"TRUE"},
	}

	return de
}

func (td ToDo) SerializeForQueryFilteredDataWithSelectFields(keys []string) GetQuery {
	conditions := td.GetConditions()

	de := GetQuery{
		keys:       keys,
		conditions: conditions,
	}

	return de
}

func (td ToDo) SerializeForQueryAllDataWithSelectFields(keys []string) GetQuery {
	de := GetQuery{
		keys:       keys,
		conditions: []string{"TRUE"},
	}

	return de
}

func (td ToDo) SerializeForCountWithFilter() CountQuery {
	conditions := td.GetConditions()

	cq := CountQuery{
		conditions: conditions,
	}

	return cq
}

func (td ToDo) SerializeForCountAll() CountQuery {
	cq := CountQuery{
		conditions: []string{"TRUE"},
	}

	return cq
}

func (td ToDo) SerializeToJson() []byte {
	b, err := json.Marshal(td)

	if err != nil {
		log.Fatal(err)
	}

	return b
}
