package usecase

type Message struct {
	Id                     uint64 `json:"id"`
	Text                   string `json:"text"`
	AnswerFormat           `json:"answer_format"`
	NextMessage            *Message                 `json:"-"`
	NextMessageByCondition map[interface{}]*Message `json:"-"`
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
