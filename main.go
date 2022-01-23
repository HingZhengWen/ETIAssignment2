package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"bufio"
	"os"
	"log"
	"database/sql"
	"github.com/gorilla/mux"
	
	_ "github.com/go-sql-driver/mysql"
)
var students map[string]studentInfo
type Students struct{
	StudentID string
	StudentName string 
	DOB string
	Address string
	PhoneNumber string
}
type Followers struct{
	StudentID string
	FollowerID string
}
type studentInfo struct {
	Title string `json:"Student"`
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/tutors", allStudents).Methods( "GET", "POST", "PUT", "DELETE")
	fmt.Println("Listening at port 8063")
	log.Fatal(http.ListenAndServe(":8063",router))
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/ETIStudentSocialDB")
	if err  != nil {
        panic(err.Error())
	} 
	defer db.Close()
	inputs(db)
}
func allStudents(w http.ResponseWriter, r *http.Request) {
	kv := r.URL.Query()
	for k, v := range kv {
		fmt.Println(k, v)
	}
	//returns all the students in JSON
	json.NewEncoder(w).Encode(students)
}
func inputs(db *sql.DB){
	fmt.Println("1. Search students")
	fmt.Println("2. List all followers")
	fmt.Println("3. List all following")
	fmt.Println("")
	var inputChoice string
	fmt.Scanln(&inputChoice)
	if inputChoice == "1"{
		searchStudents(db)
	}else if inputChoice == "2"{
		listFollowers(db)
	}else if inputChoice == "3"{
		listFollowings(db)
	}
}
func searchStudents(db *sql.DB){
	// 00001 for now as i have no way to confirm student's current logon user
	// so for now it is 00001, as the user for this test would be user 00001
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter student name: ")
	scanner.Scan()
	nameinput := scanner.Text()
	query := fmt.Sprintf("Select * FROM Students Where StudentID != '%s' && StudentName = '%s'",
	"00001",nameinput)
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
		studentProfile(db,student)
	}
}
func studentProfile(db *sql.DB,student Students){
	if checkStudentFollowed(db,student.StudentID) == true{
		fmt.Println(student.StudentID)
		fmt.Println(student.StudentName)
		fmt.Println(student.Address)
		fmt.Println("")
		fmt.Print("Unfollow student [Y/N]: ")
		var inputChoice string
		fmt.Scanln(&inputChoice)
		if inputChoice == "Y"{
			fmt.Println("Unfollowed "+student.StudentName)
			query := fmt.Sprintf("Delete FROM Followers Where followerID = '%s' && studentID = '%s'",
			"00001",student.StudentID)
			_, err := db.Query(query)  
			if err != nil {
				panic(err.Error())
			}
		}else{
			inputs(db)	
		}
	}else{
		fmt.Println(student.StudentID)
		fmt.Println(student.StudentName)
		fmt.Println(student.Address)
		fmt.Println("")
		fmt.Print("Follow student [Y/N]: ")
		var inputChoice string
		fmt.Scanln(&inputChoice)
		if inputChoice == "Y"{
			fmt.Println("Followed "+student.StudentName)
			query := fmt.Sprintf("Insert INTO Followers Values('%s','%s')",
			student.StudentID,"00001")
			_, err := db.Query(query)  
			if err != nil {
				panic(err.Error())
		}
		}else{
			inputs(db)	
		}
	}
}


//To check if searched profile is followed by user
//Return true if yes

func checkStudentFollowed(db *sql.DB,studentid string) bool{
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
		if followers.StudentID == studentid{
			studentfollowed = true
		}
	}
	return studentfollowed
}


func listFollowers(db *sql.DB){
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
	fmt.Println("Followers")
	fmt.Println("")
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
	for e:=0; e<len(studentFollowersList);e++{
		fmt.Println(studentFollowersList[e])
	}
}
func listFollowings(db *sql.DB){
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
	fmt.Println("Following")
	fmt.Println("")
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
	for e:=0; e<len(studentFollowingsList);e++{
		fmt.Println(studentFollowingsList[e])
	}
}

