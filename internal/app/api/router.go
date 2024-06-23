package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

// InitRouter initializes routes for the API server.
func (s *Server) InitRouter() *fiber.App {
	app := fiber.New()

	v1 := app.Group("api/v1")
	s.initUserRoutes(v1)
	s.initPingRoute(v1)
	s.initMenuRoutes(v1)

	// Print all routes when server starts
	printRoutes(app)

	return app
}

// initUserRoutes initializes user-related routes.
func (s *Server) initUserRoutes(group fiber.Router) {
	userRoutes := group.Group("/user")
	userRoutes.Post("/login", s.userHandler.Login)
	userRoutes.Post("/register", s.userHandler.Register)
}

// initUserRoutes initializes user-related routes.
func (s *Server) initMenuRoutes(group fiber.Router) {
	menuRoutes := group.Group("/menu/v1")
	menuRoutes.Post("/", s.menuHandler.AddMenu)
	menuRoutes.Delete("/user/:user_id/", s.menuHandler.UpdateActiveMenuBatchByUserID)
	menuRoutes.Delete("/:menu_id/", s.menuHandler.UpdateActiveMenuByMenuID)
	menuRoutes.Patch("/:menu_id/", s.menuHandler.UpdateMenuByMenuID)
	menuRoutes.Get("/:user_id/", s.menuHandler.GetMenuByUserID)
}

// initPingRoute initializes the ping route.
func (s *Server) initPingRoute(group fiber.Router) {
	group.Get("/ping", func(c fiber.Ctx) error {
		return c.SendString("Pong")
	})
}

// printRoutes prints all registered routes.
func printRoutes(app *fiber.App) {
	for _, route := range app.GetRoutes() {
		logger.Log.Infof("[%s] %s", route.Method, route.Path)
	}
}
