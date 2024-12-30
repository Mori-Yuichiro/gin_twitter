package router

import (
	"gin-twitter/controllers"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func NewRouter(uc controllers.IUserController) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			// os.Getenv("FE_URL"),
		},
		AllowHeaders: []string{
			"Origin",
			"Accept",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"X-CSRF-Token",
			// "Access-Control-Allow-Credentials",
			// "Content-Length",
			// "Accept-Encoding",
			// "Authorization",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowCredentials: true,
	}))

	r.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	// CSRF
	r.Use(csrf.Middleware(csrf.Options{
		// Secret: CSRFトークンの署名に使用するシークレットを設定します。ランダムな強力な文字列を設定してください。
		Secret: os.Getenv("SECRET"),
		// IgnoreMethods: CSRFトークンを検証しないHTTPメソッドを指定できます。
		IgnoreMethods: []string{"GET"},
		// ErrorFunc: CSRFトークンが一致しない場合に実行されるハンドラを指定できます。
		ErrorFunc: func(c *gin.Context) {
			c.JSON(http.StatusForbidden, gin.H{"error": "CSRF token mismatch"})
			c.Abort()
		},
		// TokenGetter: リクエストからCSRFトークンを取得するカスタム関数です。
		TokenGetter: func(c *gin.Context) string {
			return c.GetHeader("X-CSRF-TOKEN")
		},
	}))

	r.POST("/signup", uc.SignUp)
	r.POST("/login", uc.LogIn)
	r.GET("/csrf", uc.CsrfToken)

	return r
}
