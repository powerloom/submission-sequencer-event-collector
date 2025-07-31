package prost

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type EpochWindow struct {
	EpochID           *big.Int
	DataMarketAddress string
	StartTime         time.Time
	WindowDuration    time.Duration
	Timer             *time.Timer
	Done              chan struct{}
	StartBlockNum     int64 // Track block number when epoch was released
	EndBlockNum       int64 // Will be set when window expires
}

type WindowManager struct {
	activeWindows map[string]*EpochWindow // key: dataMarketAddress:epochID
	mu            sync.RWMutex
	done          chan struct{}

	// Add semaphore to limit concurrent windows
	windowSemaphore chan struct{}
	maxWindows      int
}

func newEpochWindowKey(dataMarketAddress string, epochID *big.Int) string {
	return fmt.Sprintf("%s:%s", dataMarketAddress, epochID.String())
}

func NewWindowManager() *WindowManager {
	maxWindows := 100 // Adjust based on your system's capacity
	return &WindowManager{
		activeWindows:   make(map[string]*EpochWindow),
		done:            make(chan struct{}),
		windowSemaphore: make(chan struct{}, maxWindows),
		maxWindows:      maxWindows,
	}
}

// GetActiveWindowCount returns the current number of active windows
func (wm *WindowManager) GetActiveWindowCount() int {
	wm.mu.RLock()
	defer wm.mu.RUnlock()
	return len(wm.activeWindows)
}

// StartSubmissionWindow starts a new submission window in its own managed goroutine
func (wm *WindowManager) StartSubmissionWindow(ctx context.Context, dataMarketAddress string, epochID *big.Int, windowDuration time.Duration, startBlockNum int64) error {
	// Quick check if window already exists - this needs the lock
	wm.mu.RLock()
	key := newEpochWindowKey(dataMarketAddress, epochID)
	exists := false
	if _, exists = wm.activeWindows[key]; exists {
		wm.mu.RUnlock()
		return fmt.Errorf("❌ submission window already active for epoch %s in market %s", epochID, dataMarketAddress)
	}
	wm.mu.RUnlock()

	// Try to acquire a semaphore slot with timeout
	select {
	case wm.windowSemaphore <- struct{}{}:
		// Got permission to proceed
		log.Infof("🎫 Acquired semaphore to open submission windown monitor for epoch %s in market %s", epochID, dataMarketAddress)
	case <-time.After(1 * time.Second):
		// Couldn't acquire semaphore in time
		log.Warnf("🚫 Failed to acquire semaphore to open submission window monitor for epoch %s in market %s", epochID, dataMarketAddress)
		return fmt.Errorf("❌ too many active windows (%d), refusing to create more for epoch %s in market %s",
			wm.GetActiveWindowCount(), epochID, dataMarketAddress)
	}

	// Create a new root context for this window's lifecycle
	windowCtx, windowCancel := context.WithCancel(context.Background())

	// Start a managed goroutine for the window creation and monitoring
	go func() {
		// Always release the semaphore when done
		defer func() {
			<-wm.windowSemaphore
			log.Infof("🎫 Released semaphore: submission window monitor for epoch %s in market %s", epochID, dataMarketAddress)
		}()

		// Recover from panics
		defer func() {
			if r := recover(); r != nil {
				log.Errorf("💥 Panic in submission window setup for epoch %s in market %s: %v",
					epochID, dataMarketAddress, r)
			}
		}()

		// Double-check that window wasn't created while we were waiting for semaphore
		wm.mu.Lock()
		if _, exists := wm.activeWindows[key]; exists {
			wm.mu.Unlock()
			windowCancel() // Cancel the context since we're not using it
			log.Warnf("⚠️ Window for epoch %s in market %s was created by another goroutine",
				epochID, dataMarketAddress)
			return
		}

		// Create the window object
		window := &EpochWindow{
			EpochID:           epochID,
			DataMarketAddress: dataMarketAddress,
			StartTime:         time.Now(),
			WindowDuration:    windowDuration,
			Done:              make(chan struct{}),
			StartBlockNum:     startBlockNum,
		}

		// Create timer for batch preparation
		window.Timer = time.NewTimer(windowDuration)
		wm.activeWindows[key] = window
		activeCount := len(wm.activeWindows)
		wm.mu.Unlock()

		log.Infof("🚀 Started submission window for epochID %s, data market %s, duration: %.2f seconds (active windows: %d)",
			epochID, dataMarketAddress, windowDuration.Seconds(), activeCount)

		// Create a channel to signal when monitoring is complete
		monitorDone := make(chan struct{})

		go func() {
			defer close(monitorDone)
			defer windowCancel()

			log.Infof("👀 Monitoring goroutine started for epoch %s in market %s",
				epochID, dataMarketAddress)

			defer func() {
				log.Infof("🧹 Cleanup triggered for submission window processing for epoch %s in market %s",
					epochID, dataMarketAddress)
				window.Timer.Stop()
				close(window.Done)
				wm.removeWindow(dataMarketAddress, epochID)

				// Log active window count for monitoring
				log.Infof("📊 Active windows remaining: %d", wm.GetActiveWindowCount())

				// Recover from panics
				if r := recover(); r != nil {
					log.Errorf("💥 Panic in submission window processing for epoch %s in market %s: %v",
						epochID, dataMarketAddress, r)
				}
			}()

			log.Infof("⏰ Waiting for timer to expire for epoch %s in market %s (duration: %v)",
				epochID, dataMarketAddress, windowDuration)

			select {
			case <-window.Timer.C:
				log.Infof("⌛ Timer expired for epoch %s in market %s", epochID, dataMarketAddress)

				// Get current block number when window expires
				// Create a new context with timeout for getting block number
				blockCtx, blockCancel := context.WithTimeout(context.Background(), 10*time.Second)
				currentBlock, err := RPCHelper.BlockNumber(blockCtx)
				blockCancel()

				if err != nil {
					log.Errorf("❓ Failed to get current block number for epoch %s in market %s: %v",
						epochID, dataMarketAddress, err)
					return
				}
				window.EndBlockNum = int64(currentBlock)

				log.Infof("🪟 Window for epoch %s in market %s begin at block %d, duration: %v ended at block %d",
					epochID, dataMarketAddress, window.StartBlockNum, windowDuration, window.EndBlockNum)

				// Create a fresh context with appropriate timeout for batch processing
				batchCtx, batchCancel := context.WithTimeout(context.Background(), batchProcessingTimeout)
				defer batchCancel()

				log.Infof("🔄 Starting batch preparation for epoch %s in market %s", epochID, dataMarketAddress)
				if err := triggerBatchPreparation(batchCtx, dataMarketAddress, epochID, window.StartBlockNum, window.EndBlockNum); err != nil {
					log.Errorf("❌ Failed to trigger batch preparation for epoch %s in market %s: %v",
						epochID, dataMarketAddress, err)
				} else {
					log.Infof("✅ Completed batch preparation for epoch %s in market %s", epochID, dataMarketAddress)
				}
			case <-windowCtx.Done():
				log.Infof("🛑 Window context canceled for epoch %s in market %s",
					epochID, dataMarketAddress)
				return
			case <-wm.done:
				log.Infof("🛑 Window manager shutdown signal received for epoch %s in market %s",
					epochID, dataMarketAddress)
				return
			}
		}()

		// Set a safety timeout to ensure the monitoring goroutine doesn't run forever
		// This is a fallback in case something goes wrong with the timer
		safetyTimeout := windowDuration + 5*time.Minute

		select {
		case <-monitorDone:
			// Monitoring completed normally
			log.Infof("✅ Monitoring completed normally for batch preparation for epoch %s in market %s",
				epochID, dataMarketAddress)
		case <-time.After(safetyTimeout):
			// Safety timeout - force cancel the window context
			log.Warnf("⏱️ Safety timeout triggered for batch preparation for epoch %s in market %s after %v",
				epochID, dataMarketAddress, safetyTimeout)
			windowCancel()

			// Wait a short time for cleanup
			select {
			case <-monitorDone:
				log.Infof("🧹 Monitoring goroutine cleaned up after safety timeout for batch preparation for epoch %s in market %s",
					epochID, dataMarketAddress)
			case <-time.After(5 * time.Second):
				log.Errorf("⚠️ Monitoring goroutine failed to clean up after safety timeout for epoch %s in market %s",
					epochID, dataMarketAddress)

				// Force cleanup as a last resort
				wm.mu.Lock()
				if window, exists := wm.activeWindows[key]; exists {
					window.Timer.Stop()
					close(window.Done)
					delete(wm.activeWindows, key)
					log.Warnf("🔨 Forced cleanup of window for epoch %s in market %s",
						epochID, dataMarketAddress)
				}
				wm.mu.Unlock()
			}
		}
	}()

	return nil
}

