package tests

import (
	_ "embed"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/utils"
)

//go:embed testdata/photo.png
var Photo []byte

// AdvertisementFields возвращает независимый набор полей с уникальным названием.
func AdvertisementFields() map[string]string {
	return map[string]string{
		"title":       constants.TitlePrefix + utils.RandomString(16),
		"description": constants.Description,
		"price":       constants.Price,
		"quantity":    constants.Quantity,
	}
}

// PhotoFiles подготавливает нужное количество фотографий для запроса.
func PhotoFiles(count int) [][]byte {
	photos := make([][]byte, count)
	for index := range photos {
		photos[index] = Photo
	}
	return photos
}

// CreateAdvertisement создаёт отдельное объявление для одного теста.
func CreateAdvertisement(t *testing.T, token string) string {
	t.Helper()
	t.Log("Подготовка: создаём тестовое объявление")
	body := advertisements.Create(t, token, AdvertisementFields(), PhotoFiles(constants.MinPhotos), http.StatusCreated)
	return gjson.Get(body, "id").String()
}

// AssertError проверяет код и сообщение ошибки API.
func AssertError(t *testing.T, body, code, message string) {
	t.Helper()
	require.Equal(t, code, gjson.Get(body, "error").String(), "Неверный код ошибки, ответ: %s", body)
	require.Equal(t, message, gjson.Get(body, "message").String(), "Неверное сообщение ошибки, ответ: %s", body)
	t.Logf("Проверка: получены ошибка %s и ожидаемое сообщение", code)
}

// AssertFavorite проверяет наличие конкретного объявления в ответе избранного.
func AssertFavorite(t *testing.T, body, id string, expected bool) {
	t.Helper()
	found := false
	for _, item := range gjson.Get(body, "items").Array() {
		if item.Get("advertisement.id").String() == id {
			found = true
			break
		}
	}
	require.Equal(t, expected, found, "Неверное наличие объявления %s в избранном, ответ: %s", id, body)
	t.Logf("Проверка: наличие объявления %s в избранном равно %t", id, expected)
}
