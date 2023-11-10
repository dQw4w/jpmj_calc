package main

import (
	"net/http"

	"github.com/dQw4w/jpmj_calc/controllers/calc"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*")

	calcController := calc.NewCalculateController()
	r.POST("/calc", calcController.CalculateResults)
	r.GET("/calc", func(c *gin.Context) {
		c.HTML(http.StatusOK, "page.html", nil)
	})
	r.Run(":8082")

	// handstr := "12340678mps11122z"
	// furolist := []string{ /*"333z", "444z"*/ } //TODO: add furo menzis here (including 暗槓, XnnXs is the format for it, n:rank,s:suit)
	// tempwin := win.CreateEmptyCommon()
	// // calc.Calculate(handstr, furolist, tempwin)
}
