package main

import (
	"encoding/json"
	"fmt"
)

var document string = `{
    "id": "20130106222315",
    "up": {
        "add_column": [{ "table_name": "seasons", "column_name" : "name", "data_type":"varchar" }, { "table_name": "seasons1", "column_name" : "name1", "data_type":"varchar1" }]
        ,
        "remove_column": [{ "table_name": "tb1", "column_name" : "col1" }, { "table_name": "tb2", "column_name" : "col2" }]
        ,
        "create_table": [{"table_name" : "tb3", "columns" : [{"fieldname": "col1", "datatype": "dt1"}, {"fieldname": "col2", "datatype": "dt2"}, {"fieldname": "col3", "datatype": "dt3"}, {"fieldname": "col4", "datatype": "dt4"}] }]
    },
    "down": {
        "remove_column": [{ "table_name": "rtb1", "column_name" : "rcol1" }, { "table_name": "rtb2", "column_name" : "rcol2" }]
    }
}`

type Migration struct {
	Id   string `jpath:"id"`
	Up   UpDown `jpath:"up"`
	Down UpDown `jpath:"down"`
}

type UpDown struct {
	Add_Column    []AddColumn    `jpath:"add_column"`
	Remove_Column []RemoveColumn `jpath:"remove_column"`
	Create_Table  []CreateTable  `jpath:"create_table"`
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

type CreateTable struct {
	Table_Name string    `jpath:"table_name"`
	Columns    []Columns `jpath:"columns"`
}

type Columns struct {
	FieldName string `jpath:"fieldname"`
	DataType  string `jpath:"datatype"`
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

	fmt.Println(len(m.Down.Remove_Column))
	for k, v := range m.Down.Remove_Column {
		fmt.Println(k, "===", v)
		fmt.Println("v. Table Name ", v.Table_Name)
		fmt.Println("v. Column Name", v.Column_Name)
	}

	fmt.Println("Create Table")
	fmt.Println(len(m.Up.Create_Table))
	for k, v := range m.Up.Create_Table {
		fmt.Println(k, "===", v)
		fmt.Println("v. Table Name ", v.Table_Name)
		for kk, vv := range v.Columns {
			fmt.Println(kk, "vvvvv ", vv.FieldName)
			fmt.Println(kk, "vvvvv ", vv.DataType)
		}
	}

	fmt.Println("=-=-=-=-=-=")
}
