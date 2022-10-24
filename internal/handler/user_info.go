package handler

type UserInfo struct{}

//// GetMyInfo get UserInfo based on the id in Cookie
//func (*UserInfo) GetMyInfo(c *gin.Context) {
//	idStr, err := cookie.GetUserId(c)
//	if err != nil {
//		c.JSON(http.StatusOK, dto.GetUserInfoResp{
//			CommonResp: dto.CommonResp{
//				Code: eh.Preset.CodeTokenError,
//				Msg:  "you're not logged in",
//			},
//		})
//		return
//	}
//	id := util.ParseU64(idStr)
//	var userInfo *dal.UserInfo
//	userInfo, err = userInfo.FindById(id)
//	errHandler := eh.JSONHandler{C: c, V: dto.GetUserInfoResp{}}
//	if err != nil {
//		errHandler.Handle(err)
//		return
//	}
//	c.JSON(http.StatusOK, dto.GetUserInfoResp{
//		CommonResp: dto.CommonResp{
//			Code: eh.Preset.CodeOK,
//			Msg:  "success",
//		},
//		UserInfo: userInfo,
//	})
//}
