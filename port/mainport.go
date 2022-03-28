package port

import (
	// "fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

func Quest2(con *gin.Context) {

	widescanresults := WideScan("localhost")

	con.JSON(http.StatusOK, widescanresults)
}
func Quest3(con *gin.Context) {
	scaner := WideScan1("localhost")
	con.JSON(http.StatusOK, scaner)
}
