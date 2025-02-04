package main

import (
	"bufio"
	"context"
	"fmt"
	"log-analyzer-go/internal/usecase/log"
	"log-analyzer-go/pkg/wpool"
	"os"
	"strconv"
	"strings"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fprintf(out, "Enter the number of workers: ")
	out.Flush()

	workerInput, _ := in.ReadString('\n')
	workerInput = strings.TrimSpace(workerInput)
	workersCount, err := strconv.Atoi(workerInput)
	if err != nil {
		fmt.Fprintln(out, "Invalid number of workers")
		out.Flush()
		return
	}

	fmt.Fprintf(out, "Enter the path to the LogFile: ")
	out.Flush()

	filePath, _ := in.ReadString('\n')
	filePath = strings.TrimSpace(filePath)

	wp := wpool.New(workersCount)
	log.ProcessLogFile(ctx, filePath, wp)
}
