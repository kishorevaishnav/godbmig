package migration

type Migration struct {
	Id   string `json:"id,omitempty"`
	Up   UpDown `json:"up,omitempty"`
	Down UpDown `json:"down,omitempty"`
}

type UpDown struct {
	Create_Table  []CreateTable  `json:"create_table,omitempty"`
	Drop_Table    []DropTable    `json:"drop_table,omitempty"`
	Add_Column    []AddColumn    `json:"add_column,omitempty"`
	Remove_Column []RemoveColumn `json:"remove_column,omitempty"`
	Add_Index     []AddIndex     `json:"add_index,omitempty"`
	Remove_Index  []RemoveIndex  `json:"remove_index,omitempty"`
	Sql           string         `json:"sql,omitempty"`
}

type CreateTable struct {
	Table_Name string    `json:"table_name,omitempty"`
	Columns    []Columns `json:"columns,omitempty"`
}

type DropTable struct {
	Table_Name string `json:"table_name,omitempty"`
}

type AddColumn struct {
	Table_Name string    `json:"table_name,omitempty"`
	Columns    []Columns `json:"columns,omitempty"`
}

type RemoveColumn struct {
	Table_Name string    `json:"table_name,omitempty"`
	Columns    []Columns `json:"columns,omitempty"`
}

type AddIndex struct {
	Table_Name string    `json:"table_name,omitempty"`
	Index_Type string    `json:"index_type,omitempty"`
	Columns    []Columns `json:"columns,omitempty"`
}

type RemoveIndex struct {
	Table_Name string    `json:"table_name,omitempty"`
	Index_Type string    `json:"index_type,omitempty"`
	Columns    []Columns `json:"columns,omitempty"`
}

type Columns struct {
	FieldName string `json:"fieldname,omitempty"`
	DataType  string `json:"datatype,omitempty"`
}
