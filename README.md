# Go Job Queue (Mini RQ-like)

A simple, lightweight **in-memory job queue** in Go inspired by Python RQ.  
Supports multiple workers, concurrent job processing, and retry logic.

---

## Features

- Define jobs as functions with payloads
- Multiple concurrent workers
- In-memory job queue using Go channels
- Retry failed jobs automatically
- Graceful shutdown with `context.Context`
- Easy to extend (priorities, persistence, delayed jobs)

---
