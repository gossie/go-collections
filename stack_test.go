package gocollections

import "testing"

func TestThatStackIsEmpty(t *testing.T) {
	s := CreateNew[string]()
	if !s.Empty() {
		t.Fatal("stack is not empty")
	}
}

func TestThatStackIsNotEmpty(t *testing.T) {
	s := CreateNew[string]()
	s.Push("eins")
	if s.Empty() {
		t.Fatal("stack is empty")
	}
}

func TestLength(t *testing.T) {
	s := CreateNew[string]()
	s.Push("eins")
	s.Push("zwei")
	if s.Len() != 2 {
		t.Fatalf("length = %v", s.Len())
	}
}

func TestPushPeek(t *testing.T) {
	s := CreateNew[string]()
	s.Push("eins")
	s.Push("zwei")
	value, err := s.Peek()

	if s.Len() != 2 {
		t.Fatalf("length = %v", s.Len())
	}

	if value != "zwei" {
		t.Fatalf("value = %v", value)
	}

	if err != nil {
		t.Fatalf("err = %v", err)
	}
}

func TestPeekError(t *testing.T) {
	s := CreateNew[string]()
	value, err := s.Pop()

	if value != "" {
		t.Fatalf("value = %v", value)
	}

	if err.Error() != "stack is empty" {
		t.Fatalf("err = %v", err)
	}
}

func TestPushPop(t *testing.T) {
	s := CreateNew[string]()
	s.Push("eins")
	s.Push("zwei")
	value, err := s.Pop()

	if s.Len() != 1 {
		t.Fatalf("length = %v", s.Len())
	}

	if value != "zwei" {
		t.Fatalf("value = %v", value)
	}

	if err != nil {
		t.Fatalf("err = %v", err)
	}
}

func TestPopError(t *testing.T) {
	s := CreateNew[string]()
	value, err := s.Pop()

	if value != "" {
		t.Fatalf("value = %v", value)
	}

	if err.Error() != "stack is empty" {
		t.Fatalf("err = %v", err)
	}
}

func TestPushAll(t *testing.T) {
	s1 := CreateNew[string]()
	s1.Push("eins")
	s1.Push("zwei")

	s2 := CreateNew[string]()
	s2.PushAll(s1)

	value1, _ := s2.Pop()
	if value1 != "zwei" {
		t.Fatalf("value1 = %v", value1)
	}

	if s2.Len() != 1 {
		t.Fatalf("length = %v", s2.Len())
	}

	value2, _ := s2.Pop()
	if value2 != "eins" {
		t.Fatalf("value2 = %v", value2)
	}

	if !s2.Empty() {
		t.Fatalf("length = %v", s2.Len())
	}

}

func TestPopMultiple(t *testing.T) {
	s1 := CreateNew[string]()
	s1.Push("eins")
	s1.Push("zwei")
	s1.Push("drei")

	s2, _ := s1.PopMultiple(2)

	if s1.Len() != 1 {
		t.Fatalf("length = %v", s1.Len())
	}

	value1, _ := s1.Pop()
	if value1 != "eins" {
		t.Fatalf("value1 = %v", value1)
	}

	if s2.Len() != 2 {
		t.Fatalf("length = %v", s2.Len())
	}

	value2, _ := s2.Pop()
	if value2 != "drei" {
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

func TestPopMultipleError(t *testing.T) {
	s1 := CreateNew[string]()
	s1.Push("eins")
	s1.Push("zwei")

	_, err := s1.PopMultiple(3)

	if err.Error() != "stack does not contain 3 elements" {
		t.Fatalf("err = %v", err)
	}
}
