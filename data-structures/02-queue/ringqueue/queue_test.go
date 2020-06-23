/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package ringqueue

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

//BenchmarkQueue_EnQueue-4        12245059               124 ns/op
func BenchmarkQueue_EnQueue(b *testing.B) {
	q := New(16)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.EnQueue(i)
	}
	b.StopTimer()
}

//BenchmarkQueue_DeQueue-4       9756049               159 ns/op
func BenchmarkQueue_DeQueue(b *testing.B) {
	q := New(16)

	for i := 0; i < b.N; i++ {
		q.EnQueue(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		q.DeQueue()
	}
	b.StopTimer()
}

//BenchmarkQueue_EnDeQueue-4      12904446                88.9 ns/op
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
