package utils

import (
	"database/sql"
	"os"
	"time"
)

const FILEPATH = "./db/arshjot_homework.db"

var db_conn *sql.DB

type Details struct {
	Id          int
	CompanyName string
	PostingAge  string
	JobID       string
	Country     string
	Location    string
	Publication string
	SalaryMax   string
	SalaryMin   string
	SalaryType  string
	JobTitle    string
	Createddate time.Time
}

func Dbinit() {

	// make direct database connection, if db file already exists
	if IsFileExist(FILEPATH) {
		// initiate database connection only
		db_conn = Make_db_connection()

	} else {
		// if db file doesn't exist then, create file, db tables and populate data to tables from scratch
		Create_db_file()

		// initiate database connection only
		db_conn = Make_db_connection()

		// initiating table creation
		CreateCompanyTable(db_conn)

		load_xlsx_data := Load_xlsx()
		var rows = load_xlsx_data[1:]

		for _, v := range rows {
			PopulateCompanyTable(db_conn, v)
		}

	}
}

// method to create db file
func Create_db_file() {
	// creating new database file
	createDb, err := os.Create(FILEPATH)
	ThrowError(err)

	// closing the file open connection
	createDb.Close()
}

// method to create and open sqlite3 database connection
func Make_db_connection() (db *sql.DB) {
	// opening sqlite3 database connection
	open_db, err := sql.Open("sqlite3", FILEPATH)
	ThrowError(err)
	return open_db
}

// method to create new company table
func CreateCompanyTable(db *sql.DB) {
	sqlStatement := `CREATE TABLE COMPANY (
		"ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"COMPANYNAME" VARCHAR500,
		"POSTINGAGE" VARCHAR250,
		"JOBID" VARCHAR250,
		"COUNTRY" VARCHAR20,
		"LOCATION" VARCHAR250,
		"PUBLICATIONDATE" VARCHAR250,
		"SALARYMAX" INTEGER,
		"SALARYMIN" INTEGER,
		"SALARYTYPE" VARCHAR250,
		"JOBTITLE" TEXT,
		"CREATEDDATE" DATETIME
	);`

	prepareStatement, err := db.Prepare(sqlStatement)
	ThrowError(err)
	prepareStatement.Exec()
}

// method to insert data to company table
func PopulateCompanyTable(db *sql.DB, details []string) {
	sqlStatement := `INSERT INTO COMPANY (
		COMPANYNAME, POSTINGAGE, JOBID, COUNTRY, LOCATION, PUBLICATIONDATE, SALARYMAX, SALARYMIN, SALARYTYPE, JOBTITLE, CREATEDDATE
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`
	prepareStatement, err := db.Prepare(sqlStatement)
	ThrowError(err)
	_, err = prepareStatement.Exec(details[0], details[1], details[2], details[3],
		details[4], details[5], details[6], details[7], details[8], details[9], time.Now())
	ThrowError(err)
}

// fetching data from comoany table
func FetchCompanyData(db *sql.DB) (details []Details) {
	execStatement, err := db.Query("SELECT * FROM COMPANY ORDER BY CREATEDDATE DESC")
	ThrowError(err)
	var detailsList []Details
	for execStatement.Next() {
		var details Details
		execStatement.Scan(&details.Id, &details.CompanyName, &details.PostingAge, &details.JobID, &details.Country, &details.Location,
			&details.Publication, &details.SalaryMax, &details.SalaryMin, &details.SalaryType, &details.JobTitle, &details.Createddate)

		detailsList = append(detailsList, Details{Id: details.Id, CompanyName: details.CompanyName, PostingAge: details.PostingAge, JobID: details.JobID,
			Country: details.Country, Location: details.Location, Publication: details.Publication, SalaryMax: details.SalaryMax, SalaryMin: details.SalaryMin,
			SalaryType: details.SalaryType, JobTitle: details.JobTitle, Createddate: details.Createddate})
	}

	return detailsList
}

func GetCompanyById(db *sql.DB, id string) (rec Details) {
	err := db.QueryRow("SELECT * FROM COMPANY WHERE id = ?", id).Scan(&rec.Id, &rec.CompanyName, &rec.PostingAge, &rec.JobID, &rec.Country, &rec.Location,
		&rec.Publication, &rec.SalaryMax, &rec.SalaryMin, &rec.SalaryType, &rec.JobTitle, &rec.Createddate)
	ThrowError(err)
	return
}

// insert single record to the company table
func InsertRecord(db *sql.DB, details Details) int64 {
	sqlStatement := `INSERT INTO COMPANY (
		COMPANYNAME, POSTINGAGE, JOBID, COUNTRY, LOCATION, PUBLICATIONDATE, SALARYMAX, SALARYMIN, SALARYTYPE, JOBTITLE, CREATEDDATE
	) VALUES (
		?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
	);`
	prepareStatement, err := db.Prepare(sqlStatement)
	ThrowError(err)
	res, err := prepareStatement.Exec(details.CompanyName, details.PostingAge, details.JobID, details.Country,
		details.Location, details.Publication, details.SalaryMax, details.SalaryMin, details.SalaryType, details.JobTitle, details.Createddate)
	ThrowError(err)
	lstInsId, err := res.LastInsertId()
	ThrowError(err)
	return lstInsId
}

// delete single record to the company table
func DeleteRecord(db *sql.DB, id string) {
	sqlStatement := `DELETE FROM COMPANY WHERE ID = ?;`
	prepareStatement, err := db.Prepare(sqlStatement)
	ThrowError(err)
	_, err = prepareStatement.Exec(id)
	ThrowError(err)
}

// update single record to the company table
func UpdateRecord(db *sql.DB, details Details, id string) {
	sqlStatement := `UPDATE COMPANY SET COMPANYNAME = ?, POSTINGAGE = ?, JOBID = ?, COUNTRY = ?, LOCATION = ?, 
	PUBLICATIONDATE = ?, SALARYMAX = ?, SALARYMIN = ?, SALARYTYPE = ?, JOBTITLE = ?
	WHERE ID = ?;`
	prepareStatement, err := db.Prepare(sqlStatement)
	ThrowError(err)
	_, err = prepareStatement.Exec(details.CompanyName, details.PostingAge, details.JobID, details.Country,
		details.Location, details.Publication, details.SalaryMax, details.SalaryMin, details.SalaryType, details.JobTitle, id)
	ThrowError(err)
}

func CheckCompanyExist(db *sql.DB, id string) (flag bool) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM COMPANY WHERE ID = ?", id).Scan(&count)
	ThrowError(err)
	if count > 0 {
		return true
	} else {
		return false
	}

}
