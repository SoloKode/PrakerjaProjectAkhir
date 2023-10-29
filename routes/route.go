package routes

import (
	"projectakhir/controllers"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	// Middleware Logger akan mencatat permintaan HTTP dan tanggapan dalam log
	e.Use(middleware.Logger())

	// Middleware RemoveTrailingSlash akan menghapus karakter trailing slash dari URL jika ada
	e.Pre(middleware.RemoveTrailingSlash())

	// Grouping Route: Membuat grup rute untuk pengguna dengan peran "member" dan "admin"
	eMemberAuth := e.Group("/member")
	eAdminAuth := e.Group("/admin")

	// Middleware JWT digunakan untuk otentikasi berdasarkan token JWT
	eMemberAuth.Use(echojwt.JWT([]byte("tokobuku")))
	eAdminAuth.Use(echojwt.JWT([]byte("admintokobuku")))

	// Rute untuk admin terkait buku
	eAdminAuth.GET("/books", controllers.GetBookController)
	eAdminAuth.POST("/books", controllers.AddBookController)
	eAdminAuth.GET("/books/:id", controllers.GetDetailBookController)

	// Rute untuk admin terkait anggota (member)
	eAdminAuth.GET("/members", controllers.GetMembersController)
	eAdminAuth.PUT("/members/:id", controllers.UpdateMemberController)
	eAdminAuth.GET("/members/:id", controllers.GetDetailMemberController)
	eAdminAuth.GET("/sewa", controllers.GetSewaController)

	// Rute untuk anggota (member) terkait buku
	eMemberAuth.GET("/books", controllers.GetBookController)
	eMemberAuth.GET("/profile", controllers.GetProfileController)
	eMemberAuth.PUT("/profile", controllers.UpdateProfileController)

	// Rute untuk anggota (member) menyewa buku
	eMemberAuth.POST("/sewa/:idbuku", controllers.SewaBukuController)

	// Rute untuk registrasi pengguna baru
	e.POST("/register", controllers.RegisterController)

	// Rute untuk masuk (login) pengguna
	e.POST("/login", controllers.LoginController)
}
