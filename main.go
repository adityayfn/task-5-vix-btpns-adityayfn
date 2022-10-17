package main

import (
	"github.com/adityayfn/task-5-vix-btpns-adityayfn/router"
)



func main (){
	r := router.InitRoutes()
  r.Run()
}