package singleton

import "testing"

func TestSingleTon(t *testing.T) {
	instance1 := GetInstance()
	println("实例对象的信息和地址: ", instance1, &instance1)

	instance2 := GetInstance()
	println("实例对象的信息和地址: ", instance2, &instance2)

	if instance1 != instance2 {
		t.Error("实例对象不一样")
	}
}
