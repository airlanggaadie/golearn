package versions

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"golearn/m/v1/common"
	"time"
)

var (
	db *pgxpool.Pool
	tableName = "versions"
)

// VersionModel /** Define model for versions table
type VersionModel struct {
	Id			int			`db:"id" json:"id"`
	App 		string		`db:"app" json:"app"`
	Version 	string		`db:"version" json:"version"`
	Code		*int		`db:"code" json:"code"`
	Description string		`db:"description" json:"description"`
	IsActive	int			`db:"is_active" json:"is_active" default:"1"`
	CreatedAt	*time.Time	`db:"created_at" json:"created_at"`
	UpdatedAt	*time.Time	`db:"updated_at" json:"updated_at"`
}

// FindOneVersionById /** Define function to get one versions by id but not complete fields. only id, app, versions, and code
func FindOneVersionById(id interface{}) (VersionModel, error) {
	db = common.GetDB()
	var model VersionModel

	sql := fmt.Sprintf("SELECT id, app, version, code FROM %s WHERE is_active=true and id=$1",tableName)
	if err := db.
		QueryRow(context.Background(), sql,id).
		Scan(&model.Id, &model.App, &model.Version, &model.Code); err != nil {
			return VersionModel{}, err
	}

	return model, nil
}

// AllVersion /** Define function to get all versions but not complete fields. only id, app, versions, and code.
func AllVersion() ([]VersionModel, error) {
	db = common.GetDB()

	sql := fmt.Sprintf("SELECT id, app, version, code FROM %s",tableName)
	rows, err := db.
		Query(context.Background(),sql)
	if err != nil {
			return []VersionModel{}, err
	}
	defer rows.Close()

	var models []VersionModel
	for rows.Next() {
		var row VersionModel
		if err := rows.Scan(&row.Id, &row.App, &row.Version, &row.Code); err != nil {
			log.Debug().Stack().Err(err).Msgf("%v",err)
			return []VersionModel{},err
		}
		models = append(models, row)
	}
	return models, nil
}

type NewVersionsParams struct {
	App 		string		`json:"app"`
	Version 	string		`json:"version"`
	Code		int			`json:"code"`
	Description string		`json:"description"`
}

func InsertVersion(params NewVersionsParams) (bool,error) {
	db = common.GetDB()
	sql := fmt.Sprintf("INSERT INTO %s (app,version,code,description) VALUES ($1,$2,$3,$4);",tableName)
	_, err :=db.Exec(context.Background(),sql,params.App,params.Version,params.Code,params.Description)
	if err != nil {
		log.Debug().Stack().Err(err).Msgf("%v",err)
		return false, err
	}
	return true, nil
}