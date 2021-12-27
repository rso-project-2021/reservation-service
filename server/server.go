package server

import (
	"reservation-service/config"
	"reservation-service/db"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	config config.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {

	gin.SetMode(config.GinMode)
	router := gin.Default()

	server := &Server{
		config: config,
		store:  store,
	}

	// Setup routing for server.
	v1 := router.Group("v1")
	{
		v1.GET("/reservations/:id", server.GetByID)
		v1.GET("/reservations", server.GetAll)
		v1.POST("/reservations", server.Create)
		v1.PUT("/reservations/:id", server.Update)
		v1.DELETE("/reservations/:id", server.Delete)
		v1.GET("/reservations/user/:id", server.GetAllByUserID)
	}

	// Setup health check routes.
	health := router.Group("health")
	{
		health.GET("/live", server.Live)
		health.GET("/ready", server.Ready)
	}

	// Setup metrics routes.
	metrics := router.Group("metrics")
	{
		metrics.GET("/", func(ctx *gin.Context) {
			handler := promhttp.Handler()
			handler.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
