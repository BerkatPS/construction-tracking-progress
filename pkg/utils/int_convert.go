package utils

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ParseInt64Param(r *http.Request) (int64, error) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
