package models

type Task struct {
	ID        int    `json:"id"`
	StartDate string `json:"start_date"`
	PlanDate  string `json:"plan_date"`
	Detail    string `json:"detail"`
	Status    bool   `json:"status"`
}
