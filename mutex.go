package nmutex

type Mutex interface {
	Lock() error

	Unlock() error
}
