package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"week04/internal/service"
)

func UserIndex(service service.UserCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// todo 入参校验这里需要加强
		uid, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
		user, err := service.Query(uid)
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading."))
			return
		}
		if err := json.NewEncoder(w).Encode(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error reading."))
		}
	})
}

// todo 这个函数有时间了提炼成为一个register function，类似option function 那种写法
func MakeUserHandlers(s *http.ServeMux, service service.UserCase) {
	s.Handle("/v1/bookmark", UserIndex(service))
}
