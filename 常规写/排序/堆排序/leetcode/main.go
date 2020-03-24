/*
维护大顶堆实现排序
 */
package main

import "fmt"

type TopHeapSorting struct {
	topHeap []int
	heapLen int
}

func (t *TopHeapSorting) AddElements(elements []int)  {
	t.topHeap = elements
	t.heapLen = len(elements)
}

func (t *TopHeapSorting) Sorting() {
	if t.heapLen <= 1 {
		return
	}
	for end := t.heapLen - 1; end > 1; end-- {
		t.rebuildTopHeap(end)
		t.topHeap[0], t.topHeap[end] = t.topHeap[end], t.topHeap[0]
	}
	if t.topHeap[0] > t.topHeap[1] {
		t.topHeap[0], t.topHeap[1] = t.topHeap[1], t.topHeap[0]
	}
}

/*
左奇右偶
 */
func (t *TopHeapSorting) rebuildTopHeap(endIndex int)  {
	var right *int
	// 奇数 左子树
	if endIndex & 1 != 0 {
		left := &t.topHeap[endIndex]
		root := &t.topHeap[(endIndex - 2) / 2]
		right = nil
		sortTree(root, left, right)
		endIndex--
	}
	for ; endIndex > 1; endIndex -= 2 {
		right = &t.topHeap[endIndex]
		left := &t.topHeap[endIndex - 1]
		root := &t.topHeap[(endIndex - 2) / 2]
		sortTree(root, left, right)
	}
}

/*
对树进行排序
返回操作数，当操作数为0时表示不需要排序
 */
func sortTree(root *int, left *int, right *int) {
	if right == nil {
		if *root < *left {
			*root, *left = *left, *root
		}
	} else {
		if *left > *right {
			if *root < *left {
				*root, *left = *left, *root
			}
		} else {
			if *root < *right {
				*root, *right = *right, *root
			}
		}
	}
}

func main() {
	var t TopHeapSorting

	var eles []int
	eles = append(eles, -100)

	t.AddElements(eles)

	t.Sorting()
	fmt.Println(t.topHeap)
}

