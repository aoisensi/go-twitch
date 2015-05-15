package twitch

import "time"

type VideosService service

type Video struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	BroadcastID   int       `json:"broadcast_id"`
	Status        string    `json:"status"`
	ID            string    `json:"_id"`
	TagList       string    `json:"tag_list"`
	RecordedAt    time.Time `json:"recorded_at"`
	Game          string    `json:"game"`
	Length        int       `json:"length"`
	Preview       string    `json:"preview"`
	URL           string    `json:"url"`
	Views         int       `json:"views"`
	BroadcastType string    `json:"broadcast_type"`
	Channel       struct {
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
	} `json:"channel"`
}

func (s *VideosService) Get(id string) (*Video, error) {
	url := "https://api.twitch.tv/kraken/videos/" + id
	var video Video
	if err := s.client.httpGet(url, nil, &video); err != nil {
		return nil, err
	}
	return &video, nil
}