func (wm *WindowManager) removeWindow(dataMarketAddress string, epochID *big.Int) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	key := newEpochWindowKey(dataMarketAddress, epochID)
	delete(wm.activeWindows, key)
	log.Infof("🗑️ Removed window for epoch %s in market %s from active windows map", epochID, dataMarketAddress)
}

// CleanupStaleWindows forces cleanup of windows that have been active for too long
func (wm *WindowManager) CleanupStaleWindows(maxAge time.Duration) int {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	now := time.Now()
	count := 0

	for key, window := range wm.activeWindows {
		// If a window has been active for longer than maxAge, force cleanup
		if now.Sub(window.StartTime) > maxAge {
			log.Warnf("⏱️ Forcing cleanup of stale window for epoch %s in market %s (age: %v)",
				window.EpochID, window.DataMarketAddress, now.Sub(window.StartTime))

			window.Timer.Stop()
			close(window.Done)
			delete(wm.activeWindows, key)
			count++
		}
	}

	if count > 0 {
		log.Infof("🧹 Cleaned up %d stale windows", count)
	}

	return count
}

func (wm *WindowManager) Shutdown() {
	log.Info("🛑 Shutting down window manager")
	close(wm.done)

	// Wait with timeout for all windows to clean up
	deadline := time.After(30 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-deadline:
			// Force cleanup any remaining windows
			wm.mu.Lock()
			remaining := len(wm.activeWindows)
			if remaining > 0 {
				log.Warnf("🔨 Forcing cleanup of %d remaining windows during shutdown", remaining)
				for key, window := range wm.activeWindows {
					window.Timer.Stop()
					close(window.Done)
					delete(wm.activeWindows, key)
				}
			}
			wm.mu.Unlock()
			log.Info("✅ Window manager shutdown complete")
			return
		case <-ticker.C:
			wm.mu.RLock()
			count := len(wm.activeWindows)
			wm.mu.RUnlock()

			if count == 0 {
				log.Info("✅ All windows cleaned up, window manager shutdown complete")
				return
			}
			log.Infof("⏳ Waiting for %d windows to clean up during shutdown", count)
		}
	}
}
