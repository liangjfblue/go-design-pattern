/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package utils

import "reflect"

//Min2 最近的2的倍数
func Min2(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}

//Compare 1 if a> b; -1 if a < b; 0 if a == b
func Compare(a, b interface{}) int {
	aT := reflect.TypeOf(a).String()
	bT := reflect.TypeOf(b).String()

	if aT != bT {
		panic("a,b must the same type")
	}

	switch a.(type) {
	case int:
		if a.(int) > b.(int) {
			return 1
		} else if a.(int) < b.(int) {
			return -1
		} else {
			return 0
		}
	case float32:
		if a.(float32) > b.(float32) {
			return 1
		} else if a.(float32) < b.(float32) {
			return -1
		} else {
			return 0
		}
	case float64:
		if a.(float64) > b.(float64) {
			return 1
		} else if a.(float64) < b.(float64) {
			return -1
		} else {
			return 0
		}
	case string:
		if a.(string) > b.(string) {
			return 1
		} else if a.(string) < b.(string) {
			return -1
		} else {
			return 0
		}
	default:
		panic("not support type")
	}
}
