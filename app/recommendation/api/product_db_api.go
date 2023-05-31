package recommendationApiService

import (
	"database/sql"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// ProductDbApi service getting product data from the database.
type ProductDbApi struct {
	gorm   *gorm.DB
	logger *fxlogger.Logger
}

// GetMany gets many products from the database.
func (p *ProductDbApi) GetMany(ids []int) ([]recommendationModel.RecommendationProduct, error) {
	// Get db connection from gorm
	db, err := p.gorm.DB()
	if err != nil {
		p.logger.Err(err).Msg("Unable to get db connection from gorm while getting recommended products.")
		return nil, err
	}

	// Join the ids into a string
	var strIds []string
	for _, id := range ids {
		strIds = append(strIds, strconv.Itoa(id))
	}
	joinedIds := strings.Join(strIds, ",")

	// Query products from the database
	query := "SELECT P.id, P.name, B.id as brandId, B.name AS brandName " +
		"FROM products AS P " +
		"INNER JOIN brands AS B ON B.id = P.brand_id " +
		"WHERE P.id IN (" + joinedIds + ")"
	p.logger.Debug().Str("query", query).Msg("Querying the database to get recommended products.")
	rows, err := db.Query(query)
	if err != nil {
		p.logger.Err(err).Msg("Unable to query the database to get recommended products.")
		return nil, err
	}
	defer rows.Close()

	return p.mapRows(rows)
}

// mapRows maps the sql rows to a list of products.
func (p *ProductDbApi) mapRows(rows *sql.Rows) ([]recommendationModel.RecommendationProduct, error) {
	var list []recommendationModel.RecommendationProduct
	for rows.Next() {
		// Get the data from the row
		var productId int
		var productName string
		var brandId int
		var brandName string
		err := rows.Scan(&productId, &productName, &brandId, &brandName)
		if err != nil {
			p.logger.Err(err).Msg("Unable to scan row while mapping products")
			return nil, err
		}

		// Create the product and append it to the list
		list = append(list, recommendationModel.RecommendationProduct{
			Id:   productId,
			Name: productName,
			Brand: recommendationModel.RecommendationProductBrand{
				Id:   brandId,
				Name: brandName,
			},
		})
	}
	if err := rows.Err(); err != nil {
		p.logger.Err(err).Msg("Unable to scan all rows while mapping products")
		return nil, err
	}

	return list, nil
}
