package clients

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"submission-sequencer-collector/config"

	log "github.com/sirupsen/logrus"
)

// Add this function to periodically dump memory profiles
func StartMemoryProfiling() {
	ticker := time.NewTicker(time.Duration(config.SettingsObj.MemoryProfilingInterval) * time.Minute)
	go func() {
		for range ticker.C {
			log.Infof("📊 Current goroutine count: %d", runtime.NumGoroutine())

			// Create a memory profile
			f, err := os.Create(fmt.Sprintf("memory_profile_%s.pprof",
				time.Now().Format("2006-01-02_15-04-05")))
			if err != nil {
				log.Errorf("❌ Could not create memory profile: %v", err)
				continue
			}

			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Errorf("❌ Could not write memory profile: %v", err)
			}
			f.Close()

			// Log memory stats
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			log.Infof("📈 Memory stats: Alloc=%v MiB, TotalAlloc=%v MiB, Sys=%v MiB, NumGC=%v",
				m.Alloc/1024/1024, m.TotalAlloc/1024/1024, m.Sys/1024/1024, m.NumGC)
		}
	}()
}
