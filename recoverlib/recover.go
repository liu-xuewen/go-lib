package recoverlib

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liu-xuewen/go-lib/stacklib"
	"github.com/liu-xuewen/logger"
)

func Recover(ctx context.Context) {
	if err := recover(); err != nil {
		stack := stacklib.FullStack() // 错误堆栈信息

		reqContext := map[string]interface{}{
			"request_data": reqData,
			"trace_stack":  string(stack),
			"recover_err":  err,
		}

		logger.Error(c.Request.Context(), "exec panic error", reqContext)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "服务异常",
		})
	}
}
