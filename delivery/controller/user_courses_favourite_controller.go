package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/shared/common"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

type UserCoursesFavouriteController struct {
	userCoursesFavouriteUsecase usecase.UserCoursesFavouriteUsecase
	rg                          *gin.RouterGroup
	authMid                     middleware.AuthMiddleware
}

func NewUserCoursesFavouriteController(userCoursesFavouriteUsecase usecase.UserCoursesFavouriteUsecase, rg *gin.RouterGroup, authMid middleware.AuthMiddleware) *UserCoursesFavouriteController {
	return &UserCoursesFavouriteController{
		userCoursesFavouriteUsecase: userCoursesFavouriteUsecase,
		rg:                          rg,
		authMid:                     authMid,
	}
}

func (u *UserCoursesFavouriteController) AddOrRemoveCourseFavourite(c *gin.Context) {
	var ucf entity.UserCoursesFavourite

	if err := c.ShouldBindJSON(&ucf); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	msg, err := u.userCoursesFavouriteUsecase.AddOrRemoveToFavourite(ucf)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSuccessResponse(c, http.StatusOK, msg)
}

func (u *UserCoursesFavouriteController) GetUserFavouriteList(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	favouriteList, err := u.userCoursesFavouriteUsecase.FindAllByUserID(uint(userID))
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, favouriteList, "Success")
}

// Routing User Course Favorite
func (u *UserCoursesFavouriteController) Route() {
	u.rg.GET(routes.GetUserCourseFavouriteList, u.authMid.RequireToken("admin", "user", "instructor"), u.GetUserFavouriteList)
	u.rg.POST(routes.PostUserCourseFavourite, u.authMid.RequireToken("admin", "user", "instructor"), u.AddOrRemoveCourseFavourite)
}
