/**
 *
 * @author liangjf
 * @create on 2020/6/23
 * @version 1.0
 */
package singlelist

import (
	"testing"
)

func TestList_AddFirst(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddFirst(i)
	}

	t.Log(l)
}

func TestList_AddLast(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}

	t.Log(l)
}

func TestList_Remove(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}

	index := 0
	del := l.Next()
	for del.Next != nil {
		index++
		if index == 4 {
			break
		}
		del = del.Next
	}

	l.Remove(del)

	t.Log(l)
}

func TestList_RemoveFirst(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}

	t.Log(l)

	for i := 0; i < 2; i++ {
		l.RemoveFirst()
	}

	t.Log(l)
}

func TestList_RemoveLast(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}

	t.Log(l)

	for i := 0; i < 2; i++ {
		l.RemoveLast()
	}

	t.Log(l)
}

func TestList_Contains(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}

	if !l.Contains(4) {
		t.Fatal("not Contains 4")
	} else {
		t.Log("Contains 4")
	}

	t.Log(l)
}

func TestList_Find(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}

	if l.Find(4) < 0 {
		t.Fatal("not Find 4")
	} else {
		t.Log("Find 4")
	}

	t.Log(l)
}

func TestList_FindAll(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}
	l.AddLast(4)

	if len(l.FindAll(4)) < 0 {
		t.Fatal("not Find 4")
	} else {
		t.Log("Find 4")
	}

	t.Log(l)
}

func TestList_RemoveElement(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}

	if !l.RemoveElement(4) {
		t.Fatal("not RemoveElement 4")
	} else {
		t.Log("RemoveElement 4")
	}

	t.Log(l)
}

func TestList_RemoveAllElement(t *testing.T) {
	l := New()

	for i := 1; i < 10; i++ {
		l.AddLast(i)
	}
	l.AddLast(4)

	t.Log(l)

	if !l.RemoveAllElement(4) {
		t.Fatal("not RemoveAllElement 4")
	} else {
		t.Log("RemoveAllElement 4")
	}

	t.Log(l)
}
