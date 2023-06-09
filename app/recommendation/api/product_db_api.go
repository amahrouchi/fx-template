package recommendationApiService

import (
	"database/sql"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// ProductDbApi service getting product data from the database.
type ProductDbApi struct {
	gorm *gorm.DB
	link LinkGeneratorContract
}

// GetMany gets many products from the database.
func (p *ProductDbApi) GetMany(ids []int, lang string) ([]*recommendationModel.RecommendationProduct, error) {
	// Get db connection from gorm
	db, err := p.gorm.DB()
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
	query :=
		"SELECT " +
			"P.id, " +
			"P.name, " +
			// product name (translated)
			"(SELECT translation FROM product_translations WHERE model_id = P.id AND `key`='name' AND language = ?) AS translatedName, " +
			"B.id as brandId, " +
			"B.name AS brandName " +
			"FROM products AS P " +
			"INNER JOIN brands AS B ON B.id = P.brand_id " +
			"WHERE P.id IN (" + joinedIds + ")" +
			"ORDER BY FIELD(P.id, " + joinedIds + ")"

	// Prepare the query
	statement, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	// Execute the query
	rows, err := statement.Query(lang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return p.mapRows(rows)
}

// mapRows maps the sql rows to a list of products.
func (p *ProductDbApi) mapRows(rows *sql.Rows) ([]*recommendationModel.RecommendationProduct, error) {
	var list []*recommendationModel.RecommendationProduct
	for rows.Next() {
		// Get the data from the row
		var productId int
		var productName string
		var productTranslatedName any
		var brandId int
		var brandName string
		err := rows.Scan(
			&productId,
			&productName,
			&productTranslatedName,
			&brandId,
			&brandName,
		)
		if err != nil {
			return nil, err
		}

		// Convert the translated name to a string
		displayedName := productName
		productTranslatedNameStr, ok := productTranslatedName.(string)
		if ok {
			displayedName = productTranslatedNameStr
		}

		// Create the product and append it to the list
		list = append(list, &recommendationModel.RecommendationProduct{
			Type: recommendationEnum.ProductRecommendation,
			Id:   productId,
			Name: displayedName,
			Link: p.link.GetProductLink(productId, displayedName, brandId, brandName),
			Brand: recommendationModel.RecommendationProductBrand{
				Id:   brandId,
				Name: brandName,
				Link: p.link.GetBrandLink(brandId, brandName),
			},
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}
