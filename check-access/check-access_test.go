package main

import "testing"

func TestHello(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Julien", "Welcome Julien!"},
		{"John", "Access forbidden for John!"},
	}
	for _, c := range cases {
		got := checkAccess(c.in, administrators)
		if got != c.want {
			t.Errorf("checkAccess(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
