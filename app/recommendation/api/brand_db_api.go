package recommendationApiService

import (
	"database/sql"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// BrandDbApi service getting brand data from the database.
type BrandDbApi struct {
	gorm *gorm.DB
}

// GetMany gets many brands from the database.
func (b *BrandDbApi) GetMany(ids []int) ([]*recommendationModel.RecommendationBrand, error) {
	// Get db connection from gorm
	db, err := b.gorm.DB()
	if err != nil {
		return nil, err
	}

	// Join the ids into a string
	var strIds []string
	for _, id := range ids {
		strIds = append(strIds, strconv.Itoa(id))
	}
	joinedIds := strings.Join(strIds, ",")

	// Query products from the database
	query := "SELECT " +
		"id, " +
		"name " +
		"FROM brands AS B " +
		"WHERE id IN (" + joinedIds + ")" +
		"ORDER BY FIELD(id, " + joinedIds + ")"

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return b.mapRows(rows)
}

// mapRows maps the sql rows to a list of products.
func (b *BrandDbApi) mapRows(rows *sql.Rows) ([]*recommendationModel.RecommendationBrand, error) {
	var list []*recommendationModel.RecommendationBrand
	for rows.Next() {
		// Get the data from the row
		var brandId int
		var brandName string
		err := rows.Scan(
			&brandId,
			&brandName,
		)
		if err != nil {
			return nil, err
		}

		// Create the product and append it to the list
		list = append(list, &recommendationModel.RecommendationBrand{
			Type: recommendationEnum.BrandRecommendation,
			Id:   brandId,
			Name: brandName,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
