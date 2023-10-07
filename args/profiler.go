package args

import (
	"log/slog"
	"os"
	"runtime/pprof"
)

func Profile() func() {
	fileCpu, err := os.Create("profilerLogs/cpu.pprof")
	if err != nil {
		slog.Error("Error creating file for cpu profiler: " + err.Error())
	}
	pprof.StartCPUProfile(fileCpu)

	fileMem, err := os.Create("profilerLogs/memory.pprof")
	if err != nil {
		slog.Error("Error creating file for memory profiler: " + err.Error())
	}
	pprof.WriteHeapProfile(fileMem)

	return func() {
		pprof.StopCPUProfile()
		fileMem.Close()
	}
}
