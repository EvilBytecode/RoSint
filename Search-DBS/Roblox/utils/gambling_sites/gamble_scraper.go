package gambling_sites


// GambleScraper is responsible for scraping gambling site data for a user
type GambleScraper struct {
	UserID  int64
	Stats   map[string]interface{}
	Headers map[string]string
}

// NewGambleScraper initializes a new GambleScraper instance
func NewGambleScraper(userID int64) *GambleScraper {
	return &GambleScraper{
		UserID: userID,
		Stats:  make(map[string]interface{}),
		Headers: map[string]string{
			"User-Agent": "Mozilla/5.0 (X11; CrOS armv7l 13597.84.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.5112.105 Safari/537.36",
		},
	}
}

// Run executes the scraping methods and returns the stats
func (gs *GambleScraper) Run() (map[string]interface{}, error) {
	if err := gs.bloxflip(); err != nil {
		return nil, err
	}
	return gs.Stats, nil
}
