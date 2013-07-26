package bubbleSort

import (
    "testing"
)

func Test_bubbleSort (t *testing.T) {
    ints := []int{1,2,3,4,5}
//    b := BubbleSort{}
    if bubbleSort(ints) == ints {
        t.Log("bubbleSort:  --test1 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}

func Test_sort_reverse_order (t *testing.T) {
    ints := []int{5,4,3,2,1}
    if bubbleSort(ints) == ints {
        t.Log("bubbleSort:  --test2 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}

func Test_sort_negative_numbers (t *testing.T) {
    if bubbleSort(ints) == ints {
        t.Log("bubbleSort:  --test3 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}

func Test_sort_alpha_chars (t *testing.T) {
    if bubbleSort(ints) == ints {
        t.Log("bubbleSort:  --test4 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}

func Test_sort_invalid_chars (t *testing.T) {
    if bubbleSort(ints) == ints {
        t.Log("bubbleSort:  --test5 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}

func Test_sort_empty_list (t *testing.T) {
    if bubbleSort(ints) == ints {
        t.Log("bubbleSort:  --test6 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}

func Test_sort_single_element (t *testing.T) {
    if bubbleSort(ints) == ints {
        t.Log("bubbleSort:  --test7 passed")
    } else {
        t.Error("Error in bubbleSort: Test_sort_already_sorted")
    }
}
