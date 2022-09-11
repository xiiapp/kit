// Package kit
// @Description: 非类型安全的map[string]interface{}，并提供一个判断是否是StrAnyMap的方法
//
package kit

type (
	StrAnyMap   map[string]any
	StrAnyMapIf interface {
		IsMe()
	}
)

func (s *StrAnyMap) IsMe() {}

func IsStrAnyMap(d any) bool {
	if _, ok := any(d).(StrAnyMapIf); ok {
		return true
	}
	return false
}
