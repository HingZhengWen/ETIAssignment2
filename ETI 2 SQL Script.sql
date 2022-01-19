CREATE database ETIStudentSocialDB;
USE ETIStudentSocialDB;
CREATE TABLE Students (StudentID varchar(5) NOT NULL PRIMARY KEY, StudentName VARCHAR(30), DOB VARCHAR(30), Address VARCHAR(50), PhoneNumber VARCHAR(8)); 
Insert INTO Students Values ("00001","Jayman Hing",08/08/2002,"Apt blk 121 lengkong tiga","87776080");
Insert INTO Students Values ("00002","Julies Liew",03/01/2002,"Apt blk 111 Punggol drive","83233431");
Insert INTO Students Values ("00003","James Lee",06/01/2002,"Paya lebar road", "97770000");
Insert INTO Students Values ("00004","Jeffrey Dahmer",03/02/2002, "Springleaf rd", "93213214");
Insert INTO Students Values ("00005","Jeremy Jones",02/09/2002,"Sengkang ave", "96661234");
CREATE TABLE Followers (studentID varchar(5), followerID varchar(5)); 