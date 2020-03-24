package main

import "fmt"

func strStr(haystack string, needle string) int {
	if haystack == needle { return 0 }
	if haystack == "" { return -1 }
	if needle == "" { return 0 }
	kmp := New(needle)
	return kmp.Search(haystack)
}

type KMP struct {
	prefix []int
	pattern string
}

func New(p string) *KMP {
	kmp := KMP{
		prefix:  nil,
		pattern: p,
	}
	kmp.prefix = append(kmp.prefix, -1)
	Len := len(kmp.pattern)
	for i := 1; i < Len; i++ {
		kmp.prefix = append(kmp.prefix, getMaxRepeat(kmp.pattern[0:i], i))
	}
	return &kmp
}

func getMaxRepeat(text string, Len int) int {
	// 计算前缀表
	// 因为要算最小的，不用分奇偶
	for i := 1; i <= Len - 1; i++ {
		if text[0 : Len - i] == text[i : Len] {
			return Len - i
		}
	}
	// 未匹配到
	return 0
}

func (k *KMP) Search(text string) int {
	patLen := len(k.pattern)
	textLen := len(text)

	if patLen == 1 {
		for i := 0; i < textLen; i++ {
			if k.pattern[0] == text[i] {
				return i
			}
		}
		return -1
	}

	// 索引号
	j := 0
	for i := 0; i < textLen; {
		// 如果匹配则继续向下
		if text[i] == k.pattern[j] {
			// 判断一下已经匹配的模式串是不是够了
			if j == patLen - 1 {
				return i - j
			}
			//  长度不够继续向后匹配
			i++; j++
		} else
		// 不匹配的话尝试移动模式串
		{
			// 得到新的索引号
			j = k.prefix[j]
			// 如果j为-1-两数之和，说明要要向下一位文本串进行匹配了
			if j == -1 {
				// 先归位j
				j = 0
				// 对齐到文本串下一位
				i++
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}