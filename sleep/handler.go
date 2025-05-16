package function

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var r *rand.Rand

func init() {

	r = rand.New(rand.NewSource(time.Now().Unix()))
}

// Handle a serverless request
// 1. When no headers are given, sleep for the environment variable: sleep_duration.
// 2. When an X-Sleep header is given, sleep for that amount of time.
// 3. When the X-Min-Sleep and X-Max-Sleep headers are given, sleep for a random amount
// of time between those two figures
func Handle(req []byte) string {

	if v := os.Getenv("Http_Path"); v == "/_/ready" {
		return "ok"
	}

	if minV, ok := os.LookupEnv("Http_X_Min_Sleep"); ok && len(minV) > 0 {
		if maxV, ok := os.LookupEnv("Http_X_Max_Sleep"); ok && len(maxV) > 0 {
			minSleep, _ := time.ParseDuration(minV)
			maxSleep, _ := time.ParseDuration(maxV)

			minMs := minSleep.Milliseconds()
			maxMs := maxSleep.Milliseconds()

			randMs := r.Int63n(maxMs-minMs) + minMs

			sleepDuration, _ := time.ParseDuration(fmt.Sprintf("%dms", randMs))

			log.Printf("Start sleep for: %fs\n", sleepDuration.Seconds())
			time.Sleep(sleepDuration)
			log.Printf("Sleep done for: %fs\n", sleepDuration.Seconds())
			return fmt.Sprintf("Slept for: %fs", sleepDuration.Seconds())
		}
	}

	sleepDuration := time.Second * 2

	if val, ok := os.LookupEnv("Http_X_Sleep"); ok && len(val) > 0 {
		sleepDuration, _ = time.ParseDuration(val)
	} else if val, ok := os.LookupEnv("sleep_duration"); ok && len(val) > 0 {
		sleepDuration, _ = time.ParseDuration(val)
	}

	log.Printf("Start sleep for: %fs\n", sleepDuration.Seconds())
	time.Sleep(sleepDuration)
	log.Printf("Sleep done for: %fs\n", sleepDuration.Seconds())

	return fmt.Sprintf("Slept for: %fs", sleepDuration.Seconds())
}
