package snowflake

import (
	"testing"
	"time"
)

func BenchmarkGenerate(b *testing.B) {
	sf := NewSnowflake(Setting{
		epoch:     time.Now(),
		machineID: 0,
	})
	for n := 0; n < b.N; n++ {
		_,_ = sf.NextID()
	}
}