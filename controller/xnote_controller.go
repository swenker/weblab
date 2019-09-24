package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"yinyushijing.cn/xnote-api/helper"
	noteservice "yinyushijing.cn/xnote-api/service"
)

// SetupRouter initializes the router with api handlers
func SetupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api.GET("/get/:xid", func(c *gin.Context) {
		xid, err := strconv.Atoi(c.Param("xid"))
		if err != nil {
			panic(err.Error())
		}
		xnoteobj := noteservice.GetNoteByID(xid)
		c.JSON(http.StatusOK, xnoteobj)
	})

	api.POST("/save", func(c *gin.Context) {
		var xnoteForm helper.XNoteForm
		err := c.BindJSON(&xnoteForm)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if *xnoteForm.ID > 0 {
			fmt.Printf("------------has id:%d \n", xnoteForm.ID)
			noteservice.UpdateXNote(helper.Form2Domain(&xnoteForm))
		} else {
			fmt.Println("id is nil")
			noteservice.CreateXNote(helper.Form2Domain(&xnoteForm))
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	api.POST("/rm", func(c *gin.Context) {
		xid, _ := strconv.Atoi(c.PostForm("xid"))
		log.Printf("id for rm ............%d\n", xid)
		noteservice.RemoveXNoteByID(xid)
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})

	api.GET("/search/*kw", func(c *gin.Context) {
		keywords := c.Param("kw") //it contains '/'
		keywords = keywords[1:]
		notes := noteservice.QueryNotesByTitle(keywords)
		c.JSON(http.StatusOK, notes)
	})

	return router
}
