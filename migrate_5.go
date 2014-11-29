package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

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
	if files := JSONMigrationFiles(); 0 < len(files) {
		for _, filename := range JSONMigrationFiles() {
			fmt.Println("Exeucting ................ ", filename)
			bs, err := ioutil.ReadFile(filename)
			if err != nil {
				return
			}
			docScript := []byte(bs)
			var mm Migration
			json.Unmarshal(docScript, &mm)
			if !ProcessNow(mm) {
				fmt.Println("Either the file is empty or not in a proper JSON Migration Format")
				// break
			}
		}
	} else {
		fmt.Println("No files in the directory")
	}
}

func JSONMigrationFiles() []string {
	files, _ := ioutil.ReadDir("./")
	var json_files []string
	for _, f := range files {
		json_files = append(json_files, f.Name())
	}
	sort.Strings(json_files)
	return json_files
}

func ProcessNow(m Migration) bool {
	db, _ := sql.Open("mysql", "root:root@/onetest")
	var query string
	nid, _ := strconv.Atoi(m.Id)
	if "Up" == "Up" {
		ud := &m.Up
	} else {
		ud := &m.Down
	}
	if nid != 0 {
		fmt.Println("=-=-=-=-=-==-=-=-=-=-==-=-=-=-=-==-=-=-=-=-=")
		fmt.Println("ID : ", m.Id)
		mmm := "CreateTable123"
		reflect.MethodByName(mmm)
		fmt.Println("Create Table")
		fmt.Println(len(m.Up.Create_Table))
		var values_array []string
		for _, v := range m.Up.Create_Table {
			fmt.Println("v. Table Name ", v.Table_Name)
			for kk, vv := range v.Columns {
				fmt.Println(kk, "vvvvv ", vv.FieldName)
				fmt.Println(kk, "vvvvv ", vv.DataType)
				values_array = append(values_array, vv.FieldName+" "+vv.DataType)
			}
			query = "CREATE TABLE " + v.Table_Name + " (" + strings.Join(values_array, ",") + ")"
		}

		fmt.Println(len(m.Up.Add_Column))
		for _, v := range m.Up.Add_Column {
			fmt.Println("v. Table Name ", v.Table_Name)
			fmt.Println("v. Column Name", v.Column_Name)
			fmt.Println("v. Data Type", v.Data_Type)
			query = "ALTER TABLE " + v.Table_Name + " CHANGE " + v.Column_Name + " " + v.Data_Type
		}
		fmt.Println(len(m.Up.Remove_Column))
		for _, v := range m.Up.Remove_Column {
			fmt.Println("v. Table Name ", v.Table_Name)
			fmt.Println("v. Column Name", v.Column_Name)
			query = "ALTER TABLE " + v.Table_Name + " DROP " + v.Column_Name
		}
		fmt.Println(len(m.Down.Remove_Column))
		for _, v := range m.Down.Remove_Column {
			fmt.Println("v. Table Name ", v.Table_Name)
			fmt.Println("v. Column Name", v.Column_Name)
			query = "ALTER TABLE " + v.Table_Name + " DROP " + v.Column_Name
		}
		fmt.Println(query)
		q, err := db.Query(query)
		if err != nil {
			log.Fatal(err)
		}
		defer q.Close()

		fmt.Println("=-=-=-=-=-==-=-=-=-=-==-=-=-=-=-==-=-=-=-=-=")
		return true
	}
	return false
}

func CreateTable123() {
	fmt.Println("test")
}
