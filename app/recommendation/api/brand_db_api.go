package recommendationApiService

import (
	"database/sql"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// BrandDbApi service getting brand data from the database.
type BrandDbApi struct {
	gorm   *gorm.DB
	link   LinkGeneratorContract
	logger *fxlogger.Logger
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
		"name, " +
		"image_squared, " +
		"image_rounded, " +
		"image_large " +
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
		var imageSquared, imageRounded, imageLarge *string
		err := rows.Scan(
			&brandId, &brandName,
			&imageSquared, &imageRounded, &imageLarge,
		)
		if err != nil {
			return nil, err
		}

		// Create the product and append it to the list
		list = append(list, &recommendationModel.RecommendationBrand{
			Type: recommendationEnum.BrandRecommendation,
			Id:   brandId,
			Name: brandName,
			Link: b.link.GetBrandLink(brandId, brandName),
			Images: &recommendationModel.RecommendationBrandImage{
				Squared: imagePath(imageSquared),
				Rounded: imagePath(imageRounded),
				Large:   imagePath(imageLarge),
			},
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

// imagePath returns the image path if the filename is not nil.
func imagePath(filename *string) *string {
	if filename == nil || *filename == "" {
		return nil
	}

	return lo.ToPtr("/brands/squared/" + *filename + ".jpg")
}
