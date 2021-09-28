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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
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

	router.Use(CORSMiddleware())

	router.POST("/query", graphqlHandler())
	router.GET("/", playgroundHandler())

	err = router.Run(defaultPort)

	if err != nil {
		log.Println(err)
	}
}
