package main

import (
	"jpmj_calc/win"

	"github.com/gin-gonic/gin"
)

func main() {

	c := gin.Default()

	c.Static("/static", "./static")

	handstr := "12340678mps11122z"
	furolist := []string{ /*"333z", "444z"*/ } //TODO: add furo menzis here (including 暗槓, XnnXs is the format for it, n:rank,s:suit)
	tempwin := win.CreateEmptyCommon()
	calc(handstr, furolist, tempwin)
}
