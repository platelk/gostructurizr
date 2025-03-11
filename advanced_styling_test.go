package gostructurizr

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/platelk/gostructurizr/tags"
)

func TestAdvancedElementStyling(t *testing.T) {
	// Create a workspace and model
	workspace := Workspace().WithName("Advanced Styling Test")
	_ = workspace.Model()

	// Create styles
	views := workspace.Views()
	styles := views.Configuration().Styles()

	// Test advanced element styling
	personStyle := styles.AddElementStyle(tags.Person)

	// Test border styling
	border := 3
	personStyle.WithBorder(border)
	assert.Equal(t, &border, personStyle.Border())

	borderStyle := Dashed
	personStyle.WithBorderStyle(borderStyle)
	assert.Equal(t, &borderStyle, personStyle.BorderStyle())

	// Test shadow
	shadow := true
	personStyle.WithShadow(shadow)
	assert.Equal(t, &shadow, personStyle.Shadow())

	// Test font styling
	fontFamily := Monospace
	personStyle.WithFontFamily(fontFamily)
	assert.Equal(t, &fontFamily, personStyle.FontFamily())

	fontStyle := "italic"
	personStyle.WithFontStyle(fontStyle)
	assert.Equal(t, &fontStyle, personStyle.FontStyle())

	// Test positioning
	zIndex := 10
	personStyle.WithZIndex(zIndex)
	assert.Equal(t, &zIndex, personStyle.ZIndex())

	rotation := 45
	personStyle.WithRotation(rotation)
	assert.Equal(t, &rotation, personStyle.Rotation())

	personStyle.WithPosition(100, 200)
	expectedPosition := [2]int{100, 200}
	assert.Equal(t, &expectedPosition, personStyle.Position())

	// Test multiple icons
	personStyle.AddIcon("icon1.png")
	personStyle.AddIcon("icon2.png")
	assert.Equal(t, []string{"icon1.png", "icon2.png"}, personStyle.MultipleIcons())
}

func TestAdvancedRelationshipStyling(t *testing.T) {
	// Create a workspace and model
	workspace := Workspace().WithName("Advanced Relationship Styling Test")
	_ = workspace.Model()

	// Create styles
	views := workspace.Views()
	styles := views.Configuration().Styles()

	// Test advanced relationship styling
	relationshipStyle := styles.AddAdvancedRelationshipStyle(tags.Synchronous)

	// Test line style
	lineStyle := DashedLine
	relationshipStyle.WithLineStyle(lineStyle)
	assert.Equal(t, &lineStyle, relationshipStyle.LineStyle())

	// Test convenience methods
	dashRelationship := styles.AddAdvancedRelationshipStyle(tags.Asynchronous)
	dashRelationship.WithDashed()
	assert.Equal(t, DashedLine, *dashRelationship.LineStyle())

	dotRelationship := styles.AddAdvancedRelationshipStyle(tags.RelationShip)
	dotRelationship.WithDotted()
	assert.Equal(t, DottedLine, *dotRelationship.LineStyle())

	// Test terminators
	startTerminator := Circle
	endTerminator := Arrow
	relationshipStyle.WithStartTerminator(startTerminator)
	relationshipStyle.WithEndTerminator(endTerminator)
	assert.Equal(t, &startTerminator, relationshipStyle.StartTerminator())
	assert.Equal(t, &endTerminator, relationshipStyle.EndTerminator())

	// Test arrow convenience method
	arrowRelationship := styles.AddAdvancedRelationshipStyle(tags.Element)
	arrowRelationship.WithArrow()
	assert.Equal(t, Arrow, *arrowRelationship.EndTerminator())

	// Test routing
	routing := Orthogonal
	relationshipStyle.WithRouting(routing)
	assert.Equal(t, &routing, relationshipStyle.Routing())

	// Test routing convenience methods
	orthogonalRouting := styles.AddAdvancedRelationshipStyle(tags.Container)
	orthogonalRouting.WithOrthogonalRouting()
	assert.Equal(t, Orthogonal, *orthogonalRouting.Routing())

	curvedRouting := styles.AddAdvancedRelationshipStyle(tags.Component)
	curvedRouting.WithCurvedRouting()
	assert.Equal(t, Curved, *curvedRouting.Routing())

	directRouting := styles.AddAdvancedRelationshipStyle(tags.SoftwareSystem)
	directRouting.WithDirectRouting()
	assert.Equal(t, Direct, *directRouting.Routing())

	// Test font styling
	fontSize := 12
	fontColor := "#FF0000"
	fontFamily := Serif
	fontStyle := "bold"

	relationshipStyle.WithFontSize(fontSize)
	relationshipStyle.WithFontColor(fontColor)
	relationshipStyle.WithFontFamily(fontFamily)
	relationshipStyle.WithFontStyle(fontStyle)

	assert.Equal(t, &fontSize, relationshipStyle.FontSize())
	assert.Equal(t, &fontColor, relationshipStyle.FontColor())
	assert.Equal(t, &fontFamily, relationshipStyle.FontFamily())
	assert.Equal(t, &fontStyle, relationshipStyle.FontStyle())

	// Test position
	position := 50
	relationshipStyle.WithPosition(position)
	assert.Equal(t, &position, relationshipStyle.Position())
}
