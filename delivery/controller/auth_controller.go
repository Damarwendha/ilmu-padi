package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/entity/dto"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

// Initialize Struct Auth Controller
type AuthController struct {
	authUsecase usecase.AuthUsecase
	rg          *gin.RouterGroup
}

// Construction to Access Auth Controller
func NewAuthController(authUsecase usecase.AuthUsecase, rg *gin.RouterGroup) *AuthController {
	return &AuthController{authUsecase: authUsecase, rg: rg}
}

// Register
func (a *AuthController) Register(c *gin.Context) {
	input := &dto.RegisterDto{}
	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.authUsecase.Register(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendCreateResponse(c, user)
}

// Login
func (a *AuthController) Login(c *gin.Context) {
	input := &dto.LoginDto{}
	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := a.authUsecase.Login(input)
	if err != nil {
		common.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Set token in cookie
	c.SetCookie("token", token, 3600, "/", "", false, true)

	common.SendSuccessResponse(c, http.StatusOK, "Login Success")
}

// Logout
func (a *AuthController) Logout(c *gin.Context) {
	if err := a.authUsecase.Logout(); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Clear the token cookie
	c.SetCookie("token", "", -1, "/", "", false, true)

	common.SendSuccessResponse(c, http.StatusOK, "Logout Success")
}

// Verify Email
func (a *AuthController) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	log.Println("token:", token)
	err := a.authUsecase.VerifyEmail(token)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, "Email Verified Success")
}

// Forgot Password
func (a *AuthController) ForgotPassword(c *gin.Context) {
	input := &dto.ForgotPasswordDto{}

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.authUsecase.ForgotPassword(input); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, "Password Reset Link Sent...")
}

// Reset Password
func (a *AuthController) ResetPassword(c *gin.Context) {
	input := &dto.ResetPasswordDto{}

	if err := c.ShouldBindJSON(&input); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := a.authUsecase.ResetPassword(input); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, "Password Has Been Reset!")
}

// Routing Auth
func (a *AuthController) Route() {
	a.rg.POST(routes.Login, a.Login)
	a.rg.POST(routes.Register, a.Register)
	a.rg.POST(routes.Logout, a.Logout)
	a.rg.POST(routes.ForgotPassword, a.ForgotPassword)
	a.rg.POST(routes.ResetPassword, a.ResetPassword)
	a.rg.POST(routes.VerifyEmail, a.VerifyEmail)
}
