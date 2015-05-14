package twitch

import "testing"

func TestChannels(t *testing.T) {
	client, err1 := NewClient(nil)
	if err1 != nil {
		t.Fatal(err1)
	}
	_, err2 := client.Channels.Get("aoisensi")
	if err2 != nil {
		t.Fatal(err2)
	}

}
