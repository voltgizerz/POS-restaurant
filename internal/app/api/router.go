package api

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	fiberLogger "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/voltgizerz/POS-restaurant/internal/app/api/middleware"
	"github.com/voltgizerz/POS-restaurant/pkg/logger"
)

// InitRouter initializes routes for the API server.
func (s *Server) InitRouter() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})
	app.Use(fiberLogger.New())

	v1 := app.Group("api/v1")
	s.initPingRoute(v1)
	s.initAuthRoutes(v1)
	s.initUserRoutes(v1)
	s.initMenuRoutes(v1)

	// Print all routes when server starts
	printRoutes(app)

	return app
}

// initAuthRoutes initializes auth-related routes.
func (s *Server) initAuthRoutes(group fiber.Router) {
	userRoutes := group.Group("/auth")
	userRoutes.Post("/login", s.authHandler.Login, middleware.Initialization)
	userRoutes.Post("/register", s.authHandler.Register, middleware.Initialization)
}

// initUserRoutes initializes user-related routes.
func (s *Server) initUserRoutes(group fiber.Router) {
	_ = group.Group("/user")
}

// initUserRoutes initializes user-related routes.
func (s *Server) initMenuRoutes(group fiber.Router) {
	menuRoutes := group.Group("/menus", )
	menuRoutes.Post("", s.menuHandler.AddMenu, s.jwtMiddleware.AuthorizeAccess())
	menuRoutes.Delete("/user/:user_id", s.menuHandler.UpdateActiveMenuBatchByUserID, s.jwtMiddleware.AuthorizeAccess())
	menuRoutes.Delete("/:menu_id", s.menuHandler.UpdateActiveMenuByMenuID, s.jwtMiddleware.AuthorizeAccess())
	menuRoutes.Patch("/:menu_id", s.menuHandler.UpdateMenuByMenuID, s.jwtMiddleware.AuthorizeAccess())
	menuRoutes.Get("/:user_id", s.menuHandler.GetMenuByUserID, s.jwtMiddleware.AuthorizeAccess())
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
		if route.Path != "/" {
			logger.Log.Infof("[%s] %s", route.Method, route.Path)
		}
	}
}
