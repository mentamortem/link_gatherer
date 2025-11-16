package collector

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Posthandler(router *gin.Context) {
	raw, err := router.GetRawData()
	format := router.GetHeader("Content-Type")
	if err != nil {
		router.Error(fmt.Errorf("bad data: %w", err))
		return
	}
	if err := ProccessData(raw, format); err != nil {
		router.Error(fmt.Errorf("processing data error: %w", err))
		return
	}
}
