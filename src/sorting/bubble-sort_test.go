package main

import (
    "testing"
)

func Test_sort_already_sorted (t *testing.T) {
    ints := []int{1,2,3,4,5}
    b := BubbleSort
    if b.sort(ints) == ints {
        t.Log("bubbleSort:  --test1 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}

func Test_sort_reverse_order (t *testing.T) {
}

func Test_sort_negative_numbers (t *testing.T) {
}

func Test_sort_alpha_chars (t *testing.T) {
}

func Test_sort_invalid_chars (t *testing.T) {
}

func Test_sort_empty_list (t *testing.T) {
}

func Test_sort_single_element (t *testing.T) {
}
