package name

import "testing"

func TestReproducibleRandomizer(t *testing.T) {
	tests := []struct {
		i int
		o int
	}{
		{10, 1},
		{11, 5},
		{12, 5},
		{13, 4},
	}

	for _, test := range tests {
		r := getRandFromString("test")
		o := r.Intn(test.i)

		if o != test.o {
			t.Fatalf("expected: %d, got %d with input: %d", test.o, o, test.i)
		}
	}
}

func TestGenerateName(t *testing.T) {
	tests := []struct {
		i string
		p string
		c string
	}{
		{"monkey", "BetterFoal", "betterFoal"},
		{"1", "MajorTurtle", "majorTurtle"},
	}

	for _, test := range tests {
		p := GeneratePascalName(test.i)
		if p != test.p {
			t.Fatalf("expected: %s, got %s with input: %s", test.p, p, test.i)
		}

		c := GenerateCamelCaseName(test.i)
		if c != test.c {
			t.Fatalf("expected: %s, got %s with input: %s", test.c, c, test.i)
		}
	}
}
