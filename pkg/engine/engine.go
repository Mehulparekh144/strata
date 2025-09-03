package engine

type StorageEngine interface {
	Set(key string, value string) error
	Get(key string) (string, bool, error)
	Del(key string) (bool, error)
	Close() error
}
