package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/leetpy/rum/util"
)

func InitRouter() {
	router := gin.Default()
	router.Static("/static", "static")
	router.Static("/ironic", "sphinx/ironic")
	router.StaticFS("/book", http.Dir("books"))
	router.LoadHTMLGlob("templates/*")
	router.GET("/books/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	router.POST("/books/upload", upload)
	router.GET("/books", showBooks)
	router.Run(":8080")
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
	books, err := util.ListDir("books", "pdf")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf(err.Error()))
		return
	}

	fmt.Println(books)
	c.HTML(http.StatusOK, "books.html", gin.H{
		"books": books,
	})
}