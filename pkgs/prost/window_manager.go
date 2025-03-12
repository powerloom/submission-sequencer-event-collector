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
}

type WindowManager struct {
	activeWindows map[string]*EpochWindow // key: dataMarketAddress:epochID
	mu            sync.RWMutex
	done          chan struct{}
}

func newEpochWindowKey(dataMarketAddress string, epochID *big.Int) string {
	return fmt.Sprintf("%s:%s", dataMarketAddress, epochID.String())
}

func NewWindowManager() *WindowManager {
	return &WindowManager{
		activeWindows: make(map[string]*EpochWindow),
		done:          make(chan struct{}),
	}
}

func (wm *WindowManager) StartSubmissionWindow(ctx context.Context, dataMarketAddress string, epochID *big.Int, windowDuration time.Duration) error {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	key := newEpochWindowKey(dataMarketAddress, epochID)
	if _, exists := wm.activeWindows[key]; exists {
		return fmt.Errorf("submission window already active for epoch %s in market %s", epochID, dataMarketAddress)
	}

	window := &EpochWindow{
		EpochID:           epochID,
		DataMarketAddress: dataMarketAddress,
		StartTime:         time.Now(),
		WindowDuration:    windowDuration,
		Done:              make(chan struct{}),
	}

	// Create timer for batch preparation
	window.Timer = time.NewTimer(windowDuration)
	wm.activeWindows[key] = window

	// Start monitoring goroutine
	go func() {
		defer func() {
			window.Timer.Stop()
			close(window.Done)
			wm.removeWindow(dataMarketAddress, epochID)
		}()

		select {
		case <-window.Timer.C:
			batchCtx, batchCancel := context.WithTimeout(context.Background(), batchProcessingTimeout)
			defer batchCancel()

			if err := triggerBatchPreparation(batchCtx, dataMarketAddress, epochID, 0, 0); err != nil {
				log.Errorf("Failed to trigger batch preparation for epoch %s in market %s: %v",
					epochID, dataMarketAddress, err)
			}
		case <-ctx.Done():
			return
		case <-wm.done:
			return
		}
	}()

	log.Infof("Started submission window for epoch %s in market %s, duration: %v",
		epochID, dataMarketAddress, windowDuration)
	return nil
}

func (wm *WindowManager) removeWindow(dataMarketAddress string, epochID *big.Int) {
	wm.mu.Lock()
	defer wm.mu.Unlock()

	key := newEpochWindowKey(dataMarketAddress, epochID)
	delete(wm.activeWindows, key)
}

func (wm *WindowManager) Shutdown() {
	close(wm.done)

	wm.mu.RLock()
	defer wm.mu.RUnlock()

	// Wait for all windows to clean up
	for _, window := range wm.activeWindows {
		<-window.Done
	}
}
