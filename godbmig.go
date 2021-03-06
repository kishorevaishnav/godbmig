package main

import (
	"encoding/json"
	"flag"
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

var Config m.Config
var ArgArr []string

func main() {
	if strings.LastIndex(os.Args[0], "godbmig") < 1 {
		fmt.Println("wrong usage")
		os.Exit(1)
	}
	switch ArgArr[0] {
	case "add", "a":
		generateMigration()
	case "up", "u":
		migrateUpDown("up")
	case "down", "d":
		migrateUpDown("down")
	case "create", "c":
		createMigration()
	default:
		panic("No or Wrong Actions provided.")
	}
	os.Exit(1)
}

func init() {
	// fmt.Println("godbmig init() it runs before other functions")
	var un, pw, dbname, host, port string
	flag.StringVar(&un, "u", "", "specify the database username")
	flag.StringVar(&pw, "p", "", "specify the database password")
	flag.StringVar(&dbname, "d", "", "specify the database name")
	flag.StringVar(&host, "h", "localhost", "specify the database hostname")
	flag.StringVar(&port, "port", "5432", "specify the database port")
	flag.Parse()
	Config.Db_username = un
	Config.Db_password = pw
	Config.Db_name = dbname
	Config.Db_hostname = host
	Config.Db_portnumber = port
	ArgArr = flag.Args()
}

func createMigration() {
	a := "mysql"
	if a == "mysql" {
		gdm_my.Init(Config)
		gdm_my.CreateMigrationTable()
	} else {
		gdm_pq.Init()
		gdm_pq.CreateMigrationTable()
	}
}

func generateMigration() {
	const layout = "20060102150405"
	t := time.Now()
	mm := m.Migration{}
	mm.Id = "3" + t.Format(layout)
	switch ArgArr[1] {
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
	// case "change_column", "cc":
	// 	fn_change_column(&mm.Up, &mm.Down)
	// case "rename_column", "rc":
	// 	fn_rename_column(&mm.Up, &mm.Down)
	// case "add_index", "ai":
	// 	fn_add_index(&mm.Up, &mm.Down)
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
	// ct := m.CreateTable{}
	// ct.Table_Name = ArgArr[2]
	// fieldArray := ArgArr[3:len(ArgArr)]
	// for key, value := range fieldArray {
	// 	fieldArray[key] = strings.Trim(value, ", ")
	// 	r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
	// 	if r.MatchString(fieldArray[key]) == true {
	// 		split := r.FindAllStringSubmatch(fieldArray[key], -1)
	// 		col := m.Columns{}
	// 		col.FieldName = split[0][1]
	// 		col.DataType = split[0][2]
	// 		ct.Columns = append(ct.Columns, col)
	// 	}
	// }
	// mm.Create_Table = append(mm.Create_Table, ct)

	// ai := m.AddIndex{}
	// ai.Table_Name = ArgArr[2]
	// fieldArray = ArgArr[3:len(ArgArr)]
	// for key, value := range fieldArray {
	// 	fieldArray[key] = strings.Trim(value, ", ")
	// 	r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
	// 	if r.MatchString(fieldArray[key]) == true {
	// 		split := r.FindAllStringSubmatch(fieldArray[key], -1)
	// 		col := m.Columns{}
	// 		col.FieldName = split[0][1]
	// 		col.DataType = split[0][2]
	// 		ai.Columns = append(ai.Columns, col)
	// 	}
	// }
	// mm_up.Add_Index = append(mm_up.Add_Index, ai)
}

func fn_change_column(mm_up *m.UpDown, mm_down *m.UpDown) {
}

func fn_rename_column(mm_up *m.UpDown, mm_down *m.UpDown) {
}

func fn_rename_table(mm_up *m.UpDown, mm_down *m.UpDown) {

}

func fn_add_column(mm *m.UpDown) {
	ac := m.AddColumn{}
	ac.Table_Name = ArgArr[2]
	fieldArray := ArgArr[3:len(ArgArr)]
	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")
		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			col := m.Columns{}
			col.FieldName = split[0][1]
			col.DataType = split[0][2]
			ac.Columns = append(ac.Columns, col)
		} else {
			ac = m.AddColumn{}
		}
	}
	mm.Add_Column = append(mm.Add_Column, ac)
}

func fn_drop_column(mm *m.UpDown) {
	dc := m.DropColumn{}
	dc.Table_Name = ArgArr[2]
	fieldArray := ArgArr[3:len(ArgArr)]
	for key, value := range fieldArray {
		fieldArray[key] = strings.Trim(value, ", ")
		r, _ := regexp.Compile(FIELD_DATATYPE_REGEXP)
		col := m.Columns{}
		if r.MatchString(fieldArray[key]) == true {
			split := r.FindAllStringSubmatch(fieldArray[key], -1)
			col.FieldName = split[0][1]
			col.DataType = split[0][2]
			dc.Columns = append(dc.Columns, col)
		} else if fieldArray[key] != "" {
			col.FieldName = fieldArray[key]
			dc.Columns = append(dc.Columns, col)
		} else {
			dc = m.DropColumn{}
		}
	}
	mm.Drop_Column = append(mm.Drop_Column, dc)
}

func fn_drop_table(mm *m.UpDown) {
	dt := m.DropTable{}
	dt.Table_Name = ArgArr[2]
	mm.Drop_Table = append(mm.Drop_Table, dt)
}

func fn_create_table(mm *m.UpDown) {
	ct := m.CreateTable{}
	ct.Table_Name = ArgArr[2]
	fieldArray := ArgArr[3:len(ArgArr)]
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
				gdm_my.Init(Config)
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
