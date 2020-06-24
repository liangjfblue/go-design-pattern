/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package priorityqueue

import (
	_5_heap "go-design-pattern/data-structures/05-heap"
	big_heap "go-design-pattern/data-structures/05-heap/big-heap"
	"reflect"
	"testing"
)

func Test_queue_Cap(t *testing.T) {
	type fields struct {
		maxHeap _5_heap.Heaper
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Cap",
			fields: fields{
				maxHeap: big_heap.New(16),
			},
			want: 32,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				maxHeap: tt.fields.maxHeap,
			}
			for i := 0; i < 20; i++ {
				q.EnQueue(i)
			}
			if got := q.Cap(); got != tt.want {
				t.Errorf("Cap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_DeQueue(t *testing.T) {
	type fields struct {
		maxHeap _5_heap.Heaper
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "DeQueue",
			fields: fields{
				maxHeap: big_heap.New(16),
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				maxHeap: tt.fields.maxHeap,
			}
			for i := 0; i < 10; i++ {
				q.EnQueue(i)
			}
			q.EnQueue(100)

			if got := q.DeQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_EnQueue(t *testing.T) {
	type fields struct {
		maxHeap _5_heap.Heaper
	}
	type args struct {
		e interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "DeQueue",
			fields: fields{
				maxHeap: big_heap.New(16),
			},
			args: args{
				e: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				maxHeap: tt.fields.maxHeap,
			}
			q.EnQueue(tt.args.e)
		})
	}
}

func Test_queue_Front(t *testing.T) {
	type fields struct {
		maxHeap _5_heap.Heaper
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "DeQueue",
			fields: fields{
				maxHeap: big_heap.New(16),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				maxHeap: tt.fields.maxHeap,
			}
			for i := 0; i < 11; i++ {
				q.EnQueue(i)
			}
			if got := q.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_IsEmpty(t *testing.T) {
	type fields struct {
		maxHeap _5_heap.Heaper
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "DeQueue",
			fields: fields{
				maxHeap: big_heap.New(16),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				maxHeap: tt.fields.maxHeap,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Size(t *testing.T) {
	type fields struct {
		maxHeap _5_heap.Heaper
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "DeQueue",
			fields: fields{
				maxHeap: big_heap.New(16),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				maxHeap: tt.fields.maxHeap,
			}
			for i := 0; i < 10; i++ {
				q.EnQueue(i)
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_String(t *testing.T) {
	q := &queue{
		maxHeap: big_heap.New(16),
	}
	for i := 0; i < 5; i++ {
		q.EnQueue(i)
	}

	t.Log(q)
}
