package xutil

// ContainStr check if string target in string array
func ContainStr(source []string, target string) bool {
	for _, elem := range source {
		if elem == target {
			return true
		}
	}
	return false
}

// ContainInt check if int target in int array
func ContainInt(source []int, target int) bool {
	for _, elem := range source {
		if elem == target {
			return true
		}
	}
	return false
}
