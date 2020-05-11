/*
@Time : 2020/5/11 22:29
@Author : liangjiefan
*/
package factory_method

import "fmt"

type ToyInfo struct {
	Color string
	High  int
	Wide  int
}

type ToyA struct {
	ToyInfo
}

func NewToyA(info ToyInfo) *ToyA {
	return &ToyA{
		ToyInfo: info,
	}
}

func (a *ToyA) Shape() {
	fmt.Printf("ToyA Shape: Color:%s, High:%d, Wide:%d\n", a.Color, a.High, a.Wide)
}

func (a *ToyA) Function() {
	fmt.Println("ToyA Function: 超人打怪兽")
}

type ToyB struct {
	ToyInfo
}

func NewToyB(info ToyInfo) *ToyB {
	return &ToyB{
		ToyInfo: info,
	}
}

func (a *ToyB) Shape() {
	fmt.Printf("ToyB Shape: Color:%s, High:%d, Wide:%d\n", a.Color, a.High, a.Wide)
}

func (a *ToyB) Function() {
	fmt.Println("ToyB Function: 龟波气功")
}

type ToyC struct {
	ToyInfo
}

func NewToyC(info ToyInfo) *ToyC {
	return &ToyC{
		ToyInfo: info,
	}
}

func (a *ToyC) Shape() {
	fmt.Printf("ToyC Shape: Color:%s, High:%d, Wide:%d\n", a.Color, a.High, a.Wide)
}

func (a *ToyC) Function() {
	fmt.Println("ToyC Function: 萨瓦迪cpu")
}
