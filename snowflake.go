package snowflake

import (
	"sync"
	"time"
)

const (
	bitLenMachineID = 10
	bitLenSequence  = 12
	sequenceMask = -1 ^ (-1 << bitLenSequence)
)

type Setting struct {
	epoch time.Time
	machineID uint64
}

type Snowflake struct {
	mutex       sync.Mutex
	epoch       time.Time
	elapsedTime int64
	machineID   uint64
	sequence    uint64
}

func NewSnowflake(st Setting) *Snowflake {
	sf := &Snowflake{
		epoch:       st.epoch,
		machineID:   st.machineID,
		sequence:    0,
	}
	return sf
}

func (sf *Snowflake) NextID() (uint64, error) {
	sf.mutex.Lock()
	defer sf.mutex.Unlock()
	curTime := getCurTime(sf.epoch)

	if sf.elapsedTime == curTime {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			for curTime <= sf.elapsedTime {
				curTime = getCurTime(sf.epoch)
			}
		}
	} else {
		sf.sequence = 0
	}
	sf.elapsedTime = curTime

	sfID := (uint64(sf.elapsedTime) << bitLenMachineID << bitLenSequence) | (sf.machineID << bitLenSequence) | sf.sequence

	return sfID, nil
}

func getCurTime(epoch time.Time) int64 {
	return time.Since(epoch).Milliseconds()
}