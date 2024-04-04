package utils

import (
	"fmt"
	"net/http"
)

func ThrowError(err error) {
	if err != nil {
		fmt.Println("Something went wrong !!! Please try again.", err)
		return
	}
}

func ThrowHttpError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
