package main

import "github.com/gin-gonic/gin"

var mux *gin.Engine

type ChildNode struct {
	Id       string `json:"id"`
	Flashing bool
}

var Children []ChildNode

func init() {
	mux = gin.Default()
	Children = make([]ChildNode, 0)
}

func main() {
	mux.GET("/", Index)
	mux.POST("/add", AddChildNode)
	mux.GET("/flashing/:id", Flashing)
	mux.GET("/flashing_off/:id", FlashingOff)
	mux.Run(":8080")
}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{"Nodes": Children})
}

func AddChildNode(c *gin.Context) {
	var childNode ChildNode
	c.BindJSON(&childNode)
	Children = append(Children, childNode)
	c.JSON(200, gin.H{"status": "ok"})
}

func Flashing(c *gin.Context) {
	id := c.Params.ByName("id")
	for i, _ := range Children {
		if Children[i].Id == id {
			Children[i].Flashing = true
		}
	}
	c.JSON(200, gin.H{"status": "ok"})
}

func FlashingOff(c *gin.Context) {
	id := c.Params.ByName("id")
	for i, _ := range Children {
		if Children[i].Id == id {
			Children[i].Flashing = false
		}
	}
	c.JSON(200, gin.H{"status": "ok"})
}
