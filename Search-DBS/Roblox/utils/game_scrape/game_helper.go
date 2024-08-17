package game_scrape

import (
	"encoding/json"
	"net/http"
	"fmt"
)

// fetchBadgeIDs is responsible for fetching badge IDs associated with the user
func fetchBadgeIDs(gs *GameScraper) error {
	cursor := ""
	for {
		url := buildBadgeURL(gs.UserID, cursor)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var result struct {
			Data           []struct{ ID int64 `json:"id"` } `json:"data"`
			NextPageCursor string                             `json:"nextPageCursor"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return err
		}

		for _, badge := range result.Data {
			gs.BadgeIDs = append(gs.BadgeIDs, badge.ID)
			if len(gs.BadgeIDs) >= gs.GameLimit {
				return nil
			}
		}

		if result.NextPageCursor == "" {
			break
		}
		cursor = result.NextPageCursor
	}

	return nil
}

// getGame fetches the game information related to a badge ID
func getGame(gs *GameScraper, badgeID int64) error {
	url := buildBadgeInfoURL(badgeID)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var game Game
	if err := json.NewDecoder(resp.Body).Decode(&game); err != nil {
		return err
	}

	gs.mu.Lock()
	defer gs.mu.Unlock()
	if _, exists := gs.Games[game.Name]; !exists {
		gs.Games[game.Name] = struct{}{}
	}

	return nil
}

// mapKeysToSlice converts map keys to a slice of strings
func mapKeysToSlice(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// buildBadgeURL constructs the URL for fetching badges
func buildBadgeURL(userID int64, cursor string) string {
	return fmt.Sprintf("https://badges.roblox.com/v1/users/%d/badges?limit=100&sortOrder=Desc&cursor=%s", userID, cursor)
}

// buildBadgeInfoURL constructs the URL for fetching badge information
func buildBadgeInfoURL(badgeID int64) string {
	return fmt.Sprintf("https://badges.roblox.com/v1/badges/%d", badgeID)
}
