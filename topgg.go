package topgg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jonas747/dshardmanager"
)

// BotStatsURL is url for sending stats
const BotStatsURL = "https://top.gg/api/bots/%d/stats"

// SendStats sends bot stats to top.gg
func SendStats(manager *dshardmanager.Manager, token string) {
	data := make(map[string]interface{})
	data["server_count"] = manager.GetFullStatus().NumGuilds

	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(BotStatsURL, manager.Session(0).State.User.ID),
		bytes.NewBuffer(b),
	)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Println(resp.Status)
	}
}
