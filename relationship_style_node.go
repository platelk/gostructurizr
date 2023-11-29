package gostructurizr

import "github.com/platelk/gostructurizr/routing"

type RelationShipStyleNode struct {
	routing *routing.Routing
	dashed  bool
}

func (n *RelationShipStyleNode) WithRouting(r routing.Routing) {
	n.routing = &r
}

func (n *RelationShipStyleNode) Routing() *routing.Routing {
	return n.routing
}

func (n *RelationShipStyleNode) WithDash(b bool) {
	n.dashed = b
}

func (n *RelationShipStyleNode) Dash() bool {
	return n.dashed
}
