package favorites

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	base "api-tests-template/tests"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
	ownerData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

// SetupSuite авторизует пользователя избранного и отдельного владельца объявлений.
func (s *TestSuite) SetupSuite() {
	base.SetupSuite()
	base.Precondition("Авторизация двух пользователей для тестов избранного")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
	s.ownerData = auth.Login(s.T(), os.Getenv("TEST_OWNER_LOGIN"), os.Getenv("TEST_OWNER_PASSWORD"))
	require.NotEqual(s.T(), s.loginData.User.Id, s.ownerData.User.Id, "Для избранного нужны два разных пользователя")
}

// TearDownSuite завершает работу набора тестов.
func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}
