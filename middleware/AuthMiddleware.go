package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Unoauthorized", http.StatusUnauthorized)
			return
		}

		//
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r, ps)

	}
}
