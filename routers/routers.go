package routers

import (
	"gin-blog/middleware/jwt"
	"gin-blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"gin-blog/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.Server.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		apiv1.GET("/articles", v1.GetArticles)         // 获得文章列表
		apiv1.GET("/articles/:id", v1.GetArticle)      // 获得指定文章
		apiv1.POST("/articles", v1.AddArticle)         // 新建文章
		apiv1.PUT("/articles/:id", v1.EditArticle)    // 更新指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)// 删除指定文章
	}

	return r
}
