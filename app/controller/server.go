package controller

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"time"
)

// Server 本地文件服务
type Server struct {
	Base
	path string
	port int
}

func NewServer() *Server {
	return &Server{}
}

// get请求
func (s *Server) get(c *gin.Context) {
	//假设token失效
	s.log(fmt.Sprintf("Token : %s", c.Request.Header.Get("token")))
	//c.JSON(http.StatusOK, s.res(NoLoginCode, "请先登录", nil))
	c.JSON(http.StatusOK, s.success("这个是本地自定义的get请求，你也可以调用第三方API"))
}

// post请求
func (s *Server) post(c *gin.Context) {
	c.JSON(http.StatusOK, s.success("这个是本地自定义的post请求，你也可以调用第三方API"))
}

// file 本地文件转为http流
func (s *Server) file(c *gin.Context) {
	c.File(c.Query("filename"))
}

// upload
func (s *Server) upload(c *gin.Context) {
	// 单文件
	file, _ := c.FormFile("file")
	saveName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	dst := s.pathConvert(fmt.Sprintf("%s\\%s", s.path, saveName))
	_ = c.SaveUploadedFile(file, dst)
	c.JSON(200, s.success(gin.H{
		"name": file.Filename,
		"dst":  dst,
		"ext":  path.Ext(file.Filename),
		"path": saveName,
		"url":  fmt.Sprintf("http://localhost:%d/preview/%s", s.port, saveName),
		"size": file.Size,
	}))
}

// start 启动时加载
func (s *Server) start(port int) *Server {
	s.path = s.pathExist(s.pathConvert(fmt.Sprintf("%s\\files", s.getAppPath())))
	s.port = port
	go func() {
		r := gin.Default()
		r.Use(cors.New(cors.Config{
			AllowAllOrigins:  true,
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"*"},
			ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour},
		))

		//API路由
		r.GET("/get", s.get)
		r.POST("/post", s.post)
		r.Any("/file", s.file)
		r.Static("/preview", s.path)
		r.POST("/upload", s.upload)

		_ = r.Run(fmt.Sprintf(":%d", s.port))
	}()
	return s
}
