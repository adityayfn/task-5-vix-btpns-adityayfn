package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	""
)

func main (){
	r := router.InitRoutes()
  r.Run()
}