package models



type ErrorResponse struct{
	Success bool   `json:"success"`
    Message string `json:"message"`
}

type SuccessResponse struct{
	Success bool	`json:"success"`
	Message string `json:"message"`
	Data UserData	`json:"data"`
	

}
