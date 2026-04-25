package player

type State int

const (
	StateStopped State = iota
	StatePlaying
	StateError
)

type Player interface {
	Play(url string) error
	Stop() error
	State() State
}
