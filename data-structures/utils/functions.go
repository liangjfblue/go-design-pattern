/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package utils

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
