package task

type Config struct {
	Name          string
	AttachStdin   bool
	AttachStdout  bool
	AttachStderr  bool
	Cmd           []string
	Image         string
	Memory        string
	Disk          int64
	Env           []string
	RestartPolicy string
}
