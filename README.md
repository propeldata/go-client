# Propel Go client

## Installing
Use `go get` to add it to your project's Go module dependencies.
```shell
go get github.com/propeldata/go-client
```

## Usage
1. Import the Propel go client package
```go
import (
	"github.com/propeldata/go-client"
)
```
2. Generate an OAuth token. This step requires having a Propel account and an application. Learn more about the API authentication [here](https://www.propeldata.com/docs/api/authentication).
```go
oauthClient := client.NewOauthClient()

oauthToken, err := oauthClient.OAuthToken(ctx, "APPLICATION_ID", "APPLICATION_SECRET")
if err != nil {
    return errors.New("Invalid Propel credentials")
}
```
3. Initialize the Propel client with the retrieved token.
```go
apiClient := client.NewApiClient(oauthToken.AccessToken)
```
4. Perform any API request.
```go
dataSource, err := apiClient.FetchDataSource(ctx, "dataSourceUniqueName")
if err != nil {
	return err
}
```