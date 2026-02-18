package services

import (
	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(student models.Student) error {
	return s.Repo.Create(student)
}

func (s *StudentService) UpdateStudent(id string, student models.Student) (models.Student, error) {
	err := s.Repo.UpdateStudent(id, student)
	if err != nil {
		return models.Student{}, err
	}

	student.Id = id
	return student, nil
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.DeleteStudent(id)
}
