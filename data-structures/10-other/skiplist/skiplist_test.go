/**
 *
 * @author liangjf
 * @create on 2020/7/17
 * @version 1.0
 */
package skiplist

import (
	"math/rand"
	"testing"
	"time"
)

func TestSkipList_Insert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var err error

	sl := New()

	for i := 0; i < 100; i++ {
		err = sl.Insert(&Node{
			Key:   i,
			Value: i,
		})
	}

	if err != nil {
		t.Fatal(err)
	}

	t.Log(sl)
}

func TestSkipList_Delete(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var err error

	sl := New()

	for i := 99; i >= 0; i-- {
		err = sl.Insert(&Node{
			Key:   i,
			Value: i,
		})
	}

	if err != nil {
		t.Fatal(err)
	}

	t.Log(sl)

	err = sl.Delete(4)

	t.Log(sl)
}

func TestSkipList_Search(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var err error

	sl := New()

	for i := 99; i >= 0; i-- {
		err = sl.Insert(&Node{
			Key:   i,
			Value: i,
		})
	}

	if err != nil {
		t.Fatal(err)
	}

	t.Log(sl)

	t.Log(sl.Search(5))
}

func BenchmarkSearch(b *testing.B) {
	var err error

	sl := New()
	for i := b.N - 1; i >= 0; i-- {
		err = sl.Insert(&Node{
			Key:   i,
			Value: i,
		})
	}
	if err != nil {
		b.Fatal(err)
	}

	b.StartTimer()
	defer b.StopTimer()
	for i := 0; i < b.N; i++ {
		sl.Search(i)
	}
}
