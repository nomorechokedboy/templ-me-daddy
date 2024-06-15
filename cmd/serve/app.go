package serve

import (
	"templ-me-daddy/internals/todo"
	"templ-me-daddy/web/pages/about"
	"templ-me-daddy/web/pages/root"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func healthz(c *fiber.Ctx) error {
	return c.JSON(getHealthCheckResonse())
}

func (a *App) RunHttpServer() {
	app := fiber.New()

	todos := []*todo.Todo{{Id: "uuid-here", Description: "Lmao"}}
	app.Use(func(c *fiber.Ctx) error {
		useLayout := string(c.Request().Header.Peek("hx-request")) != "true"
		c.Locals("useLayout", useLayout)
		return c.Next()
	})
	app.Get("/healthz", healthz)
	app.Get("/", func(c *fiber.Ctx) error {
		rootPage := root.Root(todos)
		return render(c, root.Index(rootPage))
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		aboutPage := about.About()
		useLayout, ok := c.Locals("useLayout").(bool)
		if !ok {
			return render(c, about.Index(aboutPage))
		}

		page := aboutPage
		if useLayout {
			page = about.Index(aboutPage)
		}

		return render(c, page)
	})

	app.Listen(":3000")
}

func render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}

	return adaptor.HTTPHandler(componentHandler)(c)
}
