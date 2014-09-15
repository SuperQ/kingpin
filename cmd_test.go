package kingpin

import (
	"github.com/stretchrcom/testify/assert"

	"testing"
)

func TestNestedCommands(t *testing.T) {
	app := New("app", "")
	sub1 := app.Command("sub1", "")
	sub1.Flag("sub1", "")
	subsub1 := sub1.Command("sub1sub1", "")
	subsub1.Command("sub1sub1end", "")

	sub2 := app.Command("sub2", "")
	sub2.Flag("sub2", "")
	sub2.Command("sub2sub1", "")

	tokens := Tokenize([]string{"sub1", "sub1sub1", "sub1sub1end"})
	tokens, selected, err := app.parse(tokens)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tokens))
	assert.Equal(t, "sub1 sub1sub1 sub1sub1end", selected)
}

func TestNestedCommandsWithArgs(t *testing.T) {
	app := New("app", "")
	cmd := app.Command("a", "").Command("b", "")
	a := cmd.Arg("a", "").String()
	b := cmd.Arg("b", "").String()
	tokens := Tokenize([]string{"a", "b", "c", "d"})
	tokens, selected, err := app.parse(tokens)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(tokens))
	assert.Equal(t, "a b", selected)
	assert.Equal(t, "c", *a)
	assert.Equal(t, "d", *b)
}