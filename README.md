Hing Zheng Wen


Introduction 
 
This project was made for ETI Assignment 2, and the project that i made was for the social feed for the Edufi App.
This platform was made with the combined efforts of many different students from my class, each of us having different
requirements and goals to implement
for my part, my requirements were as so

Social feed (3.20)
    Follow, unfollow students
    List all followers
    List all followings
    Post, update, delete feeds
    List feeds, activities of followings
    
This microservice that i was to implement was a part of the Messaging sector of the platform,
along with other possible features like messaging, special interest groups or activities within groups.

The purpose of my microservice was for students to interact with other students, by being able
to view the activites and follow other students as they choose.


Design Considerations

In order to meet the requirements given, i knew i had to have multiple tables in order to accomodate 
followers,students as well as feeds
To do this, i had 3 tables in my database, 
1 for all the student's details
1 for students and their followers, and finally
1 for the students and their feeds.
This used normalization in order to make it more neat.


Set-Up

To run it on the putty server,
first pull the three docker repositories

 - hingzhengwen/socialfeeddb 
 - hingzhengwen/socialfeed
 - hingzhengwen/socialfeedapi

After which, 
run the containers on the server using these commands in this order
 - docker run -d -p 8062:3306 -e MYSQL_ROOT_PASSWORD=test hingzhengwen/socialfeeddb
 - docker run -d -p 8060:8060 hingzhengwen/socialfeed

After this is done, the link to the website can be navigated through the student page or if ran locally, can be accessed through the link
 - http://localhost:8060/
