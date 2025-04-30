package span

type Status uint8

const (
	StatusUnset Status = iota
	StatusOK
	StatusError
)

const StatusUnsetString = "unset"
const StatusOKString = "ok"
const StatusErrorString = "error"
const StatusUnknownString = "unknown"

func (s Status) String() string {
	switch s {
	case StatusUnset:
		return StatusUnsetString
	case StatusOK:
		return StatusOKString
	case StatusError:
		return StatusErrorString
	default:
		return StatusUnknownString
	}
}
