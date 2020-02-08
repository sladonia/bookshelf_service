package responses

type ResponseCreated struct {
	Message   string `json:"message"`
	CreatedId int64  `json:"created_id"`
}

type ResponseDeleted struct {
	Message   string `json:"message"`
	DeletedId int64  `json:"deleted_id"`
}
