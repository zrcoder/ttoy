package common

type State string

const (
	StateSucceed State = "succeed"
	StateFailed  State = "failed"
	StatePlaying State = "playing"
)
