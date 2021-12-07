package res

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/maurana/nuswantara/core/model"
	"github.com/maurana/nuswantara/core/config"
	"github.com/maurana/nuswantara/core/constant"
	"github.com/go-chi/chi"
)

// Response
func MonkFeedback(w http.ResponseWriter, code int, feedback interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(feedback)
}

func MonkError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(model.ErrorResponse{Message: err.Error()})
}

func MonkErrorException(w http.ResponseWriter, code int, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}

// URL Handle
func GetUrlPathString(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func GetUrlPathInt(r *http.Request, key string) (int, error) {
	i, err := strconv.Atoi(chi.URLParam(r, key))
	if err != nil {
		return 0, constant.ErrUrlPathParameter
	}
	return i, nil
}

func GetUrlPathInt64(r *http.Request, key string) (int64, error) {
	i, err := strconv.ParseInt(chi.URLParam(r, key), 10, 64)
	if err != nil {
		return 0, constant.ErrUrlPathParameter
	}
	return i, nil
}

func GetUrlQueryString(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetUrlQueryInt(r *http.Request, key string) (int, error) {
	i, err := strconv.Atoi(key)
	if err != nil {
		return 0, constant.ErrUrlQueryParameter
	}
	return i, nil
}

func GetUrlQueryInt64(r *http.Request, key string) (int64, error) {
	i, err := strconv.ParseInt(r.URL.Query().Get(key), 10, 64)
	if err != nil {
		return 0, constant.ErrUrlQueryParameter
	}
	return i, nil
}

func GetPagination(r *http.Request) (limit, offset int, err error) {
	limitQuery := r.URL.Query().Get("limit")
	offsetQuery := r.URL.Query().Get("offset")

	if limitQuery == "" {
		limit = config.Cfg().PaginationLimit
	} else {
		limit, err = GetUrlQueryInt(r, limitQuery)
		if err != nil {
			return 0, 0, err
		} else if limit < 0 {
			return 0, 0, constant.ErrUrlQueryParameter
		} else if limit > config.Cfg().PaginationLimit {
			limit = config.Cfg().PaginationLimit
		}
	}

	if offsetQuery == "" {
		offset = 0
	} else {
		offset, err = GetUrlQueryInt(r, offsetQuery)
		if err != nil {
			return 0, 0, err
		} else if offset < 0 {
			return 0, 0, constant.ErrUrlQueryParameter
		}
	}
	return
}
