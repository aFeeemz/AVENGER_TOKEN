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
			http.Error(w, "Unoauthorized : Token Empty", http.StatusUnauthorized)
			return
		}

		//
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte("kue_apeeeemz"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized : Token Invalid", http.StatusUnauthorized)
			return
		}

		next(w, r, ps)

	}
}
