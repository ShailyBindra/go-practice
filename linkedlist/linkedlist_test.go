package linkedlist

import "testing"

func TestAddElementAtFront(t *testing.T) {
	cases := []struct {
		in, want int
	}{
		{1, 1},
	}

	ll := NewLinkedList()
	ll.AddElementAtFront(1)

	for _, c := range cases {
		got := ll.head.data
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}