package response

import (
	"cake-api/core"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Message struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func IsRecordNotFoundErr(err error) bool {
	unwrapped := errors.Unwrap(err)
	return unwrapped == sql.ErrNoRows || err == sql.ErrNoRows ||
		unwrapped == core.ErrRecordNotFound || err == core.ErrRecordNotFound
}

func IsValidationErr(err error) bool {
	unwrapped := errors.Unwrap(err)
	return unwrapped == core.ErrDataValidation || err == core.ErrDataValidation
}

func Error(c *gin.Context, err error) {
	log.WithError(err).WithField("url", c.Request.URL).Debug()

	if IsRecordNotFoundErr(err) {
		c.JSON(http.StatusNotFound, Message{
			Error: "error no record found on data store",
		})
		return
	}

	if IsValidationErr(err) {
		c.JSON(http.StatusBadRequest, Message{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, Message{
		Error: "internal server error",
	})
}
