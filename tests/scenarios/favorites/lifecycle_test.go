package favorites

import (
	"net/http"

	"github.com/tidwall/gjson"

	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestFavoriteLifecycle() {
	t := s.T()
	id := base.CreateAdvertisement(t, s.ownerData.Token)
	title := gjson.Get(advertisements.Get(t, id, http.StatusOK), "title").String()
	t.Log("Шаг: добавляем объявление другого пользователя в избранное")
	favorites.Add(t, s.loginData.Token, id, http.StatusCreated)
	body := favorites.List(t, s.loginData.Token, title, http.StatusOK)
	base.AssertFavorite(t, body, id, true)

	t.Log("Шаг: удаляем объявление из избранного")
	favorites.Remove(t, s.loginData.Token, id, http.StatusNoContent)
	body = favorites.List(t, s.loginData.Token, title, http.StatusOK)
	base.AssertFavorite(t, body, id, false)

	t.Log("Шаг: добавляем объявление повторно перед проверкой удаления владельцем")
	favorites.Add(t, s.loginData.Token, id, http.StatusCreated)
	body = favorites.List(t, s.loginData.Token, title, http.StatusOK)
	base.AssertFavorite(t, body, id, true)

	t.Log("Шаг: владелец удаляет само объявление")
	advertisements.Delete(t, s.ownerData.Token, id, http.StatusNoContent)
	body = favorites.List(t, s.loginData.Token, title, http.StatusOK)
	base.AssertFavorite(t, body, id, false)
}
