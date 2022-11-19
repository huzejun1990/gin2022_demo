package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"log"
	"net/http"
)

// 自定义GO中间件 拦截器
func myHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 通过自定义中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
		context.Set("usersession", "userid-1")
		context.Next() //放行
		//context.Abort()	//阻止
	}
}

//需要导包

func main() {

	//创建一个服务
	ginServer := gin.Default()
	ginServer.Use(favicon.New("./favicon.ico"))

	//加载静态页面
	ginServer.LoadHTMLGlob("templates/*")
	//ginServer.LoadHTMLFiles("templates/index.html")

	//加载资源文件
	ginServer.Static("/static", "./static")

	// gin RestFul 十分简单
	ginServer.GET("/hello", func(context *gin.Context) {
		//context.JSON(200,gin.H{"msg":"hello,world"})
		context.JSON(http.StatusOK, gin.H{"msg": "hello,world"})
	})
	/*	ginServer.POST("/user", func(c *gin.Context) {
			c.JSON(200,gin.H{"msg":"post,user"})
		})
		ginServer.PUT("/user")
		ginServer.DELETE("/user")*/

	//响应一个页面给前端
	ginServer.GET("/index", func(context *gin.Context) {
		//context.JSON()	//json数据
		context.HTML(http.StatusOK, "index.html", gin.H{
			"msg": "这是go后台传递来的数据",
		})
	})

	//接收前端传递过来的参数
	// usl? userid=xxx&username=dream
	ginServer.GET("/user/info", myHandler(), func(context *gin.Context) {

		//取出中间件中的值
		usersession := context.MustGet("usersession").(string)
		log.Println("=========================", usersession)

		userid := context.Query("userid")
		username := context.Query("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})

	})

	// web - java
	// /user/info/1/dream
	//ginServer.Run("/user/info/:userid/:username",func)

	// /user/info/1/dream
	ginServer.GET("/user/info/:userid/:username", func(context *gin.Context) {
		userid := context.Param("userid")
		username := context.Param("username")
		context.JSON(http.StatusOK, gin.H{
			"userid":   userid,
			"username": username,
		})
	})

	// 掌握技术后面的应用 - 掌握基础知识，加以了解web开发
	//前端给后端传递json
	ginServer.POST("/json", func(context *gin.Context) {
		// request.body
		data, _ := context.GetRawData()

		var m map[string]interface{}
		// 包装为json数据 []byte
		_ = json.Unmarshal(data, &m)
		context.JSON(http.StatusOK, m)
	})

	// 支持函数式编程 =>
	ginServer.POST("/user/add", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.JSON(http.StatusOK, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})

	})

	//路由 301
	ginServer.GET("/test", func(context *gin.Context) {
		// 重定向
		// 	StatusMovedPermanently  = 301
		context.Redirect(http.StatusMovedPermanently, "https://github.com/huzejun1990")
	})

	// 404 NoRoute
	ginServer.NoRoute(func(context *gin.Context) {
		context.HTML(http.StatusNotFound, "404.html", nil)
	})

	//路由组 /user/add
	userGroup := ginServer.Group("/user")
	{
		userGroup.GET("/add")
		userGroup.POST("/login")
		userGroup.POST("/logout")
	}

	orderGroup := ginServer.Group("/order")
	{
		orderGroup.GET("/add")
		orderGroup.DELETE("/delete")
	}

	//服务器端口
	ginServer.Run(":8082")

	//连接数据库的代码

	//访问地址，处理我们的请求	Request Response

}
