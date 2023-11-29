package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"io"
	"strings"
)

func renderModel(m *gostructurizr.ModelNode, writer io.Writer, level int) error {
	rendered := strings.Builder{}
	var line []string

	line = append(line, dsl.Model, dsl.Space, dsl.OpenBracket)
	writeLine(&rendered, level, line...)

	for _, p := range m.Persons() {
		if err := renderPerson(p, &rendered, level+1); err != nil {
			return fmt.Errorf("can't render person: %w", err)
		}
	}
	//rendered.WriteString(dsl.NewLine)
	for _, s := range m.SoftwareSystems() {
		if err := renderSoftwareSystem(s, &rendered, level+1); err != nil {
			return fmt.Errorf("can't render softwareSystem: %w", err)
		}
	}
	rendered.WriteString(dsl.NewLine)
	for _, u := range m.RelationShip() {
		if err := renderRelationShip(u, &rendered, level+1); err != nil {
			return fmt.Errorf("can't render relationship: %w", err)
		}
	}

	writeLine(&rendered, level, dsl.CloseBracket)

	_, err := writer.Write([]byte(rendered.String()))
	if err != nil {
		return fmt.Errorf("can't write model: %w", err)
	}

	return nil
}
