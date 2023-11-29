package gostructurizr

import (
	"github.com/platelk/gostructurizr/dsl"
)

type Identifier string

const (
	All Identifier = dsl.All
)

func (i *Identifier) Name() string {
	return string(*i)
}

type Fromer interface {
	WithOn(on Namer) *ExpressionViewNode
}

type ExpressionViewNode struct {
	on, from, to       Namer
	afferent, efferent bool
}

func Expression(on Namer) *ExpressionViewNode {
	return &ExpressionViewNode{on: on}
}

func On(n Namer) *ExpressionViewNode {
	return Expression(n)
}

func From(n Namer) Fromer {
	return Expression(nil).WithFrom(n)
}

func (e *ExpressionViewNode) WithFrom(from Namer) *ExpressionViewNode {
	if from != nil {
		e.afferent = true
	}
	e.from = from
	return e
}

func (e *ExpressionViewNode) From() Namer {
	return e.from
}

func (e *ExpressionViewNode) WithTo(to Namer) *ExpressionViewNode {
	if to != nil {
		e.efferent = true
	}
	e.to = to
	return e
}

func (e *ExpressionViewNode) To() Namer {
	return e.to
}

func (e *ExpressionViewNode) WithOn(on Namer) *ExpressionViewNode {
	e.on = on
	return e
}

func (e *ExpressionViewNode) On() Namer {
	return e.on
}

func (e *ExpressionViewNode) WithAfferent(b bool) *ExpressionViewNode {
	e.afferent = b
	return e
}

func (e *ExpressionViewNode) Afferent() bool {
	return e.afferent
}

func (e *ExpressionViewNode) WithEfferent(b bool) *ExpressionViewNode {
	e.efferent = b
	return e
}

func (e *ExpressionViewNode) Efferent() bool {
	return e.efferent
}
