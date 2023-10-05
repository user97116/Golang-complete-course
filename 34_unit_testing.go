package main

import (
    "testing"
)

func add(a, b int) int {
    return a + b
}

func TestAdd(t *testing.T) {
    result := add(3, 5)
    if result != 8 {
        t.Errorf("Expected 8, got %d", result)
    }
}
