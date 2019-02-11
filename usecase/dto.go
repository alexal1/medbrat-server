package usecase

import (
	"log"
	"reflect"
)

// -------------------------------------------------- INTERFACE --------------------------------------------------------

type BloodGeneral interface {
	ForEach(action func(index int, name string))
	Get(name string) (value float64)
	Set(name string, value float64)
	SetByIndex(index int, value float64)
	GetTextKeys(name string) (keys []string)
}

type Message struct {
	Id                 uint64 `json:"id"`
	Text               string `json:"text"`
	PossibleAnswers    []Answer
	NextMessage        *Message                                       `json:"-"`
	NextMessageByValue func(value interface{}) (nextMessage *Message) `json:"-"`
}

type Answer string

const (
	Yes           Answer = "yes"
	No            Answer = "no"
	StartAgain    Answer = "start_again"
	ChangedMyMind Answer = "changed_my_mind"
	AllRight      Answer = "all_right"
)

// ---------------------------------------------------------------------------------------------------------------------

type bloodGeneral struct {
	Hemoglobin                        float64
	Erythrocytes                      float64
	ColorIndicator                    float64
	AverageHemoglobinInOneErythrocyte float64
	Reticulocytes                     float64
	Platelets                         float64
	Leukocytes                        float64
	Myelocytes                        float64
	Metamyelocytes                    float64
	Bandcells                         float64
	Segmented                         float64
	Eosinophils                       float64
	Basophils                         float64
	Lymphocytes                       float64
	Monocytes                         float64
	PlasmaCells                       float64
	ErythrocytesSedimentationRate     float64
}

func NewBloodGeneral(components ...float64) BloodGeneral {
	blood := bloodGeneral{
		Hemoglobin:                        -1,
		Erythrocytes:                      -1,
		ColorIndicator:                    -1,
		AverageHemoglobinInOneErythrocyte: -1,
		Reticulocytes:                     -1,
		Platelets:                         -1,
		Leukocytes:                        -1,
		Myelocytes:                        -1,
		Metamyelocytes:                    -1,
		Bandcells:                         -1,
		Segmented:                         -1,
		Eosinophils:                       -1,
		Basophils:                         -1,
		Lymphocytes:                       -1,
		Monocytes:                         -1,
		PlasmaCells:                       -1,
		ErythrocytesSedimentationRate:     -1,
	}

	componentsLen := len(components)
	if componentsLen > 0 {
		blood.ForEach(func(index int, name string) {
			if index < componentsLen {
				blood.SetByIndex(index, components[index])
			}
		})
	}

	return &blood
}

// Iterate through all fields (using reflection).
func (blood *bloodGeneral) ForEach(action func(index int, name string)) {
	e := reflect.ValueOf(blood).Elem()
	t := e.Type()

	for i := 0; i < e.NumField(); i++ {
		action(i, t.Field(i).Name)
	}
}

// Get value by component name.
func (blood *bloodGeneral) Get(name string) (value float64) {
	e := reflect.ValueOf(blood).Elem()
	field := e.FieldByName(name)
	if field.IsValid() && field.Kind() == reflect.Float64 {
		value = field.Float()
	} else {
		panic("Cannot find " + name + " in BloodGeneral struct")
	}

	return
}

// Set given value by component name.
func (blood *bloodGeneral) Set(name string, value float64) {
	e := reflect.ValueOf(blood).Elem()
	field := e.FieldByName(name)
	if field.IsValid() && field.CanSet() && field.Kind() == reflect.Float64 {
		field.SetFloat(value)
	} else {
		panic("Cannot find " + name + " in BloodGeneral struct")
	}

	return
}

// Set given value by component index.
func (blood *bloodGeneral) SetByIndex(index int, value float64) {
	e := reflect.ValueOf(blood).Elem()
	field := e.Field(index)
	if field.IsValid() && field.CanSet() && field.Kind() == reflect.Float64 {
		field.SetFloat(value)
	} else {
		log.Panicf("Cannot find element with index %d in BloodGeneral struct", index)
	}

	return
}

func (blood *bloodGeneral) GetTextKeys(name string) (keys []string) {
	switch name {
	case "Hemoglobin":
		return []string{"Гемоглобин"}
	case "Erythrocytes":
		return []string{"Эритроциты"}
	case "ColorIndicator":
		return []string{"Цветовой показатель"}
	case "AverageHemoglobinInOneErythrocyte":
		return []string{"Среднее содержание гемоглобина в 1 эритроците", "Среднее содержание Hb в эритроците"}
	case "Reticulocytes":
		return []string{"Ретикулоциты"}
	case "Platelets":
		return []string{"Тромбоциты"}
	case "Leukocytes":
		return []string{"Лейкоциты"}
	case "Myelocytes":
		return []string{"Миелоциты"}
	case "Metamyelocytes":
		return []string{"Метамиелоциты"}
	case "Bandcells":
		return []string{"Палочкоядерные"}
	case "Segmented":
		return []string{"Сегментоядерные"}
	case "Eosinophils":
		return []string{"Эозинофилы"}
	case "Basophils":
		return []string{"Базофилы"}
	case "Lymphocytes":
		return []string{"Лимфоциты"}
	case "Monocytes":
		return []string{"Моноциты"}
	case "PlasmaCells":
		return []string{"Плазматические клетки"}
	case "ErythrocytesSedimentationRate":
		return []string{"Скорость (реакция) оседания эритроцитов"}
	default:
		panic("Cannot find " + name + " in BloodGeneral struct")
	}
}
