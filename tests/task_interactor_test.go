package tests

import (
	"Portfolio_Nodes/app"
	"Portfolio_Nodes/domain"
	"Portfolio_Nodes/interactors"
	"Portfolio_Nodes/repos"
	"github.com/stretchr/testify/require"
	"testing"
)

func initTaskIter() *interactors.TaskInteractor {
	return interactors.NewTaskInteractor(repos.NewTaskRepo(app.DB))
}

func initSliceForId() []domain.Task {
	s := []domain.Task{
		{0, "342", 2, true},
		{0, "242", 3, true},
		{0, "2345", 1, true},
	}
	return s
}

func init() {
	err := app.InitApp()
	if err != nil {
		panic(err)
	}

	//// Initialize Database
	_, err = app.InitDatabase()
	if err != nil {
		panic(err)
	}
	err = app.RunMigrations()
	if err != nil {
		panic(err)
	}

}

func TestGetAllTasks(t *testing.T) {
	r := initTaskIter()
	all, err := r.GetAllTasks()
	require.NoError(t, err, "should not be error while getting all tasks")
	require.NotEqual(t, all, nil)
}

func TestToggleTask(t *testing.T) {
	r := initTaskIter()
	s := initSliceForId()
	for _, task := range s {
		err := r.ToggleTask(task.Id)
		require.NoError(t, err, "error! Run debug")
		require.Equal(t, err, nil)
	}
}

func TestCreateTask(t *testing.T) {
	r := initTaskIter()
	s := initSliceForId()
	for _, task := range s {
		err := r.CreateTask(domain.TaskForms(task))
		require.NoError(t, err, "error! Run debug")
		require.Equal(t, err, nil)
	}
}

func TestGetTaskById(t *testing.T) {
	r := initTaskIter()
	s := initSliceForId()
	for _, task := range s {
		taska, err := r.GetTaskById(task.Id)

		require.NoError(t, err, "error! Run debug")
		require.NotEqual(t, taska, nil)
	}
}
