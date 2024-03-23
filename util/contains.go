package util

func ContainsString(list []string, value string) bool {
	for _, v := range list {
		if value == v {
			return true
		}
	}
	return false
}
func ContainsInt(list []int, value int) bool {
	for _, v := range list {
		if value == v {
			return true
		}
	}
	return false
}
