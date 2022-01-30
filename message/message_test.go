package message

import "testing"

func TestHelloPerson(t *testing.T) {
	hoge := HelloPerson("hoge")
	if hoge != "Hello hoge!" {
		t.Error("not friendly enough")
	}
}
