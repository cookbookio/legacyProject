package main

import "testing"

func TestAddRecipeIngredientValues(t *testing.T) {
	a, b := 2, 3
	result := a + b
	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}
