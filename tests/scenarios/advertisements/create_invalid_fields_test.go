package advertisements

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestCreateWithInvalidParameters() {
	cases := []struct {
		name    string
		field   string
		value   string
		photos  int
		message string
	}{
		{"long_title", "title", strings.Repeat("x", constants.MaxTitleLength+1), constants.MinPhotos,
			fmt.Sprintf("title must not exceed %d characters provided %d", constants.MaxTitleLength, constants.MaxTitleLength+1)},
		{"long_description", "description", strings.Repeat("x", constants.MaxDescriptionLength+1), constants.MinPhotos,
			"description must be between 1 and 500 characters"},
		{"negative_price", "price", "-1", constants.MinPhotos, constants.PriceRangeMessage},
		{"invalid_price", "price", "abc", constants.MinPhotos, "cannot convert string value abc to float"},
		{"high_price", "price", strconv.Itoa(constants.MaxPrice + 1), constants.MinPhotos, constants.PriceRangeMessage},
		{"negative_quantity", "quantity", "-1", constants.MinPhotos, constants.QuantityRangeMessage},
		{"invalid_quantity", "quantity", "abc", constants.MinPhotos, "cannot convert string value to int"},
		{"high_quantity", "quantity", strconv.Itoa(constants.MaxQuantity + 1), constants.MinPhotos, constants.QuantityRangeMessage},
		{"too_many_photos", "", "", constants.MaxPhotos + 1, "maximum 3 photos allowed"},
	}
	for _, item := range cases {
		s.Run(item.name, func() {
			t := s.T()
			t.Log("Подготовка: изменяем один параметр запроса")
			fields := base.AdvertisementFields()
			if item.field != "" {
				fields[item.field] = item.value
			}
			t.Log("Шаг: отправляем некорректные параметры объявления")
			body := advertisements.Create(t, s.loginData.Token, fields, base.PhotoFiles(item.photos), http.StatusBadRequest)
			base.AssertError(t, body, constants.InvalidRequest, item.message)
		})
	}
}
