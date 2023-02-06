package repos

import (
	"Portfolio_Nodes/domain"
	"database/sql"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (r *TaskRepo) FetchTasks() ([]domain.Task, error) {
	var tasks []domain.Task

	query := `SELECT id, name, priority,done FROM tasks ORDER BY id;`
	raws, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	for raws.Next() {
		var task domain.Task
		err = raws.Scan(&task.Id, &task.Name, &task.Priority, &task.Done)
		tasks = append(tasks, domain.Task{
			Id:       task.Id,
			Name:     task.Name,
			Priority: task.Priority,
			Done:     task.Done,
		})
	}
	return tasks, nil
}

func (r *TaskRepo) FetchTaskById(taskId int) (domain.Task, error) {
	var task domain.Task
	query := `SELECT * FROM tasks WHERE id=$1`
	raw, err := r.db.Query(query, taskId)
	if err != nil {
		return task, err
	}
	raw.Scan(&task.Id, &task.Name, &task.Done, &task.Priority)
	return task, nil
}

func (r *TaskRepo) InsertTask(task domain.Task) error {
	query := "INSERT INTO tasks (name, done, priority) VALUES ($1,$2,$3)"
	_, err := r.db.Exec(query, task.Name, task.Done, task.Priority)
	return err
}

func (r *TaskRepo) UpdateTask(taskId int) error {
	query := "UPDATE tasks SET done=NOT done WHERE id=$1"
	_, err := r.db.Exec(query, taskId)
	if err != nil {
		return err
	}
	return err
}
