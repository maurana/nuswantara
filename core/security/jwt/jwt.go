package jwt

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
        "github.com/maurana/nuswantara/core/config"
	"github.com/maurana/nuswantara/core/constant"
	"github.com/maurana/nuswantara/util/res"
	jwt "github.com/dgrijalva/jwt-go"
)

// Authentication
func JWTVerifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenHeader := r.Header.Get(constant.API_KEY_HEADER)
		if tokenHeader == "" {
			web.MonkError(w, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		tokenParse, err := jwt.Parse(tokenHeader, func(jwtToken *jwt.Token) (interface{}, error) {
			if jwtToken.Method != jwt.SigningMethodHS256 {
				return nil, constant.ErrUnauthorized
			}
			return []byte(config.Cfg().JwtSecretKey), nil
		})

		if err != nil || !tokenParse.Valid {
			web.MonkError(w, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		claims := tokenParse.Claims.(jwt.MapClaims)
		claimsID, err := strconv.ParseInt(fmt.Sprint(claims["id"]), 10, 64)
		if err != nil {
			web.MonkError(w, http.StatusUnauthorized, constant.ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), claimsIDKey, claimsID)))
	})
}

// Authorization
type key string

const claimsIDKey = key("id")

func GetClaimsID(ctx context.Context) (int64, bool) {
	claimsID, valid := ctx.Value(claimsIDKey).(int64)
	return claimsID, valid
}

func IsMe(ctx context.Context, id int64) bool {
	claimsID, valid := GetClaimsID(ctx)
	return valid && claimsID == id
}
