package myAdvertisement

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/constants"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/managers/myAdvertisements"

	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Авторизация пользователя с кредами из переменных окружения и получение его параметров")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// Positive: Проверяем получение собственных объявлений в профиле
func (s *TestSuite) TestGetMyAdvertisementsPositive() {
	createdID := base.CreateAdvertisement(s.T(), s.loginData.Token)
	var advertisementsBody string
	s.Run("Получаем список собственных объявлений", func() {
		advertisementsBody = myAdvertisements.GetMyAdvertisements(s.T(), s.loginData.Token, http.StatusOK, constants.PageSize)
	})

	var items gjson.Result
	s.Run("Проверяем, что собственное тестовое объявление есть в списке", func() {
		items = gjson.Get(advertisementsBody, "items")
		require.NotEmpty(s.T(), items.Array())
	})

	s.Run("Проверяем, что наши объявления имеют принадлежность к нашему пользователю и объявления имеют НЕ пустые параметры", func() {
		var createdItem gjson.Result
		for _, item := range items.Array() {
			if item.Get("id").String() == createdID {
				createdItem = item
				break
			}
		}
		require.True(s.T(), createdItem.Exists(), "Созданное тестом объявление отсутствует в списке")
		require.Equal(s.T(), s.loginData.User.Id, createdItem.Get("user_id").String(), "user_id не совпадает")
		require.NotEmpty(s.T(), createdItem.Get("created_at").String(), "created_at пустое")
		require.NotEmpty(s.T(), createdItem.Get("updated_at").String(), "updated_at пустое")
		require.NotEmpty(s.T(), createdItem.Get("title").String(), "Название пустое")
		require.NotEmpty(s.T(), createdItem.Get("description").String(), "Описание пустое")
		require.NotEmpty(s.T(), createdItem.Get("photos").Array(), "Присутствуют пустые фотографии")
	})
}

// Negative: Проверяем 401 ошибку при попытке получить собственные объявления без токена
func (s *TestSuite) TestGetMyAdvertisementsIncorrectToken() {
	var advertisementsBody string
	s.Run("Получаем список собственных объявлений с неправильным авторизационным токеном", func() {
		advertisementsBody = myAdvertisements.GetMyAdvertisements(s.T(), constants.InvalidToken, http.StatusUnauthorized, constants.PageSize)
	})

	s.Run("Проверяем, что показывается правильная ошибка", func() {
		require.Equal(s.T(), constants.Unauthorized, gjson.Get(advertisementsBody, "error").String(),
			"Некорректное значение error параметра")
		require.Equal(s.T(), constants.InvalidAuthorizationMessage, gjson.Get(advertisementsBody, "message").String(),
			"Некорректное значение message параметра")
	})
}
