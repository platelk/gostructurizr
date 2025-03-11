package gostructurizr

import (
	"github.com/platelk/gostructurizr/shapes"
	"github.com/platelk/gostructurizr/tags"
)

// BorderStyle represents the style of borders for elements
type BorderStyle string

const (
	Solid  BorderStyle = "Solid"
	Dashed BorderStyle = "Dashed"
	Dotted BorderStyle = "Dotted"
)

// FontType represents the font family for elements
type FontType string

const (
	DefaultFont FontType = "Default"
	Monospace   FontType = "Monospace"
	Serif       FontType = "Serif"
	SansSerif   FontType = "SansSerif"
)

type ElementStyleNode struct {
	tag            tags.Tag
	height, width  *int
	background     *string
	stroke         *string
	color          *string
	fontSize       *int
	shape          *shapes.Shape
	icon           *string
	opacity        *int
	metadata       *bool
	description    *bool
	
	// Advanced styling options
	strokeWidth    *int
	borderStyle    *BorderStyle
	border         *int
	shadow         *bool
	fontFamily     *FontType
	fontStyle      *string
	multipleIcons  []string
	zIndex         *int
	rotation       *int
	position       *[2]int // [x, y] coordinates
}

func elementStyle(tag tags.Tag) *ElementStyleNode {
	return &ElementStyleNode{tag: tag}
}

func (e *ElementStyleNode) Tag() tags.Tag {
	return e.tag
}

func (e *ElementStyleNode) Height() *int {
	return e.height
}

func (e *ElementStyleNode) WithHeight(h int) *ElementStyleNode {
	e.height = &h
	return e
}

func (e *ElementStyleNode) Width() *int {
	return e.width
}

func (e *ElementStyleNode) WithWidth(w int) *ElementStyleNode {
	e.width = &w
	return e
}

func (e *ElementStyleNode) Background() *string {
	return e.background
}

func (e *ElementStyleNode) WithBackground(b string) *ElementStyleNode {
	e.background = &b
	return e
}

func (e *ElementStyleNode) Stroke() *string {
	return e.stroke
}

func (e *ElementStyleNode) WithStroke(s string) *ElementStyleNode {
	e.stroke = &s
	return e
}

func (e *ElementStyleNode) Color() *string {
	return e.color
}

func (e *ElementStyleNode) WithColor(c string) *ElementStyleNode {
	e.color = &c
	return e
}

func (e *ElementStyleNode) FontSize() *int {
	return e.fontSize
}

func (e *ElementStyleNode) WithFontSize(s int) *ElementStyleNode {
	e.fontSize = &s
	return e
}

func (e *ElementStyleNode) Shape() *shapes.Shape {
	return e.shape
}

func (e *ElementStyleNode) WithShape(s shapes.Shape) *ElementStyleNode {
	e.shape = &s
	return e
}

func (e *ElementStyleNode) Icon() *string {
	return e.icon
}

func (e *ElementStyleNode) WithIcon(i string) *ElementStyleNode {
	e.icon = &i
	return e
}

func (e *ElementStyleNode) Opacity() *int {
	return e.opacity
}

func (e *ElementStyleNode) WithOpacity(o int) *ElementStyleNode {
	e.opacity = &o
	return e
}

func (e *ElementStyleNode) Metadata() *bool {
	return e.metadata
}

func (e *ElementStyleNode) WithMetadata(m bool) *ElementStyleNode {
	e.metadata = &m
	return e
}

func (e *ElementStyleNode) Description() *bool {
	return e.description
}

func (e *ElementStyleNode) WithDescription(m bool) *ElementStyleNode {
	e.description = &m
	return e
}

// StrokeWidth returns the stroke width of the element
func (e *ElementStyleNode) StrokeWidth() *int {
	return e.strokeWidth
}

// WithStrokeWidth sets the stroke width of the element
func (e *ElementStyleNode) WithStrokeWidth(width int) *ElementStyleNode {
	e.strokeWidth = &width
	return e
}

// BorderStyle returns the border style of the element
func (e *ElementStyleNode) BorderStyle() *BorderStyle {
	return e.borderStyle
}

// WithBorderStyle sets the border style of the element
func (e *ElementStyleNode) WithBorderStyle(style BorderStyle) *ElementStyleNode {
	e.borderStyle = &style
	return e
}

// Border returns the border width of the element
func (e *ElementStyleNode) Border() *int {
	return e.border
}

// WithBorder sets the border width of the element
func (e *ElementStyleNode) WithBorder(width int) *ElementStyleNode {
	e.border = &width
	return e
}

// Shadow returns whether the element has a shadow
func (e *ElementStyleNode) Shadow() *bool {
	return e.shadow
}

// WithShadow sets whether the element has a shadow
func (e *ElementStyleNode) WithShadow(shadow bool) *ElementStyleNode {
	e.shadow = &shadow
	return e
}

// FontFamily returns the font family of the element
func (e *ElementStyleNode) FontFamily() *FontType {
	return e.fontFamily
}

// WithFontFamily sets the font family of the element
func (e *ElementStyleNode) WithFontFamily(fontFamily FontType) *ElementStyleNode {
	e.fontFamily = &fontFamily
	return e
}

// FontStyle returns the font style of the element
func (e *ElementStyleNode) FontStyle() *string {
	return e.fontStyle
}

// WithFontStyle sets the font style of the element (normal, italic, bold, etc.)
func (e *ElementStyleNode) WithFontStyle(style string) *ElementStyleNode {
	e.fontStyle = &style
	return e
}

// MultipleIcons returns the multiple icons of the element
func (e *ElementStyleNode) MultipleIcons() []string {
	return e.multipleIcons
}

// AddIcon adds an additional icon to the element
func (e *ElementStyleNode) AddIcon(icon string) *ElementStyleNode {
	e.multipleIcons = append(e.multipleIcons, icon)
	return e
}

// ZIndex returns the z-index of the element
func (e *ElementStyleNode) ZIndex() *int {
	return e.zIndex
}

// WithZIndex sets the z-index of the element
func (e *ElementStyleNode) WithZIndex(zIndex int) *ElementStyleNode {
	e.zIndex = &zIndex
	return e
}

// Rotation returns the rotation of the element in degrees
func (e *ElementStyleNode) Rotation() *int {
	return e.rotation
}

// WithRotation sets the rotation of the element in degrees
func (e *ElementStyleNode) WithRotation(rotation int) *ElementStyleNode {
	e.rotation = &rotation
	return e
}

// Position returns the position of the element as [x, y] coordinates
func (e *ElementStyleNode) Position() *[2]int {
	return e.position
}

// WithPosition sets the position of the element as [x, y] coordinates
func (e *ElementStyleNode) WithPosition(x, y int) *ElementStyleNode {
	e.position = &[2]int{x, y}
	return e
}
