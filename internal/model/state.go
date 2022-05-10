package model

//type MsgState int

const (
	Invalid  = "Invalid"
	Queued   = "Queued"
	Accepted = "Accepted"
	Rejected = "Rejected"
)

//func NewMsgFromString(s string) MsgState {
//	switch s {
//	case "Queued":
//		return Queued
//	case "Accepted":
//		return Accepted
//	case "Rejected":
//		return Rejected
//	}
//	return Invalid
//}
//
//func (s MsgState) String() string {
//	switch s {
//	case Invalid:
//		return "Invalid"
//	case Queued:
//		return "Queued"
//	case Accepted:
//		return "Accepted"
//	case Rejected:
//		return "Rejected"
//	}
//	return "Unknown"
//}
