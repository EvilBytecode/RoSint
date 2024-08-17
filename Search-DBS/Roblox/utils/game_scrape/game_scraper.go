package game_scrape

import (
	"fmt"
	"sync"
)

// FetchBadgeIDs fetches badge IDs associated with the user
func (gs *GameScraper) FetchBadgeIDs() error {
	return fetchBadgeIDs(gs)
}

// GetGame retrieves the game associated with a given badge ID
func (gs *GameScraper) GetGame(badgeID int64) error {
	return getGame(gs, badgeID)
}

// FetchGames fetches games related to the badge IDs using concurrent requests
func (gs *GameScraper) FetchGames() {
	var wg sync.WaitGroup

	for _, badgeID := range gs.BadgeIDs {
		go func(id int64) {
			defer wg.Done()
			if err := gs.GetGame(id); err != nil {
				fmt.Printf("[-] Error fetching game with badge ID %d: %v\n", id, err)
			}
		}(badgeID)
	}

	wg.Wait()
}


// Run executes the game scraping process and returns the list of unique game names
func (gs *GameScraper) Run() ([]string, error) {
	if err := gs.FetchBadgeIDs(); err != nil {
		return nil, err
	}
	gs.FetchGames()
	return mapKeysToSlice(gs.Games), nil
}
