package usecase

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type ManualUseCase interface {
	Start() *Message
}

// ---------------------------------------------------------------------------------------------------------------------

type manual struct {
	firstMessage *Message
	blood        *map[BloodComponent]float32
}

func NewManual(blood *map[BloodComponent]float32) ManualUseCase {
	return &manual{
		blood: blood,
	}
}

func (v *manual) Start() *Message {
	panic("implement me")
}
