package model

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,min=5,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=50"`
}

type UpdateUserRequest struct {
	Name     *string `json:"name,omitempty" binding:"omitempty,min=5,max=30"`
	Email    *string `json:"email,omitempty" binding:"omitempty,email"`
	Password *string `json:"password,omitempty" binding:"omitempty,min=8,max=50"`
}

type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required,min=3,max=255"`
	Description string     `json:"description" binding:"required"`
	Status      TaskStatus `json:"status" binding:"required,oneof=pending done"`
}

type UpdateTaskRequest struct {
	Title       *string     `json:"title,omitempty" binding:"omitempty,min=3,max=255"`
	Description *string     `json:"description,omitempty" binding:"omitempty"`
	Status      *TaskStatus `json:"status,omitempty" binding:"omitempty,oneof=pending done"`
}
