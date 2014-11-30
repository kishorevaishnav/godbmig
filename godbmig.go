package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	m "./migration"

	gdm_my "./mysql"
	gdm_pq "./postgres"
)

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
			var mm m.Migration
			json.Unmarshal(docScript, &mm)
			a := "asmysql"
			if a == "mysql" {
				gdm_my.ProcessNow(mm)
			} else {
				gdm_pq.ProcessNow(mm)
			}
			// if !ProcessNow(mm) {
			// 	fmt.Println("Either the file is empty or not in a proper JSON Migration Format")
			// 	// break
			// }
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

func ProcessNow(m m.Migration) bool {
	a := "mysql"
	if a == "mysql" {
		gdm_my.Init()
	} else {
		gdm_pq.Init()
	}
	// gdm.Db, _ = sql.Open("mysql", "root:root@/onetest")

	// var query string
	return false
}
