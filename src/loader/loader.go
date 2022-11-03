package loader

type Loader interface {
	Load() ([]string, error)
}
