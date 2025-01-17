package collections_test

import (
	"testing"

	"github.com/gossie/collections"
)

func TestThatQueueIsEmpty(t *testing.T) {
	s := collections.NewQueue[string]()
	if !s.Empty() {
		t.Fatal("Queue is not empty")
	}
}

func TestThatQueueIsNotEmpty(t *testing.T) {
	s := collections.NewQueue[string]()
	s.Push("eins")
	if s.Empty() {
		t.Fatal("Queue is empty")
	}
}

func TestQueueLength(t *testing.T) {
	s := collections.NewQueue[string]()
	s.Push("eins")
	s.Push("zwei")
	if s.Len() != 2 {
		t.Fatalf("length = %v", s.Len())
	}
}

func TestQueuePushPeek(t *testing.T) {
	s := collections.NewQueue[string]()
	s.Push("eins")
	s.Push("zwei")
	value, err := s.Peek()

	if s.Len() != 2 {
		t.Fatalf("length = %v", s.Len())
	}

	if value != "eins" {
		t.Fatalf("value = %v", value)
	}

	if err != nil {
		t.Fatalf("err = %v", err)
	}
}

func TestQueuePeekError(t *testing.T) {
	s := collections.NewQueue[string]()
	value, err := s.Peek()

	if value != "" {
		t.Fatalf("value = %v", value)
	}

	if err.Error() != "queue is empty" {
		t.Fatalf("err = %v", err)
	}
}

func TestQueuePushPop(t *testing.T) {
	s := collections.NewQueue[string]()
	s.Push("eins")
	s.Push("zwei")
	value, err := s.Pop()

	if s.Len() != 1 {
		t.Fatalf("length = %v", s.Len())
	}

	if value != "eins" {
		t.Fatalf("value = %v", value)
	}

	if err != nil {
		t.Fatalf("err = %v", err)
	}
}

func TestQueuePopError(t *testing.T) {
	s := collections.NewQueue[string]()
	value, err := s.Pop()

	if value != "" {
		t.Fatalf("value = %v", value)
	}

	if err.Error() != "queue is empty" {
		t.Fatalf("err = %v", err)
	}
}

func TestQueuePushAll(t *testing.T) {
	s1 := collections.NewQueue[string]()
	s1.Push("eins")
	s1.Push("zwei")

	s2 := collections.NewQueue[string]()
	s2.PushAll(s1)

	value1, _ := s2.Pop()
	if value1 != "eins" {
		t.Fatalf("value1 = %v", value1)
	}

	if s2.Len() != 1 {
		t.Fatalf("length = %v", s2.Len())
	}

	value2, _ := s2.Pop()
	if value2 != "zwei" {
		t.Fatalf("value2 = %v", value2)
	}

	if !s2.Empty() {
		t.Fatalf("length = %v", s2.Len())
	}

}

func TestQueuePopMultiple(t *testing.T) {
	s1 := collections.NewQueue[string]()
	s1.Push("eins")
	s1.Push("zwei")
	s1.Push("drei")

	s2, _ := s1.PopMultiple(2)

	if s1.Len() != 1 {
		t.Fatalf("length = %v", s1.Len())
	}

	value1, _ := s1.Pop()
	if value1 != "drei" {
		t.Fatalf("value1 = %v", value1)
	}

	if s2.Len() != 2 {
		t.Fatalf("length = %v", s2.Len())
	}

	value2, _ := s2.Pop()
	if value2 != "eins" {
		t.Fatalf("value2 = %v", value2)
	}

	if s2.Len() != 1 {
		t.Fatalf("length = %v", s2.Len())
	}

	value3, _ := s2.Pop()
	if value3 != "zwei" {
		t.Fatalf("value3 = %v", value3)
	}

	if !s2.Empty() {
		t.Fatalf("length = %v", s2.Len())
	}
}

func TestQueuePopMultipleError(t *testing.T) {
	s1 := collections.NewQueue[string]()
	s1.Push("eins")
	s1.Push("zwei")

	_, err := s1.PopMultiple(3)

	if err.Error() != "queue does not contain 3 elements" {
		t.Fatalf("err = %v", err)
	}
}
