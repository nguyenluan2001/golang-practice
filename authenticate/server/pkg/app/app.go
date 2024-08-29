package app

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/nguyenluan2001/golang-authenticate/model"
	"github.com/nguyenluan2001/golang-authenticate/pkg/utils"
	"github.com/valyala/fasthttp"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
}

func (app *App) InitRouter() *router.Router {
	r := router.New()
	r.HandleOPTIONS = true
	group := r.Group("/api")
	group.GET("/test", app.Test)
	group.POST("/sign-up", app.SignUpApi)
	group.POST("/sign-in", app.SignInApi)
	group.GET("/sign-out", app.SignOutApi)
	group.GET("/profile", app.ProfileApi)

	todoGroup := group.Group("/todo")
	todoGroup.POST("/create", app.CreateTodoApi)
	group.GET("/todos", app.GetTodosApi)
	return r
}
func (app *App) Test(ctx *fasthttp.RequestCtx) {
	fmt.Fprintln(ctx, "hello world")
}
func (app *App) SignUpApi(ctx *fasthttp.RequestCtx) {
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalln("Generate password failed")
	}
	hashPassword := string(bytes)
	app.DB.Create(&model.UserSchema{Email: string(email), Password: hashPassword})
	user := &model.UserSchema{}
	app.DB.First(user, "email = ?", string(email))
	response, _ := json.Marshal(user)
	fmt.Fprintln(ctx, string(response))
}
func (app *App) SignInApi(ctx *fasthttp.RequestCtx) {
	defer utils.EnableCORS(ctx)
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")
	fmt.Println("email", string(email))
	fmt.Println("password", string(password))
	user := &model.UserSchema{}
	app.DB.First(user, "email = ?", string(email))
	if *&user.Id == 0 {
		fmt.Fprintln(ctx, "User is not existed")
		utils.Response(ctx, "500", "User is not existed")
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(*&user.Password), password)
	if err != nil {
		fmt.Fprintln(ctx, "Password is not correct")
		utils.Response(ctx, "500", "Password is not correct")
		return
	}
	// response, _ := json.Marshal(user)
	jwt, err := utils.EncodeJWT(string(email))
	if err != nil {
		fmt.Println("error", err)
		utils.Response(ctx, "500", "Generate token failed")
		return
	}
	utils.SetCookie(ctx, "token", jwt, 60)
	utils.Response(ctx, "200", "Sign in successfully")
}
func (app *App) ProfileApi(ctx *fasthttp.RequestCtx) {
	defer utils.EnableCORS(ctx)
	cookie := string(ctx.Request.Header.Cookie("token"))
	if len(cookie) == 0 {
		utils.Response(ctx, "403", "Unauthorized")
		return
	}
	data, err := utils.DecodeJWT(cookie)
	if err != nil {
		utils.Response(ctx, "500", "Invalid token")
		return
	}
	utils.Response(ctx, "200", data)
}
func (app *App) SignOutApi(ctx *fasthttp.RequestCtx) {
	defer utils.EnableCORS(ctx)
	cookie := string(ctx.Request.Header.Cookie("token"))
	if len(cookie) == 0 {
		utils.Response(ctx, "403", "Unauthorized")
		return
	}
	_, err := utils.DecodeJWT(cookie)
	if err != nil {
		utils.Response(ctx, "500", "Invalid token")
		return
	}
	utils.DeleteCookie(ctx, "token")
}
func (app *App) CreateTodoApi(ctx *fasthttp.RequestCtx) {
	defer utils.EnableCORS(ctx)
	cookie := string(ctx.Request.Header.Cookie("token"))
	if len(cookie) == 0 {
		utils.Response(ctx, "403", "Unauthorized")
		return
	}
	data, err := utils.DecodeJWT(cookie)
	if err != nil {
		utils.Response(ctx, "500", "Invalid token")
		return
	}
	fmt.Println(data)
	var user model.UserSchema
	userResult := app.DB.Where("email = ?", data.Email).First(&user)
	if userResult.Error != nil {
		utils.Response(ctx, "500", "User is not existed")
		return
	}
	title := ctx.FormValue("title")
	todo := model.TodoSchema{
		Title:  string(title),
		Status: "IN_PROGRESS",
		UserId: uint(user.Id),
	}
	todoResult := app.DB.Create(&todo)
	if todoResult.Error != nil {
		utils.Response(ctx, "500", "Create failed")
		return
	}
	fmt.Println("user", user)
	utils.Response(ctx, "200", todo)
}
func (app *App) GetTodosApi(ctx *fasthttp.RequestCtx) {
	defer utils.EnableCORS(ctx)
	cookie := string(ctx.Request.Header.Cookie("token"))
	if len(cookie) == 0 {
		utils.Response(ctx, "403", "Unauthorized")
		return
	}
	data, err := utils.DecodeJWT(cookie)
	if err != nil {
		utils.Response(ctx, "500", "Invalid token")
		return
	}
	fmt.Println(data)
	var user model.UserSchema
	userResult := app.DB.Where("email = ?", data.Email).First(&user)
	if userResult.Error != nil {
		utils.Response(ctx, "500", "User is not existed")
		return
	}
	var todos []model.TodoSchema
	todosResult := app.DB.Model(&user).Association("Todos").Find(&todos)
	fmt.Println("todos", todos)
	if todosResult != nil {
		utils.Response(ctx, "500", "Get all todos failed")
		return
	}
	utils.Response(ctx, "200", todos)
}
