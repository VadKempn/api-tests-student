package favorites

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestFavoritesLimit() {
	t := s.T()
	t.Log("Подготовка: определяем количество уже существующих записей")
	initial := int(gjson.Get(favorites.List(t, s.loginData.Token, "", http.StatusOK), "total").Int())
	require.LessOrEqual(t, initial, constants.FavoritesLimit, "Количество избранных уже превышает лимит")
	for index := initial; index < constants.FavoritesLimit; index++ {
		id := base.CreateAdvertisement(t, s.ownerData.Token)
		favorites.Add(t, s.loginData.Token, id, http.StatusCreated)
	}
	count := gjson.Get(favorites.List(t, s.loginData.Token, "", http.StatusOK), "total").Int()
	require.Equal(t, int64(constants.FavoritesLimit), count, "Не удалось подготовить заполненное избранное")
	id := base.CreateAdvertisement(t, s.ownerData.Token)
	t.Log("Шаг: добавляем ещё одно объявление после достижения лимита")
	body := favorites.Add(t, s.loginData.Token, id, http.StatusBadRequest)
	base.AssertError(t, body, "limit_reached", fmt.Sprintf("Favorites limit of %d reached", constants.FavoritesLimit))
	actualCount := gjson.Get(favorites.List(t, s.loginData.Token, "", http.StatusOK), "total").Int()
	require.Equal(t, count, actualCount, "После отказа количество избранных изменилось")
	t.Log("Проверка: лимит соблюдается, чужие ранее добавленные записи не изменены")
}
