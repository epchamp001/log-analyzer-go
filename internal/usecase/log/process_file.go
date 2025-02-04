package log

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"log-analyzer-go/internal/models"
	"log-analyzer-go/pkg/wpool"
	"os"
	"strconv"
)

func ProcessLogFile(ctx context.Context, path string, wp *wpool.WorkerPool) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	jobs := []wpool.Job{}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if !FilterLog(line) {
			continue
		}

		job := wpool.Job{
			Descriptor: wpool.JobDescriptor{
				ID: strconv.Itoa(i),
			},
			ExecFn: func(ctx context.Context, args interface{}) (interface{}, error) {
				return ParseLog(args.(string))
			},
			Args: line,
		}

		jobs = append(jobs, job)
	}

	go wp.Start(ctx)
	go wp.GenerateJobs(jobs)

	entries := []models.LogEntry{}

	for r := range wp.Results() {
		if r.Err == nil {
			entries = append(entries, r.Value.(models.LogEntry))
		}
	}

	stats := AggregateLogs(entries)
	fmt.Println("IP statistics:")
	for ip, count := range stats {
		fmt.Printf("%s: %d requests\n", ip, count)
	}
}
