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

// Initialize Struct Course Controller
type CourseController struct {
	courseUsecase usecase.CourseUsecase
	rg            *gin.RouterGroup
	authMid       middleware.AuthMiddleware
}

// Construction to Access Course Controller
func NewCourseController(courseUsecase usecase.CourseUsecase, rg *gin.RouterGroup, authMid middleware.AuthMiddleware) *CourseController {
	return &CourseController{courseUsecase: courseUsecase, rg: rg, authMid: authMid}
}

// Create Course
func (crs *CourseController) CreateCourse(c *gin.Context) {
	var payload entity.Course
	user := c.MustGet("user").(string)

	payload.UserId = user
	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	course, err := crs.courseUsecase.CreateCourse(payload, user)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, course, "Success")
}

// Get All Courses
func (crs *CourseController) GetAllCourses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	user := c.MustGet("user").(string)

	courses, paging, err := crs.courseUsecase.GetAllCourses(page, size, user)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var interfaceSlice = make([]interface{}, len(courses))
	for i, v := range courses {
		interfaceSlice[i] = v
	}

	common.SendPagedResponse(c, interfaceSlice, paging, "Success")
}

func (crs *CourseController) GetAllCoursesByCategory(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	category := c.Param("category")

	user := c.MustGet("user").(string)

	courses, paging, err := crs.courseUsecase.GetAllCoursesByCategory(category, page, size, user)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var interfaceSlice = make([]interface{}, len(courses))
	for i, v := range courses {
		interfaceSlice[i] = v
	}

	common.SendPagedResponse(c, interfaceSlice, paging, "Success")
}

// Get Course By ID
func (crs *CourseController) GetCourseByID(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user := c.MustGet("user").(string)

	course, err := crs.courseUsecase.GetCourseByID(courseId, user)
	if err != nil {
		common.SendErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	common.SendSingleResponse(c, course, "Success")
}

// Update Course
func (crs *CourseController) UpdateCourse(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId := c.MustGet("user").(string)

	var payload entity.Course
	payload.UserId = userId
	if err := c.ShouldBindUri(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	course, err := crs.courseUsecase.UpdateCourse(courseId, payload, userId)
	if err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	common.SendSingleResponse(c, course, "Success")
}

// Delete Course
func (crs *CourseController) DeleteCourse(c *gin.Context) {
	courseId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		common.SendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet("user").(string)
	if err := crs.courseUsecase.DeleteCourse(courseId, user); err != nil {
		common.SendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendSuccessResponse(c, http.StatusOK, "Course Delete Success")
}

// Routing Course
func (crs *CourseController) Route() {
	crs.rg.GET(routes.GetAllCourse, crs.authMid.RequireToken("admin", "instructor", "user"), crs.GetAllCourses)
	crs.rg.GET(routes.GetCourseByID, crs.authMid.RequireToken("admin", "instructor", "user"), crs.GetCourseByID)
	crs.rg.GET(routes.GetCourseByCategory, crs.authMid.RequireToken("admin", "instructor", "user"), crs.GetAllCoursesByCategory)
	crs.rg.POST(routes.PostCourse, crs.authMid.RequireToken("admin", "instructor"), crs.CreateCourse)
	crs.rg.PUT(routes.PutCourse, crs.authMid.RequireToken("admin", "instructor"), crs.UpdateCourse)
	crs.rg.DELETE(routes.DelCourse, crs.authMid.RequireToken("admin"), crs.DeleteCourse)
}
