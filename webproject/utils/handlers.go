package utils

import (
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type pageInfo struct {
	Code string
	Info string
}

type homePageStruct struct {
	PageInfo    pageInfo
	CompanyList []Details
}

type companyPageStruct struct {
	PageInfo      pageInfo
	SingleCompany Details
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	info := r.URL.Query().Get("msg")
	code := r.URL.Query().Get("code")

	PageInfo1 := pageInfo{
		Code: code,
		Info: info,
	}
	// Query all users from the database
	details := FetchCompanyData(db_conn)
	homePageData := homePageStruct{
		PageInfo:    PageInfo1,
		CompanyList: details,
	}
	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/index.html")
	ThrowHttpError(w, err)

	// Execute the template
	err = tmpl.Execute(w, struct {
		HomePageData homePageStruct
	}{homePageData})
	ThrowHttpError(w, err)
}

// handler
func CompanyHandler(w http.ResponseWriter, r *http.Request) {
	cId := r.URL.Query().Get("c_id")
	PageInfo := pageInfo{}
	// Check if the request is a POST request
	if r.Method == http.MethodPost {
		// Parse the form data
		err := r.ParseForm()
		ThrowHttpError(w, err)

		cIdInt, err := strconv.Atoi(r.Form.Get("compid"))
		ThrowHttpError(w, err)

		cIdStr := strconv.Itoa(cIdInt)

		var newCompanyDetails Details
		newCompanyDetails.Id = cIdInt
		newCompanyDetails.CompanyName = r.Form.Get("companyname")
		newCompanyDetails.PostingAge = r.Form.Get("postingage")
		newCompanyDetails.JobID = r.Form.Get("jobid")
		newCompanyDetails.Country = r.Form.Get("country")
		newCompanyDetails.Location = r.Form.Get("location")
		newCompanyDetails.Publication = r.Form.Get("publication")
		newCompanyDetails.SalaryMax = r.Form.Get("salarymax")
		newCompanyDetails.SalaryMin = r.Form.Get("salarymin")
		newCompanyDetails.SalaryType = r.Form.Get("salarytype")
		newCompanyDetails.JobTitle = r.Form.Get("jobtitle")

		if CheckCompanyExist(db_conn, cIdStr) {
			PageInfo = pageInfo{
				Code: "2",
				Info: "Company has been updated successfully",
			}
			UpdateRecord(db_conn, newCompanyDetails, cIdStr)
		} else {
			PageInfo = pageInfo{
				Code: "1",
				Info: "Error: Company doesnot exist in our system",
			}
		}
		redirectUrl := "?msg=" + PageInfo.Info + "&code=" + PageInfo.Code + "&c_id=" + cIdStr

		http.Redirect(w, r, "/company"+redirectUrl, http.StatusSeeOther)

	}

	info := r.URL.Query().Get("msg")
	code := r.URL.Query().Get("code")

	PageInfo1 := pageInfo{
		Code: code,
		Info: info,
	}

	detailRec := GetCompanyById(db_conn, cId)

	companyPageData := companyPageStruct{
		PageInfo:      PageInfo1,
		SingleCompany: detailRec,
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/company.html")
	ThrowHttpError(w, err)

	// Execute the template
	err = tmpl.Execute(w, struct {
		CompanyPageData companyPageStruct
	}{companyPageData})
	ThrowHttpError(w, err)
}

// handler
func NewCompanyHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a POST request
	if r.Method == http.MethodPost {
		PageInfo := pageInfo{}
		// Parse the form data
		err := r.ParseForm()
		ThrowHttpError(w, err)

		var newCompanyDetails Details
		newCompanyDetails.CompanyName = r.Form.Get("companyname")
		newCompanyDetails.PostingAge = r.Form.Get("postingage")
		newCompanyDetails.JobID = r.Form.Get("jobid")
		newCompanyDetails.Country = r.Form.Get("country")
		newCompanyDetails.Location = r.Form.Get("location")
		newCompanyDetails.Publication = r.Form.Get("publication")
		newCompanyDetails.SalaryMax = r.Form.Get("salarymax")
		newCompanyDetails.SalaryMin = r.Form.Get("salarymin")
		newCompanyDetails.SalaryType = r.Form.Get("salarytype")
		newCompanyDetails.JobTitle = r.Form.Get("jobtitle")
		newCompanyDetails.Createddate = time.Now()

		insId := InsertRecord(db_conn, newCompanyDetails)

		if insId > 0 {
			PageInfo = pageInfo{
				Code: "2",
				Info: "Company has been inserted successfully",
			}
		} else {
			PageInfo = pageInfo{
				Code: "1",
				Info: "Error when inserting company",
			}
		}

		http.Redirect(w, r, "?msg="+PageInfo.Info+"&code="+PageInfo.Code, http.StatusSeeOther)
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/new_company.html")
	ThrowHttpError(w, err)

	// Execute the template
	err = tmpl.Execute(w, nil)
	ThrowHttpError(w, err)
}

// handler
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	PageInfo := pageInfo{}

	cId := r.URL.Query().Get("c_id")
	if CheckCompanyExist(db_conn, cId) {
		DeleteRecord(db_conn, cId)
		PageInfo = pageInfo{
			Code: "2",
			Info: "Company has been deleted successfully",
		}

	} else {
		PageInfo = pageInfo{
			Code: "1",
			Info: "Error: Company doesnot exist in our system",
		}
	}
	redirectUrl := "?msg=" + PageInfo.Info + "&code=" + PageInfo.Code
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)

}

func MapHandler(w http.ResponseWriter, r *http.Request) {

	// Parse the HTML template
	tmpl, err := template.ParseFiles("templates/map.html")
	ThrowHttpError(w, err)

	// Execute the template
	err = tmpl.Execute(w, IMAGEPATH)
	ThrowHttpError(w, err)
}
