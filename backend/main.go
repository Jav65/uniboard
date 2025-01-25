package main

import (
	"backend/comment"
	"backend/handler"
	"backend/tag"
	"backend/thread"
	"backend/threadtags"
	"backend/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//Database
	//dbURL := "postgresql://postgres:postgres@localhost:5432/uniboard"
	dbURL := "postgresql://uniboard_user:9UpSnmtA4EJ5tLNqubnU7JdbbfaGbUcg@dpg-cua70jdsvqrc73dmj3p0-a.oregon-postgres.render.com/uniboard"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	fmt.Println("connected to database")

	db.AutoMigrate(&thread.Thread{})
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&tag.Tag{})
	db.AutoMigrate(&threadtags.ThreadTags{})
	db.AutoMigrate(&comment.Comment{})

	r := gin.Default()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"https://uniboard-lemon.vercel.app", "http://localhost:5173"})
	credentials := handlers.AllowCredentials()

	api := r.Group("/api")
	//Tag
	tagRepository := tag.NewRepository(db)
	tagService := tag.NewService(tagRepository)
	tagHandler := handler.NewTagHandler(tagService)

	api.GET("/tags/name", tagHandler.GetAllNameTagsHandler)
	api.GET("/tags", tagHandler.GetAllTagsHandler)
	api.POST("/tags", tagHandler.CreateTagHandler)

	//ThreadTagJoints
	threadTagRepository := threadtags.NewRepository(db)
	threadtagsService := threadtags.NewService(threadTagRepository, tagService)
	_ = handler.NewThreadTagsHandler(threadtagsService)

	//User
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/user", userHandler.User)
	api.POST("/logout", userHandler.Logout)

	//Thread
	threadRepository := thread.NewRepository(db)
	threadService := thread.NewService(threadRepository, threadtagsService, userService)
	threadHandler := handler.NewThreadHandler(threadService)

	api.GET("/thread/:id", threadHandler.GetThreadByIDHandler)
	api.GET("/threads", threadHandler.GetSortedThreadsHandler)
	api.POST("/thread", threadHandler.CreateThreadHandler)
	api.DELETE("/thread/:id", threadHandler.DeleteThreadHandler)

	//Comment
	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository, userService)
	commentHandler := handler.NewCommentHandler(commentService)

	api.GET("/comments/:id", commentHandler.GetCommentByThreadIDHandler)
	api.POST("/comment/:id", commentHandler.CreateCommentHandler)

	http.ListenAndServe(":8080", handlers.CORS(origins, headers, methods, credentials)(r))
}
