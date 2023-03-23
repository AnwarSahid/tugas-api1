package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"book_id"`
	Name   string `json:"name"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var bookDatas = []Book{}

func GetAllBook(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookDatas,
	})

}
func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON((&newBook)); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("%d", len(bookDatas)+1)
	bookDatas = append(bookDatas, newBook)
	ctx.JSON(http.StatusCreated, gin.H{
		"data": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range bookDatas {
		if bookID == book.BookID {
			condition = true
			bookDatas[i] = updatedBook
			bookDatas[i].BookID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":   "data not found",
			"error_messsage": fmt.Sprintf("book with id not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been succesfully update", bookID),
	})

}

func GetBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookData Book

	for i, book := range bookDatas {
		if bookID == book.BookID {
			condition = true
			bookData = bookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":        "Data Not Found",
			"eror_message": fmt.Sprintf("book with id %v is not found", bookID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})

}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookId")
	condition := false
	var bookIndex int

	for i, book := range bookDatas {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":        "Data Not Found",
			"eror_message": fmt.Sprintf("book with id %v is not found", bookID),
		})
		return
	}

	copy(bookDatas[bookIndex:], bookDatas[bookIndex+1:])
	bookDatas[len(bookDatas)-1] = Book{}
	bookDatas = bookDatas[:len(bookDatas)-1]
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book succesfully Deleted"),
	})
}
