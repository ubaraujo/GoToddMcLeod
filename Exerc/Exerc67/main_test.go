package main

import "testing"

func TestGetUser(t *testing.T) {
	md := &MockDatastore{
		Users: map[int]User{
			3: {ID: 3, First: "Tom"},
		},
	}
	s := &Service{
		ds: md,
	}

	u, err := s.GetUser(3)
	if err != nil {
		t.Errorf("go error: %v\n", err)
	}

	if u.First != "Tom" {
		t.Errorf("got: %s, want: %s", u.First, "Tom")
	}
}
