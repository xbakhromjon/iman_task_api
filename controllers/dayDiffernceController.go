package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"time"
)

func CalculateDateDifference(c *gin.Context) {
	current := time.Now()
	specificTime := time.Date(2025, 1, 1, 1, 1, 1, 1, time.Local)
	diff := specificTime.Sub(current)
	c.IndentedJSON(http.StatusOK, math.Round(diff.Hours()/24))
}
