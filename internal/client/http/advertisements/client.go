package advertisements

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strconv"
	"testing"

	"github.com/steinfletcher/apitest"
	"github.com/stretchr/testify/require"

	"api-tests-template/internal/constants"
	"api-tests-template/internal/constants/path"
	apiRunner "api-tests-template/internal/helpers/api-runner"
)

// HttpPostAdvertisement отправляет поля и фотографии нового объявления.
func HttpPostAdvertisement(t *testing.T, token string, fields map[string]string, photos [][]byte) *http.Response {
	t.Helper()
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	for name, value := range fields {
		require.NoError(t, writer.WriteField(name, value), "Не удалось записать поле %s", name)
	}
	for index, photo := range photos {
		header := make(textproto.MIMEHeader)
		header.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"photos\"; filename=\"photo-%d.png\"", index+1))
		header.Set("Content-Type", constants.PhotoContentType)
		part, err := writer.CreatePart(header)
		require.NoError(t, err, "Не удалось добавить фотографию")
		_, err = part.Write(photo)
		require.NoError(t, err, "Не удалось записать фотографию")
	}
	require.NoError(t, writer.Close(), "Не удалось завершить multipart-запрос")
	return apiRunner.GetRunner().Auth(token).Create().Post(path.AdvertisementPath).
		ContentType(writer.FormDataContentType()).Body(body.String()).Expect(t).End().Response
}

// HttpGetAdvertisement получает объявление по идентификатору.
func HttpGetAdvertisement(t *testing.T, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementPath).
		Query("id", id).Expect(t).End().Response
}

// HttpGetAdvertisements ищет объявления по полному названию.
func HttpGetAdvertisements(t *testing.T, title string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Create().Get(path.AdvertisementsPath).
		Query("search", title).Query("limit", strconv.Itoa(constants.PageSize)).Expect(t).End().Response
}

// HttpGetPhotos получает список фотографий объявления.
func HttpGetPhotos(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Get(fmt.Sprintf(path.AdvertisementPhotosPath, id)).
		Expect(t).End().Response
}

// HttpGetPhotoFile скачивает фотографию по адресу из ответа API.
func HttpGetPhotoFile(t *testing.T, photoURL string) *http.Response {
	t.Helper()
	return apitest.New().EnableNetworking().Get(photoURL).Expect(t).End().Response
}

// HttpDeleteAdvertisement удаляет объявление его владельцем.
func HttpDeleteAdvertisement(t *testing.T, token, id string) *http.Response {
	t.Helper()
	return apiRunner.GetRunner().Auth(token).Create().Delete(path.AdvertisementPath).
		Query("id", id).Expect(t).End().Response
}
