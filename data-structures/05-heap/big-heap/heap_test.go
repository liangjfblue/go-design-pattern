/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package big_heap

import (
	_1_array "go-design-pattern/data-structures/01-array"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_heap_ExtractMax(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			want: 29,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			//for i := 9; i >= 0; i-- {
			//	h.Insert(i)
			//}

			for i := 0; i < 10; i++ {
				h.Insert(i)
			}
			h.Insert(29)

			//for i := 0; i < 11; i++ {
			//	t.Logf("%d\t", h.ExtractMax())
			//}
			if got := h.ExtractMax(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_BuildHeap(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	h := &heap{
		data: _1_array.New(16),
	}

	arr := make([]interface{}, 0)
	for i := 0; i < 8; i++ {
		arr = append(arr, rand.Intn(100))
	}

	h.BuildHeap(arr)

	for !h.IsEmpty() {
		t.Log(h.ExtractMax())
	}
}

func Test_heap_Insert(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
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
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			args: args{
				10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			h.Insert(tt.args.e)
		})
	}
}

func Test_heap_IsEmpty(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			if got := h.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_Replace(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	type args struct {
		e interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			args: args{
				e: 20,
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			for i := 0; i < 10; i++ {
				h.Insert(i)
			}
			if got := h.Replace(tt.args.e); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_Size(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			for i := 0; i < 5; i++ {
				h.Insert(i)
			}
			if got := h.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_heap_getMax(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			for i := 0; i < 5; i++ {
				h.Insert(i)
			}
			if got := h.getMax(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_leftChild(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			args: args{idx: 1},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			for i := 0; i < 5; i++ {
				h.Insert(i)
			}
			if got := h.leftChild(tt.args.idx); got != tt.want {
				t.Errorf("leftChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_parent(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			args: args{idx: 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			for i := 0; i < 5; i++ {
				h.Insert(i)
			}
			if got := h.parent(tt.args.idx); got != tt.want {
				t.Errorf("rightChild() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heap_String(t *testing.T) {
	h := New(16)

	num := 30
	for i := 0; i < num; i++ {
		h.Insert(rand.Intn(50))
	}
	h.Insert(100)

	for i := 0; i < num+1; i++ {
		t.Logf("%d\t", h.ExtractMax())
	}
}

func Test_heap_rightChild(t *testing.T) {
	type fields struct {
		data _1_array.Arrayer
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "ExtractMax",
			fields: fields{
				data: _1_array.New(16),
			},
			args: args{idx: 1},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &heap{
				data: tt.fields.data,
			}
			for i := 0; i < 5; i++ {
				h.Insert(i)
			}
			if got := h.rightChild(tt.args.idx); got != tt.want {
				t.Errorf("rightChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
