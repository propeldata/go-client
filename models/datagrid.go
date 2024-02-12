package models

import "time"

type SortOrder string
type RelativeTimeRange string

const (
	SortAsc  = "ASC"
	SortDesc = "DESC"
)

type TimeRangeInput struct {
	Relative string    `json:"relative,omitempty"`
	N        int       `json:"n,omitempty"`
	Start    time.Time `json:"start,omitempty"`
	Stop     time.Time `json:"stop,omitempty"`
}

type DataGridInput struct {
	DataPool  DataPoolInput  `json:"dataPool"`
	Columns   []string       `json:"columns"`
	First     int            `json:"first,omitempty"`
	After     string         `json:"after,omitempty"`
	Last      int            `json:"last,omitempty"`
	Before    string         `json:"before,omitempty"`
	Sort      SortOrder      `json:"sort,omitempty"`
	TimeRange TimeRangeInput `json:"timeRange,omitempty"`
}

type PageInfoResponse struct {
	StartCursor     string `json:"startCursor"`
	EndCursor       string `json:"endCursor"`
	HasNextPage     bool   `json:"hasNextPage"`
	HasPreviousPage bool   `json:"hasPreviousPage"`
}

type DataGridResponse struct {
	Headers  []string         `json:"headers"`
	Rows     [][]*string      `json:"rows"`
	PageInfo PageInfoResponse `json:"pageInfo"`
}

type DataGrid struct {
	DataGrid *DataGridResponse `graphql:"dataGrid(input: $input)"`
}
