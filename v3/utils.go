package twitch

type Error struct {
	Code    string `json:"error"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

type service struct {
	client *Client
}
