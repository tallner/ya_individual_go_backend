package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenAgeIs19AndLocationIsSystemetThenNotAllowed(t *testing.T) {
	//arrange
	age := 25
	location := "Systemet"
	promille := float32(0.01)

	//act
	result := CanIBuyBeer(age, location, promille)

	//assert
	//assert.Equal(t, 123, i)
	assert.False(t, result)

}

func TestPromilleIsGreaterThanOnePointFiveThenNotAllowed(t *testing.T) {
	//arrange
	age := 50
	location := "Systemet"
	promille := float32(1.6)

	//act
	result := CanIBuyBeer(age, location, promille)

	//assert
	//assert.Equal(t, 123, i)
	assert.False(t, result)

}

func TestWhenAgeIs50AtSystemetThenAllowed(t *testing.T) {
	//arrange
	age := 50
	location := "Systemet"
	promille := float32(0.0)

	//act
	result := CanIBuyBeer(age, location, promille)

	//assert
	//assert.Equal(t, 123, i)
	assert.True(t, result)
}
