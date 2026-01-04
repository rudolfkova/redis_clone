package redis

type Storage interface {
	Set(key string, value string)
	Get(key string) (string, bool)
	Delete(key string)
}
