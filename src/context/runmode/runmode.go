package runmode

type RunMode int

const (
	FILE_TO_FILE RunMode = 1
	GIT_TO_FILE  RunMode = 2
	GIT_TO_GIT   RunMode = 3
)
