package main

import "testing"

func TestSum(t *testing.T) {
    total := Sum(40, 4)
    if total != 33 {
       t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
    }
}