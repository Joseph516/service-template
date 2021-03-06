package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service-template/global"
	"service-template/internal/middleware"
	"service-template/internal/routers/api"
)

func NewRouters() *gin.Engine {
	r := gin.New()

	/* 中间件设置 */
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// 跨域问题
	r.Use(middleware.Cors())

	// 注册与登陆
	register := api.NewRegister()
	login := api.NewLogin()
	user := r.Group("/user")
	{
		user.POST("/register", register.Register)
		user.POST("/login", login.LoginIn)
		user.DELETE("/login", login.Exit)
	}
	
	// 文件管理
	file := r.Group("file")
	// 提供静态资源访问，当访问/static，实际上访问的是global.AppSetting.UploadSavePath文件下的文件
	file.StaticFS("static", http.Dir(global.AppSetting.UploadSavePath))
	// JWT用户权限
	file.Use(middleware.JWT())
	{
		upload := api.NewUpload()
		file.POST("/upload", upload.UploadFile)
		file.GET("/upload", upload.List)
		file.GET("/upload:filename", upload.GetFile)
	}

	return r
}
