package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	m "github.com/kishorevaishnav/godbmig/migration"

	gdm_my "github.com/kishorevaishnav/godbmig/mysql"
	gdm_pq "github.com/kishorevaishnav/godbmig/postgres"
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

// type migTemplate struct {
// 	id       int64
// 	mig_type string
// 	tbl_name string
// 	cols     []string
// }

func generateMigration() {
	const layout = "20060102150405"
	t := time.Now()
	mm := m.Migration{}
	mm.Id = "3" + t.Format(layout)
	switch os.Args[2] {
	case "create_table", "ct":
		fn_create_table(&mm.Up)
		fn_drop_table(&mm.Down)
	case "drop_table", "dt":
		fn_drop_table(&mm.Up)
		fn_create_table(&mm.Down)
	// case "rename_table", "rt":
	// 	fn_rename_table(&mm.Up)
	// 	fn_rename_table(&mm.Down)
	case "add_column", "ac":
		fn_add_column(&mm.Up)
		fn_remove_column(&mm.Down)
	case "remove_column", "rc":
		fn_remove_column(&mm.Up)
		fn_add_column(&mm.Down)
	default:
		panic("No or wrong Actions provided.")
	}
	b, _ := json.MarshalIndent(mm, " ", "  ")
	fmt.Println(string(b))

	// Write to a new File.
	filename := mm.Id + ".rm.json"
	file1, _ := os.Create(filename)
	file1.Write(b)
	file1.Close()

	os.Exit(1)
}

func fn_add_column(mm *m.UpDown) {
	fieldArray := os.Args[4:len(os.Args)]
	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")
		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			ac := m.AddColumn{}
			ac.Table_Name = os.Args[3]
			ac.Column_Name = split[0][1]
			ac.Data_Type = split[0][2]
			mm.Add_Column = append(mm.Add_Column, ac)
		}
	}
}

func fn_remove_column(mm *m.UpDown) {
	rc := m.RemoveColumn{}
	rc.Table_Name = os.Args[3]
	fieldArray := os.Args[4:len(os.Args)]
	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")
		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			rc.Column_Name = split[0][1]
		}
	}
	mm.Remove_Column = append(mm.Remove_Column, rc)
}

func fn_drop_table(mm *m.UpDown) {
	dt := m.DropTable{}
	dt.Table_Name = os.Args[3]
	mm.Drop_Table = append(mm.Drop_Table, dt)
}

func fn_create_table(mm *m.UpDown) {
	ct := m.CreateTable{}
	ct.Table_Name = os.Args[3]
	fieldArray := os.Args[4:len(os.Args)]
	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")
		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			col := m.Columns{}
			col.FieldName = split[0][1]
			col.DataType = split[0][2]
			ct.Columns = append(ct.Columns, col)
		}
	}
	mm.Create_Table = append(mm.Create_Table, ct)
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
