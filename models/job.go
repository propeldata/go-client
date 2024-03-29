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

type SetColumnInput struct {
	Column     string `json:"column"`
	Expression string `json:"expression"`
}

type CreateUpdateDataPoolRecordsJobInput struct {
	DataPool string           `json:"dataPool"`
	Filters  []FilterInput    `json:"filters"`
	Set      []SetColumnInput `json:"set"`
}

type CreateUpdateDataPoolRecordsJob struct {
	CreateUpdateDataPoolRecordsJob struct {
		Job *Job `graphql:"job"`
	} `graphql:"createUpdateDataPoolRecordsJob(input: $input)"`
}

type FetchDeletionJob struct {
	Job *Job `graphql:"deletionJob (id: $id)"`
}

type FetchAddColumnJob struct {
	Job *Job `graphql:"addColumnToDataPoolJob (id: $id)"`
}

type FetchUpdateJob struct {
	Job *Job `graphql:"updateDataPoolRecordsJob (id: $id)"`
}
