package advertisements

import (
	"net/http"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestCreateWithoutRequiredFields() {
	cases := []struct {
		name    string
		field   string
		omit    bool
		photos  int
		message string
	}{
		{"missing_title", "title", true, constants.MinPhotos, "title is required"},
		{"empty_title", "title", false, constants.MinPhotos, "title is required"},
		{"missing_description", "description", true, constants.MinPhotos, "description is required"},
		{"empty_description", "description", false, constants.MinPhotos, "description is required"},
		{"missing_photos", "", false, 0, "at least 1 photo is required"},
	}
	for _, item := range cases {
		s.Run(item.name, func() {
			t := s.T()
			t.Log("Подготовка: запрос без обязательного поля")
			fields := base.AdvertisementFields()
			if item.field != "" {
				if item.omit {
					delete(fields, item.field)
				} else {
					fields[item.field] = ""
				}
			}
			t.Log("Шаг: отправляем запрос создания объявления")
			body := advertisements.Create(t, s.loginData.Token, fields, base.PhotoFiles(item.photos), http.StatusBadRequest)
			base.AssertError(t, body, constants.InvalidRequest, item.message)
		})
	}
}
