package data

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

// LastUpdated reports the date of the latest commit on the repo, so the site
// shows when it last changed without anyone editing a date by hand. The result
// is cached for an hour; on error the last good value (or "") is returned.
const commitsURL = "https://api.github.com/repos/leeclza/portoleru/commits?per_page=1"

var (
	updMu      sync.Mutex
	updValue   string
	updFetched time.Time
	updClient  = &http.Client{Timeout: 2 * time.Second}
	wib        = time.FixedZone("WIB", 7*3600)
)

func LastUpdated() string {
	updMu.Lock()
	defer updMu.Unlock()
	if time.Since(updFetched) < time.Hour {
		return updValue
	}
	updFetched = time.Now()
	if t, err := fetchLastCommit(); err == nil {
		updValue = t.In(wib).Format("02 Jan 2006")
	}
	return updValue
}

func fetchLastCommit() (time.Time, error) {
	req, _ := http.NewRequest(http.MethodGet, commitsURL, nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	resp, err := updClient.Do(req)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()
	var commits []struct {
		Commit struct {
			Committer struct {
				Date time.Time `json:"date"`
			} `json:"committer"`
		} `json:"commit"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil || len(commits) == 0 {
		return time.Time{}, http.ErrMissingFile
	}
	return commits[0].Commit.Committer.Date, nil
}
