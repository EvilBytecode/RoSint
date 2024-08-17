package gambling_sites

import (
	"encoding/json"
	"fmt"
)

// bloxflip fetches data from the bloxflip API and updates Stats
func (gs *GambleScraper) bloxflip() error {
	url := fmt.Sprintf("https://api.bloxflip.com/user/lookup/%d", gs.UserID)
	body, err := gs.getRequest(url)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	//fmt.Printf("Response Body: %s\n", body)

	var response map[string]interface{}
	if err := json.Unmarshal([]byte(body), &response); err != nil {
		return fmt.Errorf("JSON unmarshal error: %w", err)
	}

	if success, ok := response["success"].(bool); !ok || !success {
		return nil
	}

	delete(response, "username")
	delete(response, "success")
	gs.Stats["bloxflip"] = response

	return nil
}
