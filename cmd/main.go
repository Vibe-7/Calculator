package main

import (
	calculationserver "Calculator/interal/calculationServer"
	"Calculator/interal/db"
	"Calculator/interal/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	databse, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	caclRepo := calculationserver.NewCalculationRepository(databse)
	caclService := calculationserver.NewCalculationService(caclRepo)
	caclHanlers := handlers.NewCalculationHandler(caclService)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello from Go backend!</h1>")
	})
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.PATCH("/calculations/:id", caclHanlers.PatchCalcuator)
	e.GET("/calculations", caclHanlers.GetCalculation)
	e.POST("/calculations", caclHanlers.PostCalculation)
	e.DELETE("/calculations/:id", caclHanlers.DeleteCalculator)
	e.Start(":8080")
}
