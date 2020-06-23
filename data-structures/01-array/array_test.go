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

	for i := 8; i < 20; i++ {
		arr.AddFirst(i)
	}

	for i := 1; i < 20; i++ {
		arr.RemoveFirst()
	}

	t.Log(arr.String())
}

func BenchmarkArray_Add(b *testing.B) {
	arr := New(10)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		arr.AddLast(i)
	}
	b.StopTimer()
}

//BenchmarkArray_RemoveFirst-4       67414             72225 ns/op
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

//BenchmarkArray_RemoveLast-4     10619440               157 ns/op
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
