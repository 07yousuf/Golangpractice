package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, r *http.Request) {
	SecondDigit := time.Now().Second()
	picture1 := []byte(`
	<html>
	<head>
	<strong>Israeli Genozide on Palastinian School</strong>
	</head>
	<body>
	<img src="https://www.aljazeera.com/wp-content/uploads/2024/08/AFP__20240810__36CW93C__v4__HighRes__PalestinianIsraelConflict-1723275736.jpg?resize=770%2C513&quality=80" width="500" height="300"/>
	</body>
	</html>`)
	picture2 := []byte(`<html><head><strong>Israeli Genozide on Palastinian Hospital</strong></head><body><img src="https://www.aljazeera.com/wp-content/uploads/2023/10/5Q8B0151-copy-1697566650.jpg?resize=1170%2C780&quality=80" width="500" height="300" /></body></html>`)
	picture3 := []byte(`<html><head><strong>Israeli Genozide on Palastinian Hospital</strong<body><img src="https://www.aljazeera.com/wp-content/uploads/2023/12/343K6VP-highres-1703365130.jpg?fit=1170%2C780&quality=80" width="500" height="300"/></body></html>`)
	picture4 := []byte(`<html><head><strong>Israeli Genozide on Palastinian Refugee Camp</strong<body><img src="https://www.aljazeera.com/wp-content/uploads/2023/12/2023-12-04T110336Z_896420382_RC2P34A28JZ0_RTRMADP_3_YEAR-END-ISRAEL-PALESTINIANS-1703365513.jpg?fit=1170%2C780&quality=80" width="500" height="300" /></body></html>`)
	picture5 := []byte(`<html><head><strong>Israeli Genozide on PRESS </strong<body><img src="https://www.thenation.com/wp-content/uploads/2023/12/Funeral_in_Gaza-getty.jpg" width="500" height="300" /></body></html>`)
	picture6 := []byte(`<html><head><strong>Israeli Genozide on Palastinian Hospital</strong><body><img src="https://www.aljazeera.com/wp-content/uploads/2024/04/5-1714035766.png?fit=1170%2C780&quality=80" width="500" height="300" /></body></html>`)

	if SecondDigit <= 10 {
		w.Write(picture1)
	} else if SecondDigit > 10 && SecondDigit <= 20 {
		w.Write(picture2)
	} else if SecondDigit > 20 && SecondDigit <= 30 {
		w.Write(picture3)
	} else if SecondDigit > 30 && SecondDigit <= 40 {
		w.Write(picture4)
	} else if SecondDigit > 40 && SecondDigit <= 50 {
		w.Write(picture5)
	} else if SecondDigit > 50 && SecondDigit <= 60 {
		w.Write(picture6)
	} else {
		fmt.Println("There is no chance")
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/genozide", handler).Methods("GET")

	fmt.Println("Server starting port at 8001.")
	log.Fatal(http.ListenAndServe("127.0.0.1:8001", r))
}
