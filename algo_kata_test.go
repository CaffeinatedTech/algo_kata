package main

import (
  "testing"
)

var test_languages = []string{"Python", "Go"}
var test_algorithms = []Algorithm{{Name: "Bubble Sort", Type: "sort", Sorted: false}, {Name: "Shell Sort", Type: "sort", Sorted: false}}

// TestShellSort tests the shell sort algorithm
func TestShellSort(t *testing.T) {
  arr := []int{12, 34, 54, 2, 3}
  shellSort(arr)
  if arr[0] != 2 || arr[1] != 3 || arr[2] != 12 || arr[3] != 34 || arr[4] != 54 {
    t.Error("Shell sort failed")
  }
}

// TestCompareArrays tests the compareArrays function
func TestCompareArrays(t *testing.T) {
  a := []int{1, 2, 3}
  b := []int{1, 2, 3}
  if !compareArrays(a, b) {
    t.Error("compareArrays failed")
  }
}

// TestGetConfig tests the getConfig function
func TestGetConfig(t *testing.T) {
  c, l, _ := getConfig()
  if len(c) == 0 || len(l) == 0 {
    t.Error("getConfig failed")
  }
}
