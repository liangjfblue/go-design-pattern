/**
 *
 * @author liangjf
 * @create on 2020/9/11
 * @version 1.0
 */
package chain_list

import (
	"fmt"
	"testing"
)

func TestChainList(t *testing.T) {
	content := "我是 黄色, 很暴力, 不喜欢色情的人, 哈哈"

	y := new(YellowFilter)
	b := new(BaoLiFilter)
	s := new(SeQingFilter)

	y.Next(b)
	b.Next(s)

	fmt.Println("before chain list filter")
	fmt.Println(content)
	fmt.Println("after chain list filter")
	fmt.Println(y.Do(content))
}
