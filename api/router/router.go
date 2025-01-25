package router

import (
	"gin-twitter/controllers"
	"gin-twitter/middlewares"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

func NewRouter(
	uc controllers.IUserController,
	tc controllers.ITweetController,
	cc controllers.ICommentController,
	rc controllers.IRetweetController,
	fc controllers.IFavoriteController,
) *gin.Engine {
	r := gin.Default()

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			os.Getenv("FE_URL"),
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Access-Control-Allow-Headers",
			"X-CSRF-Token",
			"Access-Control-Allow-Origin",
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
			"OPTIONS",
		},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	api.Use(gin.Logger())

	api.Use(sessions.Sessions("mysession", cookie.NewStore([]byte("secret"))))
	// CSRF
	api.Use(csrf.Middleware(csrf.Options{
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
			// クッキーからトークンを取得
			csrfToken, err := c.Cookie("_csrf")
			if err == nil && csrfToken != "" {
				return csrfToken
			}

			return c.GetHeader("X-CSRF-Token")
		},
	}))

	api.POST("/signup", uc.SignUp)
	api.POST("/login", uc.LogIn)
	api.POST("/logout", uc.LogOut)
	api.GET("/csrf", uc.CsrfToken)

	user := api.Group("/users")
	user.Use(middlewares.AuthMiddleware)
	{
		user.GET("", uc.GetUserIdByToken)
		user.GET("/:userId", uc.GetUserByUserId)
	}

	tweet := api.Group("/tweets")
	tweet.Use(middlewares.AuthMiddleware)
	{
		tweet.POST("", tc.CreateTweet)
		tweet.GET("", tc.GetAllTweet)
		twid := tweet.Group("/:tweetId")
		twid.GET("", tc.GetTweetById)
		twid.DELETE("", tc.DeleteTweet)
		twid.POST("/retweet", rc.CreateRetweet)
		twid.DELETE("/retweet", rc.DeleteRetweet)
		twid.POST("/favorite", fc.CreateFavorite)
		twid.DELETE("/favorite", fc.DeleteFavorite)
	}

	comment := api.Group("/comment")
	comment.Use(middlewares.AuthMiddleware)
	{
		comment.POST("", cc.CreateComment)
	}

	return r
}
