package main

func main() {
	// Create a new Pug engine
	app := fiber.New()

	engine := pug.New("./index", ".pug")

	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine, // set template engine for rendering
	})

	app.Static(
		"/",        // mount address
		"./public", // path to the file folder
	)

	// Create a new endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		// rendering the "index" template with content passing
		return c.Render("index", fiber.Map{
			"Title":   "Hey!",
			"Message": "This is the index template.",
		})
	})
}
