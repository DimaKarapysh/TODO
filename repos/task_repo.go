package repos

import (
	"Portfolio_Nodes/domain"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) FetchTasks() ([]domain.Task, error) {
	var tasks []domain.Task
	result := r.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (r *TaskRepo) FetchTaskById(taskId int) (domain.Task, error) {
	var task domain.Task
	result := r.db.First(&task, taskId)
	if result.Error != nil {
		return task, result.Error
	}
	return task, nil
}

func (r *TaskRepo) InsertTask(task domain.Task) error {

	result := r.db.Create(&task)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}

func (r *TaskRepo) UpdateTask(taskId int) error {
	var task domain.Task
	result := r.db.Model(&task).Where("id", taskId).Update("done", !task.Done)
	if result.Error != nil {
		return result.Error
	}
	return result.Error
}
