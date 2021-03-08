package main

import "testing"

func BenchmarkSetAdd1090Has(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(200)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
		b.SetParallelism(1800)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetAdd5050Has(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetAdd9010Has(b *testing.B) {
	var set = NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1800)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
		b.SetParallelism(200)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetRWAdd1090Has(b *testing.B) {
	var set = NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(200)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
		b.SetParallelism(1800)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetRWAdd5050Has(b *testing.B) {
	var set = NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkSetRW0Add9010Has(b *testing.B) {
	var set = NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1800)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1)
			}
		})
		b.SetParallelism(200)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}
