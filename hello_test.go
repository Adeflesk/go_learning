package hello

import "testing"

func TestHello(t *testing.T) {
	want := "Hello, test!"
	got := Say([]string{"test"})
	if want != got {
		t.Errorf("wanted %s, got %s", want, got)
	}
}
