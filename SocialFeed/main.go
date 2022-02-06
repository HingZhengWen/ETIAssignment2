package main

import (
	"encoding/json"
	"fmt"
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)
type Feed struct{
	StudentID string
	Feeddata string
	FeedID string
}
func main(){
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/socialfeed/{studentid}", feed).Methods("GET","PUT","POST","DELETE")
	router.HandleFunc("/",test)
	fmt.Println("Listening on port 8061")
	http.ListenAndServe(":8061", router)
}
func test(w http.ResponseWriter, r *http.Request){
	feed := "test"
	b, err := json.Marshal(feed)
		if err != nil{
			fmt.Println(err)
		}
	w.Write([]byte(string(b)))
}
func feed(w http.ResponseWriter, r *http.Request ){	
	params := mux.Vars(r)
	db, err := sql.Open("mysql", "user:password@tcp(10.31.11.11:8062)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	query := fmt.Sprintf("Select * FROM Feed where StudentID = '%s'",params["studentid"])
	results, err := db.Query(query)
	w.Header().Set("Content-Type","application/json")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var feeds Feed
		var errormsg []string
		var feedspass []Feed
		err = results.Scan(&feeds.StudentID,&feeds.Feeddata,&feeds.FeedID)
		if err != nil {
			panic(err.Error()) 
		}     
		feedspass = append(feedspass,Feed{StudentID:feeds.StudentID,Feeddata:feeds.Feeddata,FeedID:feeds.FeedID})
		type Feeds []Feed
		feed := Feeds{
		Feed{
			StudentID:feeds.StudentID,Feeddata:feeds.Feeddata,FeedID:feeds.FeedID},
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader((http.StatusOK))
		b, err := json.Marshal(feed)
		if err != nil{
			fmt.Println(err)
		}
		w.Write([]byte(string(b)))
			errormsg = append(errormsg)
	}
}