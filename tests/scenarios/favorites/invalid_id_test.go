package favorites

import (
	"net/http"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestInvalidFavoriteID() {
	cases := []struct {
		name    string
		id      string
		message string
	}{
		{"empty", "", "invalid UUID length: 0"},
		{"invalid", constants.InvalidAdvertisementID, "invalid UUID length: 9"},
	}
	for _, item := range cases {
		s.Run(item.name, func() {
			t := s.T()
			t.Log("Подготовка: некорректный идентификатор объявления")
			t.Log("Шаг: добавляем объявление с некорректным id")
			body := favorites.Add(t, s.loginData.Token, item.id, http.StatusBadRequest)
			base.AssertError(t, body, constants.InvalidRequest, item.message)
		})
	}
	s.Run("remove_invalid", func() {
		t := s.T()
		t.Log("Подготовка: некорректный идентификатор объявления")
		t.Log("Шаг: удаляем запись с некорректным id")
		body := favorites.Remove(t, s.loginData.Token, constants.InvalidAdvertisementID, http.StatusBadRequest)
		base.AssertError(t, body, constants.InvalidRequest, "Invalid advertisement ID")
	})
}
