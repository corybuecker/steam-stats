package configuration

import (
	"testing"

	"github.com/corybuecker/steam-stats/test"
)

func TestConfiguration(t *testing.T) {
	fakeDatabase := test.FakeDatabase{}

	fakeDatabase.Rows = map[string]interface{}{
		"steamApiKey":     "1",
		"steamId":         "1",
		"giantBombApiKey": "1",
	}

	c := Configuration{}
	if err := c.Load(&fakeDatabase); err != nil {
		t.Error(err)
	}
	if c.SteamID != "1" {
		t.Errorf("expected %s, got %s", "1", c.SteamID)
	}
	if c.GiantBombAPIKey != "1" {
		t.Errorf("expected %s, got %s", "1", c.SteamID)
	}
	if c.SteamAPIKey != "1" {
		t.Errorf("expected %s, got %s", "1", c.SteamID)
	}
}
