package constructor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFruit(t *testing.T) {
	
	apple := NewFruit("Apple", "Red")

	require.Equal(t, "Apple", apple.GetName())
	require.Equal(t, "Red", apple.GetColor())

}