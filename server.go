package main

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/veritem/api/graph"
	"github.com/veritem/api/graph/generated"
)

const defaultPort = ":8080"

// graphql handler

func graphqlHandler() gin.HandlerFunc {

	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground handler

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	fmt.Println("hi")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	router := gin.Default()
	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run(defaultPort)
}
