package actions

import "github.com/gobuffalo/buffalo"

// ConnectionConnectionHandler default implementation.
func ConnectionConnectionHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("connection/connection_handler.html"))
}

