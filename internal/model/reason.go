package model

//type Reason int

const (
	Unset                  = "Unset"
	ManualValidationNeeded = "ManualValidationNeeded"
	TextValidationFailed   = "TextValidationFailed"
	LinkValidationFailed   = "LinkValidationFailed"
	ImageValidationFailed  = "ImageValidationFailed"
)

//func (r Reason) String() string {
//	switch r {
//	case Unset:
//		return "Unset"
//	case ManualValidationNeeded:
//		return "ManualValidationNeeded"
//	case TextValidationFailed:
//		return "TextValidationFailed"
//	case LinkValidationFailed:
//		return "LinkValidationFailed"
//	case ImageValidationFailed:
//		return "ImageValidationFailed"
//	}
//	return "Unknown"
//}
