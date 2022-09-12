// Package kit
// @Description: 非类型安全的[]string，并提供一个判断是否是StrSlice的方法
//
package kit

type (
	IntSlice   []int
	IntSliceIf interface {
		IsMe()
	}
)

func (s *IntSlice) IsMe() {}

func IsIntSlice(d any) bool {
	if _, ok := any(d).(IntSliceIf); ok {
		return true
	}
	return false
}
