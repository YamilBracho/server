package repository

import (
	"database/sql"
	"errors"

	"github.com/YamilBracho/server/models"
	_ "github.com/glebarez/go-sqlite"
)

var DB *sql.DB

// --------------------------------------------------------------------
// SetConnection
// Opens Database connection and sets DB global var
//
// --------------------------------------------------------------------
func SetConnection() error {
	db, err := sql.Open("sqlite", "../data/data.db")
	if err != nil {
		return err
	}

	DB = db

	return nil
}

// --------------------------------------------------------------------
// GetProvincias
// Gets list of Provincias
// --------------------------------------------------------------------
func GetProvincias() ([]*models.Provincia, error) {
	rows, err := DB.Query("select Id, Nombre from provincias")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Provincia
	for rows.Next() {
		var row models.Provincia

		err := rows.Scan(&row.Id, &row.Nombre)
		if err != nil {
			return nil, err
		}

		result = append(result, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// --------------------------------------------------------------------
// GetDistritosByProvincia
// Get list of Distritos by Provincia
// --------------------------------------------------------------------
func GetDistritosByProvincia(request models.DistritoRequest) ([]*models.Distrito, error) {

	query := "select Id, ProvinciaId, Nombre from distritos where ProvinciaId = ?"
	rows, err := DB.Query(query, request.ProvinciaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Distrito
	for rows.Next() {
		var row models.Distrito

		err := rows.Scan(&row.Id, &row.ProvinciaId, &row.Nombre)
		if err != nil {
			return nil, err
		}

		result = append(result, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("No data found.")
	}

	return result, nil
}

// --------------------------------------------------------------------
// GetCorregimientosByDistrito
// Get list of Corregimientos by Distrito
// --------------------------------------------------------------------
func GetCorregimientosByDistrito(request models.CorregimientoRequest) ([]*models.Corregimiento, error) {

	query := "select Id, DistritoId, Nombre from corregimientos where DistritoId = ?"
	rows, err := DB.Query(query, request.DistritoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Corregimiento
	for rows.Next() {
		var row models.Corregimiento

		err := rows.Scan(&row.Id, &row.DistritoId, &row.Nombre)
		if err != nil {
			return nil, err
		}

		result = append(result, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("No data found.")
	}

	return result, nil

}

// --------------------------------------------------------------------
// GetBarriosByCorregimiento
// Get list of BArrios by Corregimiento
// --------------------------------------------------------------------
func GetBarriosByCorregimiento(request models.BarrioRequest) ([]*models.Barrio, error) {
	query := "select Id, CorregimientoId, Nombre, FullName from barrios where CorregimientoId = ?"
	rows, err := DB.Query(query, request.CorregimientoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Barrio
	for rows.Next() {
		var row models.Barrio

		err := rows.Scan(&row.Id, &row.CorregimientoId, &row.Nombre, &row.FullName)
		if err != nil {
			return nil, err
		}

		result = append(result, &row)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("No data found.")
	}

	return result, nil

}
