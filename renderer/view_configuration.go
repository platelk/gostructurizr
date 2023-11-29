package renderer

import (
	"github.com/platelk/gostructurizr"
	"strings"
)

func renderViewConfiguration(c *gostructurizr.ViewConfiguration, renderer *strings.Builder, level int) error {
	return renderViewStyles(c.Styles(), renderer, level)
}
