package inheritance

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFruit(t *testing.T) {

	strawberry := NewStrawberry("Strawberry", "Red")

	require.Equal(t, "Strawberry", strawberry.Fruit.Name)
	require.Equal(t, "Red", strawberry.Fruit.Color)
	require.Equal(t, "I am a strawberry", strawberry.Message())

}