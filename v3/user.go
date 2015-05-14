package twitch

import "time"

type UsersService service

type User struct {
	Type          string          `json:"type"`
	Name          string          `json:"name"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	Logo          string          `json:"logo"`
	ID            int             `json:"_id"`
	DisplayName   string          `json:"display_name"`
	EMail         string          `json:"email"`
	Partnerd      bool            `json:"partnerd"`
	Bio           string          `json:"bio"`
	Notifications map[string]bool `json:"notifications"`
}

func (s *UsersService) Get(user string) (*User, error) {
	return s.get("https://api.twitch.tv/kraken/users/test_user1/" + user)
}

func (s *UsersService) GetSelf() (*User, error) {
	return s.get("https://api.twitch.tv/kraken/user")
}

func (s *UsersService) get(url string) (*User, error) {
	var user User
	if err := s.client.httpGet(url, &user); err != nil {
		return nil, err
	}
	return &user, nil
}
