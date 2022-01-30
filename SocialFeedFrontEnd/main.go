package main

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	"strconv"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)
type FollowingList struct{
	FollowingList []string
}
type FollowerList struct{
	FollowerList []string
}
type PostList struct{
	PostList []string
}
type SelectedPost struct{
	SelectedPost []string
}
type FeedIDList struct{
	FeedIDList []string
}
type Students struct{
	StudentID string
	StudentName string 
	DOB string
	Address string
	PhoneNumber string
}
type SearchStudents struct{
	StudentID string
	StudentName string 
	DOB string
	Address string
	PhoneNumber string
	Status bool
}
type Followers struct{
	StudentID string
	FollowerID string
}
type Feed struct{
	StudentID string
	Feeddata string
	FeedID string
}
var currentstudent Students
//Landing page
func index(w http.ResponseWriter, r *http.Request) {
	currentstudent.StudentID = "00001"
	currentstudent.StudentName = "Jayman Hing"
	currentstudent.DOB = "0.000499500499500499"
	currentstudent.Address = "Apt blk 121 lengkong tiga"
	currentstudent.PhoneNumber = "87776080"
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}
	listFollowings()
}
var lastsearchedstudent SearchStudents
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index)
	router.HandleFunc("/student_search", search_student)
	router.HandleFunc("/student_profile", student_profile)
	router.HandleFunc("/followinglist", listfollowing)
	router.HandleFunc("/followerlist", listfollowers)
	router.HandleFunc("/feedpage", feedpage)
	router.HandleFunc("/postpage", postpage)
	router.HandleFunc("/myposts", myposts)
	router.HandleFunc("/viewpost", alterposts)
	router.HandleFunc("/updatepostpage", updatepost)
	fmt.Println("Listening on port 8060")
	http.ListenAndServe(":8060", router)
}
func search_student(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/student_search.html"))
		tmpl.Execute(w, nil)
	} else{
		db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
		query := fmt.Sprintf("Select * FROM Students Where StudentID != '%s' && StudentName = '%s'",
		"00001",r.FormValue("s_name"))
		results, err := db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		for results.Next() {
			var student Students
			err = results.Scan(&student.StudentID, &student.StudentName, 
							   &student.DOB, &student.Address,&student.PhoneNumber)
			if err != nil {
				panic(err.Error()) 
			}
			lastsearchedstudent.StudentID = student.StudentID
			lastsearchedstudent.StudentName = student.StudentName
			lastsearchedstudent.DOB = student.DOB
			lastsearchedstudent.Address = student.Address
			lastsearchedstudent.PhoneNumber = student.PhoneNumber
			lastsearchedstudent.Status = false
			http.Redirect(w,r,"/student_profile",http.StatusFound)
		}
	// 00001 for now as i have no way to confirm student's current logon user
	// so for now it is 00001, as the user for this test would be user 00001
	}
}
func student_profile(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	if checkStudentFollowed() == true{
		if r.Method == "GET" {
			lastsearchedstudent.Status = true
			tmpl := template.Must(template.ParseFiles("./web/student_profile.html"))
			tmpl.Execute(w, lastsearchedstudent)
		}else{
			query := fmt.Sprintf("Delete FROM Followers Where followerID = '%s' && studentID = '%s'",
			"00001",lastsearchedstudent.StudentID)
			_, err := db.Query(query)  
			if err != nil {
				panic(err.Error())
			}
			http.Redirect(w,r,"/student_profile",http.StatusFound)
		}
	}else{
		if r.Method == "GET" {
			lastsearchedstudent.Status = false
			tmpl := template.Must(template.ParseFiles("./web/student_profile.html"))
			tmpl.Execute(w, lastsearchedstudent)
		}else{
			query := fmt.Sprintf("Insert INTO Followers Values('%s','%s')",
			lastsearchedstudent.StudentID,"00001")
			_, err := db.Query(query)  
			if err != nil {
				panic(err.Error())
			}
			http.Redirect(w,r,"/student_profile",http.StatusFound)
		}
	}
}
func checkStudentFollowed() bool{
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	var studentfollowed bool
	studentfollowed = false
	query := fmt.Sprintf("Select * FROM Followers Where followerID = '%s'",
	"00001")
	results, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var followers Followers
		err = results.Scan(&followers.StudentID, &followers.FollowerID)
		if err != nil {
			panic(err.Error()) 
		}
		if followers.StudentID == lastsearchedstudent.StudentID{
			studentfollowed = true
		}
	}
	return studentfollowed
}
var followinglist FollowingList
func listFollowings(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	var followinglist []Followers
	query := fmt.Sprintf("Select * FROM Followers Where followerID = '%s'",
	"00001")
	results, err := db.Query(query)
	if err != nil {
        panic(err.Error())
    }
    for results.Next() {
        var followings Followers
        err = results.Scan(&followings.StudentID, &followings.FollowerID)
        if err != nil {
            panic(err.Error()) 
        }
		followinglist = append(followinglist,followings)
	}
	getFollowingNames(followinglist,db)
}
func getFollowingNames(followings []Followers, db *sql.DB){
	var studentFollowingsList []string
	for i:= 0; i < len(followings); i++ {
		query := fmt.Sprintf("Select * FROM Students Where studentID = '%s'",
		followings[i].StudentID)
		results, err := db.Query(query)
		if err != nil {
        	panic(err.Error())
		}
		for results.Next() {
			var studentFollowings Students
			err = results.Scan(&studentFollowings.StudentID, &studentFollowings.StudentName,
				 &studentFollowings.DOB, &studentFollowings.Address, &studentFollowings.PhoneNumber)
			if err != nil {
				panic(err.Error()) 
			}
			studentFollowingsList = append(studentFollowingsList,studentFollowings.StudentName)
		}
	}
	followinglist.FollowingList = studentFollowingsList
}
func listfollowing(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/followinglist.html"))
		tmpl.Execute(w, followinglist)
	}
}
var followerlist FollowerList
func listFollowers(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	var followerlist []Followers
	query := fmt.Sprintf("Select * FROM Followers Where studentID = '%s'",
	"00001")
	results, err := db.Query(query)
	if err != nil {
        panic(err.Error())
    }
    for results.Next() {
        var followers Followers
        err = results.Scan(&followers.StudentID, &followers.FollowerID)
        if err != nil {
            panic(err.Error()) 
        }
		followerlist = append(followerlist,followers)
	}
	getFollowerNames(followerlist,db)
}
func getFollowerNames(followers []Followers, db *sql.DB){
	var studentFollowersList []string
	for i:= 0; i < len(followers); i++ {
		query := fmt.Sprintf("Select * FROM Students Where studentID = '%s'",
		followers[i].FollowerID)
		results, err := db.Query(query)
		if err != nil {
        	panic(err.Error())
		}
		for results.Next() {
			var studentFollowers Students
			err = results.Scan(&studentFollowers.StudentID, &studentFollowers.StudentName,
				 &studentFollowers.DOB, &studentFollowers.Address, &studentFollowers.PhoneNumber)
			if err != nil {
				panic(err.Error()) 
			}
			studentFollowersList = append(studentFollowersList,studentFollowers.StudentName)
		}	
   	}
	followerlist.FollowerList = studentFollowersList
}
func listfollowers(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/followerlist.html"))
		tmpl.Execute(w, followerlist)
	}
}
func feedpage(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/feedpage.html"))
		tmpl.Execute(w, followerlist)
	}
}
func postpage(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/postpage.html"))
		tmpl.Execute(w, followerlist)
	}else if r.Method == "POST" {
		var count string 
		db.QueryRow("Select Count(*)+1 FROM Feed").Scan(&count)
		query := fmt.Sprintf("Insert INTO Feed Values('%s','%s','%s')",
		"00001",r.FormValue("feeddata"),count)
		_, err := db.Query(query)  
			if err != nil {
			panic(err.Error())
		}		
	}	
}
var postlist PostList
var feedidlist []Feed
func getposts(){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	var feedlist []string
	
	query := fmt.Sprintf("Select * FROM Feed Where studentID = '%s'",
	"00001")
	results, err := db.Query(query)
	if err != nil {
        panic(err.Error())
    }
    for results.Next() {
        var feeds Feed
        err = results.Scan(&feeds.StudentID, &feeds.Feeddata,&feeds.FeedID)
        if err != nil {
            panic(err.Error()) 
        }
		feedlist = append(feedlist,feeds.Feeddata)
		feedidlist = append(feedidlist,feeds)
	}
	postlist.PostList = feedlist
}
var postselected string
func myposts(w http.ResponseWriter, r *http.Request){
	getposts()
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/myposts.html"))
		tmpl.Execute(w, postlist)
	}else{
		postselected = r.FormValue("Post_Number")
		http.Redirect(w,r,"/viewpost",http.StatusFound)
	}
}
var selectedpostlist SelectedPost
var myvar int
func alterposts(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	var selectedlist []string
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/viewpost.html"))
		intpostselected, err:= strconv.Atoi(postselected)
		if err != nil {
			panic(err.Error())
		}
		myvar = intpostselected
		var selectedpost string
		selectedpost = postlist.PostList[intpostselected-1]
		selectedlist = append(selectedlist,selectedpost)
		selectedpostlist.SelectedPost = selectedlist
		tmpl.Execute(w, selectedpostlist)
	}else{
		feedid := feedidlist[myvar-1].FeedID
		query := fmt.Sprintf("Delete From Feed Where feedID = '%s'",
		feedid)
		_, err := db.Query(query)  
			if err != nil {
			panic(err.Error())
		}
		http.Redirect(w,r,"/feedpage",http.StatusFound)		
	}
}
func updatepost(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err != nil {
		panic(err.Error())
	}
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("./web/updatepostpage.html"))
		tmpl.Execute(w, myvar)
	}else{
		newvar := strconv.Itoa(myvar)
		query := fmt.Sprintf("Update Feed Set feeddata = '%s' Where feedID = '%s'",
		r.FormValue("updatedpost"),newvar)
		_, err := db.Query(query)  
		if err != nil {
			panic(err.Error())
		}
		http.Redirect(w,r,"/feedpage",http.StatusFound)		
	}
}