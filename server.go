package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/veritem/api/pkg/db"
	"github.com/veritem/api/pkg/graph"
	"github.com/veritem/api/pkg/graph/generated"
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

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// if os.Getenv("APP_ENV") != "production" {
	// 	err := godotenv.Load()

	// 	if err != nil {
	// 		log.Fatal("Error while loading .env")
	// 	}

	// }

	db.Connect()

	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())
	router.Run(defaultPort)
}
