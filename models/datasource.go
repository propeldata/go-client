package models

type HttpBasicAuthInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type HttpBasicAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type WebhookDataSourceColumnInput struct {
	Name         string     `json:"name"`
	Type         PropelType `json:"type"`
	Nullable     bool       `json:"nullable"`
	JsonProperty string     `json:"jsonProperty"`
}

type WebhookConnectionSettingsInput struct {
	BasicAuth *HttpBasicAuthInput             `json:"basicAuth,omitempty"`
	Columns   []*WebhookDataSourceColumnInput `json:"columns,omitempty"`
	Timestamp string                          `json:"timestamp"`
	UniqueID  *string                         `json:"uniqueId,omitempty"`
}

type CreateWebhookDataSourceInput struct {
	UniqueName         string                         `json:"uniqueName,omitempty"`
	Description        string                         `json:"description,omitempty"`
	ConnectionSettings WebhookConnectionSettingsInput `json:"connectionSettings"`
}

type WebhookConnectionSettings struct {
	WebhookURL string          `json:"webhookUrl"`
	BasicAuth  *HttpBasicAuth  `json:"basicAuth"`
	Columns    []WebhookColumn `json:"columns"`
	Timestamp  string          `json:"timestamp"`
	UniqueID   string          `json:"uniqueID"`
}

type WebhookColumn struct {
	Name         string     `json:"name"`
	Type         PropelType `json:"type"`
	Nullable     bool       `json:"nullable"`
	JsonProperty string     `json:"jsonProperty"`
}

type ConnectionSettings struct {
	WebhookConnectionSettings WebhookConnectionSettings `graphql:"... on WebhookConnectionSettings"`
}

type DataSource struct {
	ID                 string             `json:"id"`
	Status             string             `json:"status"`
	UniqueName         string             `json:"uniqueName"`
	ConnectionSettings ConnectionSettings `json:"connectionSettings"`
}

type CreateWebhookDataSource struct {
	DataSourceResponse struct {
		DataSource *DataSource `graphql:"dataSource"`
	} `graphql:"createWebhookDataSource(input: $input)" json:"dataSourceResponse"`
}

type FetchDataSourceByName struct {
	DataSource *DataSource `graphql:"dataSourceByName (uniqueName: $dataSourceUniqueName)"`
}
