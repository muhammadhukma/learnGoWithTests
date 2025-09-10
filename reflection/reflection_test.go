package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})
	// halo nama saya adalah hukma ulil

	if len(got) != 1 {
		t.Errorf("wrong number of function call, got %d want %d", len(got), 1)
	}
}
