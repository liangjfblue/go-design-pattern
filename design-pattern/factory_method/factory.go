/*
@Time : 2020/5/11 22:29
@Author : liangjiefan
*/
package factory_method

type FactoryToyA struct {
}

func NewFactoryToyA() *FactoryToyA {
	return &FactoryToyA{}
}

func (a *FactoryToyA) CreateToy(info ToyInfo) IToy {
	return NewToyA(info)
}

type FactoryToyB struct {
}

func NewFactoryToyB() *FactoryToyB {
	return &FactoryToyB{}
}

func (a *FactoryToyB) CreateToy(info ToyInfo) IToy {
	return NewToyB(info)
}

type FactoryToyC struct {
}

func NewFactoryToyC() *FactoryToyC {
	return &FactoryToyC{}
}

func (a *FactoryToyC) CreateToy(info ToyInfo) IToy {
	return NewToyC(info)
}
