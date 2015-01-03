package main

import (
	"encoding/json"
	"fmt"
	"io"
)

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

func Decode(r io.Reader) (x *Migration, err error) {
	x = new(Migration)
	err = json.NewDecoder(r).Decode(x)
	return
}
