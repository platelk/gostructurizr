package renderer

import (
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/platelk/gostructurizr/dsl"
	"strings"
)

func renderSoftwareSystem(s *gostructurizr.SoftwareSystemNode, renderer *strings.Builder, level int) error {
	var line []string
	line = append(line, generateVarName(s.Name()), dsl.Space, dsl.Equal, dsl.Space, dsl.SoftwareSystem, dsl.Space, generateStringIdentifier(s.Name()))
	if s.Description() != nil {
		line = append(line, dsl.Space, generateStringIdentifier(*s.Description()))
	}
	containers := s.Containers()
	if (s.Tags() == nil || len(s.Tags().Values()) == 0) && (containers == nil || len(containers) == 0) {
		writeLine(renderer, level, line...)
		return nil
	}
	line = append(line, dsl.Space, dsl.OpenBracket)
	writeLine(renderer, level, line...)
	if s.Tags() != nil && len(s.Tags().Values()) > 0 {
		if err := renderTags(s.Tags(), renderer, level+1); err != nil {
			return fmt.Errorf("can't render tag of system: %w", err)
		}
	}
	for _, container := range containers {
		if err := renderContainer(container, renderer, level+1); err != nil {
			return fmt.Errorf("can't render container: %w", err)
		}
	}
	writeLine(renderer, level, dsl.CloseBracket)
	return nil
}
