package utils

import (
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

// ReadBody читает тело ответа и закрывает его.
func ReadBody(t *testing.T, response *http.Response) []byte {
	t.Helper()
	require.NotNil(t, response, "Сервер не вернул ответ")
	defer func() {
		require.NoError(t, response.Body.Close(), "Не удалось закрыть тело ответа")
	}()
	body, err := io.ReadAll(response.Body)
	require.NoError(t, err, "Не удалось прочитать тело ответа")
	return body
}

// ReadResponse проверяет статус и возвращает тело ответа.
func ReadResponse(t *testing.T, response *http.Response, expectedStatus int) string {
	t.Helper()
	body := string(ReadBody(t, response))
	require.Equal(t, expectedStatus, response.StatusCode, "Неожиданный HTTP-статус, ответ: %s", body)
	return body
}
