package interactors

import "Portfolio_Nodes/domain"

type TaskInteractor struct {
	taskRepo domain.TasksRepo
}

func NewTaskInteractor(taskRepo domain.TasksRepo) *TaskInteractor {
	return &TaskInteractor{taskRepo: taskRepo}
}

func (i *TaskInteractor) GetAllTasks() ([]domain.Task, error) {
	all, err := i.taskRepo.FetchTasks()
	if err != nil {
		return nil, err
	}
	return all, err
}

func (i *TaskInteractor) ToggleTask(taskId int) error {
	err := i.taskRepo.UpdateTask(taskId)
	return err
}

func (i *TaskInteractor) GetTaskById(taskId int) (domain.Task, error) {
	t, err := i.taskRepo.FetchTaskById(taskId)
	if err != nil {
		return t, err
	}
	return t, err
}

func (i *TaskInteractor) CreateTask(task domain.TaskForms) error {

	d := domain.Task{
		Name:     task.Name,
		Priority: task.Priority,
		Done:     task.Done,
	}
	err := i.taskRepo.InsertTask(d)
	return err
}
