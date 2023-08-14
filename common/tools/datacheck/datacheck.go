package datacheck

import "regexp"

// 检查UID是否合规
func UIDCompliance(keyword string) bool {
	re := regexp.MustCompile(`^\d{5,6}$`)
	return re.MatchString(keyword)
}
