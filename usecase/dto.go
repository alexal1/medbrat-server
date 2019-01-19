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

type BloodComponent string

const (
	Hemoglobin                        BloodComponent = "Гемоглобин"
	Erythrocytes                      BloodComponent = "Эритроциты"
	ColorIndicator                    BloodComponent = "Цветовой показатель"
	AverageHemoglobinInOneErythrocyte BloodComponent = "Среднее содержание гемоглобина в 1 эритроците"
	Reticulocytes                     BloodComponent = "Ретикулоциты"
	Platelets                         BloodComponent = "Тромбоциты"
	Leukocytes                        BloodComponent = "Лейкоциты"
	Myelocytes                        BloodComponent = "Миелоциты"
	Metamyelocytes                    BloodComponent = "Метамиелоциты"
	Bandworms                         BloodComponent = "Палочкоядерные"
	Segmented                         BloodComponent = "Сегментоядерные"
	Eosinophils                       BloodComponent = "Эозинофилы"
	Basophils                         BloodComponent = "Базофилы"
	Lymphocytes                       BloodComponent = "Лимфоциты"
	Monocytes                         BloodComponent = "Моноциты"
	PlasmaCells                       BloodComponent = "Плазматические клетки"
	ErythrocytesSedimentationRate     BloodComponent = "Скорость (реакция) оседания эритроцитов"
)
