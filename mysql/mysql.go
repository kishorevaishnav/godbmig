package mysql

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

// func Init() {
// 	Db, _ = sql.Open("mysql", "root:root@/onetest")
// }

func execQuery(query string) {
	fmt.Println(query)
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
	// sort.Strings(field)
	// fmt.Println(field)
	tmp_index_name := strings.Join(field, ',') + "_index"
	tmp_index_name = strings.ToLower(tmp_index_name)
	query := "CREATE " + strings.ToUpper(index_type) + " INDEX " + tmp_index_name + table_name + " ON " + strings.Join(field, ',')
	execQuery(query)
	return
}

func RemoveIndex(table_name string, index_type string, field []string) {
	sort.Strings(field)
	tmp_index_name := strings.ToLower(strings.Join(field, ',') + "_index")
	query := ""
	if index_type != "" && index_type != nil {
		query = "ALTER TABLE " + table_name + " DROP " + strings.ToUpper(index_type)
	} else {
		query = "ALTER TABLE " + table_name + " DROP INDEX " + tmp_index_name
	}
	execQuery(query)
	return
}
