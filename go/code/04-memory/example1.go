package main

import (
	"fmt"
	"runtime"
)

func openMemory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	beforeAlloc := m.Mallocs
	var a int
	fmt.Println(a)
	runtime.ReadMemStats(&m)
	afterAlloc := m.Mallocs
	fmt.Printf("内存块是否分配： %t\n 分配前：%d\n 分配后：%d\n", afterAlloc > beforeAlloc, beforeAlloc, afterAlloc)
}

// 下面这行是为了防止f函数的调用被内联。
//
//go:noinline
func f(i int) byte {
	var a [1 << 20]byte // 使栈增长
	return a[i]
}

func addressChange() {
	var x int
	println(&x)
	f(100)
	println(&x)
}

var p *int

func gc() {
	done := make(chan bool)
	// done通道将被使用在主协程和下面将要
	// 创建的新协程中，所以它将被开辟在堆上。

	go func() {
		x, y, z := 123, 456, 789
		_ = z  // z可以被安全地开辟在栈上。
		p = &x // 因为x和y都会将曾经被包级指针p所引用过，
		p = &y // 因此，它们都将开辟在堆上。

		// 到这里，x已经不再被任何其它值所引用。或者说承载
		// 它的内存块已经不再被使用。此内存块可以被回收了。

		p = nil
		// 到这里，y已经不再被任何其它值所引用。
		// 承载它的内存块可以被回收了。

		done <- true
	}()

	<-done
	// 到这里，done已经不再被任何其它值所引用。一个
	// 聪明的编译器将认为承载它的内存块可以被回收了。

	// ...
}
