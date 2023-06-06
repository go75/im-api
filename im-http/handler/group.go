package handler

import (
	"context"
	"im-api/im-http/common/e"
	"im-api/im-http/common/res"
	"im-api/im-http/global"
	"im-api/im-http/log"
	"im-api/im-http/pb"
	"im-api/im-http/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary 群聊头像
// @Description 获取群聊头像
// @Tags 群聊模块
// @Param filename path string true "群聊名.png"
// @Accept application/x-www-form-urlencoded
// @Produce image/png
// @Success 200 {string} GroupHead
// @Router /group/head/{filename} [get]
func GroupHead(c *gin.Context) {
	data, err := os.ReadFile("./head/group/" + c.Param("filename"))
	if err != nil {
		log.Info.Println("群聊头像读取失败:", err)
		res.Err(c, "群聊头像读取失败")
		return
	}

	c.Writer.Header().Add("Content-Type", "image/png")
	c.Status(http.StatusOK)
	c.Writer.Write(data)
}

// @Summary 群聊注册
// @Description 群聊注册服务
// @Tags 群聊模块
// @Param name formData string true "群聊名称"
// @Param introduce formData string true "群聊介绍"
// @Param head formData file true "群聊头像"
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Success 200 {string} UserLogin
// @Router /group/regist [post]
func GroupRegist(c  *gin.Context) {

	// 通过http cookie中的token解析来认证
	token, err := c.Cookie("im-token")
	if err != nil {
		c.JSON(http.StatusBadRequest, "token解析错误")
		c.Abort()
		return
	}

	// 未携带token直接返回
	if token == "" {
		c.JSON(http.StatusBadRequest, "未携带token")
		c.Abort()
		return
	}

	// 初始化一个JWT对象实例，并根据结构体方法来解析token
	// 解析token中包含的相关信息(有效载荷)
	claims, err := utils.ParseToken(token)

	if err != nil {
		// token过期
		if err == e.ExpiredTokenErr {
			c.JSON(http.StatusBadRequest, "token过期")
			c.Abort()
			return
		}

		// 其他错误
		c.JSON(http.StatusBadRequest, "身份认证失败")
		c.Abort()
		return
	}

	groupName := c.PostForm("name")
	introduce := c.PostForm("introduce")
	masterId :=  claims.ID

	resp, err := global.RPC.GroupRegist(context.Background(), &pb.GroupRegistInfo{
		MasterId: uint32(masterId),
		Name: groupName,
		Introduce: introduce,	
	})
	
	if err != nil {
		res.Err(c, e.DialServerErr.Error())
		log.Info.Println("服务器连接出错")
		return
	}

	if !resp.Ok {
		// regist err
		log.Info.Println("群聊创建失败:", err)
		c.JSON(http.StatusInternalServerError, "创建失败")
		return
	}

	var filepath string
	var content []byte

	// 存储用户头像, 若出现错误,则用户使用默认头像
	content, err = utils.ReadFormFile(c, "head")
	if err != nil {
		log.Info.Println("群聊头像读取失败:", err)
		goto END
	}

	filepath = "./head/group/" + groupName + ".png"

	err = os.WriteFile(filepath, content, 0664)
	if err != nil {
		log.Info.Println("群聊头像填充失败:", err)
		goto END
	}
	END:
	c.JSON(http.StatusOK,  resp.GroupId)
}