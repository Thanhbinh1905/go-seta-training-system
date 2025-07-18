package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	sqlc "github.com/Thanhbinh1905/seta-training-system/internal/db/sqlc"
	"github.com/Thanhbinh1905/seta-training-system/internal/graph"
	"github.com/Thanhbinh1905/seta-training-system/pkg/config"
	"github.com/Thanhbinh1905/seta-training-system/pkg/logger"

	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	logger.InitLogger(true)

	cfg := config.LoadConfig()
	if cfg == nil {
		logger.Log.Fatal("failed to load configuration")
	}

	port := cfg.Port
	if port == "" {
		port = defaultPort
	}

	logger.Log.Info("Starting server on port", zap.String("port", port))

	// Connect to the database
	dbUrl := cfg.DatabaseURL
	if dbUrl == "" {
		logger.Log.Fatal("DATABASE_URL is not set")
	}

	conn, err := sqlc.Connect(dbUrl)
	if err != nil {
		logger.Log.Error("failed to connect to database", zap.Error(err))
	}
	defer sqlc.Close()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Queries:   conn,
		JWTSecret: cfg.JWTSecret,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	r := gin.Default()

	// Health check
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// GraphQL Playground
	r.GET("/", func(c *gin.Context) {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Writer, c.Request)
	})

	// GraphQL Query Handler
	r.POST("/query", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	logger.Log.Info("Starting server on port " + port)
	if err := r.Run(":" + port); err != nil {
		logger.Log.Fatal("failed to start server", zap.Error(err))
	}
}
