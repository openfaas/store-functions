package function

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	random          *rand.Rand
	defaultDuration time.Duration = time.Second * 2
)

func init() {
	random = rand.New(rand.NewSource(time.Now().Unix()))

	if val, ok := os.LookupEnv("sleep_duration"); ok && len(val) > 0 {
		var err error
		defaultDuration, err = time.ParseDuration(val)
		if err != nil {
			log.Fatalf("Error parsing sleep_duration environment variable: %v", err)
		}
	}
}

// Handle a serverless request
// 1. When no headers are given, sleep for the environment variable: sleep_duration.
// 2. When an X-Sleep header is given, sleep for that amount of time.
// 3. When the X-Min-Sleep and X-Max-Sleep headers are given, sleep for a random amount
// of time between those two figures
func Handle(w http.ResponseWriter, r *http.Request) {
	if minV := r.Header.Get("X-Min-Sleep"); len(minV) > 0 {
		if maxV := r.Header.Get("X-Max-Sleep"); len(maxV) > 0 {
			minSleep, err := time.ParseDuration(minV)
			if err != nil {
				log.Printf("Error parsing X-Min-Sleep header: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Error parsing X-Min-Sleep header: %v", err)
				return
			}
			maxSleep, err := time.ParseDuration(maxV)
			if err != nil {
				log.Printf("Error parsing X-Max-Sleep header: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Error parsing X-Max-Sleep header: %v", err)
				return
			}

			minMs := minSleep.Milliseconds()
			maxMs := maxSleep.Milliseconds()

			randMs := random.Int63n(maxMs-minMs) + minMs

			sleepDuration, _ := time.ParseDuration(fmt.Sprintf("%dms", randMs))

			log.Printf("Start sleep for: %fs\n", sleepDuration.Seconds())
			time.Sleep(sleepDuration)
			log.Printf("Sleep done for: %fs\n", sleepDuration.Seconds())

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Slept for: %fs", sleepDuration.Seconds())
			return
		}
	}

	sleepDuration := defaultDuration

	var err error
	if val := r.Header.Get("X-Sleep"); len(val) > 0 {
		sleepDuration, err = time.ParseDuration(val)
		if err != nil {
			log.Printf("Error parsing X-Sleep header: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error parsing X-Sleep header: %v", err)
			return
		}
	}

	log.Printf("Start sleep for: %fs\n", sleepDuration.Seconds())
	time.Sleep(sleepDuration)
	log.Printf("Sleep done for: %fs\n", sleepDuration.Seconds())

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Slept for: %fs", sleepDuration.Seconds())
}
