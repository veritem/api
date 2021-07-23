package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/veritem/api/pkg/db"
	"github.com/veritem/api/pkg/graph"
	"github.com/veritem/api/pkg/graph/generated"
)

const defaultPort = ":8080"

// graphql handler
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}),
	//	generated.IntrospectionEnabled(false),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	gotenv.Load() //nolint:errcheck,gocritic,nolintlint

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	err := db.Connect()

	if err != nil {
		log.Println(err)
	}

	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())

	err = router.Run(defaultPort)

	if err != nil {
		log.Println(err)
	}
}
