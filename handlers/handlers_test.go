package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	m "takehome/matrix"
	middlewares "takehome/middlewares"
)

var matrix = &m.Matrix{
	Data: [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	},
}

var wrongMatrix = &m.Matrix{
	Data: [][]string{
		{"1", "2", "3"},
		{"c", "5", "6"},
		{"7", "8", "9"},
	},
}

func TestEcho(t *testing.T) {

	req, err := http.NewRequest("POST", "/echo", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.Handler(RootHandler(Echo))

	t.Run("no matrix provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := fmt.Sprintf(`{"detail":"%s%s"}`, BadRequestErrorFormat, MatrixNotProvidedError)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		ctxWithMatrix := context.WithValue(req.Context(), middlewares.RequestFileMatrixKey, matrix)
		rWithMatrix := req.WithContext(ctxWithMatrix)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, rWithMatrix)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "1,2,3\n4,5,6\n7,8,9\n"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestInvert(t *testing.T) {
	req, err := http.NewRequest("POST", "/invert", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.Handler(RootHandler(Invert))

	t.Run("no matrix provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := fmt.Sprintf(`{"detail":"%s%s"}`, BadRequestErrorFormat, MatrixNotProvidedError)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		ctxWithMatrix := context.WithValue(req.Context(), middlewares.RequestFileMatrixKey, matrix)
		rWithMatrix := req.WithContext(ctxWithMatrix)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, rWithMatrix)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "1,4,7\n2,5,8\n3,6,9\n"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestFlatten(t *testing.T) {
	req, err := http.NewRequest("POST", "/flatten", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.Handler(RootHandler(Flatten))

	t.Run("no matrix provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := fmt.Sprintf(`{"detail":"%s%s"}`, BadRequestErrorFormat, MatrixNotProvidedError)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		ctxWithMatrix := context.WithValue(req.Context(), middlewares.RequestFileMatrixKey, matrix)
		rWithMatrix := req.WithContext(ctxWithMatrix)

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, rWithMatrix)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "1,2,3,4,5,6,7,8,9"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestSum(t *testing.T) {
	req, err := http.NewRequest("POST", "/sum", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.Handler(RootHandler(Sum))

	t.Run("no matrix provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := fmt.Sprintf(`{"detail":"%s%s"}`, BadRequestErrorFormat, MatrixNotProvidedError)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("wrong matrix provided", func(t *testing.T) {
		ctxWithMatrix := context.WithValue(req.Context(), middlewares.RequestFileMatrixKey, wrongMatrix)
		rWithMatrix := req.WithContext(ctxWithMatrix)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, rWithMatrix)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := fmt.Sprintf(`{"detail":"%s%s"}`, BadRequestErrorFormat, NonDigitFoundError)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		ctxWithMatrix := context.WithValue(req.Context(), middlewares.RequestFileMatrixKey, matrix)
		rWithMatrix := req.WithContext(ctxWithMatrix)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, rWithMatrix)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "45"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}

func TestMultiply(t *testing.T) {
	req, err := http.NewRequest("POST", "/multiply", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.Handler(RootHandler(Multiply))

	t.Run("no matrix provided", func(t *testing.T) {
		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := fmt.Sprintf(`{"detail":"%s%s"}`, BadRequestErrorFormat, MatrixNotProvidedError)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("wrong matrix provided", func(t *testing.T) {
		ctxWithMatrix := context.WithValue(req.Context(), middlewares.RequestFileMatrixKey, wrongMatrix)
		rWithMatrix := req.WithContext(ctxWithMatrix)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, rWithMatrix)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}

		expected := fmt.Sprintf(`{"detail":"%s%s"}`, BadRequestErrorFormat, NonDigitFoundError)
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("happy path", func(t *testing.T) {
		ctxWithMatrix := context.WithValue(req.Context(), middlewares.RequestFileMatrixKey, matrix)
		rWithMatrix := req.WithContext(ctxWithMatrix)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, rWithMatrix)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := "362880"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
