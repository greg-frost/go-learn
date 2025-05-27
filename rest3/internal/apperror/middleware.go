package apperror

import (
	"errors"
	"net/http"
)

// Функция "обработчик приложения"
type appHandler func(w http.ResponseWriter, r *http.Request) error

// Промежуточный слой
func Middleware(next appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := next(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")

			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
					return
				}
				if errors.Is(err, ErrNotAuth) {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write(ErrNotAuth.Marshal())
					return
				}
				appErr = err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(appErr.Marshal())
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(systemError(err).Marshal())
		}
	}
}
