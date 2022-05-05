package model

type MsgState int

const (
	Invalid MsgState = iota
	Queued
	Accepted
	Rejected
)

func (s MsgState) String() string {
	switch s {
	case Invalid:
		return "Invalid"
	case Queued:
		return "Queued"
	case Accepted:
		return "Accepted"
	case Rejected:
		return "Rejected"
	}
	return "Unknown"
}
