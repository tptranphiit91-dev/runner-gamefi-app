package main

import (
	"fmt"
	"log"

	"booking/config"
	"booking/delivery/http"
	"booking/delivery/http/handler"
	"booking/infrastructure/database"
	"booking/infrastructure/observer"
	"booking/usecase/user"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	fmt.Printf("🔧 Database Type: %s\n", cfg.DatabaseType)

	// Initialize Observer Pattern
	subject := observer.NewSubject()

	// Attach observers
	logger := observer.NewUserEventLogger()
	notifier := observer.NewUserEventNotifier()
	subject.Attach(logger)
	subject.Attach(notifier)

	fmt.Println("✅ Observers attached")

	// Initialize Database Factory (Factory Pattern for Database Selection)
	dbFactory := database.NewDatabaseFactory(cfg, subject)
	defer dbFactory.Close()

	// Create user repository using factory
	userRepo, err := dbFactory.CreateUserRepository()
	if err != nil {
		log.Fatal("Failed to create user repository:", err)
	}

	fmt.Printf("✅ Database connected successfully (%s)\n", dbFactory.GetDatabaseType())

	// Initialize password hasher (Strategy Pattern)
	passwordHasher := user.NewBcryptHasher(10)

	// Initialize use cases with Functional Options Pattern
	userUseCase := user.NewUserUseCase(
		userRepo,
		passwordHasher,
		user.WithEmailValidation(true),
		user.WithPasswordValidation(true),
		user.WithPasswordLength(8, 72),
	)

	fmt.Println("✅ Use cases initialized")

	// Initialize handler factory (Factory Pattern)
	handlerFactory := handler.NewHandlerFactory(userUseCase)

	// Initialize router
	router := http.NewRouter(handlerFactory)
	router.SetupRoutes()

	fmt.Println("✅ Routes configured")

	// Start server
	addr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("🚀 Server starting on %s\n", addr)
	fmt.Println("📚 API Documentation:")
	fmt.Println("   - Health Check: GET /health")
	fmt.Println("   - Create User:  POST /api/v1/users")
	fmt.Println("   - List Users:   GET /api/v1/users")
	fmt.Println("   - Get User:     GET /api/v1/users/:id")
	fmt.Println("   - Update User:  PUT /api/v1/users/:id")
	fmt.Println("   - Delete User:  DELETE /api/v1/users/:id")

	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
