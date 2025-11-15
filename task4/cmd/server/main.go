package main

import (
	"blog/internal/handler"
	mylog "blog/internal/log"
	_ "blog/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(mylog.GetLogger())
	handler.Router(r)
	r.Run() // listen and serve on 0.0.0.0:8080

}
