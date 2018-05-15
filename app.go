package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DeanThompson/ginpprof"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
)

var now = time.Now()
var election = time.Date(2018, time.May, 25, 16, 0, 0, 0, time.UTC)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	//router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:    []string{"Authorization", "Access-Control-Allow-Origin", "Content-Type, Accept"},
	}))

	router.GET("/ping", func(ginContext *gin.Context) {
		ginContext.String(200, "pong")
	})

	memGroup := router.Group("/user/")
	{
		memGroup.GET("/list", userList)
	}

	exitGroup := router.Group("/exit/")
	{
		exitGroup.GET("/date", exitDate)
	}

	HTTP := &http.Server{
		Addr:           ":8046",
		Handler:        router,
		ReadTimeout:    100000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ginpprof.Wrap(router)

	err := HTTP.ListenAndServe()
	//bugsnag.Notify(err)
	if err != nil {
		log.Fatal(err)
	}
}

func userList(g *gin.Context) {

	userList := []string{"Dom", "Matias", "Elan", "Mary", "Jacques"}
	g.JSON(200, userList)
}

func exitDate(g *gin.Context) {
	//get duration between election date and now
	tillElection := election.Sub(now)
	//get duration in nanoseconds
	toNanoseconds := tillElection.Nanoseconds()
	//calculate hours from toNanoseconds
	hours := toNanoseconds / 3600000000000
	remainder := toNanoseconds % 3600000000000
	//derive minutes from remainder of hours
	minutes := remainder / 60000000000
	remainder = remainder % 60000000000
	//derive seconds from remainder of minutes
	seconds := remainder / 1000000000
	//calculate days and get hours left from remainder
	days := hours / 24
	hoursLeft := hours % 24

	g.JSON(200, fmt.Sprintf("%v Days %v Hours %v Minutes %v Seconds", days, hoursLeft, minutes, seconds))

}
