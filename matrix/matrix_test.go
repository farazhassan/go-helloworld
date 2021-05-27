package matrix

import (
	"testing"
)

var matrix = &Matrix{
	[][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	},
}

var matrixError = &Matrix{
	[][]string{
		{"1", "2", "3"},
		{"c", "5", "6"},
		{"7", "8", "9"},
	},
}

func TestEcho(t *testing.T) {

	testEcho := func(t *testing.T, matrix *Matrix, want string) {
		t.Helper()
		got := matrix.Echo()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("correct matrix", func(t *testing.T) {
		want := "1,2,3\n4,5,6\n7,8,9\n"
		testEcho(t, matrix, want)
	})

	t.Run("error matrix", func(t *testing.T) {
		want := "1,2,3\nc,5,6\n7,8,9\n"
		testEcho(t, matrixError, want)
	})
}

func TestInvert(t *testing.T) {

	testInvert := func(t *testing.T, matrix *Matrix, want string) {
		t.Helper()
		got := matrix.Invert()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("correct matrix", func(t *testing.T) {
		want := "1,4,7\n2,5,8\n3,6,9\n"
		testInvert(t, matrix, want)
	})

	t.Run("error matrix", func(t *testing.T) {
		want := "1,c,7\n2,5,8\n3,6,9\n"
		testInvert(t, matrixError, want)
	})
}

func TestFlatten(t *testing.T) {

	testFlatten := func(t *testing.T, matrix *Matrix, want string) {
		t.Helper()
		got := matrix.Flatten()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("correct matrix", func(t *testing.T) {
		want := "1,2,3,4,5,6,7,8,9"
		testFlatten(t, matrix, want)
	})

	t.Run("error matrix", func(t *testing.T) {
		want := "1,2,3,c,5,6,7,8,9"
		testFlatten(t, matrixError, want)
	})
}

func TestSum(t *testing.T) {

	testSum := func(t *testing.T, matrix *Matrix, want int, errorWanted string) {
		t.Helper()
		got, error := matrix.Sum()
		if error != nil && error.Error() != errorWanted {
			t.Errorf("got %v want %v", error, errorWanted)
			return
		}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("correct matrix", func(t *testing.T) {
		want := 45
		testSum(t, matrix, want, "")
	})

	t.Run("error matrix", func(t *testing.T) {
		want := -1
		errorWanted := "Non number value found."
		testSum(t, matrixError, want, errorWanted)
	})
}

func TestMultiply(t *testing.T) {

	testMultiply := func(t *testing.T, matrix *Matrix, want int, errorWanted string) {
		t.Helper()
		got, error := matrix.Multiply()
		if error != nil && error.Error() != errorWanted {
			t.Errorf("got %v want %v", error, errorWanted)
			return
		}
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("correct matrix", func(t *testing.T) {
		want := 362880
		testMultiply(t, matrix, want, "")
	})

	t.Run("error matrix", func(t *testing.T) {
		want := -1
		errorWanted := "Non number value found."
		testMultiply(t, matrixError, want, errorWanted)
	})
}
