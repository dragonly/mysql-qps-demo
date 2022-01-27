package router

import (
	"database/sql"
	"net/http"

	"qps_demo/dao"

	"github.com/gin-gonic/gin"
)

func registerAPIs(r *gin.RouterGroup, db *sql.DB) {
	r.GET("/db_stats", func(c *gin.Context) {
		type Res struct {
			CountStar    int
			SumTimerWait int
		}
		var res Res
		row := dao.DB.QueryRow("select COUNT_STAR, SUM_TIMER_WAIT from events_statements_summary_by_digest where DIGEST_TEXT like '%select%sbtest%'")
		row.Scan(&res.CountStar, &res.SumTimerWait)

		c.JSON(http.StatusOK, gin.H{
			"message": res,
		})
	})
}
