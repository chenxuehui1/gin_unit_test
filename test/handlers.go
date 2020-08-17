package test

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

type User struct {
	UserName string `json:"user_name" form:"user_name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age" binding:"required"`
}

func GetAgeHandler(c *gin.Context) {
	type UserNameAndPassword struct {
		UserName string `json:"user_name" form:"user_name" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	req := &UserNameAndPassword{}
	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": err.Error(),
		})
		return
	}

	// judge the password and username
	if req.UserName != "Valiben" || req.Password != "123456" {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": "password or username is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "",
		"age":    18,
	})
}

func GetPasswordHandler(c *gin.Context) {
	type UserName struct {
		UserName string `json:"user_name" form:"user_name" binding:"required"`
	}
	req := &UserName{}
	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": err.Error(),
		})
		return
	}

	// judge the password and username
	if req.UserName != "Valiben" {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": "username not exists",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "",
		"data":   "123456",
	})
}

func LoginHandler(c *gin.Context) {
	req := &User{}
	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": err.Error(),
		})
		return
	}

	// judge the password and username
	if req.UserName != "Valiben" || req.Password != "123456" {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": "password or username is wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "login success",
	})
}

func DeleteUserHandler(c *gin.Context) {
	req := &User{}
	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": fmt.Sprintf("delete user:%+v", req),
	})
}

func AddUserHandler(c *gin.Context) {
	req := &User{}
	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": fmt.Sprintf("add user:%+v", req),
	})
}

type FileRequest struct {
	FileName   string `json:"file_name" form:"file_name" binding:"required"`
	UploadName string `json:"upload_name" form:"upload_name" binding:"required"`
}

func SaveFileHandler(c *gin.Context) {

	req := &FileRequest{}

	if err := c.Bind(req); err != nil {
		log.Printf("err:%v", err)
		c.JSON(http.StatusOK, gin.H{
			"errno":  "1",
			"errmsg": err.Error(),
		})
		return
	}

	// get the file of the request
	file, _, _ := c.Request.FormFile("file")
	if file == nil {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": "file is nil",
		})
		return
	}

	out, err := os.Create("test2.txt")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": err.Error(),
		})
		return
	}

	// copy the content of the file to the out
	_, err = io.Copy(out, file)
	defer file.Close()
	defer out.Close()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"errno":  "2",
			"errmsg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"errno":  "0",
		"errmsg": "save file success",
	})
}
