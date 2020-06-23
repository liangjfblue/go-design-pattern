/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package singlequeue

import "testing"

func TestQueue_String(t *testing.T) {
	q := New(20)

	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}

	for !q.IsEmpty() {
		t.Log(q.DeQueue())
	}

	//t.Log(q)
}

//BenchmarkQueue_EnQueue-4        12002340               122 ns/op
func BenchmarkQueue_EnQueue(b *testing.B) {
	q := New(16)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.EnQueue(i)
	}
	b.StopTimer()
}

//BenchmarkQueue_EnDeQueue-4       9836065               149 ns/op
func BenchmarkQueue_DeQueue(b *testing.B) {
	q := New(16)

	for i := 0; i < b.N; i++ {
		q.EnQueue(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.DeQueue()

		if i%4 == 0 {
			q.DeQueue()
		}
	}
	b.StopTimer()
}

//BenchmarkQueue_EnDeQueue-4      16968393                92.3 ns/op
func BenchmarkQueue_EnDeQueue(b *testing.B) {
	q := New(16)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.EnQueue(i)
		if i%3 == 0 {
			q.DeQueue()
		}
	}
	b.StopTimer()
}
