package constants

const (
	TitlePrefix                 = "Lab_"
	Description                 = "Описание тестового объявления"
	Price                       = "1500.50"
	Quantity                    = "3"
	MaxTitleLength              = 50
	MaxDescriptionLength        = 500
	MaxPrice                    = 1000000
	MaxQuantity                 = 100
	MinPhotos                   = 1
	MaxPhotos                   = 3
	PageSize                    = 50
	FavoritesLimit              = 50
	PhotoContentType            = "image/png"
	InvalidToken                = "incorrect_token"
	MissingAdvertisementID      = "00000000-0000-4000-8000-000000000001"
	InvalidAdvertisementID      = "not-an-id"
	InvalidRequest              = "invalid_request"
	Unauthorized                = "unauthorized"
	Forbidden                   = "forbidden"
	NotFound                    = "not_found"
	AlreadyExists               = "already_exists"
	MissingTokenMessage         = "Authorization header required"
	InvalidAuthorizationMessage = "Invalid or expired token"
	PriceRangeMessage           = "price must be between 0 and 1000000 inclusive"
	QuantityRangeMessage        = "quantity must be between 1 and 100 inclusive"
)
