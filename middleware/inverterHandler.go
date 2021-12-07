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

func CreateInverter(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var inverter models.Inverter

	err := json.NewDecoder(r.Body).Decode(&inverter)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	insertID := insertInverter(inverter)

	res := response{
		ID:      insertID,
		Message: "Inverter data created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetInverter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	inverter, err := getInverter(int64(id))

	if err != nil {
		log.Fatalf("Unable to get inverter. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(inverter)
}

func GetAllInverters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	inverters, err := getAllInverters()

	if err != nil {
		log.Fatalf("Unable to get all inverters. %v", err)
	}

	json.NewEncoder(w).Encode(inverters)
}

func UpdateInverter(w http.ResponseWriter, r *http.Request) {

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

	var inverter models.Inverter

	err = json.NewDecoder(r.Body).Decode(&inverter)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := updateInverter(int64(id), inverter)

	// format the message string
	msg := fmt.Sprintf("Inverter updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteInverter(w http.ResponseWriter, r *http.Request) {

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

	deletedRows := deleteInverter(int64(id))

	// format the message string
	msg := fmt.Sprintf("Inverter updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------

func insertInverter(inverter models.Inverter) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	sqlStatement := `INSERT INTO inverters (Date, Energy_per_Inverter_kWh, Energy_per_inverter_per_kWp, Total_system_kWh) 
		VALUES ($1, $2, $3, $4) RETURNING Inv_id`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, time.Now(), inverter.Energy_per_Inverter_kWh,
		inverter.Energy_per_inverter_per_kWp, inverter.Total_system_kWh).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

func getInverter(id int64) (models.Inverter, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var inverter models.Inverter

	// create the select sql query
	sqlStatement := `SELECT * FROM inverters WHERE inv_id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&inverter.ID, &inverter.Date, &inverter.Energy_per_Inverter_kWh,
		&inverter.Energy_per_inverter_per_kWp, &inverter.Total_system_kWh)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return inverter, nil
	case nil:
		return inverter, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return inverter, err
}

func getAllInverters() ([]models.Inverter, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var inverters []models.Inverter

	// create the select sql query
	sqlStatement := `SELECT * FROM inverters`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var inverter models.Inverter

		err = rows.Scan(&inverter.ID, &inverter.Date, &inverter.Energy_per_Inverter_kWh,
			&inverter.Energy_per_inverter_per_kWp, &inverter.Total_system_kWh)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		inverters = append(inverters, inverter)

	}

	return inverters, err
}

func updateInverter(id int64, inverter models.Inverter) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE inverters SET Date=$2, Energy_per_Inverter_kWh=$3, 
		Energy_per_inverter_per_kWp=$4, Total_system_kWh=$5 WHERE inv_id=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, inverter.Date, inverter.Energy_per_Inverter_kWh,
		inverter.Energy_per_inverter_per_kWp, inverter.Total_system_kWh)

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

func deleteInverter(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM inverters WHERE inv_id=$1`

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
