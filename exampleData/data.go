package exampleData

type ExampleDataFormat struct {
	Etunimi       string `json:"etunimi"`
	Sukunimi      string `json:"sukunimi"`
	Henkilotunnus string `json:"henkilotunnus"`
	Verenpaine    string `json:"verenpaine"`
	Paino         string `json:"paino"`
	Pituus        string `json:"pituus"`
	Ika           string `json:"ika"`
}

func getExampleDataFormat(etunimi string, sukunimi string, henkilotunnus string, verenpaine string, paino string, pituus string, ika string) *ExampleDataFormat {
	return &ExampleDataFormat{
		Etunimi:       etunimi,
		Sukunimi:      sukunimi,
		Henkilotunnus: henkilotunnus,
		Verenpaine:    verenpaine,
		Paino:         paino,
		Pituus:        pituus,
		Ika:           ika,
	}
}

func GetData() []ExampleDataFormat {
	listOfRandomData := make([]ExampleDataFormat, 0, 10)
	for i := 1; i < 10; i++ {
		randomData := ExampleDataFormat{
			Etunimi:       GenerateRandomFirstName(),
			Sukunimi:      GenerateRandomLastName(),
			Henkilotunnus: GenerateRandomSocialSecurityNumberName(),
			Verenpaine:    GenerateRandomBloodPressure(),
			Paino:         GenerateRandomWeight(),
			Pituus:        GenerateRandomHeight(),
			Ika:           GenerateRandomAge(),
		}
		listOfRandomData = append(listOfRandomData, randomData)
	}
	return listOfRandomData
}
