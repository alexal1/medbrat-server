package usecase

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type VisionUseCase interface {
	Start(imageBase64 *string) *Message
}

// ---------------------------------------------------------------------------------------------------------------------

type vision struct {
	firstMessage *Message
	blood        *map[BloodComponent]float32
	ocr          *OCR
}

func NewVision(blood *map[BloodComponent]float32, ocr *OCR) VisionUseCase {
	return &vision{
		blood: blood,
		ocr:   ocr,
	}
}

func (v *vision) Start(imageBase64 *string) *Message {
	panic("implement me")
}
