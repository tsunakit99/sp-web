package middleware

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func SupabaseAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    tokenStr := strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")

    token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
      return []byte(os.Getenv("SUPABASE_JWT_SECRET")), nil
    })

    if err != nil || !token.Valid {
      return echo.ErrUnauthorized
    }

    claims := token.Claims.(jwt.MapClaims)
    c.Set("userID", claims["sub"].(string))

    return next(c)
  }
}
