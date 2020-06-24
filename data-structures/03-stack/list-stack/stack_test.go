/**
 *
 * @author liangjf
 * @create on 2020/6/24
 * @version 1.0
 */
package list_stack

import (
	_3_stack "go-design-pattern/data-structures/03-stack"
	_4_list "go-design-pattern/data-structures/04-list"
	"go-design-pattern/data-structures/04-list/singlelist"
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want _3_stack.Stacker
	}{
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

func Test_stack_IsEmpty(t *testing.T) {
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
			s := stack{
				data: tt.fields.data,
			}
			if got := s.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Peek(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "IsEmpty",
			fields: fields{
				data: singlelist.New(),
			},
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				data: tt.fields.data,
			}
			s.Push(101)
			if got := s.Peek(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Pop(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   interface{}
	}{
		{
			name: "IsEmpty",
			fields: fields{
				data: singlelist.New(),
			},
			want: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				data: tt.fields.data,
			}
			s.Push(999)
			if got := s.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_Push(t *testing.T) {
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
			name: "IsEmpty",
			fields: fields{
				data: singlelist.New(),
			},
			args: args{
				100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				data: tt.fields.data,
			}
			s.Push(tt.args.e)
		})
	}
}

func Test_stack_Size(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "IsEmpty",
			fields: fields{
				data: singlelist.New(),
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				data: tt.fields.data,
			}
			for i := 0; i < 4; i++ {
				s.Push(i)
			}
			if got := s.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stack_String(t *testing.T) {
	type fields struct {
		data _4_list.Lister
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "IsEmpty",
			fields: fields{
				data: singlelist.New(),
			},
			want: "head=>1=>2=>3=>NULL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				data: tt.fields.data,
			}
			for i := 1; i <= 3; i++ {
				s.Push(i)
			}
			if got := s.String(); strings.Compare(got, tt.want) == 0 {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
