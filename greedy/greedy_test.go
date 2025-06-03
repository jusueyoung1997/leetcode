package greedy

import "testing"

func TestScheduleCourse(t *testing.T) {
	ret := scheduleCourse([][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}})
	if ret != 3 {
		t.Errorf("expect 3 but get %d", ret)
	}
}

func TestMaxEvents(t *testing.T) {
	maxEvents([][]int{{1, 2}, {1, 2}, {3, 3}, {1, 5}, {1, 5}})

	// maxEvents([][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}})
}

func TestFindMaximizedCapital(t *testing.T) {
	// if findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}) != 6 {
	// 	t.Errorf("expect 6 but get %d", findMaximizedCapital(3, 0, []int{1, 2, 3}, []int{0, 1, 2}))
	// }

	ret := findMaximizedCapital(5, 2, []int{1, 2, 3, 7, 2, 19, 5, 3}, []int{0, 1, 2, 2, 6, 8, 6, 7})
	if ret != 39 {
		t.Errorf("expect 39 but get %d", ret)
	}
}
