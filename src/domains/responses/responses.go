package responses

type ResponseCreated struct {
	Message   string `json:"message"`
	CreatedId int64  `json:"created_id"`
}
