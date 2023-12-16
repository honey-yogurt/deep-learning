package main

// Map 是一个具名类型，但是 map[string]string 是一个无名类型
type Map map[string]string

// SetVale 可以编译通过
func (m Map) SetVale(key, value string) {
	m[key] = value
}

// 编译失败
// Invalid receiver type 'map[string]string' ('map[string]string' is an unnamed type)
// 无效的接收器类型'map[string]string' ('map[string]string'是未命名的类型)
//func (m map[string]string)  SetVale(key,value string) {
//	m[key] = value
//}
