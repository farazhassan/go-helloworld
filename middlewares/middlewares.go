package middlewares

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"

	m "takehome/matrix"
)

type contextKey int

const RequestFileMatrixKey contextKey = 0

type FileToMatrixMiddleware struct {
	handler http.Handler
}

func (ftm *FileToMatrixMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, `{"error": "File not found."}`, http.StatusBadRequest)
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		http.Error(w, `{"error": "Incorrect file data."}`, http.StatusBadRequest)
		return
	}

	var matrix m.Matrix
	for _, row := range records {
		// var introw []int
		for _, val := range row {
			_, err := strconv.Atoi(val)
			if err != nil {
				http.Error(w, fmt.Sprintf(`{"error": "Item '%s' is not an integer."}`, val), http.StatusBadRequest)
				return
			}
		}
		matrix.Data = append(matrix.Data, row)
	}

	ctxWithMatrix := context.WithValue(r.Context(), RequestFileMatrixKey, &matrix)
	rWithMatrix := r.WithContext(ctxWithMatrix)

	ftm.handler.ServeHTTP(w, rWithMatrix)
}

func NewFileToMatrixMiddleware(handlerToWrap http.Handler) *FileToMatrixMiddleware {
	return &FileToMatrixMiddleware{handlerToWrap}
}

type POSTMethodOnlyMiddleware struct {
	handler http.Handler
}

func (pmom *POSTMethodOnlyMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pmom.handler.ServeHTTP(w, r)
}

func NewPOSTMethodOnlyMiddleware(handlerToWrap http.Handler) *POSTMethodOnlyMiddleware {
	return &POSTMethodOnlyMiddleware{handlerToWrap}
}
