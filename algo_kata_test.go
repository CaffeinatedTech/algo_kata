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
  c, l := getConfig()
  if len(c) == 0 || len(l) == 0 {
    t.Error("getConfig failed")
  }
}

// TestCheckPracticeNumber tests the checkPracticeNumber function
func TestCheckPracticeNumber(t *testing.T) {
  s := Session{Num: 100, Algorithms: test_algorithms, Languages: test_languages}
  // Test too large
  s = checkPracticeNumer(s)
  if s.Num != 4 {
    t.Error("checkPracticeNumber failed")
  }
  // Test Less than 1
  s.Num = 0
  s = checkPracticeNumer(s)
  if s.Num != 1 {
    t.Error("checkPracticeNumber failed")
  }
  // Test normal
  s.Num = 3
  s = checkPracticeNumer(s)
  if s.Num != 3 {
    t.Error("checkPracticeNumber failed")
  }
}
