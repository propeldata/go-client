package client

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/hasura/go-graphql-client"
	"github.com/pkg/errors"

	"github.com/propeldata/go-client/models"
)

const (
	apiURL = "https://api.us-east-2.propeldata.com/graphql"
)

type ApiClient struct {
	client *graphql.Client
}

func NewApiClient(accessToken string) *ApiClient {
	httpClient := newHttpClient(Options{
		Timeout: 3 * time.Second,
		Retries: 3,
		Delay:   5 * time.Millisecond,
		Transport: &withHeaders{
			headers:   map[string]string{"Authorization": "Bearer " + accessToken},
			transport: http.DefaultTransport,
		},
	})

	c := graphql.NewClient(apiURL, httpClient)

	return &ApiClient{client: c}
}

type withHeaders struct {
	headers   map[string]string
	transport http.RoundTripper
}

func (wh *withHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range wh.headers {
		req.Header.Add(k, v)
	}
	return wh.transport.RoundTrip(req)
}

type CreateDataSourceOpts struct {
	Name      string
	BasicAuth *models.HttpBasicAuthInput
	Columns   []*models.WebhookDataSourceColumnInput
	Timestamp string
	UniqueID  string
}

func (c *ApiClient) CreateDataSource(ctx context.Context, opts CreateDataSourceOpts) (*models.DataSource, error) {
	variables := map[string]any{
		"input": models.CreateWebhookDataSourceInput{
			UniqueName:  opts.Name,
			Description: "Fivetran Propel destination",
			ConnectionSettings: models.WebhookConnectionSettingsInput{
				BasicAuth: opts.BasicAuth,
				Columns:   opts.Columns,
				Timestamp: opts.Timestamp,
				UniqueID:  &opts.UniqueID,
			},
		},
	}

	var mutation models.CreateWebhookDataSource
	if err := c.client.Mutate(ctx, &mutation, variables); err != nil {
		return nil, err
	}

	return mutation.DataSourceResponse.DataSource, nil
}

func (c *ApiClient) FetchDataSource(ctx context.Context, uniqueName string) (*models.DataSource, error) {
	variables := map[string]any{
		"dataSourceUniqueName": uniqueName,
	}

	query := models.FetchDataSourceByName{}

	if err := c.client.Query(ctx, &query, variables); err != nil {
		return nil, err
	}

	return query.DataSource, nil
}

func (c *ApiClient) FetchDataPool(ctx context.Context, uniqueName string) (*models.DataPool, error) {
	variables := map[string]any{
		"dataPoolUniqueName": uniqueName,
	}

	query := models.FetchDataPoolByName{}

	if err := c.client.Query(ctx, &query, variables); err != nil {
		return nil, err
	}

	return query.DataPool, nil
}

func (c *ApiClient) CreateDeletionJob(ctx context.Context, dataPoolId string, filters []models.FilterInput) (*models.Job, error) {
	variables := map[string]any{
		"input": models.CreateDeletionJobInput{
			DataPool: dataPoolId,
			Filters:  filters,
		},
	}

	var mutation models.CreateDeletionJob
	if err := c.client.Mutate(ctx, &mutation, variables); err != nil {
		return nil, err
	}

	return mutation.DeletionJobResponse.Job, nil
}

func (c *ApiClient) FetchDeletionJob(ctx context.Context, id string) (*models.Job, error) {
	variables := map[string]any{
		"id": id,
	}

	query := models.FetchDeletionJob{}

	if err := c.client.Query(ctx, &query, variables); err != nil {
		return nil, err
	}

	return query.Job, nil
}

func (c *ApiClient) FetchRecordsByUniqueId(ctx context.Context, dataPoolName string, uniqueIds []string, columnNames []string) (*models.RecordsByUniqueIdResponse, error) {
	variables := map[string]any{
		"input": models.RecordsByUniqueIdInput{
			DataPool:  models.DataPoolInput{Name: dataPoolName},
			Columns:   columnNames,
			UniqueIDs: uniqueIds,
		},
	}

	query := models.RecordsByUniqueId{}

	if err := c.client.Query(ctx, &query, variables); err != nil {
		return nil, err
	}

	return query.Records, nil
}

func (c *ApiClient) CreateAddColumnJob(ctx context.Context, dataPoolId string, columnName string, columnType string) (*models.Job, error) {
	variables := map[string]any{
		"input": models.CreateAddColumnToDataPoolJobInput{
			DataPool:     dataPoolId,
			ColumnName:   columnName,
			ColumnType:   columnType,
			JsonProperty: columnName,
		},
	}

	var mutation models.CreateAddColumnToDataPoolJob
	if err := c.client.Mutate(ctx, &mutation, variables); err != nil {
		return nil, err
	}

	return mutation.CreateAddColumnToDataPoolJob.Job, nil
}

func (c *ApiClient) FetchAddColumnJob(ctx context.Context, id string) (*models.Job, error) {
	variables := map[string]any{
		"id": id,
	}

	query := models.FetchAddColumnJob{}

	if err := c.client.Query(ctx, &query, variables); err != nil {
		return nil, err
	}

	return query.Job, nil
}

func NotFoundError(resourceName string, err error) bool {
	var graphqlErrs graphql.Errors

	if errors.As(err, &graphqlErrs) {
		for _, graphqlErr := range graphqlErrs {
			if strings.HasPrefix(graphqlErr.Message, resourceName) && strings.HasSuffix(graphqlErr.Message, " not found") {
				for key, value := range graphqlErr.Extensions {
					if key == "code" && value == "NOT_FOUND" {
						return true
					}
				}
			}
		}
	}

	return false
}
