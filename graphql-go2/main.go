package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./schema"

	"github.com/graphql-go/graphql"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema.TodoSchema)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8080", nil)
}
