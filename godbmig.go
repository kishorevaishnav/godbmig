package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	gdm "./mysql"

	_ "github.com/go-sql-driver/mysql"
)

type Migration struct {
	Id   string `jpath:"id"`
	Up   UpDown `jpath:"up"`
	Down UpDown `jpath:"down"`
}

type UpDown struct {
	Create_Table  []CreateTable  `jpath:"create_table"`
	Drop_Table    []DropTable    `jpath:"drop_table"`
	Add_Column    []AddColumn    `jpath:"add_column"`
	Remove_Column []RemoveColumn `jpath:"remove_column"`
	Add_Index     []AddIndex     `jpath:"add_index"`
	Remove_Index  []RemoveIndex  `jpath:"remove_index"`
	Sql           string         `jpath:"sql"`
}

type CreateTable struct {
	Table_Name string    `jpath:"table_name"`
	Columns    []Columns `jpath:"columns"`
}

type DropTable struct {
	Table_Name string `jpath:"table_name"`
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

type AddIndex struct {
	Table_Name string    `jpath:"table_name"`
	Index_Type string    `jpath:"index_type"`
	Columns    []Columns `jpath:"columns"`
}

type RemoveIndex struct {
	Table_Name string    `jpath:"table_name"`
	Columns    []Columns `jpath:"columns"`
}

type Columns struct {
	FieldName string `jpath:"fieldname"`
	DataType  string `jpath:"datatype"`
}

func main() {
	if files := JSONMigrationFiles(); 0 < len(files) {
		for _, filename := range files {
			fmt.Println("Exeucting ................ ", filename)
			bs, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println(err)
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
		if !f.IsDir() && strings.Contains(f.Name(), ".rm.json") {
			json_files = append(json_files, f.Name())
		}
	}
	sort.Strings(json_files)
	return json_files
}

func ProcessNow(m Migration) bool {
	gdm.Init()
	// gdm.Db, _ = sql.Open("mysql", "root:root@/onetest")

	// var query string
	nid, _ := strconv.Atoi(m.Id)
	if nid != 0 {
		fmt.Println("ID : ", m.Id)

		var values_array []string
		for _, v := range m.Up.Create_Table {
			for _, vv := range v.Columns {
				values_array = append(values_array, vv.FieldName+" "+vv.DataType)
			}
			gdm.CreateTable(v.Table_Name, values_array)
		}
		for _, v := range m.Up.Add_Column {
			gdm.AddColumn(v.Table_Name, v.Column_Name, v.Data_Type)
		}
		for _, v := range m.Up.Remove_Column {
			gdm.RemoveColumn(v.Table_Name, v.Column_Name)
		}
		for _, v := range m.Up.Add_Index {
			gdm.AddIndex(v.Table_Name, v.Index_Type, v.Columns)
		}
		for _, v := range m.Up.Remove_Index {
			gdm.RemoveIndex(v.Table_Name, v.Index_Type, v.Columns)
		}
		return true
	}
	return false
}
