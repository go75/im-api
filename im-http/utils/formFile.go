package utils

import (
	"io"

	"github.com/gin-gonic/gin"
)

func ReadFormFile(c *gin.Context, name string) ([]byte, error) {
	file, _, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	file.Close()
	return io.ReadAll(file)
}