package response

import (
	"errors"
)

var errorStatusMap = map[error]int{
	errs.ErrAnimalNotFound: http.StatusNotFound,
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
