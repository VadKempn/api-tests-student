package favorites

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	favoriteClient "api-tests-template/internal/client/http/favorites"
	"api-tests-template/internal/utils"
)

// Add добавляет объявление в избранное и регистрирует очистку.
func Add(t *testing.T, token, id string, expectedStatus int) string {
	t.Helper()
	response := favoriteClient.HttpPostFavorite(t, token, id)
	body := string(utils.ReadBody(t, response))
	if response.StatusCode == http.StatusCreated {
		t.Cleanup(func() {
			RemoveIfExists(t, token, id)
		})
	}
	require.Equal(t, expectedStatus, response.StatusCode, "Неожиданный статус добавления в избранное, ответ: %s", body)
	return body
}

// List возвращает избранное пользователя и проверяет статус ответа.
func List(t *testing.T, token, search string, expectedStatus int) string {
	t.Helper()
	return utils.ReadResponse(t, favoriteClient.HttpGetFavorites(t, token, search), expectedStatus)
}

// Remove удаляет запись из избранного и проверяет статус ответа.
func Remove(t *testing.T, token, id string, expectedStatus int) string {
	t.Helper()
	return utils.ReadResponse(t, favoriteClient.HttpDeleteFavorite(t, token, id), expectedStatus)
}

// RemoveIfExists очищает добавленную тестом запись, если она ещё существует.
func RemoveIfExists(t *testing.T, token, id string) {
	t.Helper()
	response := favoriteClient.HttpDeleteFavorite(t, token, id)
	body := utils.ReadBody(t, response)
	require.Contains(t, []int{http.StatusNoContent, http.StatusNotFound}, response.StatusCode,
		"Не удалось очистить избранное для объявления %s: %s", id, body)
	t.Logf("Очистка: объявление %s отсутствует в избранном", id)
}
