package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()

  // 認証付きエンドポイント（例）
  e.GET("/me", func(c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]string{
      "message": "Hello, authenticated user!",
    })
  })

  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }

  e.Logger.Fatal(e.Start(":" + port))
}
