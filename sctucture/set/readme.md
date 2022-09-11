# Set

实现了2种不同的Set：

> 1. 基于泛型实现的Set
> 2. 基于map实现的Set

## 1.1. 基于泛型实现的Set
每次只能添加一种类型，使用方法看container_set_test.go

## 1.2. 基于map实现的Set,但同时支持不同类型，且类型安全,支持的类型是有限制的！
可以添加不同类型的元素，使用方法看container_set_test.go



# 注意：
这2种类型的set添加的元素都必须是实现了compare接口的类型，否则会panic。

> map key 必须是可比较的类型，语言规范中定义了可比较的类型：boolean, numeric, string, pointer, channel, interface, 以及仅包含这些类型的struct和array 。
> 不能作为map key的类型有：slice，map, function。
> 
> As mentioned earlier, map keys may be of any type that is comparable.
> The language spec defines this precisely, but in short, comparable types are
> boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only those types.
> Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, and may not be used as map keys.


# Todo
-[ ] 无限制类型的Set