package migration

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
	Index_Type string    `jpath:"index_type"`
	Columns    []Columns `jpath:"columns"`
}

type Columns struct {
	FieldName string `jpath:"fieldname"`
	DataType  string `jpath:"datatype"`
}
