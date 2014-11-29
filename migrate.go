package main

import (
	"encoding/json"
	"fmt"
	"io"
)

var document string = `{
    "id": 20130106222315,
    "up": {
        "add_column": [ ["seasons", "name", "varchar"], ["seasons1", "name1", "varchar1"] ]
    },
    "down": "down_asdf"
}`

var action = "up"

type Migration struct {
	Id   int    `jpath:"id"`
	Down string `jpath:"down"`
	Up   Up     `jpath:"up"`
}

type Up struct {
	Add_Column []AddColumn  `jpath:"add_column"`
	DropColumn []DropColumn `jpath:"drop_column"`
}

type AddColumn struct {
	Table_Name []string
}

type DropColumn struct {
	TableName  string `jpath:"table_name"`
	ColumnName string `jpath:"column_name"`
}

type up Up
type addColumn AddColumn
type dropColumn DropColumn

func (u *Up) UnmarshalJSON(b []byte) (err error) {
	if action == "up" {
		fmt.Println("in up")
		j, s, n := up{}, addColumn{}, dropColumn{}
		if err = json.Unmarshal(b, &j); err == nil {
			*u = Up(j)
			return
		}
		if err = json.Unmarshal(b, &s); err == nil {
			return
		}
		if err = json.Unmarshal(b, &n); err == nil {
			return
		} else {
			fmt.Println("in err...")
		}
	}
	return
}

func (ac *AddColumn) UnmarshalJSON(b []byte) (err error) {
	// ac.Table_Name = b
	fmt.Println("in add column")
	json.Unmarshal(b, &ac.Table_Name)
	fmt.Println(string(b))
	fmt.Println("---")
	fmt.Println(ac)
	fmt.Println("---")

	// a, tn, cn, dt := addColumn{}, "", "", ""
	a, ac_Array := addColumn{}, ""
	if err = json.Unmarshal(b, &a); err == nil {
		*ac = AddColumn(a)
		fmt.Println("========================")
		fmt.Println("Add Column - Migrating")
		fmt.Println(" Table : ", ac.Table_Name)
		fmt.Println("========================")
		return
	}
	fmt.Println(ac)
	if err = json.Unmarshal(b, &ac_Array); err == nil {
		// ac.Table_Name = ac_Array
		return
	}
	// if err = json.Unmarshal(b, &cn); err == nil {
	// 	ac.Column_Name = cn
	// 	return
	// }
	// if err = json.Unmarshal(b, &dt); err == nil {
	// 	ac.Data_Type = dt
	// 	return
	// }
	return
}

func Decode(r io.Reader) (x *Migration, err error) {
	x = new(Migration)
	err = json.NewDecoder(r).Decode(x)
	return
}

func main() {
	docScript := []byte(document)
	// docMap := map[string]interface{}{}
	var mm Migration

	json.Unmarshal(docScript, &mm)

	fmt.Println(mm)
	fmt.Println(len(mm.Up.Add_Column))
	for k, v := range mm.Up.Add_Column {
		fmt.Println(k, v)
	}
	err := len(mm.Up.DropColumn)
	if err != 0 {
		fmt.Println(err)
		// for k, v := range mm.Up.Drop_Column {
		// 	fmt.Println(k, v)
		// }
	}
	fmt.Println(err)
}
