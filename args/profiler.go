package args

import (
	"log/slog"
	"os"
	"runtime/pprof"
	"time"
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

	fileInfo, err := os.Create("profilerLogs/info.log")
	if err != nil {
		slog.Error("Error creating file for info of profiler: " + err.Error())
	}
	startTime := time.Now()

	return func() {
		pprof.StopCPUProfile()
		fileMem.Close()

		fileInfo.WriteString("Run Duration: " + time.Since(startTime).String() + "\n")
		fileInfo.WriteString("Start Time: " + startTime.String() + "\n")
		fileInfo.WriteString("End Time: " + time.Now().String() + "\n")
		fileInfo.Close()
	}
}
