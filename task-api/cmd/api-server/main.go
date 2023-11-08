package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"task-api/internal/adapters/in/rest"
	localMiddlewares "task-api/internal/adapters/in/rest/middleware"
	"task-api/internal/adapters/out/database/inmemory"
	"task-api/internal/adapters/out/encryption"
	"task-api/internal/adapters/out/notification"
	"task-api/internal/adapters/out/token"
	"task-api/internal/core/usecase"
)

func newRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	return router
}

func main() {
	var (
		inMemoryUserDB    = inmemory.NewUserInMemDBRepository()
		jwtManager        = token.NewJWTManager("ULTRA_SECRET_VALUE")
		fakeNotifier      = notification.NewFakeNotificationSender()
		passwordEncryptor = encryption.NewHashPasswordEncryptor()
	)

	useCaseRegister := usecase.NewAuthUseCase(inMemoryUserDB, jwtManager, fakeNotifier, passwordEncryptor)
	authHandler := rest.NewAuthHandler(useCaseRegister)
	jwtMiddleware := localMiddlewares.NewJWTMiddleware(jwtManager)

	router := newRouter()

	router.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/register", authHandler.Register)
		authRouter.Post("/login", authHandler.Authenticate)
	})

	router.Route("/api/tasks", func(taskRouter chi.Router) {
		taskRouter.Use(jwtMiddleware.Validate)
	})

	http.ListenAndServe("localhost:8080", router)
}
