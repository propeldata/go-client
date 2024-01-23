package models

type DataPoolInput struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RecordsByUniqueIdInput struct {
	DataPool  DataPoolInput `json:"dataPool"`
	Columns   []string      `json:"columns"`
	UniqueIDs []string      `json:"uniqueIds"`
}

type RecordsByUniqueIdResponse struct {
	Columns []string   `json:"columns"`
	Values  [][]string `json:"values"`
}

type RecordsByUniqueId struct {
	Records *RecordsByUniqueIdResponse `graphql:"recordsByUniqueId (input: $input)"`
}
