package usecase

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type VisionUseCase interface {
	Start(imageBase64 *string) *Message
}

// ---------------------------------------------------------------------------------------------------------------------

type vision struct {
	firstMessage *Message
	blood        *BloodGeneral
	ocr          *OCR
}

func NewVision(blood *BloodGeneral, ocr *OCR) VisionUseCase {
	return &vision{
		blood: blood,
		ocr:   ocr,
	}
}

func (v *vision) Start(imageBase64 *string) *Message {
	panic("implement me")
}
