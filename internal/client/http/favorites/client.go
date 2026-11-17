package favorites

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpPostFavorite добавляет объявление в избранное.
func HttpPostFavorite(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Post(path.FavoritesPath).
		ContentType("application/json").JSON(map[string]string{"advertisement_id": id}).Expect(t).End().Response
}

// HttpGetFavorites получает избранное пользователя с поиском по названию.
func HttpGetFavorites(t *testing.T, token, search string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(path.FavoritesPath).
		Query("search", search).Query("page_size", strconv.Itoa(constants.PageSize)).Expect(t).End().Response
}

// HttpDeleteFavorite удаляет объявление из избранного.
func HttpDeleteFavorite(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(fmt.Sprintf(path.FavoritePath, id)).
		Expect(t).End().Response
}
