package collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Posthandler(c *gin.Context) {
	raw, err := c.GetRawData()
	if err != nil {
		c.Error(fmt.Errorf("bad data: %w", err))
		return
	}
	ProccessData(raw)
}
