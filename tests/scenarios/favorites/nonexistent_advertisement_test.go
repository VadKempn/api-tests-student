package favorites

import (
	"net/http"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestNonexistentAdvertisement() {
	s.Run("add", func() {
		t := s.T()
		t.Log("Подготовка: корректный UUID несуществующего объявления")
		t.Log("Шаг: добавляем несуществующее объявление в избранное")
		body := favorites.Add(t, s.loginData.Token, constants.MissingAdvertisementID, http.StatusNotFound)
		base.AssertError(t, body, constants.NotFound, "Advertisement not found")
	})
	s.Run("remove", func() {
		t := s.T()
		t.Log("Подготовка: корректный UUID несуществующего объявления")
		t.Log("Шаг: удаляем несуществующую запись из избранного")
		body := favorites.Remove(t, s.loginData.Token, constants.MissingAdvertisementID, http.StatusNotFound)
		base.AssertError(t, body, constants.NotFound, "Favorite entry not found")
	})
}
