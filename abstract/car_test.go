package abstract

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCar(t *testing.T) {
	
	var Audi ICar =NewAudi( "Audi")
	var Volvo ICar =NewVolvo( "Volvo")
	var Citron ICar =NewCitron( "Citron")

	require.Equal(t, "Choose German quality! I'm an Audi", Audi.Intro())
	require.Equal(t, "Proud to be Swedish! I'm a Volvo", Volvo.Intro())
	require.Equal(t, "French extravagance! I'm a Citron", Citron.Intro())
}