package advertisements

import (
	"net/http"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestCreateWithoutAuthorization() {
	cases := []struct {
		name    string
		token   string
		message string
	}{
		{"missing_token", "", constants.MissingTokenMessage},
		{"invalid_token", constants.InvalidToken, constants.InvalidAuthorizationMessage},
	}
	for _, item := range cases {
		s.Run(item.name, func() {
			t := s.T()
			t.Log("Подготовка: корректное объявление без действующей авторизации")
			t.Log("Шаг: отправляем запрос создания")
			body := advertisements.Create(t, item.token, base.AdvertisementFields(), base.PhotoFiles(constants.MinPhotos), http.StatusUnauthorized)
			base.AssertError(t, body, constants.Unauthorized, item.message)
		})
	}
}
