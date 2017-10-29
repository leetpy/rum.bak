package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/rum/conf"
	"github.com/leetpy/rum/util"
)

func InitRouter() {
	router := gin.Default()
	router.Static("/static", "static")
	for _, v := range conf.Conf.Sphinx {
		router.Static(v, path.Join("sphinx", v))
	}

	router.StaticFS("/book", http.Dir("books"))
	router.LoadHTMLGlob("templates/*")
	router.GET("/books/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	// login
	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})

	// Authorization group
	//authorized := router.Group("/env")
	//authorized.Use()
	//{
	//	authorized.POST("/login", )
	//}

	// index
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.POST("/books/upload", upload)
	router.GET("/books", showBooks)
	router.Run(":8080")
}

//func login(c *gin.Context) {
//	var json Login
//	if err := c.ShouldBindJSON(&json); err == nil {
//
//	}
//}

type Login struct {
	Email    string `json:"user" bniding:"required"`
	Password string `json "password" binding:"required"`
}

func upload(c *gin.Context) {
	// file, err := c.FormFile("file")
	file, _ := c.FormFile("file")

	src, err := file.Open()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer src.Close()

	// Destination
	destName := path.Join("books", file.Filename)
	dst, err := os.Create(destName)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		fmt.Println(err.Error())
	}
	c.Redirect(http.StatusMovedPermanently, "/books")
	// c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

func showBooks(c *gin.Context) {
	suffix := conf.Conf.EnabledSuffix
	books, err := util.ListDir("books", suffix)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf(err.Error()))
		return
	}

	fmt.Println(books)
	c.HTML(http.StatusOK, "books.html", gin.H{
		"books": books,
	})
}
