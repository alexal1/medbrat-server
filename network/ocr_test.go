package network

import (
	"medbrat-server/configs"
	"medbrat-server/usecase"
	"testing"
)

var trueBlood = usecase.NewBloodGeneral(143, 5.08, 0.84, 28.1, -1, 175, 5.17, -1, -1, -1, -1, 0.24, 0.04, 2.45, 0.5, -1, -1)

func TestParseBlood(t *testing.T) {
	sampleRecognizedText := "107031 г. Москва, Рождественский бульвар д. 21, стр. 2 \t\r\n• ГЕМОТЕСТ \tтел, (495) 532-13-1 З, 8(800) 550-13-13 \t\r\nмодицинскдп павоедторис• \twww.qemotest.ru \t\r\nN? направления: \t43952204 \tдата: \t08.11.2017 \tФамилия: \tМищенко \t\r\nЛПУ: 70001 70001 ул. Ярцевская д, 11, \tкорп. 1 \tИмя: \tАлександр Дмитриевич \t\r\nАдрес пациента: \tдата рождения: \t24.11.1994 \t\r\nВрач: \tПол: \tмужской \t\r\nОтделение: \t\r\nN? страх. полиса: \t\r\nНормальные \t\r\nНаименование исследования \tРезультат \tЕд. изм. \t\r\nзначения \t\r\nОБЩИЙ АНАЛИЗ КРОВИ \t\r\nГемоглобин \t143 \tг/л \t132 173 \t\r\nЭритроциты \t5.08 \tх10*12/л \t4.3 -57 \t\r\nГематокрит \t41,5 \t39 - 49 \t\r\nСредний объем эритроцитов (MCV) \t82 \tфл \t80 100 \t\r\nСреднее содержание НЬ в эритроците (МСН) \t28.1 \t27 34 \t\r\nСредняя концентрация НЬ в эритроцитах \t\r\n345 \tг/л \tзоо - 380 \t\r\n(мснс) \t\r\nЦветовой показатель \t0.84 \t\r\nТромбоциты \t175 \tх10*9/л \t180 - 320 \t\r\nЛейкоциты \t5.17 \tх10*9/л \t- 11 з \t\r\nНезрелые гранулоциты \t0.01 \t10*9/л \to - 0.09 \t\r\nНезрелые гранулоциты % \t0.2 \t\r\nНейтрофилы сегментоядерные \t1.94 \tх10*9/л \t1,60 7,90 \t\r\nНейтрофилы сегментоядерные % \t37.5 \t47 - 72 \t\r\nЭозинофилы \t0.24 \tх10*9/л \t0.02 0.30 \t\r\nЭозинофилы % \t4.6 \t1-5 \t\r\nБазофилы \t0.04 \tх10*9/л \to 0.07 \t\r\nБазофилы % \t0.8 \t0-1 \t\r\nМоноциты \t0.50 \tх10*9/л \t0.09 - 0 60 \t\r\nМоноциты % \t9.7 \t3-11 \t\r\nЛимфоциты \t2.45 \tх10*9/л \t1.20 - 3.00 \t\r\nЛимфоциты % \t47.4 \t19 37 \t\r\nСОЭ (по Вестергрену) \t2 \tмм/час \tо- 15 \t\r\nИММУНОЛОГИЯ \t\r\nИммуноглобулин IgE общий (Immulite) \t41.2 \tМЕ/мл \t0.0 87 0 \t\r\nдпя 33 ия • \t\r\nВРАЧ \tрезультотоо \t\r\nСилкинаТЛ. \t(ЭГЕМОТЕСТ \t\r\nОСИ \t\r\nТОРУ ЧЫХ \t\r\nисследоидний \t\r\nЭлектронная подпись: врач Тамара Салкина \t\r\n• чос \t\r\nсреда, 8. Ноябрь 2017 20:17 \tСтраница 1 \t\r\n"
	maxCorrectionsCount := 2

	parsedBlood := usecase.NewBloodGeneral()
	ParseBlood(&parsedBlood, &sampleRecognizedText, maxCorrectionsCount)

	allValues := 0
	matchedValues := 0
	parsedBlood.ForEach(func(index int, name string) {
		parsedValue := parsedBlood.Get(name)
		trueValue := trueBlood.Get(name)

		allValues++
		if parsedValue != trueValue {
			t.Logf("Mismatch detected for %s: parsedValue = %f, trueValue = %f", name, parsedValue, trueValue)
		} else {
			matchedValues++
		}
	})

	t.Logf("%d/%d values are correct", matchedValues, allValues)
	if matchedValues < allValues {
		t.Fail()
	}
}

func TestOcrSpace_RecognizeGeneralBloodTest(t *testing.T) {
	ocrInstance := NewOCR()
	parsedBlood := usecase.NewBloodGeneral()
	ocrInstance.RecognizeGeneralBloodTest(&parsedBlood, &configs.SampleBase64Image)

	allValues := 0
	matchedValues := 0
	parsedBlood.ForEach(func(index int, name string) {
		parsedValue := parsedBlood.Get(name)
		trueValue := trueBlood.Get(name)

		allValues++
		if parsedValue != trueValue {
			t.Logf("Mismatch detected for %s: parsedValue = %f, trueValue = %f", name, parsedValue, trueValue)
		} else {
			matchedValues++
		}
	})

	t.Logf("%d/%d values are correct", matchedValues, allValues)
	if matchedValues < allValues {
		t.Fail()
	}
}
