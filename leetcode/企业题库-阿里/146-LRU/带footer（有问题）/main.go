package main

import "fmt"

type LRUCache struct {
	cacheMap map[int]*Instance
	Header *Instance
	Current *Instance // 相当于Footer
	Length int
	Size int
}

type Instance struct {
	Last *Instance
	key int
	value int
	Next *Instance
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{}
	c.cacheMap = make(map[int]*Instance, capacity)
	c.Length = capacity
	return c
}


func (this *LRUCache) Get(key int) int {
	// 先从哈希表中找到位置
	cache, ok := this.cacheMap[key]
	if ok == false {
		return -1
	}

	// 在链表头部，不需要操作
	if cache.Last == nil {
		return cache.value
	}

	// 在链表尾部，需要放到开始
	if cache.Next == nil {
		// 清理倒数第二个Next
		cache.Last.Next = nil

		// Current/Footer 归位
		this.Current = cache.Last

		// 放到前边，要备份之前的Header
		this.Header.Last = cache
		oldHeader := this.Header
		this.Header = cache
		this.Header.Last = nil
		this.Header.Next = oldHeader

		return this.Header.value
	}

	// 不在两边的
	// 链接前后两个节点
	cache.Last.Next = cache.Next
	cache.Next.Last = cache.Last

	// 移到前边
	oldHeader := this.Header
	this.Header.Last = cache
	this.Header = cache
	this.Header.Last = oldHeader.Last
	this.Header.Next = oldHeader

	return cache.value
}


func (this *LRUCache) Put(key int, value int)  {
	cache, ok := this.cacheMap[key]
	if ok {
		cache.value = value
		_ = this.Get(key)
		return
	}

	// 当链表还是空的时候
	if this.Header == nil {
		this.Current = &Instance{
			Last:  nil,
			key:   key,
			value: value,
			Next:  nil,
		}
		this.Header = this.Current
		this.cacheMap[key] = this.Header
		this.Size++
		return
	}

	// 当链表已经满了的时候
	if this.Size == this.Length {
		// 如果只有一个元素
		if this.Length == 1 {
			delete(this.cacheMap, this.Header.key)
			this.Header = &Instance{
				Last:  nil,
				key:   key,
				value: value,
				Next:  nil,
			}
			this.cacheMap[key] = this.Header
			return
		}

		// 移除末尾的
		needDelete := this.Current
		this.Current = this.Current.Last
		this.Current.Next = nil
		delete(this.cacheMap, needDelete.key)

		// 新增到前边的
		ins := &Instance{
			Last:  nil,
			key:  key,
			value: value,
			Next:  this.Header.Next,
		}
		oldHeader := this.Header
		this.Header = ins
		oldHeader.Last = ins
		this.cacheMap[key] = this.Header

		return
	} else
	// 当链表还没满但是已经有头的时候，增加到前边
	{
		ins := Instance{
			Last:  nil,
			key:   key,
			value: value,
			Next:  this.Header,
		}
		this.Header = &ins
		this.Header.Next.Last = &ins
		this.cacheMap[key] = this.Header
		this.Size++
	}
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	var g int
	lru := Constructor(2)
	g = lru.Get(2)
	lru.Put(2, 6)
	g = lru.Get(1)
	lru.Put(1, 5)
	lru.Put(1, 2)
	g = lru.Get(1)
	g = lru.Get(2)
	fmt.Println(g)
}