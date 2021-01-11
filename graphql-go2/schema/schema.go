package schema

import (
	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"text": &graphql.Field{
			Type: graphql.String,
		},
		"done": &graphql.Field{
			Type: graphql.Boolean,
		},
	},
})

var todoQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{

		"todoList": &graphql.Field{
			Type: graphql.NewList(todoType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return []Todo{
					Todo{ID: "a", Text: "A todo not to forget", Done: false},
					Todo{ID: "b", Text: "the most important", Done: false},
					Todo{ID: "c", Text: "Please do this or else", Done: false},
				}, nil
			},
		},
	},
})

var TodoSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: todoQuery,
})
