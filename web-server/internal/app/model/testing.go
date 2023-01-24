package model

import "testing"

func TestUser(t *testing.T) *User {
	t.Helper()

	return &User{
		Name: "Alex",
		Address: "New York",
	}
}