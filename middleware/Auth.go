package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		errEnv := godotenv.Load()

		if errEnv != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Can't Load Env"})
			return
		}

		Secret := os.Getenv("SECRET")
		author := ctx.Request.Header.Get("Authorization")
		tokenString := strings.Replace(author, "Bearer ", "", -1)

		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": http.StatusText(http.StatusUnauthorized)})
			return
		}

		token, errTok := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != t.Method {
				return nil, errors.New("status forbidden")
			}

			return []byte(Secret), nil
		})

		if errTok != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Status": errTok.Error()})
			return
		}

		Claims, ok := token.Claims.(jwt.MapClaims)

		if !token.Valid || !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Status": http.StatusText(http.StatusForbidden)})
			return
		}

		exp, _ := Claims["exp"].(float64)

		if time.Now().After(time.Unix(int64(exp), 0)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Status": "Token Is Expired"})
			return
		}

		role, _ := Claims["Role"].(string)

		if role != "keuangan" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": http.StatusText(http.StatusUnauthorized)})
			return
		}

		id, _ := Claims["Uid"].(float64)

		ctx.Set("id", int(id))
		ctx.Next()
	}
}

func AuthGetByNim() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		errEnv := godotenv.Load()

		if errEnv != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Can't Load Env"})
			return
		}

		Secret := os.Getenv("SECRET")
		author := ctx.Request.Header.Get("Authorization")
		tokenString := strings.Replace(author, "Bearer ", "", -1)

		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": http.StatusText(http.StatusUnauthorized)})
			return
		}

		token, errTok := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if jwt.GetSigningMethod("HS256") != t.Method {
				return nil, errors.New("status forbidden")
			}

			return []byte(Secret), nil
		})

		if errTok != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Status": errTok.Error()})
			return
		}

		Claims, ok := token.Claims.(jwt.MapClaims)

		if !token.Valid || !ok {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Status": http.StatusText(http.StatusForbidden)})
			return
		}

		exp, _ := Claims["exp"].(float64)

		if time.Now().After(time.Unix(int64(exp), 0)) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"Status": "Token Is Expired"})
			return
		}

		role, _ := Claims["Role"].(string)

		if role != "keuangan" {
			if role != "mahasiswa" {
				if role != "admin" {
					ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Status": http.StatusText(http.StatusUnauthorized)})
					return
				}
			}
		}

		id, _ := Claims["Uid"].(float64)

		ctx.Set("id", int(id))
		ctx.Next()
	}
}
