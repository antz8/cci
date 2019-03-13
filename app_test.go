package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(42, 4)
    if total != 46 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}