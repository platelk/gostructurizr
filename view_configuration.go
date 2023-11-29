package gostructurizr

type ViewConfiguration struct {
	styles *StylesNode
}

func NewViewConfiguration() *ViewConfiguration {
	return &ViewConfiguration{styles: styles()}
}

func (c *ViewConfiguration) Styles() *StylesNode {
	return c.styles
}
