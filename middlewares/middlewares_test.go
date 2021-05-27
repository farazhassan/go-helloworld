package middlewares

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	m "takehome/matrix"
)

func TestServeHTTPFile(t *testing.T) {
	nextHandlerFile := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(RequestFileMatrixKey).(*m.Matrix)
		if val == nil {
			t.Error("Matrix not provided.")
		}
	})

	handlerToTestFileToMatrixMiddleware := NewFileToMatrixMiddleware(nextHandlerFile)

	t.Run("missing file test", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/testing", nil)
		w := httptest.NewRecorder()

		handlerToTestFileToMatrixMiddleware.ServeHTTP(w, r)

		if w.Code != http.StatusBadRequest {
			t.Errorf("got %v want %v", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("correct file data test", func(t *testing.T) {
		//Set up a pipe to avoid buffering
		pr, pw := io.Pipe()
		//This writers is going to transform
		//what we pass to it to multipart form data
		//and write it to our io.Pipe
		writer := multipart.NewWriter(pw)

		go func() {
			defer writer.Close()
			//we create the form data field 'fileupload'
			//wich returns another writer to write the actual file
			part, err := writer.CreateFormFile("file", "matrix.csv")
			if err != nil {
				t.Error(err)
			}

			data := "1,2,3\n4,5,6\n7,8,9"

			buf := bytes.NewBufferString(data)
			io.Copy(part, buf)
		}()

		r := httptest.NewRequest("POST", "/testing", pr)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		handlerToTestFileToMatrixMiddleware.ServeHTTP(w, r)
		if w.Code != http.StatusOK {
			t.Errorf("got %v want %v", w.Code, http.StatusOK)
		}
	})

	t.Run("incorrect character file data test", func(t *testing.T) {
		//Set up a pipe to avoid buffering
		pr, pw := io.Pipe()
		//This writers is going to transform
		//what we pass to it to multipart form data
		//and write it to our io.Pipe
		writer := multipart.NewWriter(pw)

		go func() {
			defer writer.Close()
			//we create the form data field 'fileupload'
			//wich returns another writer to write the actual file
			part, err := writer.CreateFormFile("file", "matrix.csv")
			if err != nil {
				t.Error(err)
			}

			data := "1,2,3\nc,5,6\n7,8,9"

			buf := bytes.NewBufferString(data)
			io.Copy(part, buf)
		}()

		r := httptest.NewRequest("POST", "/testing", pr)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		handlerToTestFileToMatrixMiddleware.ServeHTTP(w, r)
		if w.Code != http.StatusBadRequest {
			t.Errorf("got %v want %v", w.Code, http.StatusBadRequest)
		}
	})

	t.Run("incorrect row items file data test", func(t *testing.T) {
		pr, pw := io.Pipe()
		writer := multipart.NewWriter(pw)

		go func() {
			defer writer.Close()
			part, err := writer.CreateFormFile("file", "matrix.csv")
			if err != nil {
				t.Error(err)
			}

			data := "1,2,3\nc,5,6,7\n7,8,9"

			buf := bytes.NewBufferString(data)
			io.Copy(part, buf)
		}()

		r := httptest.NewRequest("POST", "/testing", pr)
		r.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		handlerToTestFileToMatrixMiddleware.ServeHTTP(w, r)
		if w.Code != http.StatusBadRequest {
			t.Errorf("got %v want %v", w.Code, http.StatusBadRequest)
		}
	})
}

func TestServeHTTP(t *testing.T) {
	nextHandlerMethod := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	handlerToTestPOSTMiddleware := NewPOSTMethodOnlyMiddleware(nextHandlerMethod)

	t.Run("non POST method test", func(t *testing.T) {
		r := httptest.NewRequest("GET", "/testing", nil)
		w := httptest.NewRecorder()

		handlerToTestPOSTMiddleware.ServeHTTP(w, r)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("got %v want %v", w.Code, http.StatusMethodNotAllowed)
		}
	})

	t.Run("POST method test", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/testing", nil)
		w := httptest.NewRecorder()

		handlerToTestPOSTMiddleware.ServeHTTP(w, r)

		if w.Code == http.StatusMethodNotAllowed {
			t.Errorf("got %v want %v", w.Code, http.StatusMethodNotAllowed)
		}
	})
}
