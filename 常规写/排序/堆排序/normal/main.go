/*
维护大顶堆实现排序
*/
package main

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

type TopHeapSorting struct {
	topHeap []uint64
	heapLen uint64
}

/*
添加元素函数
接口收，可以容纳各种数据类型（整数的）
下边有做类型断言和转换
*/
func (t *TopHeapSorting) AddElements(elements ...interface{})  {
	for _, e := range elements {
		if u64Ele, ok := e.(uint64); ok {
			t.add(u64Ele)
		} else if iEle, ok := e.(int); ok {
			t.add(uint64(iEle))
		} else {
			panic("Type error")
		}
	}
}

/*
打散添加并统计
之所以逐个统计是考虑到数据量如果太大，len返回的int可能出现溢出，因此独立实现了长度统计
*/
func (t *TopHeapSorting) add(element uint64)  {
	t.topHeap = append(t.topHeap, element)
	t.heapLen++
}

/*
排序主函数
*/
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
重建大顶堆
左子树为奇数索引，右子树为偶数索引
判断一下是不是只有左子树，如果是就只排左子树
*/
func (t *TopHeapSorting) rebuildTopHeap(endIndex uint64)  {
	// 奇数 左子树
	if endIndex & 1 != 0 {
		sortTree(&t.topHeap[(endIndex - 2) / 2],
			&t.topHeap[endIndex],
			nil)
		endIndex--
	}
	for ; endIndex > 1; endIndex -= 2 {
		sortTree(&t.topHeap[(endIndex - 2) / 2],
			&t.topHeap[endIndex - 1],
			&t.topHeap[endIndex])
	}
}

/*
对当前树进行排序
*/
func sortTree(root *uint64, left *uint64, right *uint64) {
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
	go func() {
		fmt.Println(http.ListenAndServe("localhost:7000", nil))
	}()
	var t TopHeapSorting
	for i := 0; i < 500000; i++ {
		r := rand.Uint64()
		t.AddElements(r)
	}
	t.AddElements(100, 450, 1, 1000)
	fmt.Println("添加完成")
	t.Sorting()
}

