package models

type Timestamp struct {
	ColumnName string `json:"columnName"`
}

type UniqueId struct {
	ColumnName string `json:"columnName"`
}

type DataPool struct {
	ID        string    `json:"id"`
	Timestamp Timestamp `json:"timestamp"`
	UniqueID  UniqueId  `json:"uniqueId"`
}

type FetchDataPoolByName struct {
	DataPool *DataPool `graphql:"dataPoolByName (uniqueName: $dataPoolUniqueName)"`
}
