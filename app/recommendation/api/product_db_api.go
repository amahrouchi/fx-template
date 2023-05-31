package recommendationApiService

import (
	"database/sql"
	recommendationModel "github.com/ekkinox/fx-template/app/recommendation/model"
	"github.com/ekkinox/fx-template/modules/fxlogger"
	"gorm.io/gorm"
)

// ProductDbApi service getting product data from the database.
type ProductDbApi struct {
	gorm   *gorm.DB
	logger *fxlogger.Logger
}

// GetMany gets many products from the database.
func (p *ProductDbApi) GetMany(ids []int) ([]recommendationModel.RecommendationProduct, error) {
	// get db connection from gorm
	db, err := p.gorm.DB()
	if err != nil {
		p.logger.Err(err).Msg("Unable to get db connection from gorm while getting recommended products.")
		return nil, err
	}

	// Query products from the database
	rows, err := db.Query("SELECT id, title FROM posts")
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
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			p.logger.Err(err).Msg("Unable to scan row while mapping products")
			return nil, err
		}

		list = append(list, recommendationModel.RecommendationProduct{
			Id:   id,
			Name: name,
		})
	}
	if err := rows.Err(); err != nil {
		p.logger.Err(err).Msg("Unable to scan all rows while mapping products")
		return nil, err
	}

	return list, nil
}
