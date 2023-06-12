package recommendationApiService

import (
	"database/sql"
	recommendationEnum "github.com/ekkinox/fx-template/app/recommendation/enum"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// ProductDbApi service getting product data from the database.
type ProductDbApi struct {
	gorm   *gorm.DB
	link   LinkGeneratorContract
	logger *fxlogger.Logger
}

// GetMany gets many products from the database.
func (p *ProductDbApi) GetMany(productIds []int, lang string) ([]*recommendationModel.RecommendationProduct, error) {
	// Get db connection from gorm
	db, err := p.gorm.DB()
	if err != nil {
		return nil, err
	}

	// Join the productIds into a string
	var strIds []string
	for _, id := range productIds {
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

	// Map the rows to a list of products
	productList, err := p.mapRows(rows)
	if err != nil {
		return nil, err
	}

	// Populate images
	err = p.fetchImages(productIds, productList)
	if err != nil {
		return nil, err
	}

	return productList, nil
}

// mapRows maps the sql rows to a list of products.
func (p *ProductDbApi) mapRows(rows *sql.Rows) ([]*recommendationModel.RecommendationProduct, error) {
	var list []*recommendationModel.RecommendationProduct
	for rows.Next() {
		// Get the data from the row
		var productId, brandId int
		var productName, brandName string
		var productTranslatedName any
		err := rows.Scan(
			&productId, &productName, &productTranslatedName,
			&brandId, &brandName,
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

// fetchImages populates the images into the product list.
func (p *ProductDbApi) fetchImages(
	productIds []int,
	productRecos []*recommendationModel.RecommendationProduct,
) error {
	// Get db connection from gorm
	db, err := p.gorm.DB()
	if err != nil {
		return err
	}

	// Join the productIds into a string
	var strProductIds []string
	for _, id := range productIds {
		strProductIds = append(strProductIds, strconv.Itoa(id))
	}
	joinedIds := strings.Join(strProductIds, ",")

	// Query products from the database
	query := "SELECT P.id, PI.filename " +
		"FROM products AS P " +
		"INNER JOIN product_images AS PI ON PI.product_id = P.id " +
		"WHERE P.id IN (" + joinedIds + ") " +
		"AND PI.product_variant_id IS NULL " +
		"AND PI.deleted_at IS NULL " +
		"ORDER BY PI.product_id, PI.`order`"

	// Prepare the query
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Map the images per product
	imageMap := make(map[int][]string)
	for rows.Next() {
		// Get the data from the row
		var productId int
		var filename string
		err := rows.Scan(&productId, &filename)
		if err != nil {
			return err
		}

		// Append the filename to the product
		imageMap[productId] = append(imageMap[productId], filename)
	}

	// Populate the images into the product list
	for _, product := range productRecos {
		product.Images = imageMap[product.Id]
	}

	return nil
}
