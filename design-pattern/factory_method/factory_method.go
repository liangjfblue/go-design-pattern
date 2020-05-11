/*
@Time : 2020/5/11 22:24
@Author : liangjiefan
*/
package factory_method

type IToy interface {
	//外形
	Shape()
	//功能
	Function()
}

type IFactory interface {
	CreateToy(info ToyInfo) IToy
}
