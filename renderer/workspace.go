package renderer

import (
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderWorkspace(w *gostructurizr.WorkspaceNode, renderer *strings.Builder, level int) error {
	var line []string

	line = append(line, dsl.Workspace, dsl.Space)

	if w.Extend() != nil {
		line = append(line, dsl.Extends, dsl.Space, generateStringIdentifier(*w.Extend()), dsl.Space)
	} else {
		if w.Name() != nil {
			line = append(line, generateStringIdentifier(*w.Name()), dsl.Space)
		}
		if w.Desc() != nil {
			if w.Name() == nil {
				line = append(line, dsl.EmptyIdentifier, dsl.Space)
			}
			line = append(line, generateStringIdentifier(*w.Desc()), dsl.Space)
		}
	}
	line = append(line, dsl.OpenBracket)

	writeLine(renderer, level, line...)

	err := renderModel(w.Model(), renderer, level+1)
	if err != nil {
		return err
	}
	err = renderView(w.Views(), renderer, level+1)
	if err != nil {
		return err
	}

	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}
