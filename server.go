package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func server() {

	duration := time.Now()

	// Use Gin as the HTTP router
	gin.SetMode(gin.ReleaseMode)
	recover := gin.New()
	recover.Use(gin.Recovery())
	ginroute := recover

	// Custom NotFound handler
	ginroute.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, fmt.Sprintln("[404] NOT FOUND"))
	})

	// Print homepage
	ginroute.GET("/", func(c *gin.Context) {
		runtime.ReadMemStats(&mem)
		NumGCMem = fmt.Sprintf("%v", mem.NumGC)
		timeElapsed = fmt.Sprintf("%v", time.Since(duration))

		latestLog = fmt.Sprintf("\n •===========================• \n • [ SERVER STATUS ] \n • Last Modified: %v \n • Completed GC Cycles: %v \n • Time Elapsed: %v \n •===========================• \n\n", time.Now().UTC().Format(time.RFC850), NumGCMem, timeElapsed)

		c.String(http.StatusOK, fmt.Sprintf("%v", latestLog))
	})

	ginroute.POST("/register", func(c *gin.Context) {
		runtime.ReadMemStats(&mem)
		NumGCMem = fmt.Sprintf("%v", mem.NumGC)
		timeElapsed = fmt.Sprintf("%v", time.Since(duration))

		pesan := "Berhasil membuat akun baru!"

		latestLog = fmt.Sprintf("\n •===========================• \n • [ SERVER STATUS ] \n • Last Modified: %v \n • Completed GC Cycles: %v \n • Time Elapsed: %v \n •===========================• \n\n %v \n", time.Now().UTC().Format(time.RFC850), NumGCMem, timeElapsed, string(pesan))

		c.String(http.StatusOK, fmt.Sprintf("%v", latestLog))
	})

	ginroute.POST("/login", func(c *gin.Context) {
		runtime.ReadMemStats(&mem)
		NumGCMem = fmt.Sprintf("%v", mem.NumGC)
		timeElapsed = fmt.Sprintf("%v", time.Since(duration))

		pesan := "Logged in!"

		latestLog = fmt.Sprintf("\n •===========================• \n • [ SERVER STATUS ] \n • Last Modified: %v \n • Completed GC Cycles: %v \n • Time Elapsed: %v \n •===========================• \n\n %v \n", time.Now().UTC().Format(time.RFC850), NumGCMem, timeElapsed, string(pesan))

		c.String(http.StatusOK, fmt.Sprintf("%v", latestLog))
	})

	httpserverGin := &http.Server{
		Addr:              fmt.Sprintf(":%v", hostPortGin),
		Handler:           ginroute,
		TLSConfig:         tlsConf,
		MaxHeaderBytes:    64 << 10, // 64k
		ReadTimeout:       timeoutTr,
		ReadHeaderTimeout: timeoutTr,
		WriteTimeout:      timeoutTr,
		IdleTimeout:       timeoutTr,
	}
	httpserverGin.SetKeepAlivesEnabled(true)

	notifyGin := fmt.Sprintf("Server is running on %v", fmt.Sprintf(":%v", hostPortGin))

	fmt.Println()
	fmt.Println(notifyGin)
	fmt.Println()
	httpserverGin.ListenAndServe()
	// httpserverGin.ListenAndServeTLS(CertFilePath, KeyFilePath)

}
