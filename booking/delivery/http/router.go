package http

import (
	"booking/delivery/http/handler"
	"booking/delivery/http/middleware"
	
	"github.com/gin-gonic/gin"
)

// Router sets up HTTP routes
type Router struct {
	engine         *gin.Engine
	handlerFactory *handler.HandlerFactory
}

// NewRouter creates a new router
func NewRouter(handlerFactory *handler.HandlerFactory) *Router {
	engine := gin.Default()
	
	// Apply global middleware
	engine.Use(middleware.CORS())
	engine.Use(middleware.Logger())
	
	return &Router{
		engine:         engine,
		handlerFactory: handlerFactory,
	}
}

// SetupRoutes configures all routes
func (r *Router) SetupRoutes() {
	// Health check
	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Booking service is running",
		})
	})
	
	// API v1 routes
	v1 := r.engine.Group("/api/v1")
	{
		// User routes
		userHandler := r.handlerFactory.GetUserHandler()
		users := v1.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("", userHandler.ListUsers)
			users.GET("/:id", userHandler.GetUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}
}

// Run starts the HTTP server
func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}

// GetEngine returns the Gin engine
func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}

