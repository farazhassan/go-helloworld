package errors

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const (
	cause  = "Unknown error cause"
	detail = "Unknown error detail"
)

var errorWithCause = &HTTPError{
	Cause:  errors.New(cause),
	Detail: detail,
	Status: http.StatusInternalServerError,
}

var errorWithoutCause = &HTTPError{
	Cause:  nil,
	Detail: detail,
	Status: http.StatusInternalServerError,
}

var errorWithoutDetail = &HTTPError{
	Cause: nil,
}

func TestError(t *testing.T) {

	t.Run("correct error with cause", func(t *testing.T) {
		want := fmt.Sprintf("%s : %s", detail, cause)
		got := errorWithCause.Error()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("correct error without cause", func(t *testing.T) {
		want := fmt.Sprintf("%s", detail)
		got := errorWithoutCause.Error()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestResponseBody(t *testing.T) {

	t.Run("correct error with cause", func(t *testing.T) {
		want := []byte(fmt.Sprintf(`{"detail":"%s"}`, detail))
		got, _ := errorWithCause.ResponseBody()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", string(got), string(want))
		}
	})
}

func TestResponseHeaders(t *testing.T) {

	t.Run("correct error with cause", func(t *testing.T) {
		want := http.StatusInternalServerError
		got, _ := errorWithCause.ResponseHeaders()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestNewHTTPError(t *testing.T) {
	t.Run("correct error with cause", func(t *testing.T) {
		want := errors.New(detail)
		got := NewHTTPError(nil, http.StatusBadRequest, detail)
		if got.Error() != want.Error() {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
