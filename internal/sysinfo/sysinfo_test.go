package sysinfo

import (
	"testing"

	"github.com/kourtnet/dummyfetch/internal/entities"
)

func fetchNArgs(b *testing.B, argNames []string) {
	for b.Loop() {
		Fetch(argNames)
		b.StopTimer()

		for _, v := range argNames {
			arg := entities.ArgsMap[v]
			arg.Contents = ""
			entities.ArgsMap[v] = arg
		}

		b.StartTimer()
	}
}

func BenchmarkFetch1Arg(b *testing.B) {
	argNames := []string{"kernel"}
	fetchNArgs(b, argNames)
}

func BenchmarkFetch5Args(b *testing.B) {
	argNames := []string{"kernel", "os", "terminal", "uptime", "shell"}
	fetchNArgs(b, argNames)
}

func BenchmarkFetch1ArgRepeatedly(b *testing.B) {
	argNames := []string{"kernel"}
	for b.Loop() {
		Fetch(argNames)
	}
}
