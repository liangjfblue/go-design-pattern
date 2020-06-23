/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package _1_array

import "testing"

func TestArray_String(t *testing.T) {
	arr := New(10)
	for i := 0; i < 12; i++ {
		arr.AddLast(i)
	}

	for i := 8; i < 10; i++ {
		arr.AddFirst(i)
	}

	arr.Set(1, 100)

	t.Log(arr.String())
	t.Log(arr.FindAll(0))
}

func BenchmarkArray_Add(b *testing.B) {
	arr := New(10)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		arr.AddLast(i)
	}
	b.StopTimer()
}

//BenchmarkArray_RemoveFirst-4          93024               110025 ns/op
func BenchmarkArray_RemoveFirst(b *testing.B) {
	arr := New(10)

	for i := 0; i < b.N; i++ {
		arr.AddLast(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		arr.RemoveFirst()
	}
	b.StopTimer()
}

//BenchmarkArray_RemoveLast-4          2714934               501 ns/op
func BenchmarkArray_RemoveLast(b *testing.B) {
	arr := New(10)

	for i := 0; i < b.N; i++ {
		arr.AddLast(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		arr.RemoveLast()
	}
	b.StopTimer()
}
