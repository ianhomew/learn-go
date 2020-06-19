package main

import "fmt"

// 定義一個介面 裡面有兩個方法
type Usb interface {
	Start()
	Stop()
	// Using() int    後面可以定義返回類型
}

type Phone struct {
}

// 讓 Phone 實現 Usb 接口的方法
func (p Phone) Start() {
	fmt.Println("phone is working")
}

// 讓 Phone 實現 Usb 接口的方法
func (p Phone) Stop() {
	fmt.Println("phone is stopped")
}

type Camera struct {
}

// 讓 Phone 實現 Usb 接口的方法
func (c Camera) Start() {
	fmt.Println("camera is working")
}

// 讓 Phone 實現 Usb 接口的方法
func (c Camera) Stop() {
	fmt.Println("camera is stopped")
}

// Computer 給 Phone and Camera 插入的

type Computer struct {
}

// 接收一個 Usb 這個介面  ->  (usb Usb)
// 這裡其實就是多型了
// 哪一個結構體有 Working 這個方法呢?  ->  *Computer
func (c *Computer) Working(usb Usb) {
	// 通過 usb 介面調用兩個方法
	usb.Start()
	usb.Stop()
}

// 讓 Phone 實現 Usb 接口的方法
func (c Computer) Start() {
	fmt.Println("Computer is working")
}

// 讓 Phone 實現 Usb 接口的方法
func (c Computer) Stop() {
	fmt.Println("Computer is stopped")
}

func main() {

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// key point key point key point key point key point key point
	computer.Working(phone)
	computer.Working(camera)

	// 鬆耦合 高內聚

}
