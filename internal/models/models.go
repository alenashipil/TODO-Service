package models

type Task struct {
	ID      string `json:"id"`
	Details string `json:"details"`
	Status  bool   `json:"status"`
}
