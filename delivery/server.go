package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelompok-2/ilmu-padi/config"
	"github.com/kelompok-2/ilmu-padi/config/routes"
	"github.com/kelompok-2/ilmu-padi/delivery/controller"
	"github.com/kelompok-2/ilmu-padi/delivery/middleware"
	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/service"
	"github.com/kelompok-2/ilmu-padi/usecase"
)

type Server struct {
	authUsc        usecase.AuthUsecase
	userUc         usecase.UserUsecase
	courseUc       usecase.CourseUsecase
	userFavoriteUc usecase.UserCoursesFavouriteUsecase
	paymentUc      usecase.PaymentUsecase
	jwtService     service.JwtService
	engine         *gin.Engine
	host           string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(routes.ApiGroup)
	authMid := middleware.NewAuthMiddleware(s.jwtService)
	controller.NewAuthController(s.authUsc, rg).Route()
	controller.NewUserController(s.userUc, rg, authMid).Route()
	controller.NewCourseController(s.courseUc, rg, authMid).Route()
	controller.NewUserCoursesFavouriteController(s.userFavoriteUc, rg, authMid).Route()
	controller.NewPaymentController(s.paymentUc, rg, authMid).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPassword, cfg.DbName)

	db, err := gorm.Open(cfg.DbDriver, dsn)
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Course{}, &entity.Subscription{}, &entity.UserCoursesFavourite{}, &entity.Transaction{})

	jwtService := service.NewJwtService(cfg.TokenConfig)
	mailService := service.NewMailService(cfg.SmtpConfig)

	// Setup Configuration Layer Repo
	authRepo := repository.NewAuthRepository(db)
	userRepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	userCourseFavoriteRepo := repository.NewUserCoursesFavouriteRepository(db)
	// mailRepo := repository.New
	paymentRepo := repository.NewPaymentRepository(db)
	paymentService := service.NewService(paymentRepo, courseRepo)

	// Setup Configuration Layer Usecase
	authUsecase := usecase.NewAuthUsecase(authRepo, jwtService, mailService)
	userUsecase := usecase.NewUserUsecase(userRepo)
	courseUsecase := usecase.NewCourseUsecase(courseRepo)
	userCourseFavotiteUsecase := usecase.NewUserCoursesFavouriteUsecase(userCourseFavoriteRepo)
	paymentUsecase := usecase.NewPaymentUsecase(paymentRepo, courseRepo, paymentService)

	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		authUsc:        authUsecase,
		userUc:         userUsecase,
		courseUc:       courseUsecase,
		userFavoriteUc: userCourseFavotiteUsecase,
		paymentUc:      paymentUsecase,
		jwtService:     jwtService,
		engine:         engine,
		host:           host,
	}
}
