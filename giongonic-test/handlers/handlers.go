package handlers

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

type Handlers struct{}

func NewHandler() *Handlers {
	return &Handlers{}
}

func (h *Handlers) IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello message from golang gin " + name,
	})
}

func (h *Handlers) XMLHandler(c *gin.Context) {
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		FirstName string   `xml:"firstname"`
		LastName  string   `xml:"lastname"`
	}
	c.XML(200, Person{FirstName: "Danail", LastName: "Krasimirov"})
}
