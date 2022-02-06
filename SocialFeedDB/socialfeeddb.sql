CREATE database ETIStudentSocialDB;
USE ETIStudentSocialDB;

CREATE USER 'user'@'%' IDENTIFIED BY 'password';

GRANT ALL PRIVILEGES
ON ETIStudentSocialDB.*
TO 'user'@'%';

CREATE TABLE Students (StudentID varchar(5) NOT NULL PRIMARY KEY, StudentName VARCHAR(30), DOB VARCHAR(10), Address VARCHAR(50), PhoneNumber VARCHAR(8)); 
CREATE TABLE Followers (studentID varchar(5), followerID varchar(5)); 
CREATE TABLE Feed(studentID varchar(5), feeddata varchar(100), feedID varchar(5) NOT NULL PRIMARY KEY);

INSERT INTO Students (StudentID, StudentName, DOB, Address, PhoneNumber) VALUES ("1", "Jake", "20-10-1985", "1 Temasek Avenue", "99998888");
INSERT INTO Students (StudentID, StudentName, DOB, Address, PhoneNumber) VALUES ("2", "Amy", "21-09-1997", "1057 Eunos Avenue", "88889999");
INSERT INTO Students (StudentID, StudentName, DOB, Address, PhoneNumber) VALUES ("3", "Raymond", "15-12-1995", "1 Bilal Lane", "77776666");
INSERT INTO Students (StudentID, StudentName, DOB, Address, PhoneNumber) VALUES ("4", "Terry", "14-02-1976", "Kallang Puddin Rd", "66665555");
INSERT INTO Students (StudentID, StudentName, DOB, Address, PhoneNumber) VALUES ("5", "Charles", "09-08-1965", "163 Tanglin Rd", "55554444");