package server

import (
	"github.com/labstack/echo/v4"
	"new-go-api/internal/database"
	"new-go-api/internal/handler"
	"new-go-api/internal/service"
)

type App struct {
	*echo.Echo
	Q           *database.Queries
	handlers    *handler.Handler
	UserService service.UserService
}

func NewApp(q *database.Queries) *App {
	e := echo.New()
	u := service.NewUserService(q)
	h := handler.NewHandler(u)
	srv := &App{
		Echo:        e,
		Q:           q,
		handlers:    h,
		UserService: u,
	}
	return srv
}

func (app *App) Start(addr string) {
	app.routes()
	app.Logger.Fatal(app.Echo.Start(addr))
}

func (app *App) routes() {
	app.GET("/", app.handlers.HelloWorld)
	app.POST("/register", app.handlers.RegisterUser)
	app.POST("/login", app.handlers.Login)
}
