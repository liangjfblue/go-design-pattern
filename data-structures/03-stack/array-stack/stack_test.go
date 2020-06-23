/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package array_stack

import "testing"

func TestStack_String(t *testing.T) {
	s := New(8)

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	t.Log(s)
	for !s.IsEmpty() {
		t.Log(s.Pop())
	}
}

//BenchmarkStack_Push-4           13333258                84.8 ns/op
func BenchmarkStack_Push(b *testing.B) {
	s := New(8)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	b.StopTimer()
}

//BenchmarkStack_PushPop-4        10619412               144 ns/op
func BenchmarkStack_PushPop(b *testing.B) {
	s := New(8)

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
	b.StopTimer()
}
