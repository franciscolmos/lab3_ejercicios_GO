package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


	func Sumar(c *gin.Context)  {
	num1, _ := strconv.ParseInt(c.Param("num1"), 10, 64)
	num2, _ := strconv.ParseInt(c.Param("num2"), 10, 64)
	result := num1 + num2
		c.JSON(http.StatusOK, gin.H{"Result": result})
	}

	func Restar(c *gin.Context) {
	num1, _ := strconv.ParseInt(c.Param("num1"), 10, 64)
	num2, _ := strconv.ParseInt(c.Param("num2"), 10, 64)
	result := num1 - num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
	}

	func Multiplicar(c *gin.Context) {
	num1, _ := strconv.ParseInt(c.Param("num1"), 10, 64)
	num2, _ := strconv.ParseInt(c.Param("num2"), 10, 64)
	result := num1 * num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
	}

	func Dividir(c *gin.Context) {
	num1, _ := strconv.ParseFloat(c.Param("num1"),64)
	num2, _ := strconv.ParseFloat(c.Param("num2"),64)
	result  := num1 / num2
	c.JSON(http.StatusOK, gin.H{"Result": result})
	}

