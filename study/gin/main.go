package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 默认引擎
	c := gin.Default() //自带logger和recovery中间件
	//c = gin.New()      //不使用默认中间件

	//RESTful接口风格 GET表示获取资源 POST表示新增资源 PUT表示更新资源 DELETE表示删除资源
	c.GET("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "get",
		})
	})
	c.POST("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "post",
		})
	})
	c.PUT("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "put",
		})
	})
	c.DELETE("/book", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"method": "delete",
		})
	})

	//数据返回 结构体 map
	//结构体要注意大小写访问
	type Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
		Num  int    `json:"num"`
	}
	var stu = Student{Name: "张三", Age: 18, Num: 101}
	c.GET("/student", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, stu)
	})
	c.GET("/student2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"name": "张三",
			"age":  18,
			"num":  102,
		})
	})

	//URL参数获取
	c.GET("/query", func(ctx *gin.Context) {
		var name = ctx.Query("name")
		// GetQuery 尝试获取
		if name, ok := ctx.GetQuery("name"); ok {
			fmt.Printf("name: %v\n", name)
		}
		// DefaultQuery 默认值获取 取不到给默认值
		name = ctx.DefaultQuery("name", "缺少姓名")
		var age = ctx.Query("age")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	c.POST("/query", func(ctx *gin.Context) {
		//var name = ctx.PostForm("name")
		//var age = ctx.PostForm("age")
		var name string
		if s, b := ctx.GetPostForm("name"); b {
			name = s
		}
		age := ctx.DefaultPostForm("age", "0")
		ctx.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
			"note": "PostForm",
		})
	})

	//url路径参数
	c.GET("/user/:name/:age", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name":   ctx.Param("name"),
			"age":    ctx.Param("age"),
			"params": ctx.Params,
			"notes":  "Url路径参数",
		})
	})

	//数据绑定 form标签将前端传入的数据反射到变量
	type User struct {
		Name string `form:"name" json:"name"`
		Age  int    `form:"age" json:"age"`
		Note string `form:"note" json:"note"`
	}
	//表单数据
	c.POST("/user2", func(ctx *gin.Context) {
		var user User
		ctx.ShouldBind(&user)
		ctx.JSON(http.StatusOK, user)
	})

	//文件上传
	c.POST("/file", func(ctx *gin.Context) {
		if fh, err := ctx.FormFile("file"); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
		} else {
			if err := ctx.SaveUploadedFile(fh, fmt.Sprintf("C:/temp/%s", fh.Filename)); err != nil {
				ctx.JSON(http.StatusBadRequest, err)
			} else {
				ctx.JSON(http.StatusOK, "上传成功")
			}
		}
	})

	//地址重定向 登录跳转
	c.GET("/redirect", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "http://www.baidu.com")
	})

	//路由转发 a不处理，由b处理
	c.GET("/route/a", func(ctx *gin.Context) {
		ctx.Request.URL.Path = "/route/b"
		c.HandleContext(ctx)
	})
	c.GET("/route/b", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	//任意路由
	c.Any("/route", func(ctx *gin.Context) {
		switch ctx.Request.Method {
		case http.MethodGet:
			ctx.JSON(http.StatusOK, gin.H{
				"method": http.MethodGet,
			})
		case http.MethodPost:
			ctx.JSON(http.StatusOK, gin.H{
				"method": http.MethodPost,
			})
		case http.MethodPut:
			ctx.JSON(http.StatusOK, gin.H{
				"method": http.MethodPut,
			})
		case http.MethodDelete:
			ctx.JSON(http.StatusOK, gin.H{
				"method": http.MethodDelete,
			})
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"method": ctx.Request.Method,
			})
		}
	})

	//没找到路由
	c.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, "找不到页面")
	})

	//路由组 一组路由 叠加词缀
	var member = c.Group("/member")
	{
		// member.GET("/index", func(ctx *gin.Context) {
		// 	ctx.JSON(http.StatusOK, ctx.Request.URL)
		// })
		member.GET("/list", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, ctx.Request.URL)
		})
		member.Any("/index", func(ctx *gin.Context) {
			switch ctx.Request.Method {
			default:
				ctx.JSON(http.StatusOK, gin.H{
					"method": ctx.Request.Method,
					"path":   ctx.Request.URL.Path,
				})
			}
		})
	}

	//中间件
	c.GET("/middle", TimeCount, testHandle(), func(ctx *gin.Context) {
		fmt.Println("执行函数。。。")
		ctx.JSON(http.StatusOK, "这是中间件测试")
	})
	//全局使用中间件
	//c.Use(TimeCount)
	//路由组使用中间件
	//member.Use(TimeCount)

	c.Run(":8080")
}

// 中间件
func TimeCount(ctx *gin.Context) {
	fmt.Println("计时开始：")
	var start = time.Now()
	ctx.Set("name", "middle") //上下文传递
	ctx.Next()                //运行下一个处理器
	//ctx.Abort() //停止运行下一个处理器
	fmt.Println("计时结束：", time.Since(start))
}

func testHandle() gin.HandlerFunc {
	fmt.Println("testHandle执行中")
	return func(ctx *gin.Context) {
		fmt.Println("匿名测试开始")
		var name = ctx.GetString("name")
		fmt.Printf("name: %v\n", name)
		ctx.Next()
		fmt.Println("匿名测试结束")
	}
}