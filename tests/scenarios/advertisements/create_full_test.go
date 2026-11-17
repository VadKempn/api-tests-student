package advertisements

import (
	"net/http"
	"strconv"

	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/managers/advertisements"
	base "api-tests-template/tests"
)

func (s *TestSuite) TestCreateAdvertisementWithAllParameters() {
	t := s.T()
	t.Log("Подготовка: поля объявления и три фотографии")
	fields := base.AdvertisementFields()
	created := advertisements.Create(t, s.loginData.Token, fields, base.PhotoFiles(constants.MaxPhotos), http.StatusCreated)
	id := gjson.Get(created, "id").String()

	t.Log("Шаг: получаем созданное объявление по id")
	actual := advertisements.Get(t, id, http.StatusOK)
	for _, body := range []string{created, actual} {
		require.Equal(t, fields["title"], gjson.Get(body, "title").String(), "Название не сохранилось")
		require.Equal(t, fields["description"], gjson.Get(body, "description").String(), "Описание не сохранилось")
		price, err := strconv.ParseFloat(fields["price"], 64)
		require.NoError(t, err)
		require.Equal(t, price, gjson.Get(body, "price").Float(), "Цена не сохранилась")
		quantity, err := strconv.Atoi(fields["quantity"])
		require.NoError(t, err)
		require.Equal(t, int64(quantity), gjson.Get(body, "quantity").Int(), "Количество не сохранилось")
		require.Equal(t, s.loginData.User.Id, gjson.Get(body, "user_id").String(), "Неверный владелец объявления")
		require.NotEmpty(t, gjson.Get(body, "created_at").String(), "Нет времени создания")
		require.NotEmpty(t, gjson.Get(body, "updated_at").String(), "Нет времени изменения")
	}
	require.Equal(t, id, gjson.Get(actual, "id").String(), "Вернулось другое объявление")

	t.Log("Шаг: ищем объявление по полному названию")
	items := gjson.Get(advertisements.Search(t, fields["title"]), "items").Array()
	require.Len(t, items, 1, "По уникальному названию должно находиться одно объявление")
	require.Equal(t, id, items[0].Get("id").String(), "Созданное объявление не найдено")
	require.Equal(t, fields["title"], items[0].Get("title").String(), "Найдено неверное название")

	t.Log("Шаг: получаем фотографии и скачиваем их по возвращённым адресам")
	photos := gjson.Parse(advertisements.GetPhotos(t, s.loginData.Token, id)).Array()
	require.Len(t, photos, constants.MaxPhotos, "Сохранились не все фотографии")
	createdPhotos := gjson.Get(created, "photos").Array()
	actualPhotos := gjson.Get(actual, "photos").Array()
	require.Len(t, createdPhotos, constants.MaxPhotos)
	require.Len(t, actualPhotos, constants.MaxPhotos)
	for index, photo := range photos {
		require.NotEmpty(t, photo.Get("id").String(), "У фотографии отсутствует id")
		require.Equal(t, createdPhotos[index].Get("id").String(), photo.Get("id").String())
		require.Equal(t, actualPhotos[index].Get("id").String(), photo.Get("id").String())
		require.Equal(t, int64(index), photo.Get("sort_order").Int(), "Неверный порядок фотографий")
		photoURL := photo.Get("url").String()
		require.NotEmpty(t, photoURL, "У фотографии отсутствует URL")
		require.Equal(t, base.Photo, advertisements.GetPhotoFile(t, photoURL), "Скачанная фотография отличается от загруженной")
	}
	t.Log("Проверка: поля сохранились, поиск находит объявление, все фотографии доступны")
}
