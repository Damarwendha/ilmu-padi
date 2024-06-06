package usecase

import (
	"log"

	"github.com/kelompok-2/ilmu-padi/entity"
	"github.com/kelompok-2/ilmu-padi/repository"
	"github.com/kelompok-2/ilmu-padi/shared/model"
)

// Initialize Struct Course Usecase
type courseUsecase struct {
	courseRepository repository.CourseRepository
}

// Initialize Interface Course Sender Usecase
type CourseUsecase interface {
	CreateCourse(courses entity.Course, user string) (entity.Course, error)
	GetAllCourses(page, size int, user string) ([]entity.Course, model.Paging, error)
	GetAllCoursesByCategory(category string, page, size int, user string) ([]entity.Course, model.Paging, error)
	GetCourseByID(ID int, user string) (entity.Course, error)
	UpdateCourse(ID int, courses entity.Course, user string) (entity.Course, error)
	DeleteCourse(ID int, user string) error
}

// Construction to Access Course Usecase
func NewCourseUsecase(courseRepository repository.CourseRepository) CourseUsecase {
	return &courseUsecase{
		courseRepository: courseRepository,
	}
}

// Create Course
func (c *courseUsecase) CreateCourse(courses entity.Course, user string) (entity.Course, error) {
	course := entity.Course{Title: courses.Title, Description: courses.Description, Category: courses.Category, Video_URL: courses.Video_URL, Duration: courses.Duration, Instructor_Name: courses.Instructor_Name, Rating: courses.Rating}
	if err := c.courseRepository.Create(&course); err != nil {
		return entity.Course{}, err
	}

	return course, nil
}

// Get All Courses
func (c *courseUsecase) GetAllCourses(page, size int, user string) ([]entity.Course, model.Paging, error) {
	return c.courseRepository.FindAll(page, size)
}

func (c *courseUsecase) GetAllCoursesByCategory(category string, page, size int, user string) ([]entity.Course, model.Paging, error) {
	return c.courseRepository.FindAllByCategory(category, page, size)
}

// Get Course By ID
func (c *courseUsecase) GetCourseByID(ID int, user string) (entity.Course, error) {
	return c.courseRepository.FindByID(ID)
}

// Update Course
func (c *courseUsecase) UpdateCourse(ID int, courses entity.Course, user string) (entity.Course, error) {
	log.Printf("Attempting to update course with ID %d by user %s", ID, user)

	// Find the existing course by ID
	course, err := c.courseRepository.FindByID(ID)
	if err != nil {
		log.Printf("Error finding course with ID %d: %v", ID, err)
		return entity.Course{}, err
	}

	// // Validate the input course data
	// if courses.Title == "" || courses.Description == "" || courses.Category == "" || courses.Video_URL == "" || courses.Duration <= 0 || courses.Instructor_Name == "" || courses.Rating < 0 {
	// 	return entity.Course{}, errors.New("invalid course data")
	// }

	// Update the course fields
	course.Title = courses.Title
	course.Description = courses.Description
	course.Category = courses.Category
	course.Video_URL = courses.Video_URL
	course.Duration = courses.Duration
	course.Instructor_Name = courses.Instructor_Name
	course.Rating = courses.Rating

	// Attempt to update the course in the repository
	if err := c.courseRepository.Update(course); err != nil {
		log.Printf("Error updating course with ID %d: %v", ID, err)
		return entity.Course{}, err
	}

	log.Printf("Successfully updated course with ID %d by user %s", ID, user)
	return course, nil
}

// Delete Course
func (c *courseUsecase) DeleteCourse(ID int, user string) error {
	return c.courseRepository.Delete(ID)
}
