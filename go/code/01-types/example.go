package main

// name 由 type 定义的定义类型，显然是一个 具名类型
type name string

func (n name) f() {

}

// 该别名的 基类型 name 显然是 具名类型，所以该别名类型也是具名类型
type myName = name

// m 方法名如果是 f ，编译器还会报错 Method redeclared 'myName.f'
func (n myName) m() {

}

// 别名源类型为 无名类型，该别名类型为 无名类型
type inventory = map[string]int

// 无法通过编译
// Invalid receiver type 'map[string]int' ('map[string]int' is an unnamed type
//func (i inventory) f()  {
//
//}

// 具名类型，但是是 组合类型
type inventory2 map[string]int

func (i inventory2) f() {

}
