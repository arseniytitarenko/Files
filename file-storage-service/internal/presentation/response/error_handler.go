package response

import (
	"errors"
	"files/internal/application/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

var errorStatusMap = map[error]int{
	errs.ErrFileNotFound:  http.StatusNotFound,
	errs.FailedToReadFile: http.StatusBadRequest,
}

func HandleError(c *gin.Context, err error) {
	for e, code := range errorStatusMap {
		if errors.Is(err, e) {
			c.JSON(code, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
