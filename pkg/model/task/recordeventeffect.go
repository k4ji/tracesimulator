package task

type RecordEventEffect struct {
	event Event
}

func NewRecordEventEffect(event Event) RecordEventEffect {
	return RecordEventEffect{
		event: event,
	}
}

func (r *RecordEventEffect) Event() Event {
	return r.event
}
