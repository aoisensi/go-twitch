package twitch

import "time"

type ChannelsService service

type Channel struct {
	Mature                       bool      `json:"mature"`
	Status                       string    `json:"status"`
	BroadcasterLaungage          string    `json:"broadcaster_laungage"`
	DisplayName                  string    `json:"display_name"`
	Game                         string    `json:"game"`
	Delay                        int       `json:"delay"`
	Laungage                     string    `json:"laungage"`
	ID                           int       `json:"_id"`
	Name                         string    `json:"name"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	Logo                         string    `json:"logo"`
	Banner                       string    `json:"banner"`
	VideoBanner                  string    `json:"video_banner"`
	Background                   string    `json:"background"`
	ProfileBanner                string    `json:"profile_banner"`
	ProfileBannerBackgroundColor string    `json:"profile_banner_background_color"`
	Partner                      bool      `json:"partner"`
	URL                          string    `json:"url"`
	Views                        int       `json:"views"`
	Followers                    int       `json:"followers"`
}

func (c *ChannelsService) GetSelf() (*Channel, error) {
	return c.get("https://api.twitch.tv/kraken/channel")
}

func (c *ChannelsService) Get(channel string) (*Channel, error) {
	return c.get("https://api.twitch.tv/kraken/channels/" + channel)
}

func (c *ChannelsService) get(url string) (*Channel, error) {
	var channel Channel
	if err := c.client.httpGet(url, nil, &channel); err != nil {
		return nil, err
	}
	return &channel, nil
}
