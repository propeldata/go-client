package models

type Job struct {
	ID       string   `json:"id"`
	Status   string   `json:"status"`
	Progress float32  `json:"progress"`
	Error    JobError `json:"error"`
}

type JobError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type FilterInput struct {
	Column   string `json:"column"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type CreateDeletionJobInput struct {
	DataPool string        `json:"dataPool"`
	Filters  []FilterInput `json:"filters"`
}

type CreateDeletionJob struct {
	DeletionJobResponse struct {
		Job *Job `graphql:"job"`
	} `graphql:"createDeletionJob(input: $input)"`
}

type CreateAddColumnToDataPoolJobInput struct {
	DataPool     string `json:"dataPool"`
	ColumnName   string `json:"columnName"`
	ColumnType   string `json:"columnType"`
	JsonProperty string `json:"jsonProperty"`
}

type CreateAddColumnToDataPoolJob struct {
	CreateAddColumnToDataPoolJob struct {
		Job *Job `graphql:"job"`
	} `graphql:"createAddColumnToDataPoolJob(input: $input)"`
}

type FetchDeletionJob struct {
	Job *Job `graphql:"deletionJob (id: $id)"`
}

type FetchAddColumnJob struct {
	Job *Job `graphql:"addColumnToDataPoolJob (id: $id)"`
}
