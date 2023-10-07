package args

import (
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime/pprof"
	"time"
)

const logDir = "profilerLogs"

func Profile() func() {
	if _, err := os.Stat(logDir); errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(logDir, os.ModePerm); err != nil {
			slog.Error(fmt.Sprintf("Error creating logger directory: %v", err.Error()))
		}
	}

	fileNameDate := time.Now().Format("2006-01-02-15:04:05")

	fileCpu, fileLatestCpu := createLog(fileNameDate, "cpu", false)
	fileMem, fileLatestMem := createLog(fileNameDate, "memory", false)
	fileInfo, fileLatestInfo := createLog(fileNameDate, "info", true)

	pprof.StartCPUProfile(fileCpu)
	pprof.WriteHeapProfile(fileMem)
	startTime := time.Now()

	return func() {
		pprof.StopCPUProfile()
		fileMem.Close()

		fileInfo.WriteString("Run Duration: " + time.Since(startTime).String() + "\n")
		fileInfo.WriteString("Start Time: " + startTime.String() + "\n")
		fileInfo.WriteString("End Time: " + time.Now().String() + "\n")
		fileInfo.Close()

		writeLatest(fileLatestCpu, "cpu", fileNameDate, false)
		writeLatest(fileLatestMem, "memory", fileNameDate, false)
		writeLatest(fileLatestInfo, "info", fileNameDate, true)
	}
}

func writeLatest(latestFile *os.File, profileType string, fileNameDate string, logExt bool) {
	var ext string
	if logExt {
		ext = "log"
	} else {
		ext = "pprof"
	}

	file, err := os.Open(fmt.Sprintf("%s/%s-%s.%s", logDir, fileNameDate, profileType, ext))
	fileError(profileType, err)

	_, err = io.Copy(latestFile, file)
	if err != nil {
		slog.Error(fmt.Sprintf("Error writing to latest %s profiler file: %v", profileType, err.Error()))
	}

	file.Close()
}

func fileError(profileType string, err error) {
	if err != nil {
		slog.Error(fmt.Sprintf("Error creating file for %s profiler: %v", profileType, err.Error()))
	}
}

func createLog(time string, profileType string, logExt bool) (log, latest *os.File) {
	var err error

	var ext string
	if logExt {
		ext = "log"
	} else {
		ext = "pprof"
	}

	log, err = os.Create(fmt.Sprintf("%s/%s-%s.%s", logDir, time, profileType, ext))
	fileError(profileType, err)
	latest, err = os.Create(fmt.Sprintf("%s/latest-%s.%s", logDir, profileType, ext))
	fileError(profileType, err)

	return
}
