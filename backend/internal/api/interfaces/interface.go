// @tg version=2.3.51
// @tg backend=PetStore
// @tg title=`PetStore API`

//go:generate tg transport --services . --out ../transport/httpserver --outSwagger swagger.yaml
package interfaces

// "context"

// PetStoreAPI
// @tg http-server jsonRPC-server log metrics
// @tg http-prefix=api/v1
type PetStoreApi interface {
	// GetBrands
	// @tg http-method=GET
	// @tg http-path=/getBrands
	// @tg summary=`Позволяет получить список брендов`
	// @tg desc=`Позволяет получить список брендов`
	// GetBrands(ctx context.Context) (brands []models.Brand, err error)

	// GetBrandByID
	// @tg http-method=GET
	// @tg http-path=/getBrandByID
	// @tg summary=`Позволяет получить бренд по айди`
	// @tg desc=`Позволяет получить бренд по айди`
	// GetBrandByID(ctx context.Context, brandID int) (brands []models.Brand, err error)
}