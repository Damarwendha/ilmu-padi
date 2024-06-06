package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

// Initialize Struct User Controller
type UserController struct {
	userUsecase usecase.UserUsecase
	rg          *gin.RouterGroup
	authMid     middleware.AuthMiddleware
}

// Construction to Access User Controller
func NewUserController(userUsecase usecase.UserUsecase, rg *gin.RouterGroup, authMid middleware.AuthMiddleware) *UserController {
	return &UserController{userUsecase: userUsecase, rg: rg, authMid: authMid}
}

// Find All
func (payload *UserController) GetAllProfile(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	user := c.MustGet("user").(string)

	users, paging, err := payload.userUsecase.FindAll(page, size, user)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var interfaceSlice = make([]interface{}, len(users))
	for i, v := range users {
		interfaceSlice[i] = v
	}
	common.SendPagedResponse(c, interfaceSlice, paging, "Success")
}

// Get Profile By ID
func (ctrl *UserController) GetProfileByID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	users := c.MustGet("user").(string)
	user, err := ctrl.userUsecase.GetProfileByID(uint(userID), users)
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(c, user, "Success")
}

// Routing User
func (u *UserController) Route() {
	u.rg.GET(routes.GetAllProfile, u.authMid.RequireToken("admin"), u.GetAllProfile)
	u.rg.GET(routes.GetProfileByID, u.authMid.RequireToken("admin"), u.GetProfileByID)
	// u.rg.GET(routes.GetProfileByEmail, u.authMid.RequireToken("user"), u.GetAllProfile)
	// u.rg.GET(routes.GetProfileBySubscriptionStatus, u.authMid.RequireToken("user"), u.GetAllProfile)
	// u.rg.GET(routes.GetProfileByCourseName, u.authMid.RequireToken("user"), u.GetAllProfile)
}
