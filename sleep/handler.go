package function

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Handle a serverless request
func Handle(req []byte) string {

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
