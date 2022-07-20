package nmutex

type Session interface {
	NewMutex(key string) Mutex
}
