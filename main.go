package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Сервер запущен")

	a := echo.New()
	a.Use(CheckUserRole)
	a.GET("/status", getStatus)

	err := a.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getStatus(ctx echo.Context) error {
	targetDate := time.Date(2025, time.May, 27, 0, 0, 0, 0, time.UTC)
	daysUntil := int(time.Until(targetDate).Hours() / 24)
	return ctx.String(http.StatusOK, fmt.Sprintf("Количество дней: %d", daysUntil))
}

func CheckUserRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userRole := ctx.Request().Header.Get("User-role")
		if strings.EqualFold(userRole, "admin") {
			log.Println("Красная кнопка пользователя обнаружена")
		}
		return next(ctx)
	}
}
