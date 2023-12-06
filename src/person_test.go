package src

import "testing"

func Test_SayHello(t *testing.T) {
	p := NewPerson("John", 30)
	p.SayHello()
}

func Test_Aging(t *testing.T) {
	p := NewPerson("John", 30)
	p.Aging()

	if p.Age != 31 {
		t.Errorf("Expected 31, got %d", p.Age)
	}
}

func Test_SetName(t *testing.T) {
	p := NewPerson("John", 30)
	p.SetName("Bob")

	if p.Name != "Bob" {
		t.Errorf("Expected Bob, got %s", p.Name)
	}
}

func Test_String(t *testing.T) {
	p := NewPerson("John", 30)
	s := p.String()

	if s != "John (30 years)" {
		t.Errorf("Expected John (30 years), got %s", s)
	}
}
