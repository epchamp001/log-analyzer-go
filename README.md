# Log Analyzer with Worker Pool

This project is a **log processing tool** that uses **Golang worker pool** to efficiently process large log files.  
It reads a log file, filters error entries, processes them in parallel using a **worker pool**, and aggregates statistics.

## 🛠 How It Works
1. The program reads a **log file** (Apache/Nginx format).
2. It filters log entries with **HTTP error codes (4xx, 5xx)**.
3. A **worker pool** processes log lines concurrently.
4. Results are collected into a **pipeline** and aggregated.
5. The program outputs **IP statistics** based on the log entries.

---

## ⚙️ Worker Pool Implementation
A **worker pool** is a concurrency pattern where multiple worker goroutines process tasks from a shared queue.

### **Worker Pool Workflow:**
1. The main goroutine **reads** the log file and **creates jobs**.
2. Jobs are **sent to a buffered channel** (`wp.jobs`).
3. Multiple **worker goroutines** read from `wp.jobs` and process the logs.
4. Each worker **parses** a log entry and sends results to `wp.results`.
5. The main goroutine **collects results** and aggregates them.