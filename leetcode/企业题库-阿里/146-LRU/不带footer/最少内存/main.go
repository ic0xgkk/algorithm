package main

import "fmt"

/*
LRU大结构体
 */
type LRUCache struct {
	cacheMap map[int]*Instance
	Header *Instance // 变量名称为Header，实际是来回变的
	Length int // 哈希表的大小
	Size int // 已存入的缓存的数量
}

/*
每个KV的实体（双向链表）
 */
type Instance struct {
	Last *Instance
	key int
	value int
	Next *Instance
}

/*
LRU初始化函数
 */
func Constructor(capacity int) LRUCache {
	c := LRUCache{}
	c.cacheMap = make(map[int]*Instance, capacity)
	c.Length = capacity
	return c
}

/*
Get元素方法
此操作会将元素（若有）移到链表顶端
 */
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

		// 放到前边
		this.transToHeader(cache)

		return cache.value
	}

	// 不在两边的
	// 链接前后两个节点
	cache.Last.Next = cache.Next
	cache.Next.Last = cache.Last

	// 移到前边
	this.transToHeader(cache)

	return cache.value
}

/*
Put函数
当覆盖写入时会隐式执行一次Get（前移使用）
当插入时会放在链表前端
 */
func (this *LRUCache) Put(key int, value int)  {
	cache, ok := this.cacheMap[key]
	if ok {
		cache.value = value
		_ = this.Get(key)
		return
	}

	// 当链表还是空的时候
	if this.Header == nil {
		this.Header = &Instance{
			Last:  nil,
			key:   key,
			value: value,
			Next:  nil,
		}
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
				key:   key,
				value: value,
			}
			this.cacheMap[key] = this.Header
			return
		}

		// 移除末尾的
		// 把指针移到尾部
		this.goFooter()

		needDelete := this.Header
		this.Header = this.Header.Last
		this.Header.Next = nil
		delete(this.cacheMap, needDelete.key)

		// 新增到前边的
		ins := &Instance{
			key:  key,
			value: value,
		}
		this.transToHeader(ins)
		this.cacheMap[key] = ins

		return
	} else
	// 当链表还没满但是已经有头的时候，增加到前边
	{
		ins := &Instance{
			key:   key,
			value: value,
		}
		this.transToHeader(ins)
		this.cacheMap[key] = ins
		this.Size++
		return
	}
}

func (this *LRUCache) goHeader()  {
	for {
		if this.Header.Last == nil {
			break
		}
		this.Header = this.Header.Last
	}
}

func (this *LRUCache) goFooter()  {
	for {
		if this.Header.Next == nil {
			break
		}
		this.Header = this.Header.Next
	}
}

func (this *LRUCache) transToHeader(needTrans *Instance) {
	this.goHeader()
	oldHeader := this.Header
	this.Header = needTrans
	this.Header.Last = nil
	this.Header.Next = oldHeader
	oldHeader.Last = this.Header
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