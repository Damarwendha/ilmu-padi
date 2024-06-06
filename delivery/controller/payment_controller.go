package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

// Initialize Struct Payment Controller
type PaymentController struct {
	paymentUsecase usecase.PaymentUsecase
	rg             *gin.RouterGroup
	authMid        middleware.AuthMiddleware
}

// Construction to Access Payment Controller
func NewPaymentController(paymentUsecase usecase.PaymentUsecase, rg *gin.RouterGroup, authMid middleware.AuthMiddleware) *PaymentController {
	return &PaymentController{paymentUsecase: paymentUsecase, rg: rg, authMid: authMid}
}

// Get Campaign Transactions
func (h *PaymentController) GetCampaignTransactions(c *gin.Context) {
	var payload dto.GetCourseTransactionsInput

	err := c.ShouldBindUri(&payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet("user").(string)

	payload.UserId = user

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	transactions, err := h.paymentUsecase.GetTransactionsByCourseID(payload, user)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, transactions)
}

// Create Transaction
func (h *PaymentController) CreateTransaction(c *gin.Context) {
	var payload dto.CreateTransactionInput

	// Coba bind JSON dari body request
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, "Invalid JSON format: "+err.Error())
		return
	}

	// Ambil user dari context
	user, exists := c.Get("user")
	if !exists {
		common.SendErrorResponse(c, http.StatusUnauthorized, "User not found in context")
		return
	}

	// Pastikan user adalah string
	userStr, ok := user.(string)
	if !ok {
		common.SendErrorResponse(c, http.StatusUnauthorized, "Invalid user format")
		return
	}

	payload.UserId = userStr

	// Panggil usecase untuk membuat transaksi baru
	newTransaction, err := h.paymentUsecase.CreateTransaction(payload, userStr)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, newTransaction)
}

// Get Notification
func (h *PaymentController) GetNotification(c *gin.Context) {
	var payload dto.TransactionNotificationInput

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet("user").(string)

	payload.UserId = user

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.paymentUsecase.ProcessPayment(payload, user)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, payload)
}

// Routing Payment
func (pc *PaymentController) Route() {
	pc.rg.GET(routes.GetCourseTransaction, pc.authMid.RequireToken("admin", "instructor", "user"), pc.CreateTransaction)
	pc.rg.POST(routes.PostTransaction, pc.authMid.RequireToken("admin", "instructor", "user"), pc.CreateTransaction)
	pc.rg.POST(routes.GetNotification, pc.authMid.RequireToken("admin", "instructor", "user"), pc.CreateTransaction)
}
