package response

type SprintPlanResponse struct {
	ID         string `bun:"id" json:"id"`
	SprintName string `bun:"sprint_name" json:"sprint_name"`
	TeamID     string `bun:"team_id" json:"team_id"`
	StartDate  int64  `bun:"start_date" json:"start_date"`
	EndDate    int64  `bun:"end_date" json:"end_date"`
	Status     string `bun:"status" json:"status"`
	CreatedAt  int64  `bun:"created_at" json:"created_at"`
	UpdatedAt  int64  `bun:"updated_at" json:"updated_at"`
}
