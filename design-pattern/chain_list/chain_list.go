/**
 *
 * @author liangjf
 * @create on 2020/9/11
 * @version 1.0
 */
package chain_list

import "strings"

/**
过滤器
*/

type Filter interface {
	Do(string) string
	Next(Filter)
}

//黄色词语过滤
type YellowFilter struct {
	nextHandle Filter
}

func (f *YellowFilter) Do(s string) string {
	s = strings.ReplaceAll(s, "黄色", "**")
	if f.nextHandle != nil {
		return f.nextHandle.Do(s)
	}
	return s
}

func (f *YellowFilter) Next(filter Filter) {
	if filter != nil {
		f.nextHandle = filter
	}
}

//暴力词语过滤
type BaoLiFilter struct {
	nextHandle Filter
}

func (f *BaoLiFilter) Do(s string) string {
	s = strings.ReplaceAll(s, "暴力", "*")
	if f.nextHandle != nil {
		return f.nextHandle.Do(s)
	}
	return s
}

func (f *BaoLiFilter) Next(filter Filter) {
	if filter != nil {
		f.nextHandle = filter
	}
}

//色情词语过滤
type SeQingFilter struct {
	nextHandle Filter
}

func (f *SeQingFilter) Do(s string) string {
	s = strings.ReplaceAll(s, "色情", "***")
	if f.nextHandle != nil {
		return f.nextHandle.Do(s)
	}
	return s
}

func (f *SeQingFilter) Next(filter Filter) {
	if filter != nil {
		f.nextHandle = filter
	}
}
