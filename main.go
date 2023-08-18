package main

import (
	"github.com/dQw4w/jpmj_calc/controllers/calc"

	"github.com/gin-gonic/gin"
)

func main() {

	c := gin.Default()

	c.Static("/static", "./static")
	c.LoadHTMLGlob("templates/*")

	calcController := calc.NewCalculateController()
	c.GET("/calc", calcController.CalculateResults)

	c.Run(":8080")

	// handstr := "12340678mps11122z"
	// furolist := []string{ /*"333z", "444z"*/ } //TODO: add furo menzis here (including 暗槓, XnnXs is the format for it, n:rank,s:suit)
	// tempwin := win.CreateEmptyCommon()
	// // calc.Calculate(handstr, furolist, tempwin)
}
