package game_scrape

import "sync"
// Game represents a Roblox game with its ID and name
type Game struct {
	BadgeID int64  `json:"id"`
	ID      int64  `json:"id"`
	Name    string `json:"name"`
}

// GameScraper is responsible for scraping games related to a Roblox user
type GameScraper struct {
	UserID    int64
	GameLimit int
	BadgeIDs  []int64
	Games     map[string]struct{} // Use a map to ensure unique game names
	mu        sync.Mutex
}

// NewGameScraper initializes a new GameScraper instance
func NewGameScraper(userID int64, gameLimit int) *GameScraper {
	return &GameScraper{
		UserID:    userID,
		GameLimit: gameLimit,
		Games:     make(map[string]struct{}),
	}
}