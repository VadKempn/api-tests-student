package advertisements

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	advertisementClient "api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/constants"
	"api-tests-template/internal/utils"
)

// Create создаёт объявление и регистрирует удаление тестовых данных.
func Create(t *testing.T, token string, fields map[string]string, photos [][]byte, expectedStatus int) string {
	t.Helper()
	response := advertisementClient.HttpPostAdvertisement(t, token, fields, photos)
	body := string(utils.ReadBody(t, response))
	if response.StatusCode == http.StatusCreated {
		id := gjson.Get(body, "id").String()
		require.NotEmpty(t, id, "В созданном объявлении отсутствует id")
		t.Cleanup(func() {
			DeleteIfExists(t, token, id)
		})
	}
	require.Equal(t, expectedStatus, response.StatusCode, "Неожиданный статус создания объявления, ответ: %s", body)
	return body
}

// Get возвращает объявление и проверяет статус ответа.
func Get(t *testing.T, id string, expectedStatus int) string {
	t.Helper()
	return utils.ReadResponse(t, advertisementClient.HttpGetAdvertisement(t, id), expectedStatus)
}

// Search возвращает результат поиска по названию.
func Search(t *testing.T, title string) string {
	t.Helper()
	return utils.ReadResponse(t, advertisementClient.HttpGetAdvertisements(t, title), http.StatusOK)
}

// GetPhotos возвращает фотографии объявления.
func GetPhotos(t *testing.T, token, id string) string {
	t.Helper()
	return utils.ReadResponse(t, advertisementClient.HttpGetPhotos(t, token, id), http.StatusOK)
}

// GetPhotoFile проверяет доступность фотографии и возвращает её содержимое.
func GetPhotoFile(t *testing.T, photoURL string) []byte {
	t.Helper()
	response := advertisementClient.HttpGetPhotoFile(t, photoURL)
	body := utils.ReadBody(t, response)
	require.Equal(t, http.StatusOK, response.StatusCode, "Фотография недоступна: %s", photoURL)
	require.Contains(t, response.Header.Get("Content-Type"), constants.PhotoContentType, "Неверный тип фотографии")
	return body
}

// Delete удаляет объявление и проверяет статус ответа.
func Delete(t *testing.T, token, id string, expectedStatus int) {
	t.Helper()
	utils.ReadResponse(t, advertisementClient.HttpDeleteAdvertisement(t, token, id), expectedStatus)
}

// DeleteIfExists удаляет только созданное тестом объявление, если оно ещё существует.
func DeleteIfExists(t *testing.T, token, id string) {
	t.Helper()
	response := advertisementClient.HttpGetAdvertisement(t, id)
	body := utils.ReadBody(t, response)
	if response.StatusCode == http.StatusNotFound {
		t.Logf("Очистка: объявление %s уже отсутствует", id)
		return
	}
	require.Equal(t, http.StatusOK, response.StatusCode, "Не удалось проверить тестовое объявление %s: %s", id, body)
	Delete(t, token, id, http.StatusNoContent)
	t.Logf("Очистка: объявление %s удалено или уже отсутствует", id)
}
