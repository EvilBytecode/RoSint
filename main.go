package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "EByte-OSINT/Search-DBS/Roblox/utils/gambling_sites"
    "EByte-OSINT/Search-DBS/Roblox/utils/profile"
    "EByte-OSINT/Search-DBS/Roblox/utils/game_scrape"
)

func main() {
    // Replace with the actual user ID you want to scrape
    userID := int64(4746081815) // Example user ID

    gambleScraper := gambling_sites.NewGambleScraper(userID)
    profileScraper := profile.NewProfileScraper(userID)
    gameScraper := game_scrape.NewGameScraper(userID, 10) // Adjust the limit as necessary

    gamblingStats, err := gambleScraper.Run()
    if err != nil {
        log.Fatalf("Error scraping gambling stats: %v", err)
        return
    }

    profileData, err := profileScraper.ScrapeBio()
    if err != nil {
        log.Fatalf("Error scraping profile data: %v", err)
        return
    }

    onlineFriendsCount, err := profileScraper.GetTotalFriendsCount()
    if err != nil {
        log.Fatalf("Error fetching online friends count: %v", err)
        return
    }

    games, err := gameScraper.Run()
    if err != nil {
        log.Fatalf("Error scraping game data: %v", err)
        return
    }

    usernameHistory, err := profileScraper.GetNames()
    if err != nil {
        log.Fatalf("Error fetching username history: %v", err)
        return
    }

    robloxInfo := map[string]interface{}{
        "username_history":         usernameHistory,
        "user_info":              profileData,
        "games":                    games,
        "bloxflip_info":   gamblingStats,
        "total_friends":     onlineFriendsCount,
    }

    file, err := os.Create("roblox_info.json")
    if err != nil {
        log.Fatalf("Error creating file: %v", err)
        return
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    if err := encoder.Encode(robloxInfo); err != nil {
        log.Fatalf("Error encoding JSON: %v", err)
    }

    fmt.Println("Data saved to roblox_info.json")
}
