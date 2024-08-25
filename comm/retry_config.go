package comm

// Retry levels based on importance
const (
	lowImportance    = 5  // 5 retries for standard messages
	highImportance   = 12 // 12 retries for priority messages
)
