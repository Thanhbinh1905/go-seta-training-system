package main

import (
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	sqlc "github.com/Thanhbinh1905/seta-training-system/internal/db/sqlc"
	"github.com/Thanhbinh1905/seta-training-system/internal/graph"
	"github.com/Thanhbinh1905/seta-training-system/pkg/logger"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
)

const defaultPort = "8080"

func main() {
	godotenv.Load()

	logger.InitLogger(true)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger.Log.Info("Starting server on port", zap.String("port", port))

	// Connect to the database
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		logger.Log.Fatal("DATABASE_URL is not set")
	}

	conn, err := sqlc.Connect(os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Log.Error("failed to connect to database", zap.Error(err))
	}
	defer sqlc.Close()

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Queries: conn,
	}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	logger.Log.Info("Starting server on port " + port)
	logger.Log.Info("Server started successfully", zap.String("port", port))

	// 🧠 Đây là chỗ thiếu nè
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		logger.Log.Fatal("failed to start server", zap.Error(err))
	}
}
