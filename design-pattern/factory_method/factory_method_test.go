/*
@Time : 2020/5/11 22:38
@Author : liangjiefan
*/
package factory_method

import "testing"

func TestFactoryCreateToy(t *testing.T) {
	toyFactoryA := NewFactoryToyA()
	toy := toyFactoryA.CreateToy(ToyInfo{
		Color: "red",
		High:  100,
		Wide:  50,
	})
	toy.Shape()
	toy.Function()

	toyFactoryB := NewFactoryToyB()
	toy = toyFactoryB.CreateToy(ToyInfo{
		Color: "green",
		High:  200,
		Wide:  10,
	})
	toy.Shape()
	toy.Function()

	toyFactoryC := NewFactoryToyC()
	toy = toyFactoryC.CreateToy(ToyInfo{
		Color: "blue",
		High:  300,
		Wide:  60,
	})
	toy.Shape()
	toy.Function()
}
