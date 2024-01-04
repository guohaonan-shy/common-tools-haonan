package _go

import "testing"

func Benchmark_ForRange(b *testing.B) {
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		result := make(chan struct{})
		go func() {
			for j := 0; j < 100; j++ {
				_ = j
			}
			result <- struct{}{}
		}()
		<-result
	}
}

func Benchmark_GoPool(b *testing.B) {
	b.ResetTimer()
	pool := NewGoPool[any, any](4, 100)

	defer pool.Close()

	for i := 1; i <= b.N; i++ {
		pool.SyncSubmit(func(any) any {
			for j := 0; j < 100; j++ {
				_ = j
			}
			return nil
		}, nil)
	}
}
