package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iyiola-dev/numeris/internal/handlers"
	"github.com/iyiola-dev/numeris/internal/repository"
	"github.com/iyiola-dev/numeris/internal/service"
	"github.com/iyiola-dev/numeris/internal/util"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	// Initialize dependencies
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handlers.NewHandler(svc)

	// Public routes
	router.POST("/api/auth/register", h.Register)
	router.POST("/api/auth/login", h.Login)
	router.GET("/api/invoices/shared/:invoice_number", h.GetInvoiceByShareableLink)

	// Protected routes
	api := router.Group("/api")
	api.Use(util.AuthMiddleware(repo))
	{
		// Invoice routes
		invoices := api.Group("/invoices")
		{
			invoices.POST("", h.CreateInvoice)
			invoices.GET("", h.GetInvoices)
			invoices.GET("/:id", h.GetInvoice)
			invoices.PUT("/:id", h.UpdateInvoice)
			invoices.DELETE("/:id", h.DeleteInvoice)

			// Payment details routes
			invoices.POST("/:invoice_id/payment", h.CreatePaymentDetails)
			invoices.GET("/:invoice_id/payment", h.GetPaymentDetails)
			invoices.PUT("/:invoice_id/payment", h.UpdatePaymentDetails)
			invoices.DELETE("/:invoice_id/payment", h.DeletePaymentDetails)
		}
	}

	return router
}

