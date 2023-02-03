package domain

type Task struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
	Done     bool   `json:"done"`
}

type TaskForm struct {
	Id int `json:"id"binding:"required"`
}
type TaskForms struct {
	Id       int    `json:"id"binding:"required"`
	Name     string `json:"name"binding:"required"`
	Priority int    `json:"priority"binding:"required"`
	Done     bool   `json:"done"binding:"required"`
}

type TasksRepo interface {
	FetchTasks() ([]Task, error)
	FetchTaskById(taskId int) (Task, error)
	InsertTask(task Task) error
	UpdateTask(taskId int) error
}

type TasksInteractor interface {
	GetAllTasks() ([]Task, error)
	ToggleTask(taskId int) error
	GetTaskById(taskId int) (Task, error)
	CreateTask(task TaskForms) error
}
