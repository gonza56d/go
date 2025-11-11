# go
Golang asynchronism parallelism concepts.

## Study Plan
	1.	Goroutines
	•	Lightweight threads managed by the Go runtime.
	•	Created with the go keyword (e.g., go func() { ... }()).
	•	Core mechanism for asynchronous execution.
	2.	Channels
	•	Typed conduits for communication between goroutines.
	•	Enable synchronization and data sharing safely (via <- send/receive).
	3.	select statement
	•	Allows waiting on multiple channel operations simultaneously.
	•	Used for multiplexing communication between goroutines.
	4.	sync package
	•	Provides synchronization primitives like WaitGroup, Mutex, RWMutex, and Once.
	•	Useful for controlling goroutine lifecycles or protecting shared state.
	5.	context package
	•	Enables cancellation and timeout propagation across goroutines.
	•	Often used in servers, pipelines, and network calls.
	6.	Worker Pools / Pipelines
	•	Concurrency design patterns built using goroutines and channels.
	•	Used to control resource usage and manage workloads.
	7.	runtime.GOMAXPROCS
	•	Controls how many OS threads can execute Go code simultaneously — affects parallelism (not concurrency).
	8.	Third-party abstractions (optional)
	•	Libraries like errgroup, tunny, or reactive-style frameworks add structured concurrency patterns.
