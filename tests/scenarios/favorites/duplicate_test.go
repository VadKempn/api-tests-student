package favorites

import (
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestDuplicateFavorite() {
	t := s.T()
	id := base.CreateAdvertisement(t, s.ownerData.Token)
	title := gjson.Get(advertisements.Get(t, id, http.StatusOK), "title").String()
	favorites.Add(t, s.loginData.Token, id, http.StatusCreated)
	t.Log("Шаг: добавляем то же объявление в избранное ещё раз")
	body := favorites.Add(t, s.loginData.Token, id, http.StatusConflict)
	base.AssertError(t, body, constants.AlreadyExists, "Advertisement is already in favorites")
	items := gjson.Get(favorites.List(t, s.loginData.Token, title, http.StatusOK), "items").Array()
	require.Len(t, items, 1, "Повторное добавление создало дубликат")
	require.Equal(t, id, items[0].Get("advertisement.id").String())
	t.Log("Проверка: в избранном осталась одна запись")
}
