package models

type Timestamp struct {
	ColumnName string `json:"columnName"`
}

type UniqueId struct {
	ColumnName string `json:"columnName"`
}

type DataPoolColumn struct {
	ColumnName string `json:"columnName"`
	Type       string `json:"type"`
	IsNullable bool   `json:"isNullable"`
}

type DataPool struct {
	ID             string           `json:"id"`
	Timestamp      Timestamp        `json:"timestamp"`
	UniqueID       UniqueId         `json:"uniqueId"`
	TableSettings  *TableSettings   `json:"tableSettings"`
	OrderByColumns []DataPoolColumn `json:"orderByColumns,omitempty"`
}

type FetchDataPoolByName struct {
	DataPool *DataPool `graphql:"dataPoolByName (uniqueName: $dataPoolUniqueName)"`
}
