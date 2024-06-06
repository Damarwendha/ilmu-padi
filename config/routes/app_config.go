package routes

const (
	// Route Middleware
	GroupMiddleware = "/asset"

	// Route Group
	ApiGroup = "/api/v1"

	// Route Course
	GetAllCourse        = "/asset/courses"
	GetCourseByID       = "/asset/courses/:id"
	GetCourseByCategory = "/asset/courses/category/:category"
	PostCourse          = "/asset/courses"
	PutCourse           = "/asset/courses/:id"
	DelCourse           = "/asset/courses/:id"

	// Route User Course Favourite
	PostUserCourseFavourite    = "/course/favourite"
	GetUserCourseFavouriteList = "/course/favourite/:id"

	// Route Authentitacion
	Register       = "/auth/register"
	Login          = "/auth/login"
	Logout         = "/auth/logout"
	ForgotPassword = "/auth/forgot-password"
	ResetPassword  = "/auth/reset-password"
	VerifyEmail    = "/auth/verify-email"

	// Route User
	GetAllProfile                  = "/profile"
	GetProfileByID                 = "/profile/:id"
	GetProfileByEmail              = "/profile/:email"
	GetProfileBySubscriptionStatus = "/profile/:subscription-status"
	GetProfileByCourseName         = "/profile/:course"

	// Route Payment
	GetCourseTransaction = "/courses/:id/transaction"
	PostTransaction      = "/transaction/payment"
	GetNotification      = "/transaction/notification"
)
