/**
 *
 * @author liangjf
 * @create on 2020/6/28
 * @version 1.0
 */
package binary_search_tree

import (
	"testing"
)

func TestBSTAdd(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)

	b.InOrder()

	t.Log(b.Contains(5))
	t.Log(b.Contains(57))
}

func TestBSTContains(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)

	t.Log(b.Contains(5))
	t.Log(b.Contains(57))
}

func TestBSTMinimum(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)
	b.Add(6)

	t.Log(b.Minimum())
}

func TestBSTMaximum(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)
	b.Add(6)

	t.Log(b.Maximum())
}

func TestBSTRemove(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)

	b.InOrder()

	b.Remove(5)

	b.InOrder()
}

func TestBSTRemoveMin(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)

	b.InOrder()

	b.RemoveMin()

	b.InOrder()
}

func TestBSTRemoveMax(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)

	b.InOrder()

	b.RemoveMax()

	b.InOrder()
}

func TestBSTLevelOrder(t *testing.T) {
	b := New()
	b.Add(10)
	b.Add(3)
	b.Add(5)
	b.Add(1)

	b.LevelOrder()
}
