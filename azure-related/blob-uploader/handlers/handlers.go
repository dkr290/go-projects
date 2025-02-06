package handlers

import (
	"blob-uploader/pkg"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	r         *gin.Engine
	storage   string
	container string
}

type PageData struct {
	PageDataArray []string
	BlobStorage   string
	Containername string
}

func NewHandlers(r *gin.Engine, storage, container string) *Handlers {
	return &Handlers{
		r:         r,
		storage:   storage,
		container: container,
	}
}

func (h *Handlers) UploadHandler(c *gin.Context) {

	// // Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}
	files := form.File["files"]

	for _, file := range files {
		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		if err := pkg.UploadFileBlob(h.storage, h.container, file.Filename); err != nil {
			c.String(http.StatusBadRequest, "upload file to blob err: %s", err.Error())
			return
		}
		if err := pkg.DeleteLocalFiles(file.Filename); err != nil {
			c.String(http.StatusBadRequest, "delete local file err: %s", err.Error())
			return
		}

	}
	c.HTML(http.StatusOK, "upload-success.html", gin.H{
		"message":    fmt.Sprintf("Uploaded successfully %d files", len(files)),
		"redirectTo": "/",
	})

}

func (h *Handlers) GetHandler(c *gin.Context) {

	// Create a slice to store secrets
	blobs, err := pkg.GetBlobFiles(h.storage, h.container)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create a PageData struct
	pageData := PageData{
		PageDataArray: blobs,
		BlobStorage:   h.storage,
		Containername: h.container,
	}

	// Render the HTML page with the PageData struct
	c.HTML(http.StatusOK, "index.html", gin.H{
		"PageData": pageData,
	})

}
