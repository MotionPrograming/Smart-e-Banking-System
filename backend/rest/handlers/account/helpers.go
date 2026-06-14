package account

import (
	"errors"
	"net/http"
	"strconv"
)

func parseID(r *http.Request, key string) (int64, error) {
	s := r.PathValue(key)
	if s == "" {
		return 0, errors.New("missing " + key)
	}
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil || id <= 0 {
		return 0, errors.New("invalid " + key)
	}
	return id, nil
}
