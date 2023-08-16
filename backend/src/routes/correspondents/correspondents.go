package correspondents

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (r *CorrespondentRouter) GetCorrespondents(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", r.conf.FRONTEND_URL)
	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, authorization")
	c.Header("Access-Control-Allow-Methods", "GET")

	docs, err := r.dm.GetCorrespondents()
	if err != nil {
		log.Printf("Error occured: %s", err.Error())
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, docs)
}
