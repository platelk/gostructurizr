package renderer

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"io"
	"strings"
)

type DSLRenderer struct {
	writer io.Writer
}

func NewDSLRenderer(writer io.Writer) *DSLRenderer {
	return &DSLRenderer{
		writer: writer,
	}
}

func (r *DSLRenderer) Render(w *gostructurizr.WorkspaceNode) error {
	return renderWrapper(r.writer, func(renderer *strings.Builder) error {
		return renderWorkspace(w, renderer, 0)
	})
}

func renderWrapper(writer io.Writer, renderNode func(renderer *strings.Builder) error) error {
	rendered := &strings.Builder{}
	if err := renderNode(rendered); err != nil {
		return fmt.Errorf("can't render node: %w", err)
	}
	_, err := writer.Write([]byte(rendered.String()))
	if err != nil {
		return fmt.Errorf("can't write workspace: %w", err)
	}
	return nil
}

func generateStringIdentifier(s string) string {
	return dsl.DoubleQuotes + strings.ReplaceAll(s, "\"", "\\\"") + dsl.DoubleQuotes
}

func generateIdent(level int) string {
	return strings.Repeat(" ", level*4)
}

func generateVarName(s string) string {
	return strcase.ToLowerCamel(s)
}

func writeLine(renderer *strings.Builder, level int, values ...string) {
	renderer.WriteString(generateIdent(level))
	for _, value := range values {
		renderer.WriteString(value)
	}
	renderer.WriteString(dsl.NewLine)
}
