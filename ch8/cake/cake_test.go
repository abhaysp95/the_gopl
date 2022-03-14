package cake_test

import (
	"testing"
	"time"

	"the_gopl/ch8/cake"
)

var defaults = cake.Shop {
	Verbose:       testing.Verbose(),
	Cakes:         20,
	BakeTime:      10 * time.Millisecond,
	NumIcers:      1,
	IceTime:       10 * time.Millisecond,
	InscribeTime:  10 * time.Millisecond,
}

func Benchmark(b *testing.B) {
	// Baseline: one baker, one icer, one inscriber
	// Each step takes exactly 10ms. No buffers
	cakeshop := defaults
	cakeshop.Work(b.N)
}

func BenchmarkBuffers(b *testing.B) {
	// Adding buffers
	cakeshop := defaults
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N)
}

func BenchmarkVariable(b *testing.B) {
	// Adding variability to rate of each step
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeStdDev / 4
	cakeshop.Work(b.N)
}

func BenchmarkVariableBuffers(b *testing.B) {
	// Adding channel buffers reduces
	cakeshop := defaults
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N)
}

func BenchmarkSlowIcing(b *testing.B) {
	// Making the middle stage slower
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.Work(b.N)
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	// Adding more icing cooks reduces the cost of icing
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.Work(b.N)
}

func BenchmarkSlowIcingManyIcersVariableBuffers(b *testing.B) {
	cakeshop := defaults
	cakeshop.IceTime = 50 * time.Millisecond
	cakeshop.NumIcers = 5
	cakeshop.BakeStdDev = cakeshop.BakeTime / 4
	cakeshop.IceStdDev = cakeshop.IceTime / 4
	cakeshop.InscribeStdDev = cakeshop.InscribeTime / 4
	cakeshop.BakeBuf = 10
	cakeshop.IceBuf = 10
	cakeshop.Work(b.N)
}
