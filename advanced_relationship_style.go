package gostructurizr

import "github.com/platelk/gostructurizr/tags"

// LineStyle represents the style of lines for relationships
type LineStyle string

const (
	SolidLine  LineStyle = "Solid"
	DashedLine LineStyle = "Dashed"
	DottedLine LineStyle = "Dotted"
)

// RouteStyle represents the routing style for relationships
type RouteStyle string

const (
	Direct     RouteStyle = "Direct"
	Orthogonal RouteStyle = "Orthogonal"
	Curved     RouteStyle = "Curved"
)

// TerminatorStyle represents the style of the line terminator
type TerminatorStyle string

const (
	None      TerminatorStyle = "None"
	Arrow     TerminatorStyle = "Arrow"
	Circle    TerminatorStyle = "Circle"
	Diamond   TerminatorStyle = "Diamond"
	Triangle  TerminatorStyle = "Triangle"
)

// AdvancedRelationshipStyleNode provides enhanced styling options for relationships
type AdvancedRelationshipStyleNode struct {
	tag     tags.Tag
	color   *string
	opacity *int
	width   *int
	
	// Advanced styling options
	lineStyle      *LineStyle
	fontSize       *int
	fontColor      *string
	fontFamily     *FontType
	fontStyle      *string
	routing        *RouteStyle
	position       *int
	startTerminator *TerminatorStyle
	endTerminator   *TerminatorStyle
}

// AdvancedRelationshipStyle creates a new AdvancedRelationshipStyleNode
func AdvancedRelationshipStyle(tag tags.Tag) *AdvancedRelationshipStyleNode {
	return &AdvancedRelationshipStyleNode{tag: tag}
}

// Tag returns the tag of the relationship style
func (r *AdvancedRelationshipStyleNode) Tag() tags.Tag {
	return r.tag
}

// WithColor sets the color of the relationship
func (r *AdvancedRelationshipStyleNode) WithColor(c string) *AdvancedRelationshipStyleNode {
	r.color = &c
	return r
}

// Color returns the color of the relationship
func (r *AdvancedRelationshipStyleNode) Color() *string {
	return r.color
}

// WithOpacity sets the opacity of the relationship
func (r *AdvancedRelationshipStyleNode) WithOpacity(o int) *AdvancedRelationshipStyleNode {
	r.opacity = &o
	return r
}

// Opacity returns the opacity of the relationship
func (r *AdvancedRelationshipStyleNode) Opacity() *int {
	return r.opacity
}

// WithWidth sets the width of the relationship
func (r *AdvancedRelationshipStyleNode) WithWidth(w int) *AdvancedRelationshipStyleNode {
	r.width = &w
	return r
}

// Width returns the width of the relationship
func (r *AdvancedRelationshipStyleNode) Width() *int {
	return r.width
}

// LineStyle returns the line style of the relationship
func (r *AdvancedRelationshipStyleNode) LineStyle() *LineStyle {
	return r.lineStyle
}

// WithLineStyle sets the line style of the relationship
func (r *AdvancedRelationshipStyleNode) WithLineStyle(style LineStyle) *AdvancedRelationshipStyleNode {
	r.lineStyle = &style
	return r
}

// FontSize returns the font size of the relationship label
func (r *AdvancedRelationshipStyleNode) FontSize() *int {
	return r.fontSize
}

// WithFontSize sets the font size of the relationship label
func (r *AdvancedRelationshipStyleNode) WithFontSize(size int) *AdvancedRelationshipStyleNode {
	r.fontSize = &size
	return r
}

// FontColor returns the font color of the relationship label
func (r *AdvancedRelationshipStyleNode) FontColor() *string {
	return r.fontColor
}

// WithFontColor sets the font color of the relationship label
func (r *AdvancedRelationshipStyleNode) WithFontColor(color string) *AdvancedRelationshipStyleNode {
	r.fontColor = &color
	return r
}

// FontFamily returns the font family of the relationship label
func (r *AdvancedRelationshipStyleNode) FontFamily() *FontType {
	return r.fontFamily
}

// WithFontFamily sets the font family of the relationship label
func (r *AdvancedRelationshipStyleNode) WithFontFamily(family FontType) *AdvancedRelationshipStyleNode {
	r.fontFamily = &family
	return r
}

// FontStyle returns the font style of the relationship label
func (r *AdvancedRelationshipStyleNode) FontStyle() *string {
	return r.fontStyle
}

// WithFontStyle sets the font style of the relationship label
func (r *AdvancedRelationshipStyleNode) WithFontStyle(style string) *AdvancedRelationshipStyleNode {
	r.fontStyle = &style
	return r
}

// Routing returns the routing style of the relationship
func (r *AdvancedRelationshipStyleNode) Routing() *RouteStyle {
	return r.routing
}

// WithRouting sets the routing style of the relationship
func (r *AdvancedRelationshipStyleNode) WithRouting(routing RouteStyle) *AdvancedRelationshipStyleNode {
	r.routing = &routing
	return r
}

// Position returns the position percentage of the relationship label
func (r *AdvancedRelationshipStyleNode) Position() *int {
	return r.position
}

// WithPosition sets the position percentage of the relationship label (0-100)
func (r *AdvancedRelationshipStyleNode) WithPosition(position int) *AdvancedRelationshipStyleNode {
	r.position = &position
	return r
}

// StartTerminator returns the start terminator style of the relationship
func (r *AdvancedRelationshipStyleNode) StartTerminator() *TerminatorStyle {
	return r.startTerminator
}

// WithStartTerminator sets the start terminator style of the relationship
func (r *AdvancedRelationshipStyleNode) WithStartTerminator(terminator TerminatorStyle) *AdvancedRelationshipStyleNode {
	r.startTerminator = &terminator
	return r
}

// EndTerminator returns the end terminator style of the relationship
func (r *AdvancedRelationshipStyleNode) EndTerminator() *TerminatorStyle {
	return r.endTerminator
}

// WithEndTerminator sets the end terminator style of the relationship
func (r *AdvancedRelationshipStyleNode) WithEndTerminator(terminator TerminatorStyle) *AdvancedRelationshipStyleNode {
	r.endTerminator = &terminator
	return r
}

// WithArrow sets the end terminator style to Arrow
func (r *AdvancedRelationshipStyleNode) WithArrow() *AdvancedRelationshipStyleNode {
	return r.WithEndTerminator(Arrow)
}

// WithDashed sets the line style to dashed
func (r *AdvancedRelationshipStyleNode) WithDashed() *AdvancedRelationshipStyleNode {
	return r.WithLineStyle(DashedLine)
}

// WithDotted sets the line style to dotted
func (r *AdvancedRelationshipStyleNode) WithDotted() *AdvancedRelationshipStyleNode {
	return r.WithLineStyle(DottedLine)
}

// WithOrthogonalRouting sets the routing style to orthogonal
func (r *AdvancedRelationshipStyleNode) WithOrthogonalRouting() *AdvancedRelationshipStyleNode {
	return r.WithRouting(Orthogonal)
}

// WithCurvedRouting sets the routing style to curved
func (r *AdvancedRelationshipStyleNode) WithCurvedRouting() *AdvancedRelationshipStyleNode {
	return r.WithRouting(Curved)
}

// WithDirectRouting sets the routing style to direct
func (r *AdvancedRelationshipStyleNode) WithDirectRouting() *AdvancedRelationshipStyleNode {
	return r.WithRouting(Direct)
}