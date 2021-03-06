# DevOps SQL test

## Introduction

An SQL database takes time to resolve queries and at some point the app brakes

## Requisites

Linux SO
SQL database
Go script

## Installation of environment

For the proposal of this test I use an Ec2  Ubuntu 20.04 server running on AWS

Database Base Installation :

[Follow this tutorial](https://www.digitalocean.com/community/tutorials/how-to-install-mariadb-on-ubuntu-20-04)

## Instructions:

 **Create a database called "devops"**

```
CREATE DATABASE devops;
```

 
**Create a new table called "employees"**

 

    CREATE TABLE employees (
    emp_no INT NOT NULL,
    birth_date DATE NOT NULL,
    first_name VARCHAR(14) NOT NULL,
    last_name VARCHAR(16) NOT NULL,
    gender ENUM(‘M’,’F’) NOT NULL,
    hire_date DATE NOT NULL,
    PRIMARY KEY (emp_no)
    );

**Create a new table called "salaries"**

    CREATE TABLE salaries (
    emp_no INT NOT NULL,
    salary INT NOT NULL,
    from_date DATE NOT NULL,
    to_date DATE NOT NULL
    );

**Download the sample data and import it into devops database** 

    curl -LJO https://raw.githubusercontent.com/datacharmer/test_db/master/load_salaries1.dump
    
    curl -LJO https://raw.githubusercontent.com/datacharmer/test_db/master/load_employees.dump

**Import the Dada in the devops database**

    mysql -u admin -p devops < /home/ubuntu/load_salaries1.dump

**Create a new database user **senior** and password: **erplysdv** with read only access**

    CREATE USER 'user_readonly'@'localhost' IDENTIFIED BY 'secret_password';
    GRANT SELECT ON my_database.* TO 'user_readonly'@'localhost';
    FLUSH PRIVILEGES;


**Download the following go code to test the database**

    **curl -LJO https://raw.githubusercontent.com/diiegg/devop-test/main/test.go**

**Download the Golang [MySQL driver](https://github.com/go-sql-driver/mysql) and install it** 



**Run the script**


    go run test.go

**Open a tab in your browser and navigate to http://localhost:9099**

You should get something like this:

![enter image description here](https://user-images.githubusercontent.com/12648295/104135352-b2c6de80-5387-11eb-8d5a-712eedaf844f.png)



Open developer tools on your web browser and go into the network tab, refresh the website

Every time you refresh the website the request time is taking long and long. What is going on there ?

![enter image description here](https://user-images.githubusercontent.com/12648295/104135120-d852e880-5385-11eb-9fbb-eeb4fa13692a.jpg)


## Solution

After inspecting the petitions  I came across with a RAND on the petition 15 that was sending a lot of rows and the database was not able to order quickly that petitions

![enter image description here](https://user-images.githubusercontent.com/12648295/104135099-b5c0cf80-5385-11eb-8ac3-9c01de1d3c4e.jpg)


My solution was to create an index in the tables "salaries" and "employees" on the row "emp_no" getting a 32 ms response time

![enter image description here](https://user-images.githubusercontent.com/12648295/104135056-827e4080-5385-11eb-9827-d5a2ab6ab56d.jpg)

###  Create a index

**Long in into SQL and select DevOps database**

    CREATE INDEX index_emp_no ON salaries(emp_no);
    CREATE INDEX index_emp_no ON employees(emp_no);


Run the go script again and test the database

## Recommendations

**[MySQLTuner](https://github.com/major/MySQLTuner-perl)** is a script written in Perl that allows you to review a MySQL installation quickly and make adjustments to increase performance and stability.

Monitor your database with [Prometheus and Grafana](https://medium.com/schkn/complete-mysql-dashboard-with-grafana-prometheus-36e98cba1390) 

[Golang MySQL Tutorial](https://tutorialedge.net/golang/golang-mysql-tutorial/)

