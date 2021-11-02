package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
)

func main() {
	ctx := datadog.NewDefaultContext(context.Background())

	pageSize := int64(10)                    // int64 | Size for a given page. (optional) (default to 10)
	pageNumber := int64(0)                   // int64 | Specific page number to return. (optional) (default to 0)
	sort := "name"                           // string | User attribute to order results by. Sort order is ascending by default. Sort order is descending if the field is prefixed by a negative sign, for example `sort=-name`. Options: `name`, `modified_at`, `user_count`. (optional) (default to "name")
	sortDir := datadog.QuerySortOrder("asc") // QuerySortOrder | Direction of sort. Options: `asc`, `desc`. (optional) (default to "desc")
	filter := ""                             // string | Filter all users by the given string. Defaults to no filtering. (optional)
	filterStatus := "Active"                 // string | Filter on status attribute. Comma separated list, with possible values `Active`, `Pending`, and `Disabled`. Defaults to no filtering. (optional)
	optionalParams := datadog.ListUsersOptionalParameters{
		PageSize:     &pageSize,
		PageNumber:   &pageNumber,
		Sort:         &sort,
		SortDir:      &sortDir,
		Filter:       &filter,
		FilterStatus: &filterStatus,
	}

	configuration := datadog.NewConfiguration()

	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersApi.ListUsers(ctx, optionalParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.ListUsers`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUsers`: UsersResponse
	responseContent, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintf(os.Stdout, "Response from UsersApi.ListUsers:\n%s\n", responseContent)
}
