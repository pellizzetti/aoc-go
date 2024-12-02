// Code generated by aocgen; DO NOT EDIT.
package year2024

import (
	"testing"

	"aocgen/pkg/aoc"
)

func Benchmark2024Day01(b *testing.B) {
	Init()
	input := aoc.TestInput(2024, 1)
	p := aoc.NewPuzzle(2024, 1)
	b.Run("PartA", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartA(input)
		}
	})
	b.Run("PartB", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartB(input)
		}
	})
}
