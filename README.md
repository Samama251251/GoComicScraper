GoComicScrapper

A learning project to explore concurrency, channels, and the worker pool pattern in Golang
Overview
GoComicScrapper is a Golang-based web scraper that fetches comic data using a concurrent worker pool. The project uses goroutines, channels, and wait groups to efficiently retrieve and store data while handling multiple requests in parallel.
This project is designed as a learning exercise to explore key Golang concurrency concepts, including:

    Goroutines for parallel execution
    Channels for safe communication between goroutines
    Wait Groups to synchronize concurrent tasks
    Worker Pool Pattern for controlled parallelism

Features

1)Fetches comic metadata using an HTTP API
2)Implements a worker pool for efficient request handling
3)Uses channels to distribute and collect tasks
