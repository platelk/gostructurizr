package gostructurizr

type parallelFlow struct {
	start, end int
}

type DynamicViewNode struct {
	name          Namer
	key           *string
	desc          *string
	includeAll    bool
	relationShip  []*RelationShipNode
	parallelFlows []parallelFlow
}

func dynamicView() *DynamicViewNode {
	return &DynamicViewNode{}
}

func (d *DynamicViewNode) WithIdentifier(n Namer) *DynamicViewNode {
	d.name = n
	return d
}

func (d *DynamicViewNode) Identifier() Namer {
	return d.name
}

func (d *DynamicViewNode) IncludeAll() *DynamicViewNode {
	d.includeAll = true
	return d
}

func (d *DynamicViewNode) WithDescription(desc string) *DynamicViewNode {
	d.desc = &desc
	return d
}

func (d *DynamicViewNode) Description() *string {
	return d.desc
}

func (d *DynamicViewNode) WithKey(k string) *DynamicViewNode {
	d.key = &k
	return d
}

func (d *DynamicViewNode) Key() *string {
	return d.key
}

func (d *DynamicViewNode) Add(from, to Namer, others ...string) *DynamicViewNode {
	var desc string
	if len(others) >= 1 {
		desc = others[0]
	}
	d.relationShip = append(d.relationShip, Uses(from, to, desc))
	return d
}

func (d *DynamicViewNode) StartParallelSequence() *DynamicViewNode {
	var p parallelFlow
	d.parallelFlows = append(d.parallelFlows, p)
	d.parallelFlows[len(d.parallelFlows)-1].start = len(d.relationShip) - 1

	return d
}

func (d *DynamicViewNode) EndParallelSequence() *DynamicViewNode {
	d.parallelFlows[len(d.parallelFlows)-1].end = len(d.relationShip) - 1
	return d
}
