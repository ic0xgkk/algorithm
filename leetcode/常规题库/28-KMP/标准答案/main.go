package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	hL := len(haystack)
	nL := len(needle)

	for i :=0 ; i < hL; i++ {
		start := i
		if nL > hL - i {
			return -1
		}
		for j := 0; j < nL; j++ {
			if needle[j] != haystack[start] {
				break
			}
			start += 1
		}
		if start - i == nL {
			return i
		}
	}
	return -1
}
