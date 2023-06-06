package handler

import (
	"context"
	"net/http"
	"os"

	"im-api/im-http/common/dto"
	"im-api/im-http/common/e"
	"im-api/im-http/common/res"
	"im-api/im-http/global"
	"im-api/im-http/log"
	"im-api/im-http/pb"
	"im-api/im-http/utils"

	"github.com/gin-gonic/gin"
)

// @Summary 用户头像
// @Description 获取用户头像
// @Tags 用户模块
// @Param filename path string true "用户名.png"
// @Accept application/x-www-form-urlencoded
// @Produce image/png
// @Success 200 {string} UserHead
// @Router /head/{filename} [get]
func UserHead(c *gin.Context) {
	filepath := "./head/user/" + c.Param("filename")
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Info.Println("read file " + filepath + " err: " + err.Error())
		res.Err(c, "read file err")
		return
	}

	c.Writer.Header().Add("Content-Type", "image/png")
	c.Status(http.StatusOK)
	c.Writer.Write(data)
}

// @Summary 用户登录
// @Description 用户登录服务
// @Tags 用户模块
// @Param name formData string true "用户名"
// @Param pwd formData string true "用户密码"
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Success 200 {string} UserLogin
// @Router /login [post]
func UserLogin(c *gin.Context) {

	resp, err := global.RPC.UserLogin(context.Background(), &pb.UserLoginInfo{
		Name: c.PostForm("name"),
		Pwd:  c.PostForm("pwd"),
	})

	if err != nil {
		log.Info.Println("服务器连接出错")
		res.Err(c, e.DialServerErr.Error())
		return
	}

	if resp == nil {
		log.Info.Println("请求参数错误")
		res.Err(c, e.BadRequestErr.Error())
		return
	}

	res.OkWithData(c, "登录成功", dto.LoginData{
		ID: resp.ID,
		Name: resp.Name,
		Token: resp.Token,
	})
}

// @Summary 用户注册
// @Description 用户注册服务
// @Tags 用户模块
// @Param name formData string true "用户名"
// @Param pwd formData string true "用户密码"
// @Param check formData string true "确认密码"
// @Param head formData file true "用户头像"
// @Accept application/x-www-form-urlencoded
// @Produce application/json
// @Success 200 {string} UserLogin
// @Router /regist [post]
func UserRegist(c *gin.Context) {
	pwd := c.PostForm("pwd")
	if pwd != c.PostForm("check") {
		log.Warn.Println("两次密码不一致")
		res.Err(c, "两次密码不一致")
		return
	}

	name := c.PostForm("name")

	resp, err := global.RPC.UserRegist(context.Background(), &pb.UserRegistInfo{
		Name: name,
		Pwd: pwd,
	})

	if err != nil {
		log.Info.Println("服务器连接出错:", err)
		res.Err(c, e.DialServerErr.Error())
		return
	}

	if !resp.Ok {
		// regist err
		log.Info.Println("注册失败:", err)
		res.Err(c, string(resp.Msg))
		return
	}

	// regist succ
	// 存储用户头像, 若出现错误,则用户使用默认头像
	content, err := utils.ReadFormFile(c, "head")
	if err != nil {
		log.Info.Println("头像读取失败: " + err.Error())
		res.Ok(c, "注册成功,头像读取失败,请登录后切换")
		return
	}

	filepath := "./head/user/" + name + ".png"

	err = os.WriteFile(filepath, content, 0664)
	if err !=  nil {
		log.Info.Println("头像填充失败: " + err.Error())
		res.Ok(c, "注册成功,头像内容填充失败,请登录后切换")
		return
	}
	res.Ok(c, "注册成功")
}
