package binder

import "github.com/gofiber/fiber/v2"

type Binder struct {
	app     *fiber.App
	handler *Handler
	mw      *Middleware
}

func NewBinder(app *fiber.App, handler *Handler) *Binder {
	return &Binder{
		app:     app,
		handler: handler,
		mw:      New("asdqwe2131241eqeqw"),
	}
}

func (b *Binder) BindRoutes() {
	b.mapAdmin()
	b.mapClient()

}

func (b *Binder) mapAdmin() {
	client := b.app.Group("/admin")

	v1 := client.Group("/v1")

	{
		auth := v1.Group("/auth")

		auth.Post("/sign-up", b.handler.Auth.SignUp)

		auth.Post("/sign-in", b.handler.Auth.SignIn)

		auth.Post("/token", b.mw.Auth, b.handler.Auth.PostInviteToken)
	}

	{
		guide := v1.Group("/guide")

		guide.Post("/", b.mw.Auth, b.handler.Guide.CreateGuide)

		guide.Delete("/:guide_id", b.mw.Auth, b.handler.Guide.DeleteGuide)

		guide.Delete("/theme/:theme_id", b.mw.Auth, b.handler.Guide.DeleteTheme)
	}

	{
		news := v1.Group("/news")

		news.Post("/", b.mw.Auth, b.handler.News.PostNews)

	}

	{
		project := v1.Group("/project")

		project.Post("/", b.handler.Project.PostProject)

		project.Delete("/:project_id", b.mw.Auth, b.handler.Project.DeleteProject)
	}

	{
		documents := v1.Group("/documents")

		documents.Post("/", b.mw.Auth, b.handler.Documents.PostDocument)

		documents.Delete("/:document_id", b.mw.Auth, b.handler.Documents.DeleteDocument)
	}
}

func (b *Binder) mapClient() {
	client := b.app.Group("/client")

	v1 := client.Group("/v1")

	{
		guide := v1.Group("/guide")

		guide.Get("/", b.handler.Guide.GetGuide)
	}

	{
		project := v1.Group("/project")

		project.Get("/", b.handler.Project.GetProjects)

		project.Get("/:project_id", b.handler.Project.GetProject)
	}

	{
		news := v1.Group("/news")

		news.Get("/", b.handler.News.GetNews)

		news.Get("/:new_id", b.handler.News.GetNew)
	}

	{
		documents := v1.Group("/documents")

		documents.Get("/", b.handler.Documents.GetDocuments)
	}
}
