package repositories

import (
	"database/sql"
	"todolist-service/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	query := "INSERT INTO tasks (start_date, plan_date, detail, status) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.DB.QueryRow(query, task.StartDate, task.PlanDate, task.Detail, task.Status).Scan(&task.ID)
}

func (r *TaskRepository) GetTasks() ([]models.Task, error) {
	rows, err := r.DB.Query("SELECT id, start_date, plan_date, detail, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.StartDate, &task.PlanDate, &task.Detail, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskByID(id int) (*models.Task, error) {
	var task models.Task
	query := "SELECT id, start_date, plan_date, detail, status FROM tasks WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&task.ID, &task.StartDate, &task.PlanDate, &task.Detail, &task.Status)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	query := "UPDATE tasks SET start_date = $1, plan_date = $2, detail = $3, status = $4 WHERE id = $5"
	_, err := r.DB.Exec(query, task.StartDate, task.PlanDate, task.Detail, task.Status, task.ID)
	return err
}

func (r *TaskRepository) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
