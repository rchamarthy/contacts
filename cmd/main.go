package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/rchamarthy/contacts"
)

func main() {
	server := NewContactServer("./contacts")
	server.Run()
}

// ContactServer is a simple server that serves contact information
type ContactServer struct {
	r           *gin.Engine
	contactsDir string
}

func NewContactServer(contactsDir string) *ContactServer {
	r := gin.Default()
	cs := &ContactServer{contactsDir: contactsDir}
	r.GET("/contacts/:id", cs.getContact)
	r.POST("/contacts", cs.createContact)
	r.GET("/ping", cs.ping)
	return &ContactServer{r: r, contactsDir: contactsDir}
}

func (cs *ContactServer) Run() {
	e := cs.r.Run(":9999")
	fmt.Println(e)
}

func (cs *ContactServer) getContact(c *gin.Context) {
	id := c.Param("id")
	personFile := path.Join(cs.contactsDir, id+".yaml")
	contact, err := contacts.LoadPerson(personFile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contact)
}

func (cs *ContactServer) ping(c *gin.Context) {
	c.JSON(200, `{"status": "ok"}`)
}

func (cs *ContactServer) createContact(c *gin.Context) {
	p := contacts.EmptyPerson()
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	personFile := path.Join(cs.contactsDir, p.Email+".yaml")
	if err := p.Save(personFile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "created"})
}
