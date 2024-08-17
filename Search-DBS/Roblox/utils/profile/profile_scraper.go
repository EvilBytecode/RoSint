package profile

import (
    "fmt"
    "strings"
	"net/http"
    "encoding/json"
)

type ProfileScraper struct {
    UserID int64
    Stats  map[string]interface{}
}

func NewProfileScraper(userID int64) *ProfileScraper {
    return &ProfileScraper{
        UserID: userID,
        Stats:  make(map[string]interface{}),
    }
}

func (ps *ProfileScraper) GetNames() ([]string, error) {
    url := fmt.Sprintf("https://users.roblox.com/v1/users/%d/username-history?limit=100&sortOrder=Desc", ps.UserID)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Data []struct {
            Name string `json:"name"`
        } `json:"data"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    var names []string
    for _, info := range result.Data {
        names = append(names, info.Name)
    }

    return names, nil
}

func (ps *ProfileScraper) ScrapeData(count int, words []string) {
    word := words[count]
    scrapeDict := map[string]interface{}{
        "age": map[string]interface{}{
            "length": 2, "value": word, "number": true,
        },
        "gender": []map[string]interface{}{
            {"length": 999, "value": "male", "number": false, "has": []string{"he", "him", "his"}},
            {"length": 999, "value": "female", "number": false, "has": []string{"she", "her", "hers"}},
        },
    }

    // Ensure that count-1 is only accessed when count > 0
    if count > 0 {
        scrapeDict["discord"] = []map[string]interface{}{
            {"length": 4, "value": fmt.Sprintf("%s#%s", words[count-1], word), "number": true},
            {"length": 999, "value": word, "number": false, "end": map[string]interface{}{"number": true, "length": 4}},
        }
    }

    for name, requirements := range scrapeDict {
        switch req := requirements.(type) {
        case map[string]interface{}:
            if response := CheckWord(name, word, req); len(response) > 0 {
                for k, v := range response {
                    ps.Stats[k] = v
                }
            }
        case []map[string]interface{}:
            for _, requirement := range req {
                if response := CheckWord(name, word, requirement); len(response) > 0 {
                    for k, v := range response {
                        ps.Stats[k] = v
                    }
                }
            }
        }
    }
}

func (ps *ProfileScraper) ScrapeBio() (map[string]interface{}, error) {
    bio, err := ps.Bio()
    if err != nil {
        return nil, err
    }

    words := strings.Split(bio, " ")
    for count := range words {
        ps.ScrapeData(count, words)
    }

    // Add follower, following counts, and online friends count to stats
    followerCount, err := ps.FollowerCount()
    if err != nil {
        return nil, err
    }
    followingCount, err := ps.FollowingCount()
    if err != nil {
        return nil, err
    }

    profileInfo, err := ps.GetProfileInfo()
    if err != nil {
        return nil, err
    }

    ps.Stats["followers_count"] = followerCount
    ps.Stats["following_count"] = followingCount
    ps.Stats["profile_info"] = profileInfo

    return ps.Stats, nil
}
