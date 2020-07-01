package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HelloIndex default implementation.
func HelloIndex(c buffalo.Context) error {
	c.Set("greeting", "Hello Full Cycle")
	return c.Render(http.StatusOK, r.HTML("hello/index.html"))
}
