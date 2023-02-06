package rest_delivery

import (
	"Portfolio_Nodes/domain"
	"Portfolio_Nodes/interactors"
	"github.com/gin-gonic/gin"
)

type TaskDelivery struct {
	taskIter domain.TasksInteractor
}

func NewTaskDelivery(taskIter *interactors.TaskInteractor) *TaskDelivery {
	return &TaskDelivery{taskIter: taskIter}
}
func (s *TaskDelivery) Route(r *gin.RouterGroup) {
	r.GET("/tasks", s.All)
	r.POST("/tasks/byId", s.ById)
	r.POST("/task/create", s.Create)
	r.POST("/task/toggle", s.Toggle)
}
func (s *TaskDelivery) All(r *gin.Context) {
	tasks, err := s.taskIter.GetAllTasks()
	if err != nil {
		_ = r.Error(err)
		return
	}
	r.JSON(200, MakeSuccessWithData(tasks))
}

func (s *TaskDelivery) ById(r *gin.Context) {
	idForm := domain.TaskForm{}
	err := r.ShouldBindJSON(&idForm)
	if err != nil {
		_ = r.Error(NewUserError("Not Found Id", err))
	}
	task, err := s.taskIter.GetTaskById(idForm.Id)
	if err != nil {
		_ = r.Error(err)
		return
	}
	r.JSON(200, MakeSuccessWithData(task))
}

func (s *TaskDelivery) Create(r *gin.Context) {
	idForms := domain.TaskForms{}
	err := r.ShouldBindJSON(&idForms)
	if err != nil {
		_ = r.Error(NewUserError("Not Invalid Data", err))
		return
	}
	err = s.taskIter.CreateTask(idForms)
	if err != nil {
		_ = r.Error(err)
		return
	}
	r.JSON(200, MakeSuccessWithData("Created!"))
}
func (s *TaskDelivery) Toggle(r *gin.Context) {
	idForm := domain.TaskForm{}
	err := r.ShouldBindJSON(&idForm)
	if err != nil {
		_ = r.Error(NewUserError("Not Found Id", err))
	}
	err = s.taskIter.ToggleTask(idForm.Id)
	if err != nil {
		_ = r.Error(err)
		return
	}
	r.JSON(200, MakeSuccessWithData("Change!"))
}
