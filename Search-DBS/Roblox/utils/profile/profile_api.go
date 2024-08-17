package profile

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func (ps *ProfileScraper) Bio() (string, error) {
    url := fmt.Sprintf("https://users.roblox.com/v1/users/%d", ps.UserID)
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var result struct {
        Description string `json:"description"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }

    return result.Description, nil
}

func (ps *ProfileScraper) GetProfileInfo() (map[string]interface{}, error) {
    url := fmt.Sprintf("https://users.roblox.com/v1/users/%d", ps.UserID)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var result struct {
        Created              string  `json:"created"`
        IsBanned             bool    `json:"isBanned"`
        ExternalAppDisplayName *string `json:"externalAppDisplayName"`
        HasVerifiedBadge     bool    `json:"hasVerifiedBadge"`
        ID                   int64   `json:"id"`
        Name                 string  `json:"name"`
        DisplayName          string  `json:"displayName"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, err
    }

    profileInfo := map[string]interface{}{
        "created":                result.Created,
        "isBanned":               result.IsBanned,
        "externalAppDisplayName": result.ExternalAppDisplayName,
        "hasVerifiedBadge":       result.HasVerifiedBadge,
        "id":                     result.ID,
        "name":                   result.Name,
        "displayName":            result.DisplayName,
    }

    return profileInfo, nil
}

func (ps *ProfileScraper) FollowerCount() (int, error) {
    url := fmt.Sprintf("https://friends.roblox.com/v1/users/%d/followers/count", ps.UserID)
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var result struct {
        Count int `json:"count"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, err
    }

    return result.Count, nil
}

func (ps *ProfileScraper) FollowingCount() (int, error) {
    url := fmt.Sprintf("https://friends.roblox.com/v1/users/%d/followings/count", ps.UserID)
    resp, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    var result struct {
        Count int `json:"count"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return 0, err
    }

    return result.Count, nil
}

func (ps *ProfileScraper) GetTotalFriendsCount() (int, error) {
    count := 0
    url := fmt.Sprintf("https://friends.roblox.com/v1/users/%d/friends", ps.UserID)

    for {
        resp, err := http.Get(url)
        if err != nil {
            return 0, err
        }
        defer resp.Body.Close()

        var result struct {
            Data        []map[string]interface{} `json:"data"`
            NextPageURL string                    `json:"nextPageCursor"`
        }
        if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
            return 0, err
        }

        count += len(result.Data)

        if result.NextPageURL == "" {
            break
        }

        url = fmt.Sprintf("https://friends.roblox.com/v1/users/%d/friends?pageCursor=%s", ps.UserID, result.NextPageURL)
    }

    return count, nil
}
