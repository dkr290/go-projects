package packagelog

// level represents available logging level
type Level byte

const (
	//LevelDebug represents the lowest level of the log , mostly used for debugging purposes
	LevelDebug Level = iota
	// LevelInfo represents information that seems valuable

	LevelInfo

	// LevelError represents the highest logging level, using to trace errors
	LevelError
)
