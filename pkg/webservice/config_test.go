package webservice

import "testing"

func TestCreateConfig(t *testing.T) {
	u := "testuser"
	p := "testpass"

	c := NewConfig(u, p)

	if c.Username != u {
		t.Errorf("Username was not set properly, got %s, want: %s", c.Username, u)
	}

	if c.Password != p {
		t.Errorf("Password was not set properly, got %s, want: %s", c.Password, p)
	}

	// @TODO Add more fields + default values when decided
}
