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
	case "rename_table", "rt":
		fn_rename_table(&mm.Up, &mm.Down)
	case "add_column", "ac":
		fn_add_column(&mm.Up)
		fn_drop_column(&mm.Down)
	case "drop_column", "dc":
		fn_drop_column(&mm.Up)
		fn_add_column(&mm.Down)
	case "change_column", "cc":
		fn_change_column(&mm.Up, &mm.Down)
	case "rename_column", "rc":
		fn_rename_column(&mm.Up, &mm.Down)
	case "add_index", "ai":
		fn_add_index(&mm.Up, &mm.Down)
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

func fn_add_index(mm_up *m.UpDown, mm_down *m.UpDown) {

}

func fn_change_column(mm_up *m.UpDown, mm_down *m.UpDown) {
}

func fn_rename_column(mm_up *m.UpDown, mm_down *m.UpDown) {
}

func fn_rename_table(mm_up *m.UpDown, mm_down *m.UpDown) {
}

func fn_add_column(mm *m.UpDown) {
	ac := m.AddColumn{}
	ac.Table_Name = os.Args[3]
	fieldArray := os.Args[4:len(os.Args)]
	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")
		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			col := m.Columns{}
			col.FieldName = split[0][1]
			col.DataType = split[0][2]
			ac.Columns = append(ac.Columns, col)
		}
	}
	mm.Add_Column = append(mm.Add_Column, ac)
}

func fn_drop_column(mm *m.UpDown) {
	dc := m.DropColumn{}
	dc.Table_Name = os.Args[3]
	fieldArray := os.Args[4:len(os.Args)]
	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")
		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			col := m.Columns{}
			col.FieldName = split[0][1]
			col.DataType = "" // split[0][2]   // Ignore this value as its not needed for Removing Columns.
			dc.Columns = append(dc.Columns, col)
		}
	}
	mm.Drop_Column = append(mm.Drop_Column, dc)
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
			fmt.Println("Executing ................ ", filename)
			bs, err := ioutil.ReadFile(filename)
			if err != nil {
				fmt.Println(err)
				return
			}
			docScript := []byte(bs)
			var mm m.Migration
			json.Unmarshal(docScript, &mm)
			a := "mysql"
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
