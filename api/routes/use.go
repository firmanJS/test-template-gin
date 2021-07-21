package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserRoutes struct
type UseRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup user routes
func (s UserRoutes) SetupUse() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/api").Use(s.authMiddleware.Handler())
	{
		api.GET("/use", s.userController.GetUser)
		api.GET("/use/:id", s.userController.GetOneUser)
		api.POST("/use", s.userController.SaveUser)
		api.POST("/use/:id", s.userController.UpdateUser)
		api.DELETE("/use/:id", s.userController.DeleteUser)
	}
}

// NewUserRoutes creates new user controller
func NewUseRoutes(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controllers.UserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}
