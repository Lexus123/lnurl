package handlers

import (
	"net/http"
)

/*
GetPrices ...
*/
func GetPaymentRequest(w http.ResponseWriter, r *http.Request) {
	// // Read JSON body of request and check for errors. For example: buy, eur, 100
	// bodyAsByte, err := ioutil.ReadAll(r.Body)
	// defer r.Body.Close()
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	// // Unmarshal the body and check for errors
	// // var message models.GetPricesData
	// err = json.Unmarshal(bodyAsByte, &message)
	// if err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }

	output := []byte("dikke")

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}
