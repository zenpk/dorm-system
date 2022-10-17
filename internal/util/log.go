package util

// 旧代码勿用
//// LogAndReturnWarning 以 warning 级别记录错误日志，并将错误信息返回给前端
//func LogAndReturnWarning(c *gin.Context, err error) {
//	middleware.Logger.Warn(err.Error())                   // 日志记录错误信息
//	c.String(http.StatusInternalServerError, err.Error()) // 错误信息返回给前端
//}
//
//// LogAndReturnError 以 Error 级别记录错误日志，并将错误信息返回给前端
//func LogAndReturnError(c *gin.Context, err error) {
//	middleware.Logger.Error(err.Error())
//	c.String(http.StatusInternalServerError, err.Error())
//}
