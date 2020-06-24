/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package listqueue

import (
	_2_queue "go-design-pattern/data-structures/02-queue"
	_4_list "go-design-pattern/data-structures/04-list"
	"go-design-pattern/data-structures/04-list/singlelist"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want _2_queue.Queuer
	}{
		// TODO: Add test cases.
		{
			name: "new",
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Cap(t *testing.T) {

}

func Test_queue_DeQueue(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "DeQueue",
			fields: fields{
				data: singlelist.New(),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
			}

			q.EnQueue(10)

			if got := q.DeQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_EnQueue(t *testing.T) {
	type fields struct {
		data _4_list.Lister
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
			name: "EnQueue",
			fields: fields{
				data: singlelist.New(),
			},
			args: args{
				e: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
			}
			q.EnQueue(tt.args.e)
		})
	}
}

func Test_queue_Front(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "Front",
			fields: fields{
				data: singlelist.New(),
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
			}
			q.EnQueue(10)
			if got := q.Front(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Front() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_IsEmpty(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "IsEmpty",
			fields: fields{
				data: singlelist.New(),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_Size(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Size",
			fields: fields{
				data: singlelist.New(),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queue_String(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "String",
			fields: fields{
				data: singlelist.New(),
			},
			want: "head=>1=>2=>3=>NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &queue{
				data: tt.fields.data,
			}
			for i := 1; i <= 3; i++ {
				q.EnQueue(i)
			}

			if got := q.String(); strings.Compare(got, tt.want) == 0 {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
