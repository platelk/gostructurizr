package gostructurizr

import (
	"github.com/platelk/gostructurizr/shapes"
	"github.com/platelk/gostructurizr/tags"
)

type ElementStyleNode struct {
	tag           tags.Tag
	height, width *int
	background    *string
	stroke        *string
	color         *string
	fontSize      *int
	shape         *shapes.Shape
	icon          *string
	opacity       *int
	metadata      *bool
	description   *bool
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
	return e.metadata
}

func (e *ElementStyleNode) WithDescription(m bool) *ElementStyleNode {
	e.description = &m
	return e
}
