package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"solaris/models"
	"strconv"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func CreateElectricMeter(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var ElectricMeter models.ElectricMeter

	err := json.NewDecoder(r.Body).Decode(&ElectricMeter)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := insertElectricMeter(ElectricMeter)

	// format a response object
	res := response{
		ID:      insertID,
		Message: "Electric Meter created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func GetElectricMeter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	electricMeter, err := getElectricMeter(int64(id))

	if err != nil {
		log.Fatalf("Unable to get electric meter. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(electricMeter)
}

func GetAllElectricMeters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	electricMeters, err := getAllElectricMeters()

	if err != nil {
		log.Fatalf("Unable to get all electric meters. %v", err)
	}

	json.NewEncoder(w).Encode(electricMeters)
}

func UpdateElectricMeter(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	var electricMeter models.ElectricMeter

	err = json.NewDecoder(r.Body).Decode(&electricMeter)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateElectricMeter(int64(id), electricMeter)

	// format the message string
	msg := fmt.Sprintf("Electric meter updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteElectricMeter(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	deletedRows := deleteElectricMeter(int64(id))

	// format the message string
	msg := fmt.Sprintf("Electric meter updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------

func insertElectricMeter(electricMeter models.ElectricMeter) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// Insert SQL Querry
	sqlStatement := `INSERT INTO electric_meters (timestamp, import_kwh, export_kwh, comment) VALUES ($1, $2, $3, $4) RETURNING em_id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, time.Now(), electricMeter.Import_kWh, electricMeter.Export_kWh, electricMeter.Comment).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func getElectricMeter(id int64) (models.ElectricMeter, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var electricMeter models.ElectricMeter

	// create the select sql query
	sqlStatement := `SELECT * FROM electric_meters WHERE em_id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&electricMeter.ID, &electricMeter.Timestamp, &electricMeter.Import_kWh,
		&electricMeter.Export_kWh, &electricMeter.Comment)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return electricMeter, nil
	case nil:
		return electricMeter, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return electricMeter, err
}

func getAllElectricMeters() ([]models.ElectricMeter, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var electricMeters []models.ElectricMeter

	sqlStatement := `SELECT * FROM electric_meters`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var electricMeter models.ElectricMeter

		err = rows.Scan(&electricMeter.ID, &electricMeter.Timestamp, &electricMeter.Import_kWh,
			&electricMeter.Export_kWh, &electricMeter.Comment)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		electricMeters = append(electricMeters, electricMeter)

	}

	return electricMeters, err
}

func updateElectricMeter(id int64, electricMeter models.ElectricMeter) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE electric_meters SET timestamp=$2, import_kwh=$3, export_kwh=$4, comment=$5 WHERE em_id=$1`

	// (timestamp, import_kwh, export_kwh, comment) VALUES ($1, $2, $3, $4) RETURNING em_id`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, time.Now(), electricMeter.Import_kWh, electricMeter.Export_kWh, electricMeter.Comment)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func deleteElectricMeter(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM electric_meters WHERE em_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
