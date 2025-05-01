package response

type SprintTasksResponse struct {
	SprintTaskByID
	SprintID string `bun:"sprint_id" json:"sprint_id"`
}

type SprintTaskByID struct {
	ID         string `bun:"id" json:"id"`
	TaskName   string `bun:"task_name" json:"task_name"`
	TaskDetail string `bun:"task_detail" json:"task_detail"`
	AssignedTo string `bun:"assigned_to" json:"assigned_to"`
	Status     string `bun:"status" json:"status"`
	CreatedAt  int64  `bun:"created_at" json:"created_at"`
	UpdatedAt  int64  `bun:"updated_at" json:"updated_at"`
}
