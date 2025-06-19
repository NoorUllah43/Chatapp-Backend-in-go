package models



type ErrorResponse struct{
	Success bool   `json:"success"`
    Message string `json:"message"`
}

type SuccessResponse struct{
	Success bool	`json:"success"`
	Message string `json:"message"`
	UserData UserData	`json:"userData"`
	

}
type Users struct {
	Success bool	`json:"success"`
	Message string `json:"message"`
	Users []UserData	`json:"users"`
}
