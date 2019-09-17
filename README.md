# Go MySQL Task

A small Golang program that reads data from a csv file and saves it to MySql database.

Instructions:
1- Install Docker
2- Install MySQL
3- git pull
4- docker-compose up in project directory
5- Start MySQL Workbench. Start new connection named 'mysql-dev' on port 'localhost:3308' with    default schema 'testdb'. Enter password 'password'.
6- Start CMD or PowerShell and move to project directory. run command 'docker build -t my-go-app .'
7- Run application on host machine i.e 'go run main.go'