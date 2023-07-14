package main

import(
	//"jpmj_calc/test1"
	"jpmj_calc/combination"
	"log"
	//"strconv"
)

func main()  {
	
	a,err := combination.NewPair('x',1)
	if err != nil{
		log.Println(err)
		log.Panic(err)
	}
	log.Println(a.Rank)
}