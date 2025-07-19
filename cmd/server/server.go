package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/Thanhbinh1905/seta-training-system/internal/user/graph"
	"github.com/Thanhbinh1905/seta-training-system/pkg/config"
	"github.com/Thanhbinh1905/seta-training-system/pkg/db"
	"github.com/Thanhbinh1905/seta-training-system/pkg/logger"
	"github.com/Thanhbinh1905/seta-training-system/pkg/middleware"

	teamHandler "github.com/Thanhbinh1905/seta-training-system/internal/team/handler"
	teamRepository "github.com/Thanhbinh1905/seta-training-system/internal/team/repository"
	teamService "github.com/Thanhbinh1905/seta-training-system/internal/team/service"

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
	logger.Log.Info("Connecting to database", zap.String("url", dbUrl))

	if dbUrl == "" {
		logger.Log.Fatal("DATABASE_URL is not set")
	}

	conn, err := db.Connect(dbUrl)
	if err != nil {
		logger.Log.Fatal("failed to connect to database", zap.Error(err))
	}
	defer db.Close(conn)

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB:        conn,
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
		playground.Handler("GraphQL Playground", "/graphQL").ServeHTTP(c.Writer, c.Request)
	})

	r.Use(middleware.OptionalAuthMiddleware(cfg.JWTSecret))
	r.Use(middleware.ContextMiddleware())

	// GraphQL Query Handler
	r.POST("/graphQL", func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	})

	// Protected routes group: yêu cầu auth
	authGroup := r.Group("/")
	authGroup.Use(middleware.RequiredAuthMiddleware(cfg.JWTSecret))

	teamRepo := teamRepository.NewRepository(conn)
	teamSvc := teamService.NewTeamService(teamRepo)
	teamHdl := teamHandler.NewTeamHandler(teamSvc)

	// Routes cho team management chỉ dành cho manager
	teamGroup := authGroup.Group("/teams")
	teamGroup.Use(middleware.RequireManagerRole("manager"))
	{
		teamGroup.POST("/", teamHdl.CreateTeam)
		teamGroup.POST("/:teamId/members", teamHdl.AddMember)
		teamGroup.DELETE("/:teamId/members/:memberId", teamHdl.RemoveMember)
		teamGroup.POST("/:teamId/managers", teamHdl.AddManager)
		teamGroup.DELETE("/:teamId/managers/:managerId", teamHdl.RemoveManager)
	}

	logger.Log.Info("Starting server on port " + port)
	if err := r.Run(":" + port); err != nil {
		logger.Log.Fatal("failed to start server", zap.Error(err))
	}
}
