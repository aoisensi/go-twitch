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

type Preview struct {
	Small    string `json:"small"`
	Medium   string `json:"medium"`
	Large    string `json:"large"`
	Template string `json:"template"`
}
