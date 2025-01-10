package department

import (
	"errors"

	"github.com/DAR-Innovations/city-zen/internal/data"
	"github.com/DAR-Innovations/city-zen/internal/modules/tasks/department/types"
	"gorm.io/gorm"
)

type DepartmentService interface {
	GetTasks(filter types.TaskFilterDTO, page, limit int) ([]types.DepartmentTaskDTO, error)
	AssignTask(taskID, assigneeID uint) error
	ReportTask(taskID uint, input *types.ReportTaskDTO) error
}

type departmentService struct {
	db *gorm.DB
}

func NewDepartmentService(db *gorm.DB) DepartmentService {
	return &departmentService{
		db: db,
	}
}

func (s *departmentService) GetTasks(filter types.TaskFilterDTO, page, limit int) ([]types.DepartmentTaskDTO, error) {
	tasks := []data.DepartmentTask{}
	query := s.db.Model(&data.DepartmentTask{}).Preload("Department").Preload("TaskType").Preload("Issue").Preload("Assignee")

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.Urgency != "" {
		query = query.Where("urgency = ?", filter.Urgency)
	}
	if filter.Complexity != "" {
		query = query.Where("complexity = ?", filter.Complexity)
	}
	if filter.DepartmentID != 0 {
		query = query.Where("department_id = ?", filter.DepartmentID)
	}

	offset := (page - 1) * limit
	query = query.Offset(offset).Limit(limit)

	err := query.Find(&tasks).Error
	if err != nil {
		return nil, err
	}

	taskDTOs := make([]types.DepartmentTaskDTO, len(tasks))
	for i, task := range tasks {
		taskDTOs[i] = types.DepartmentTaskDTO{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Type:        task.TaskType.Name,
			ImageURL:    task.Issue.ImageURL,
			Location: types.LocationDTO{
				Longitude: task.Issue.Longitude,
				Latitude:  task.Issue.Latitude,
			},
			Status:         task.Status,
			Urgency:        task.Urgency,
			Complexity:     task.Complexity,
			DepartmentName: task.Department.Name,
			Assignee: &types.AssigneeDTO{
				ID:    task.Assignee.ID,
				Name:  task.Assignee.FirstName + task.Assignee.LastName,
				Phone: task.Assignee.Phone,
			},
		}

		if task.AssigneeID == nil {
			taskDTOs[i].Assignee = nil
		}
	}

	return taskDTOs, nil
}

func (s *departmentService) AssignTask(taskID, assigneeID uint) error {
	var task data.DepartmentTask
	err := s.db.First(&task, taskID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("task not found")
	}

	task.AssigneeID = &assigneeID
	task.Status = "IN PROGRESS"

	return s.db.Save(&task).Error
}

func (s *departmentService) ReportTask(taskID uint, input *types.ReportTaskDTO) error {
	task := &data.DepartmentTask{}
	err := s.db.Preload("Issue").First(task, taskID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("task not found")
	}

	if task.Status == "DONE" {
		return errors.New("task is already completed")
	}

	// TODO: add AI VALIDATION

	report := &data.DepartmentReport{
		IssueID:     task.IssueID,
		Description: input.Comment,
		ImageURL:    input.ImageURL,
		ReportedBy:  *task.AssigneeID,
	}

	if err := s.db.Create(report).Error; err != nil {
		return err
	}

	task.Status = "DONE"

	if err := s.db.Model(&data.Issue{}).UpdateColumn("is_completed", true).Error; err != nil {
		return err
	}

	return s.db.Save(task).Error
}
