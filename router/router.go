package router

import (
	"net/http"

	"github.com/wolv89/gotnsapp/middleware"
)


func LoadRouter() http.Handler {

	router := http.NewServeMux()
	LoadPublicRoutes(router)

	adminRouter := http.NewServeMux()
	LoadAdminRoutes(adminRouter)

	adminMiddleware := middleware.CreateStack(
		middleware.IsAuthenticated,
	)

	router.Handle("/admin/", http.StripPrefix("/admin", adminMiddleware(adminRouter)))

	publicMiddleware := middleware.CreateStack(
		middleware.EnableCors,
	)

	return publicMiddleware(router)

}