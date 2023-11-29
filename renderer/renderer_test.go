package renderer

import (
	"bytes"
	"fmt"
	"github.com/platelk/gostructurizr"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewDSLRenderer(t *testing.T) {
	w := gostructurizr.Workspace()
	buf := bytes.Buffer{}
	r := NewDSLRenderer(&buf)

	require.NoError(t, r.Render(w))

	fmt.Println(buf.String())
}

func Test_renderWorkspace(t *testing.T) {
	w := gostructurizr.Workspace()
	buf := bytes.Buffer{}
	r := NewDSLRenderer(&buf)

	require.NoError(t, r.Render(w))

	fmt.Println(buf.String())
}

func Test_renderWorkspace_with_name(t *testing.T) {
	w := gostructurizr.Workspace().WithName("test")
	buf := bytes.Buffer{}
	r := NewDSLRenderer(&buf)

	require.NoError(t, r.Render(w))

	fmt.Println(buf.String())
}

func Test_renderWorkspace_without_name_and_desc(t *testing.T) {
	w := gostructurizr.Workspace().WithDesc("test dec")
	buf := bytes.Buffer{}
	r := NewDSLRenderer(&buf)

	require.NoError(t, r.Render(w))

	fmt.Println(buf.String())
}

func Test_renderWorkspace_with_extend(t *testing.T) {
	w := gostructurizr.Workspace().WithExtend("./source.dsl")
	buf := bytes.Buffer{}
	r := NewDSLRenderer(&buf)

	require.NoError(t, r.Render(w))

	fmt.Println(buf.String())
}
