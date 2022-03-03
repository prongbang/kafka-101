package command

type PostCommand struct {
	Message string `json:"message"`
}

type ReplyCommand struct {
	Message string `json:"message"`
}
