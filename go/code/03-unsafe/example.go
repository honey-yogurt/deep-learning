package main

import "fmt"
import "unsafe"

func main() {
	//example1()
	//example2()
	example3()
}

func example1() {
	var x struct {
		a int64
		b bool
		c string
	}
	const M, N = unsafe.Sizeof(x.c), unsafe.Sizeof(x)
	fmt.Println(M, N) // 16 32

	fmt.Println(unsafe.Alignof(x.a)) // 8
	fmt.Println(unsafe.Alignof(x.b)) // 1
	fmt.Println(unsafe.Alignof(x.c)) // 8

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16
}

func example2() {
	type T struct {
		c string
	}
	type S struct {
		b bool
	}
	var x struct {
		a int64
		*S
		T
	}

	fmt.Println(unsafe.Offsetof(x.a)) // 0

	fmt.Println(unsafe.Offsetof(x.S)) // 8
	fmt.Println(unsafe.Offsetof(x.T)) // 16

	// 此行可以编译过，因为选择器x.c中的隐含字段T为非指针。
	fmt.Println(unsafe.Offsetof(x.c)) // 16

	// 此行编译不过，因为选择器x.b中的隐含字段S为指针。
	//fmt.Println(unsafe.Offsetof(x.b)) // error

	// 此行可以编译过，但是它将打印出字段b在x.S中的偏移量.
	fmt.Println(unsafe.Offsetof(x.S.b)) // 0
}

func example3() {
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a)
	// 计算数组每个元素的大小eleSize。
	eleSize := int(unsafe.Sizeof(a[0]))
	fmt.Println("eleSize:", eleSize)
	p9 := &a[9]
	fmt.Println("p9:", p9)
	// 将数组元素地址强转为unsafe.Pointer,可以直接操作整数内存。
	up9 := unsafe.Pointer(p9)
	fmt.Println("up9:", up9)
	// 通过unsafe.Add以byte为单位偏移指针,获取下标为3的元素地址,并打印值3。
	p3 := (*int)(unsafe.Add(up9, -6*eleSize))
	fmt.Println(*p3) // 3
	// 将下标为9的元素地址强转为slice首地址,从9开始长度5元素的slice,可以获取下标为9和11的元素值。这 circumvent 了slice越界检查。
	s := unsafe.Slice(p9, 5)[:3]
	fmt.Println(s)              // [9 0 11]
	fmt.Println(len(s), cap(s)) // 3 5

	// new 一个nil slice,并验证它等于nil。
	t := unsafe.Slice((*int)(nil), 0)
	fmt.Println(t == nil) // true

	// 尝试不合法的指针偏移,这可能会访问未知内存区域。
	// 下面是两个不正确的调用。因为它们
	// 的返回结果引用了未知的内存块。
	_ = unsafe.Add(up9, 7*eleSize)
	_ = unsafe.Slice(p9, 8)
}

func String2ByteSlice(str string) []byte {
	if str == "" {
		return nil
	}
	return unsafe.Slice(unsafe.StringData(str), len(str))
}

func ByteSlice2String(bs []byte) string {
	if len(bs) == 0 {
		return ""
	}
	return unsafe.String(unsafe.SliceData(bs), len(bs))
}
