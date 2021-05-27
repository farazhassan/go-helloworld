package handlers

import (
	"fmt"
	"net/http"

	err "takehome/errors"
	m "takehome/matrix"
	middlewares "takehome/middlewares"
)

const (
	BadRequestErrorFormat  = "Bad request : "
	MatrixNotProvidedError = "matrix not provided."
	NonDigitFoundError     = "non-digit character found."
)

type RootHandler func(http.ResponseWriter, *http.Request) error

func (fn RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	error := fn(w, r) // Call handler function
	if error == nil {
		return
	}
	// This is where our error handling logic starts.

	clientError, ok := error.(err.ClientError) // Check if it is a ClientError.
	if !ok {
		// If the error is not ClientError, assume that it is ServerError.
		w.WriteHeader(http.StatusInternalServerError) // return 500 Internal Server Error.
		return
	}

	body, error := clientError.ResponseBody() // Try to get response body of ClientError.
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	status, headers := clientError.ResponseHeaders() // Get http status code and headers.
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
	w.Write(body)
}

func Echo(w http.ResponseWriter, r *http.Request) error {
	matrix, ok := r.Context().Value(middlewares.RequestFileMatrixKey).(*m.Matrix)
	if !ok {
		return err.NewHTTPError(nil, http.StatusBadRequest, fmt.Sprintf("%s%s", BadRequestErrorFormat, MatrixNotProvidedError))
	}
	fmt.Fprint(w, matrix.Echo())
	return nil
}

func Invert(w http.ResponseWriter, r *http.Request) error {
	matrix, ok := r.Context().Value(middlewares.RequestFileMatrixKey).(*m.Matrix)
	if !ok {
		// http.Error(w, "Matrix not provided.", http.StatusBadRequest)
		return err.NewHTTPError(nil, http.StatusBadRequest, fmt.Sprintf("%s%s", BadRequestErrorFormat, MatrixNotProvidedError))
	}
	fmt.Fprint(w, matrix.Invert())
	return nil
}

func Flatten(w http.ResponseWriter, r *http.Request) error {
	matrix, ok := r.Context().Value(middlewares.RequestFileMatrixKey).(*m.Matrix)
	if !ok {
		return err.NewHTTPError(nil, http.StatusBadRequest, fmt.Sprintf("%s%s", BadRequestErrorFormat, MatrixNotProvidedError))
	}
	fmt.Fprint(w, matrix.Flatten())
	return nil
}

func Sum(w http.ResponseWriter, r *http.Request) error {
	matrix, ok := r.Context().Value(middlewares.RequestFileMatrixKey).(*m.Matrix)
	if !ok {
		return err.NewHTTPError(nil, http.StatusBadRequest, fmt.Sprintf("%s%s", BadRequestErrorFormat, MatrixNotProvidedError))
	}
	sum, error := matrix.Sum()
	if error != nil {
		return err.NewHTTPError(error, http.StatusBadRequest, fmt.Sprintf("%s%s", BadRequestErrorFormat, NonDigitFoundError))
	}
	fmt.Fprint(w, sum)
	return nil
}

func Multiply(w http.ResponseWriter, r *http.Request) error {
	matrix, ok := r.Context().Value(middlewares.RequestFileMatrixKey).(*m.Matrix)
	if !ok {
		return err.NewHTTPError(nil, http.StatusBadRequest, fmt.Sprintf("%s%s", BadRequestErrorFormat, MatrixNotProvidedError))
	}
	product, error := matrix.Multiply()
	if error != nil {
		return err.NewHTTPError(error, http.StatusBadRequest, fmt.Sprintf("%s%s", BadRequestErrorFormat, NonDigitFoundError))
	}
	fmt.Fprint(w, product)
	return nil
}
