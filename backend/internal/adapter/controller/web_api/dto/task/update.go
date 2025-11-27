package task

import "github.com/andreychh/coopera-backend/internal/entity"

type UpdateTaskRequest struct {
	TaskID        int32   `json:"task_id" validate:"required,gt=0"`
	Status        *string `json:"status" validate:"omitempty,oneof=open assigned completed archived"`
	AssignedTo    *int32  `json:"assigned_to" validate:"omitempty,gt=0"`
	CurrentUserID int32   `json:"current_user_id" validate:"required,gt=0"`
}

func ToEntityUpdateTaskRequest(req *UpdateTaskRequest) *entity.Task {
	task := &entity.Task{
		ID: req.TaskID,
	}

	if req.Status != nil {
		status := entity.Status(*req.Status)
		task.Status = &status
	}

	if req.AssignedTo != nil {
		task.AssignedTo = req.AssignedTo
	}

	return task
}

type UpdateTaskResponse struct {
	ID          int32   `json:"id"`
	TeamID      int32   `json:"team_id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
	Points      int32   `json:"points"`
	Status      string  `json:"status"`
	AssignedTo  *int32  `json:"assigned_to,omitempty"`
	CreatedBy   int32   `json:"created_by"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
}

func ToUpdateTaskResponse(task *entity.Task) *UpdateTaskResponse {
	resp := &UpdateTaskResponse{
		ID:        task.ID,
		TeamID:    task.TeamID,
		Title:     task.Title,
		Points:    task.Points,
		Status:    task.Status.String(),
		CreatedBy: task.CreatedBy,
		CreatedAt: task.CreatedAt.Format("2006-01-02T15:04:05Z"),
	}

	if task.Description != nil {
		resp.Description = task.Description
	}

	if task.AssignedTo != nil {
		resp.AssignedTo = task.AssignedTo
	}

	if task.UpdatedAt != nil {
		ts := task.UpdatedAt.Format("2006-01-02T15:04:05Z")
		resp.UpdatedAt = &ts
	}

	return resp
}
