package greedy

import (
	"container/heap"
	"sort"
)

type CourseHeap [][]int

func (courseHeap CourseHeap) Len() int {
	return len(courseHeap)
}

func (courseHeap CourseHeap) Less(i int, j int) bool {
	return courseHeap[i][0] > courseHeap[j][0]
}

func (courseHeap CourseHeap) Swap(i int, j int) {
	courseHeap[i], courseHeap[j] = courseHeap[j], courseHeap[i]
}

func (courseHeap *CourseHeap) Push(x any) {
	*courseHeap = append(*courseHeap, x.([]int))
}

func (courseHeap *CourseHeap) Pop() any {
	ret := (*courseHeap)[len(*courseHeap)-1]
	*courseHeap = (*courseHeap)[:len(*courseHeap)-1]
	return ret
}

func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	h := make(CourseHeap, 0)
	heap.Init(&h)

	progress := 0

	for _, course := range courses {
		if course[0] > course[1] {
			continue
		}

		if len(h) == 0 || progress+course[0] <= course[1] {
			progress += course[0]
			heap.Push(&h, course)
			continue
		}

		if course[0] < h[0][0] {
			progress -= heap.Pop(&h).([]int)[0]
			progress += course[0]
			heap.Push(&h, course)
		}
	}

	return len(h)
}

type IntHeap []int

func (heap IntHeap) Len() int {
	return len(heap)
}

func (heap IntHeap) Less(i, j int) bool {
	return heap[i] < heap[j]
}

func (heap IntHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *IntHeap) Push(x any) {
	*heap = append(*heap, x.(int))
}

func (heap *IntHeap) Pop() any {
	elem := (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	return elem
}

func maxEvents(events [][]int) int {
	if len(events) == 0 {
		return 0
	}

	// 按会议开始时间排序
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})

	// 获取会议的minDay maxDay
	minDay, maxDay, ans, index := events[0][0], events[0][1], 0, 0
	for _, event := range events {
		maxDay = max(maxDay, event[1])
	}

	// 构建小根堆
	queue := make(IntHeap, 0)
	heap.Init(&queue)

	for day := range maxDay - minDay + 1 {
		day += minDay

		// 当天可以考虑的会议丢入小根堆
		for i := index; i < len(events); i++ {
			if day == events[i][0] {
				heap.Push(&queue, events[i][1])
				index++
			} else {
				break
			}
		}
		// 清理过期的会议
		for len(queue) != 0 && queue[0] < day {
			heap.Pop(&queue)
		}

		// 选择一天开会
		if len(queue) != 0 {
			ans++
			heap.Pop(&queue)
		}
	}

	return ans
}

type MaxIntHeap []int

func (heap MaxIntHeap) Len() int {
	return len(heap)
}

func (heap MaxIntHeap) Less(i, j int) bool {
	return heap[i] > heap[j]
}

func (heap MaxIntHeap) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
}

func (heap *MaxIntHeap) Push(x any) {
	*heap = append(*heap, x.(int))
}

func (heap *MaxIntHeap) Pop() any {
	ret := (*heap)[len(*heap)-1]
	*heap = (*heap)[:len(*heap)-1]
	return ret
}

// 502. IPO
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {

	// 定义proj
	type project struct {
		profit  int
		capital int
	}

	// 构建proj列表
	index, n := 0, len(profits)
	projects := make([]project, n)
	for i := range n {
		projects[i] = project{profits[i], capital[i]}
	}

	// 排序projects
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	queue := make(MaxIntHeap, 0)
	heap.Init(&queue)

	for i := 0; i < k; i++ {
		for index < n {
			if projects[index].capital <= w {
				heap.Push(&queue, projects[index].profit)
				index++
			} else {
				break
			}
		}

		if len(queue) == 0 {
			break
		}
		w += heap.Pop(&queue).(int)
	}
	return w
}
