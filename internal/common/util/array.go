package util

// Contains 檢查字串陣列是否包含指定元素
func Contains(arr []string, target string) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

// ContainsInt 檢查整數陣列是否包含指定元素
func ContainsInt(arr []int, target int) bool {
	for _, item := range arr {
		if item == target {
			return true
		}
	}
	return false
}

// Unique 字串陣列去重
func Unique(arr []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, item := range arr {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

// UniqueInt 整數陣列去重
func UniqueInt(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, item := range arr {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}
