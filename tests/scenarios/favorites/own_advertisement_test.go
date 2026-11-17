package favorites

import (
	"net/http"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestOwnAdvertisement() {
	t := s.T()
	id := base.CreateAdvertisement(t, s.loginData.Token)
	t.Log("Шаг: добавляем собственное объявление в избранное")
	body := favorites.Add(t, s.loginData.Token, id, http.StatusForbidden)
	base.AssertError(t, body, constants.Forbidden, "Cannot add own advertisement to favorites")
	base.AssertFavorite(t, favorites.List(t, s.loginData.Token, "", http.StatusOK), id, false)
}
