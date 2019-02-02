package usecase

type Message struct {
	Id                 uint64 `json:"id"`
	Text               string `json:"text"`
	AnswerFormat       `json:"answer_format"`
	NextMessage        *Message                                       `json:"-"`
	NextMessageByValue func(value interface{}) (nextMessage *Message) `json:"-"`
}

type AnswerFormat string

const (
	YesNo  AnswerFormat = "yes_no"
	Image  AnswerFormat = "image"
	Number AnswerFormat = "number"
	None   AnswerFormat = "none"
	Ok     AnswerFormat = "ok"
)

const AnswerYes = "yes"
const AnswerNo = "no"

type BloodComponent uint8

const (
	Hemoglobin = iota
	Erythrocytes
	ColorIndicator
	AverageHemoglobinInOneErythrocyte
	Reticulocytes
	Platelets
	Leukocytes
	Myelocytes
	Metamyelocytes
	Bandworms
	Segmented
	Eosinophils
	Basophils
	Lymphocytes
	Monocytes
	PlasmaCells
	ErythrocytesSedimentationRate
)

const firstComponent = Hemoglobin
const lastComponent = ErythrocytesSedimentationRate

func ForEachBloodComponent(action func(component BloodComponent)) {
	for component := firstComponent; component <= lastComponent; component++ {
		action(BloodComponent(component))
	}
}

func (component *BloodComponent) Name() (name string) {
	switch *component {
	case Hemoglobin:
		name = "Гемоглобин"
	case Erythrocytes:
		name = "Эритроциты"
	case ColorIndicator:
		name = "Цветовой показатель"
	case AverageHemoglobinInOneErythrocyte:
		name = "Среднее содержание гемоглобина в 1 эритроците"
	case Reticulocytes:
		name = "Ретикулоциты"
	case Platelets:
		name = "Тромбоциты"
	case Leukocytes:
		name = "Лейкоциты"
	case Myelocytes:
		name = "Миелоциты"
	case Metamyelocytes:
		name = "Метамиелоциты"
	case Bandworms:
		name = "Палочкоядерные"
	case Segmented:
		name = "Сегментоядерные"
	case Eosinophils:
		name = "Эозинофилы"
	case Basophils:
		name = "Базофилы"
	case Lymphocytes:
		name = "Лимфоциты"
	case Monocytes:
		name = "Моноциты"
	case PlasmaCells:
		name = "Плазматические клетки"
	case ErythrocytesSedimentationRate:
		name = "Скорость (реакция) оседания эритроцитов"
	}

	return
}
