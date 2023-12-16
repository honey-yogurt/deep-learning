package main

type MyInt int64
type Ta *int64
type Tb *MyInt

func transform1() {
	// s 是 *int64 类型，为无名类型， 底层类型是自己，即 *int64，其基类型 int64， 基类型底层类型是 int64
	var s *int64 = new(int64)
	var s1 MyInt = 1
	// s2 是 *MyInt 类型，为无名类型， 底层类型是自己，即 *MyInt。其基类型 MyInt， 基类型底层类型是 int64
	var s2 = &s1
	// 下一行 无法编译通过，Cannot use 's' (type *int64) as the type *MyInt
	// 虽然二者都是无名类型，但是二者的底层类型不同，不满足第一条，肯定不能隐式转换的。
	// 从第一条来看，大前提不满足，是不能转换（显式和隐式），但是第一条和第二条是或的关系，所以如果满足第二条，那还是可以显示转换的
	s2 = s

	// 满足第二条，可以显式转换，二者基类型底层类型相同
	s2 = (*MyInt)(s)
}

func transform2() {
	var s *int64 = new(int64)
	var s1 Ta
	// s1 底层是 *int64 , 可以隐式转换
	s1 = s
}

func transform3() {
	var s1 MyInt = 1
	// s2 底层类型是 *MyInt， *MyInt 是无名类型
	var s2 = &s1
	var s3 Tb
	s3 = s2
}

func transform4() {
	var s1 MyInt = 1
	// s2 底层类型是 *MyInt， *MyInt 是无名类型
	var s2 = &s1

	var s3 *int64 = (*int64)(s2)
}
