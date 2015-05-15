package twitch

import "net/url"

type GamesService service

type Game struct {
	Name        string  `json:"name"`
	Box         Preview `json:"box"`
	Logo        Preview `json:"logo"`
	Id          int     `json:"_id"`
	GiantbombID int     `json:"giantbomb_id"`
}

type GamesTop struct {
	Total int `json:"_total"`
	Top   struct {
		Game     Game `json:"game"`
		Viewers  int  `json:"viewers"`
		Channels int  `json:"channels"`
	} `json:"top"`
}

type GamesGetTopParams url.Values

func (p GamesGetTopParams) SetLimit(limit int) {
	url.Values(p).Set("limit", string(limit))
}

func (p GamesGetTopParams) SetOffset(offset int) {
	url.Values(p).Set("offset", string(offset))
}

func (s *GamesService) GetTop(params GamesGetTopParams) (*GamesTop, error) {
	ep := "https://api.twitch.tv/kraken/games/top"
	var top GamesTop
	if err := s.client.httpGet(ep, url.Values(params), &top); err != nil {
		return nil, err
	}
	return &top, nil
}
