package favorites

import (
	"net/http"
	"testing"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/favorites"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestFavoritesWithoutAuthorization() {
	id := base.CreateAdvertisement(s.T(), s.ownerData.Token)
	tokens := []struct {
		name    string
		token   string
		message string
	}{
		{"missing_token", "", constants.MissingTokenMessage},
		{"invalid_token", constants.InvalidToken, constants.InvalidAuthorizationMessage},
	}
	operations := []struct {
		name string
		call func(*testing.T, string) string
	}{
		{"list", func(t *testing.T, token string) string { return favorites.List(t, token, "", http.StatusUnauthorized) }},
		{"add", func(t *testing.T, token string) string { return favorites.Add(t, token, id, http.StatusUnauthorized) }},
		{"remove", func(t *testing.T, token string) string {
			return favorites.Remove(t, token, id, http.StatusUnauthorized)
		}},
	}
	for _, token := range tokens {
		for _, operation := range operations {
			s.Run(operation.name+"_"+token.name, func() {
				t := s.T()
				t.Log("Подготовка: запрос без действующей авторизации")
				t.Log("Шаг: обращаемся к ручке избранного")
				body := operation.call(t, token.token)
				base.AssertError(t, body, constants.Unauthorized, token.message)
			})
		}
	}
}
