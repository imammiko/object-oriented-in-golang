package animal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnimal(t *testing.T) {

	dog := NewDog()
	cat := NewCat()
	mouse := NewMouse()

	require.Equal(t, "Woof", dog.MakeSound())
	require.Equal(t, "Meow", cat.MakeSound())
	require.Equal(t, "Squeak", mouse.MakeSound())


}