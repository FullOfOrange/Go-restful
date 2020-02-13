package v1

import (
	"net/http"
	"fmt"

	"github.com/FullOfOrange/devlog-server/models"
	"github.com/gomarkdown/markdown"
	"github.com/gin-gonic/gin"
)

type AddPost struct {
	Title string `json:"title" binding:"required"`
	Desc  string `json:"desc" binding:"required"`
	Body  string `json:"body" binding:"required"`
}

// CreatePost create post
func CreatePost(c *gin.Context) {
	var post AddPost
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b := []byte(post.Body)
	out := markdown.ToHTML(b, nil, nil);

	fmt.Println(out);
	c.JSON(http.StatusOK, gin.H{"status": "ok", "body": string(out)})
}

// GetPostByID parse post by object id
func GetPostByID(c *gin.Context) {
	id := c.Param("id")

	post, err := models.FindPostByObjectID(id)
	if err != nil {
		c.String(http.StatusInternalServerError, "database parse errror")
	}

	c.JSON(http.StatusOK, post)
}

// GetPosts parse all blog posts
func GetPosts(c *gin.Context) {
	posts, err := models.FindAllPost()

	if err != nil {
		c.String(http.StatusInternalServerError, "database parse errror")
	}

	c.JSON(http.StatusOK, posts)
}
