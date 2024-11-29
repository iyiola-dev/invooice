package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/iyiola-dev/numeris/internal/inputs"
	"github.com/iyiola-dev/numeris/internal/service"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

// Auth handlers
func (h *Handler) Register(c *gin.Context) {
	var input inputs.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.svc.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Handler) Login(c *gin.Context) {
	var input inputs.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.svc.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// Invoice handlers
func (h *Handler) CreateInvoice(c *gin.Context) {
	var input inputs.CreateInvoiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	input.UserID = userID.(uuid.UUID)

	invoice, err := h.svc.CreateInvoice(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

func (h *Handler) GetInvoice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
		return
	}

	invoice, err := h.svc.GetInvoiceWithItems(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (h *Handler) GetInvoices(c *gin.Context) {
	// Get user ID from context
	userID, _ := c.Get("userID")

	filters := map[string]interface{}{
		"user_id": userID.(uuid.UUID),
	}

	invoices, err := h.svc.GetInvoices(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch invoices"})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

func (h *Handler) UpdateInvoice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.UpdateInvoice(id, updates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invoice updated successfully"})
}

func (h *Handler) DeleteInvoice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
		return
	}

	err = h.svc.DeleteInvoice(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "invoice deleted successfully"})
}

// Payment Details handlers
func (h *Handler) CreatePaymentDetails(c *gin.Context) {
	var input inputs.CreatePaymentDetailsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse invoice ID from URL
	invoiceID, err := uuid.Parse(c.Param("invoice_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
		return
	}
	input.InvoiceID = invoiceID

	details, err := h.svc.CreatePaymentDetails(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, details)
}

func (h *Handler) GetPaymentDetails(c *gin.Context) {
	invoiceID, err := uuid.Parse(c.Param("invoice_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
		return
	}

	details, err := h.svc.GetPaymentDetailsByInvoiceID(invoiceID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment details not found"})
		return
	}

	c.JSON(http.StatusOK, details)
}

func (h *Handler) UpdatePaymentDetails(c *gin.Context) {
	invoiceID, err := uuid.Parse(c.Param("invoice_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
		return
	}

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.UpdatePaymentDetails(invoiceID, updates)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment details updated successfully"})
}

func (h *Handler) DeletePaymentDetails(c *gin.Context) {
	invoiceID, err := uuid.Parse(c.Param("invoice_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
		return
	}

	err = h.svc.DeletePaymentDetails(invoiceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment details deleted successfully"})
}

// Shareable link handler
func (h *Handler) GetInvoiceByShareableLink(c *gin.Context) {
	invoiceNumber := c.Param("invoice_number")
	if invoiceNumber == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice number"})
		return
	}

	// Get invoice using filters
	filters := map[string]interface{}{
		"invoice_number": invoiceNumber,
	}
	invoices, err := h.svc.GetInvoices(filters)
	if err != nil || len(invoices) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoices[0])
}

// Activity Log handlers
func (h *Handler) GetActivityLogs(c *gin.Context) {
	// Get user ID from context
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	// Get invoice ID from params if provided
	var filters map[string]interface{}
	if invoiceID := c.Query("invoice_id"); invoiceID != "" {
		id, err := uuid.Parse(invoiceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid invoice ID"})
			return
		}
		filters = map[string]interface{}{
			"user_id": userID.(uuid.UUID),
			"invoice_id": id,
		}
	} else {
		filters = map[string]interface{}{
			"user_id": userID.(uuid.UUID),
		}
	}

	logs, err := h.svc.GetActivityLogs(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch activity logs"})
		return
	}

	c.JSON(http.StatusOK, logs)
}
