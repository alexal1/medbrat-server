package usecase

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type ManualUseCase interface {
	Start() *Message
}

// ---------------------------------------------------------------------------------------------------------------------

type manual struct {
	firstMessage *Message
	blood        *BloodGeneral
}

func NewManual(blood *BloodGeneral) ManualUseCase {
	return &manual{
		blood: blood,
	}
}

func (v *manual) Start() *Message {
	return nil
}
