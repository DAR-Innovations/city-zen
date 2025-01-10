package volunteer

import (
	"errors"

	"github.com/DAR-Innovations/city-zen/internal/data"
	"github.com/DAR-Innovations/city-zen/internal/modules/tasks/volunteer/types"
	"gorm.io/gorm"
)

type VolunteerService interface {
	GetTasks(filter types.TaskFilterDTO, page, limit int) ([]types.VolunteerTaskDTO, error)
	ReportTask(taskID uint, input *types.ReportTaskDTO) error
}

type volunteerService struct {
	db *gorm.DB
}

func NewVolunteerService(db *gorm.DB) VolunteerService {
	return &volunteerService{
		db: db,
	}
}

func (s *volunteerService) GetTasks(filter types.TaskFilterDTO, page, limit int) ([]types.VolunteerTaskDTO, error) {
	tasks := []data.VolunteerTask{}

	query := s.db.Model(&data.VolunteerTask{}).Preload("Issue").Preload("Assignee")

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Urgency != "" {
		query = query.Where("urgency = ?", filter.Urgency)
	}
	if filter.Complexity != "" {
		query = query.Where("complexity = ?", filter.Complexity)
	}

	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	if err := query.Find(&tasks).Error; err != nil {
		return nil, err
	}

	taskDTOs := make([]types.VolunteerTaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = types.VolunteerTaskDTO{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			ImageURL:    task.Issue.ImageURL,
			Location: types.LocationDTO{
				Longitude: task.Issue.Longitude,
				Latitude:  task.Issue.Latitude,
			},
			Status:     task.Status,
			Urgency:    task.Urgency,
			Complexity: task.Complexity,
			Assignee: &types.AssigneeDTO{
				ID:    task.Volunteer.ID,
				Name:  task.Volunteer.FirstName + task.Volunteer.LastName,
				Phone: task.Volunteer.Phone,
			},
		}

		if task.VolunteerID == 0 {
			taskDTOs[i].Assignee = nil
		}
	}

	return taskDTOs, nil
}

func (s *volunteerService) ReportTask(taskID uint, input *types.ReportTaskDTO) error {
	task := &data.VolunteerTask{}
	if err := s.db.Preload("Issue").First(task, taskID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("task not found")
		}
		return err
	}

	if task.Status == "DONE" {
		return errors.New("task is already completed")
	}

	// TODO: Add AI Validation for the report image here

	task.Status = "DONE"

	if err := s.db.Model(&data.Issue{}).UpdateColumn("is_completed", true).Error; err != nil {
		return err
	}

	return s.db.Save(task).Error
}
