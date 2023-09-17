package medium

type Medium interface {
	Send() (err error)
}
