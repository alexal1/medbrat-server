package usecase

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type VisionUseCase interface {
	Start(imageBase64 string) *Message
}

// ---------------------------------------------------------------------------------------------------------------------

type vision struct {
	firstMessage *Message
}

func NewVision() VisionUseCase {
	return &vision{}
}

func (v *vision) Start(imageBase64 string) *Message {
	panic("implement me")
}
