ARSHJOT SAINI

FILE STRUCTURE
- Db:-        *Arshjot_xxx.db
- Assets:-    *Project.xlxs
              *Nameofimage.png
- Templates:- *index.html
              *new_company.html
              *map.html
              *company.html
- Utils:-     *errorhandling.go
              *handler.go
              *helper.go
              *maps.go
              *sqlstatment.go

- HOW TO RUN THE APPLICATION
              * To run the project you need to install these dependencies in the terminal
                go get github.com/mattn/go-sqlite3
                go get github.com/xuri/excelize/v2
                go get github.com/gorilla/mux
              * After running these dependencies type go run main.go in terminal to run the project
       (or you an just type these dependencies in the import and hover over them to import the dependencies)

- OPERATIONS
    - READ:- When you load the application you will land on the tabel that displays the list of companies.To view any of them click on view button.
    - INSERT:- Click on new company option on the top and fill the information for the new company you want to add and then press insert button at the bottom. 
    - UPDATE:- Choose any field you want to update and make changes according to the need and then click on the update button at the bottom
    - DELETE:- Click the delete button in front of any field that want you want to erase permanently.

- REFERENCES
     *https://medium.com/@kamruljpi/web-development-in-go-with-http-server-part-14-61cb95304db3
     *https://getbootstrap.com/docs/4.0/components/list-group/
     *https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/
     *https://medium.com/@ansujain/how-to-import-functions-across-files-in-a-go-project-using-modules-13bb84f896e6

- NOTE:- According to me using SQL Database is very much easier than excel beacuse we can manage the database very easily as SQL is very much powerful tool.           
         SQL is also more user friendly than excel.


