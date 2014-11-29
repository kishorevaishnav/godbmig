package main

import (
	"encoding/json"
	"fmt"
)

var document string = `{
    "id": "20130106222315",
    "up": {
        "add_column": [{ "table_name": "seasons", "column_name" : "name", "data_type":"varchar" }, { "table_name": "seasons1", "column_name" : "name1", "data_type":"varchar1" }],
        "remove_column": [{ "table_name": "tb1", "column_name" : "col1" }, { "table_name": "tb2", "column_name" : "col2" }]
    },
    "down": "down_asdf"
}`

type Migration struct {
	Id   string `jpath:"id"`
	Up   Up     `jpath:"up"`
	Down string `jpath:"down"`
}

type Up struct {
	Add_Column    []AddColumn    `jpath:"add_column"`
	Remove_Column []RemoveColumn `jpath:"remove_column"`
}

type AddColumn struct {
	Table_Name  string `jpath:"table_name"`
	Column_Name string `jpath:"column_name"`
	Data_Type   string `jpath:"data_type"`
}

type RemoveColumn struct {
	Table_Name  string `jpath:"table_name"`
	Column_Name string `jpath:"column_name"`
}

type up Up
type addColumn AddColumn
type removeColumn RemoveColumn
type migration Migration

func (m *Migration) UnmarshalJSON(b []byte) (err error) {
	fmt.Println("in migration 1")
	j := migration{}
	if err = json.Unmarshal(b, &j); err == nil {
		*m = Migration(j)
		return
	}
	return
}

func (u *Up) UnmarshalJSON(b []byte) (err error) {
	fmt.Println("in up")
	j := up{}
	if err = json.Unmarshal(b, &j); err == nil {
		*u = Up(j)
		return
	}
	return
}

func (ac *AddColumn) UnmarshalJSON(b []byte) (err error) {
	fmt.Println("in add column")
	j := addColumn{}
	if err = json.Unmarshal(b, &j); err == nil {
		*ac = AddColumn(j)
		return
	}
	return
}

func (rc *RemoveColumn) UnmarshalJSON(b []byte) (err error) {
	fmt.Println("in remove column")
	j := removeColumn{}
	if err = json.Unmarshal(b, &j); err == nil {
		*rc = RemoveColumn(j)
		return
	}
	return
}

func main() {
	docScript := []byte(document)
	var mm Migration
	json.Unmarshal(docScript, &mm)
	ProcessNow(mm)
}

func ProcessNow(m Migration) {
	fmt.Println("=-=-=-=-=-=")
	fmt.Println("ID : ", m.Id)
	fmt.Println(len(m.Up.Add_Column))
	for k, v := range m.Up.Add_Column {
		fmt.Println(k, "===", v)
		fmt.Println("v. Table Name ", v.Table_Name)
		fmt.Println("v. Column Name", v.Column_Name)
		fmt.Println("v. Data Type", v.Data_Type)
	}

	fmt.Println(len(m.Up.Remove_Column))
	for k, v := range m.Up.Remove_Column {
		fmt.Println(k, "===", v)
		fmt.Println("v. Table Name ", v.Table_Name)
		fmt.Println("v. Column Name", v.Column_Name)
	}

	fmt.Println("=-=-=-=-=-=")
}
