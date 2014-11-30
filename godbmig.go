package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"

	m "./migration"

	gdm_my "./mysql"
	gdm_pq "./postgres"
)

const FIELD_DATATYPE_REGEXP = `^([A-Za-z]{2,15}):([A-Za-z]{2,15})`

func main() {
	if strings.LastIndex(os.Args[0], "godbmig") < 1 {
		fmt.Println("wrong usage")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "add", "a":
		generateMigration()
	case "up", "u":
		migrateUpDown("up")
	case "down", "d":
		migrateUpDown("down")
	default:
		panic("No or Wrong Actions provided.")
	}
	os.Exit(1)
}

func generateMigration() {
	var gm m.Migration
	ct := []gm.Create_Table{}
	col := []gm.Columns{}
	gm.Up.Create_Table[0].Table_Name = os.Args[3]

	fieldArray := os.Args[4:len(os.Args)]

	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")

		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			fmt.Println(split[0][1])
			fmt.Println(split[0][2])
		}

		// fds, err := fld_dtype_sep(fieldArray[key])
		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// f_name, f_data_type, f_required, f_min, f_max := fds[0], fds[1], fds[2], fds[3], fds[4]
	}

}

func migrateUpDown(updown string) {
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
				gdm_my.ProcessNow(mm, updown)
			} else {
				gdm_pq.ProcessNow(mm, updown)
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
