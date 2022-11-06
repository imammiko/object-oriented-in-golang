package agredrive

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	v1000 := NewEngine("v1000")

	auto := Transmission{}

	auto.SetType("automatic")

	car := NewCar("BMW")

	car.AddEngine(v1000)
	car.AddTransmision(auto)
	car.Engine.SetHorsePower(1000)

	require.Equal(t, "automatic", car.Transmission.GetType())
	require.Equal(t, 1000, car.Engine.GetHorsePower())

}