// routes/router.go
package routes

import (
	"github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
	"github.com/saarthi123/saarthi-backend/controllers"
	"github.com/saarthi123/saarthi-backend/handlers"
	"github.com/saarthi123/saarthi-backend/middleware"
	"github.com/saarthi123/saarthi-backend/websocket"
)

func RegisterRoutes(router *gin.Engine) {
	// Public routes
	RegisterAuthRoutes(router)
	RegisterPredictionRoutes(router)
	RegisterCareerPathRoutes(router)
	RegisterSearchRoutes(router)
	RegisterWebSocketRoutes(router)
	RegisterPublicQueryRoutes(router)

	// Public Instructor Panel, Role & Draft routes
	RegisterInstructorPanelRoutes(router)
	RegisterRoleRoutes(router)
	RegisterDraftRoutes(router)

	// Protected API routes
	api := router.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		RegisterAnalyticsRoutes(api)
		RegisterAttendanceRoutes(api)
		RegisterCampusRoutes(api)
		RegisterCourseRoutes(api)
		RegisterCourseManagementRoutes(api)
		RegisterDashboardRoutes(api)
		RegisterMessagingRoutes(api)
		RegisterMailRoutes(api)
		RegisterNotificationRoutes(api)
		RegisterPaymentsRoutes(api)
		RegisterAIEmojiRoutes(api)
		RegisterUserRoutes(api)
	}
}

func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/api/login", controllers.LoginHandler)
}

func RegisterSearchRoutes(r *gin.Engine) {
	r.GET("/api/search", controllers.SearchHandler)
}

func RegisterPredictionRoutes(r *gin.Engine) {
	r.GET("/ai/predictions", handlers.GetPredictions)
}

func RegisterCareerPathRoutes(r *gin.Engine) {
	r.POST("/api/career-paths", controllers.GenerateCareerPaths)
}

func RegisterWebSocketRoutes(r *gin.Engine) {
	r.GET("/ws/public", func(c *gin.Context) {
		websocket.HandlePublicWS(c.Writer, c.Request)
	})
	r.GET("/ws/notifications", func(c *gin.Context) {
		websocket.HandleNotificationWS(c.Writer, c.Request)
	})
}

func RegisterPublicQueryRoutes(r *gin.Engine) {
	r.GET("/api/queries", controllers.GetAllQueries)
	r.POST("/api/queries", controllers.SubmitQuery)
}

func RegisterInstructorPanelRoutes(r *gin.Engine) {
	r.GET("/courses", controllers.GetCourses)
	r.POST("/courses", controllers.UploadCourse)
	r.GET("/live-sessions", controllers.GetLiveSessions)
	r.POST("/live-sessions", controllers.ScheduleSession)
	r.GET("/queries", controllers.GetQueries)
	r.POST("/queries/:id/reply", controllers.ReplyQuery)
	r.GET("/progress", controllers.GetProgress)
	r.GET("/progress/:studentId", controllers.GetStudentProgress)
}

func RegisterRoleRoutes(r *gin.Engine) {
	r.GET("/roles", controllers.GetRoles)
	r.PUT("/roles/:role", controllers.UpdateRolePermissions)
	r.GET("/roles-all", controllers.GetAllRoles)
}


func RegisterDraftRoutes(r *gin.Engine) {
	r.POST("/drafts", controllers.SaveDraftHandler)
	r.GET("/drafts", controllers.GetDraftsHandler)
	r.DELETE("/drafts/:id", controllers.DeleteDraftHandler)
}

func RegisterAnalyticsRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/analytics")
	{
		r.GET("/engagement", controllers.GetEngagement)
		r.GET("/dropoffs", controllers.GetDropOffPoints)
		r.POST("/ai-suggestions", controllers.GetAISuggestions)
	}
}

func RegisterAttendanceRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/attendance")
	{
		r.GET("", controllers.GetAttendance)
		r.POST("/mark", controllers.MarkAttendance)
	}
}

func RegisterCampusRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/campus-attendance")
	{
		r.GET("", controllers.GetCampusAttendance)
	}
}

func RegisterCourseRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/")
	{
		r.GET("courses", controllers.GetCourses)
		r.POST("enroll", controllers.EnrollCourse)
	}
}

func RegisterCourseManagementRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/")
	{
		r.GET("instructor/courses", controllers.GetInstructorCourses)
		r.PUT("course/:id", controllers.UpdateCourse)
		r.DELETE("course/:id", controllers.DeleteCourse)
		r.POST("ai/analyze", controllers.AnalyzeCourseAI)
	}
}

func RegisterDashboardRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/user")
	{
		r.GET(":id/dashboard", controllers.GetDashboardData)
		r.POST(":id/continue-course", controllers.ContinueCourse)
	}
}

func RegisterMessagingRoutes(rg *gin.RouterGroup) {
	msg := rg.Group("/messaging")
	{
		msg.GET("/inbox", controllers.GetInboxMessages)
		msg.POST("/send", controllers.SendMessage)
		// Future messaging endpoints
	}
}

func RegisterMailRoutes(rg *gin.RouterGroup) {
	rg.GET("/mails", controllers.GetMails)
	rg.POST("/mails", controllers.SendMail)
}

func RegisterNotificationRoutes(rg *gin.RouterGroup) {
	n := rg.Group("/notifications")
	{
		n.GET("", controllers.GetNotifications)
		n.POST("/preferences", controllers.UpdateNotificationPreferences)
	}
}

func RegisterPaymentsRoutes(rg *gin.RouterGroup) {
	p := rg.Group("/payments")
	{
		p.GET("", controllers.GetPayments)       // optional
		p.POST("/create", controllers.CreatePayment)
	}
}
func RegisterAIEmojiRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/")
	{
		r.POST("ai/suggestions", handlers.AISuggestionsHandler)
		r.GET("emojis/categories", handlers.GetEmojiCategories)
		r.GET("emojis", handlers.GetEmojisByCategory)
	}
}

func RegisterUserRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")
	{
		user.GET(":id", controllers.GetUserHandler)
		user.GET("/check-profile/:phone", controllers.CheckProfileHandler)
		user.GET("/profile", controllers.GetProfile)
		user.PUT("/profile", controllers.UpdateProfile)
		user.PUT("/profile/password", controllers.ChangePassword)
		user.DELETE("/profile", controllers.DeleteAccount)
	}

	security := rg.Group("/security")
	{
		security.GET("/status", controllers.GetSecurityStatus)
		security.GET("/devices", controllers.GetDeviceActivity)
		security.POST("/change-pin", controllers.ChangePin)
		security.DELETE("/device/:id", controllers.LogoutDevice)
		security.POST("/alert-toggle", controllers.ToggleAlert)
	}

	tx := rg.Group("/transactions")
	{
		tx.GET("", controllers.GetTransactions)
		tx.POST("/filter", controllers.GetFilteredTransactions)
		tx.POST("/download/pdf", controllers.DownloadTransactionsPDF)
		tx.POST("/download/csv", controllers.DownloadTransactionsCSV)
		tx.POST("/export/pdf", controllers.ExportTransactionsPDF)
		tx.POST("/export/excel", controllers.ExportTransactionsExcel)
	}

	statements := rg.Group("/statements")
	{
		statements.POST("/filter", controllers.GetFilteredTransactions)
		statements.POST("/download/pdf", controllers.DownloadPDF)
		statements.POST("/download/excel", controllers.DownloadExcel)
	}

	trading := rg.Group("/trading")
	{
		trading.GET("/portfolio", controllers.GetPortfolio)
		trading.POST("/trade", controllers.PlaceTrade)
		trading.GET("/exchange-info", controllers.GetExchangeInfo)
		trading.GET("/market-news", controllers.GetMarketNews)
		trading.GET("/history", controllers.GetTradingHistory)
		trading.GET("/history/export", controllers.ExportTradingHistory)
		trading.POST("/reset", controllers.ResetSimulation)
	}

	upload := rg.Group("/upload")
	{
		upload.POST("/", controllers.UploadFile)
	}

	upi := rg.Group("/upi")
	{
		upi.POST("/change-pin", handlers.ChangeUpiPin)
		upi.POST("/link-bank", handlers.AddBank)
		upi.GET("/history", handlers.GetUpiHistory)
	}
}


func main() {
    router := gin.Default()
router.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"*"}, // or specific frontend IP
    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    AllowCredentials: true,
}))
}