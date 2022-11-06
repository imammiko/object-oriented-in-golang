package reference

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFruit(t *testing.T) {

	apple := Fruit{
		Name:  "Apple",
		Color: "Red",
	}

	require.Equal(t, "Apple", apple.GetName())
	require.Equal(t, "Red", apple.GetColor())

	apple.SetName("Green Apple")
	apple.SetColor("Green")

	require.Equal(t, "Green Apple", apple.GetName())
	require.Equal(t, "Green", apple.GetColor())

}
