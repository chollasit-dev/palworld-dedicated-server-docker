package webhook

import (
	"encoding/json"
	"log"
	"net/http"
	"slices"
	"time"

	pal "palhook/internal/palworld"
)

const (
	URL           = "http://localhost:8212"
	MaxRetryCount = 5
	SleepSecond   = 5

	PrefixReadBody = "[Read]"
)

var (
	FailStatuses = []int{http.StatusBadRequest, http.StatusUnauthorized}
)

func main() {
	errCount := 0
	for {
		resp, err := http.Get(URL)
		defer resp.Body.Close()

		switch {
		case err != nil, slices.Contains(FailStatuses, resp.StatusCode):
			if errCount >= MaxRetryCount {
				// TODO: Dishook
				return
			}
			errCount += 1
		default:
			// TODO: JSON Decoder
			var body []byte
			var server *pal.GetPlayersResponse

			if _, err := resp.Body.Read(body); err != nil {
				log.Fatalf("%s %v\n", PrefixReadBody, err)
			}

			if err := json.Unmarshal(body, server); err != nil {
				log.Fatalf("%s %v\n", PrefixReadBody, err)
			}

			// TODO: Dishook
		}

		time.Sleep(time.Duration(SleepSecond) * time.Second)
	}
}
