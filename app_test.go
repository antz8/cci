package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(2, 4)
    if total != 6 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}