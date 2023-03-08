package service

import (
	"GinChat/models"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "index")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "welcome!!!",
	//})
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	ind.Execute(c.Writer, "register")
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "welcome!!!",
	//})
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}

	user := models.UserBasic{}
	userId, _ := strconv.Atoi(c.Query("userId"))
	user.Identity = c.Query("token")
	user.ID = uint(userId)
	ind.Execute(c.Writer, user)
	//c.JSON(http.StatusOK, gin.H{
	//	"message": "welcome!!!",
	//})
}

func Chat(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
