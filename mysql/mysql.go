package mysql

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"
	"strings"

	m "../migration"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Init() {
	fmt.Println("Inside the MySQL")
	Db, _ = sql.Open("mysql", "root:root@/onetest")
}

func ProcessNow(m m.Migration) {
	nid, _ := strconv.Atoi(m.Id)
	if nid != 0 {
		fmt.Println("ID : ", m.Id)

		for _, v := range m.Up.Create_Table {
			var values_array []string
			for _, vv := range v.Columns {
				values_array = append(values_array, vv.FieldName+" "+vv.DataType)
			}
			CreateTable(v.Table_Name, values_array)
		}
		for _, v := range m.Up.Add_Column {
			AddColumn(v.Table_Name, v.Column_Name, v.Data_Type)
		}
		for _, v := range m.Up.Remove_Column {
			RemoveColumn(v.Table_Name, v.Column_Name)
		}
		for _, v := range m.Up.Add_Index {
			var fieldname_array []string
			for _, vv := range v.Columns {
				fieldname_array = append(fieldname_array, vv.FieldName)
			}
			AddIndex(v.Table_Name, v.Index_Type, fieldname_array)
		}
		for _, v := range m.Up.Remove_Index {
			var fieldname_array []string
			for _, vv := range v.Columns {
				fieldname_array = append(fieldname_array, vv.FieldName)
			}
			RemoveIndex(v.Table_Name, v.Index_Type, fieldname_array)
		}
	}
}
func execQuery(query string) {
	fmt.Println("MySQL---" + query)
	// q, err := Db.Query(query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer q.Close()
}

func CreateTable(table_name string, field_datatype []string) {
	query := "CREATE TABLE " + table_name + " (" + strings.Join(field_datatype, ",") + ")"
	execQuery(query)
	return
}

func DropTable(table_name string) {
	query := "DROP TABLE " + table_name
	execQuery(query)
	return
}

func AddColumn(table_name string, column_name string, data_type string) {
	query := "ALTER TABLE " + table_name + " CHANGE " + column_name + " " + data_type
	execQuery(query)
	return
}

func RemoveColumn(table_name string, column_name string) {
	query := "ALTER TABLE " + table_name + " DROP " + column_name
	execQuery(query)
	return
}

func AddIndex(table_name string, index_type string, field []string) {
	sort.Strings(field)
	tmp_index_name := strings.ToLower(strings.Join(field, "_") + "_index")
	query := "CREATE " + strings.ToUpper(index_type) + " INDEX " + tmp_index_name + " ON " + table_name + "( " + strings.Join(field, ",") + " )"
	execQuery(query)
	return
}

func RemoveIndex(table_name string, index_type string, field []string) {
	sort.Strings(field)
	tmp_index_name := strings.ToLower(strings.Join(field, "_") + "_index")
	query := ""
	if index_type != "" {
		query = "ALTER TABLE " + table_name + " DROP " + strings.ToUpper(index_type)
	} else {
		query = "ALTER TABLE " + table_name + " DROP INDEX " + tmp_index_name
	}
	execQuery(query)
	return
}
