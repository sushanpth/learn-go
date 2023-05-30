package schema

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/graphql-go/graphql"
)

// Helper function to import json from file to map
func importJSONDataFromFile(fileName string, result interface{}) bool {

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}
	return true
}

var BeastList []Beast
var _ = importJSONDataFromFile("beastData.json", &BeastList)

// if err != nil {
// 	fmt.Println("Error: ", err)
// }

type Beast struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	OtherNames  []string `json:"otherNames"`
	ImageURL    string   `json:"imageUrl"`
}

var beastType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Beast",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"otherNames": &graphql.Field{
			Type: graphql.NewList(graphql.String),
		},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"beast": &graphql.Field{
			Type:        beastType,
			Description: "Get single beast",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				nameQuery, isOk := p.Args["name"].(string)
				if isOk {
					// search for el with name
					for _, beast := range BeastList {
						if beast.Name == nameQuery {
							return beast, nil
						}
					}
				}
				return Beast{}, nil
			},
		},
		"beastList": &graphql.Field{
			Type:        graphql.NewList(beastType),
			Description: "List of beasts",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return BeastList, nil
			},
		},
	},
})

var currentMaxId = 5

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addBeast": &graphql.Field{
			Type:        beastType,
			Description: "Add a new beast",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"otherNames": &graphql.ArgumentConfig{
					Type: graphql.NewList(graphql.String),
				},
				"imageUrl": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name, _ := p.Args["name"].(string)
				description, _ := p.Args["description"].(string)
				otherNames, _ := p.Args["otherNames"].([]string)
				imageUrl, _ := p.Args["imageUrl"].(string)

				newID := currentMaxId + 1
				currentMaxId++

				// fmt.Println(otherNames, "OTHER")

				newBeast := Beast{
					ID:          newID,
					Name:        name,
					Description: description,
					OtherNames:  otherNames,
					ImageURL:    imageUrl,
				}

				// fmt.Println(newBeast)

				BeastList = append(BeastList, newBeast)

				return newBeast, nil

			},
		},
	},
})

// define schema, with our rootQuery
var BeastSchema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})
