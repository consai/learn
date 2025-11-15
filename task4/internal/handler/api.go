package handler

import (
	"blog/internal/model"
	"blog/pkg/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiGroup(r *gin.RouterGroup) {
	r.GET("post/:postid", GetPost)
	r.GET("post", GetPost)
	r.POST("post", AddPost)
	r.PUT("post", UpdatePost)
	r.DELETE("post", DeletePost)

	r.POST("comment", AddComment)
	r.GET("comment/:postid", GetComments)
}

func GetPost(c *gin.Context) {

	userid := c.MustGet("userID").(uint)
	PostID, err := strconv.ParseUint(c.Param("postid"), 10, 32)
	if err != nil {
		PostID = 0
	}
	posts, err := database.GetPosts(c, model.Post{Model: gorm.Model{ID: uint(PostID)}, UserID: uint(userid)})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"posts": posts})
}

type jPost struct {
	PostID  uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func AddPost(c *gin.Context) {
	userid := c.MustGet("userID").(uint)
	post := jPost{}
	if err := c.BindJSON(&post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err := database.AddPost(c, &model.Post{Title: post.Title, Content: post.Content, UserID: userid})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func UpdatePost(c *gin.Context) {
	userid := c.MustGet("userID").(uint)
	post := jPost{}
	if err := c.BindJSON(&post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err := database.UpdatePost(c, &model.Post{Model: gorm.Model{ID: post.PostID}, Title: post.Title, Content: post.Content, UserID: userid})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func DeletePost(c *gin.Context) {
	userid := c.MustGet("userID").(uint)
	post := jPost{}
	if err := c.BindJSON(&post); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err := database.DeletePost(c, &model.Post{Model: gorm.Model{ID: post.PostID}, UserID: userid})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func AddComment(c *gin.Context) {
	userid := c.MustGet("userID").(uint)
	comment := model.Comment{}
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err := database.AddComment(c, &model.Comment{Content: comment.Content, PostID: comment.PostID, UserID: userid})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}

func GetComments(c *gin.Context) {
	PostID, err := strconv.ParseUint(c.Param("postid"), 10, 32)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	comments, err := database.GetComments(c, uint(PostID))
	if err != nil {
		c.JSON(501, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"comments": comments})
}
