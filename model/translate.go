package model

type Translate struct {
	Error struct {
		NotFound          string `json:"notFound"`
		RecordNotFound    string `json:"recordNotFound"`
		RecordUpdateErr   string `json:"recordUpdateErr"`
		ParamsNotValid    string `json:"paramsNotValid"`
		UserNotFound      string `json:"userNotFound"`
		UserNotAuthorized string `json:"userNotAuthorized"`
		UserAlreadyExists string `json:"userAlreadyExists"`
	} `json:"error"`

	Status struct {
		NotFound          int `json:"notFound"`
		RecordNotFound    int `json:"recordNotFound"`
		RecordUpdateErr   int `json:"recordUpdateErr"`
		ParamsNotValid    int `json:"paramsNotValid"`
		UserNotFound      int `json:"userNotFound"`
		UserNotAuthorized int `json:"userNotAuthorized"`
		UserAlreadyExists int `json:"userAlreadyExists"`
	} `json:"status"`
}
