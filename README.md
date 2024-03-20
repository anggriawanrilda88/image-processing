# Dating App Documentation
Welcome to the Dating App! This application is testing app that designed to connect individuals in a fun and interactive way, allowing users to discover potential matches based on shared interests, preferences, and geographical proximity. With a user-friendly interface and a range of features, the Dating App provides a platform for users to engage, build connections, and find meaningful relationships.

# Domain-Driven Design (DDD) Structure
This application is developed following the principles of Domain-Driven Design (DDD). DDD promotes a modular and structured approach to building software by organizing code around the business domain.

# Show Of Structure

dating-app <br />
----postman_test <br />
----migrations <br />
----src <br />
--------app <br />
------------dto <br />
------------usecase <br />
--------domain <br />
------------entities <br />
------------repositories <br />
------------value_object <br />
----------------user <br />
--------infra <br />
------------auth <br />
----------------jwt <br />
------------constants <br />
------------helpers <br />
------------models <br />
------------persistence <br />
----------------postgresql <br />
--------interface <br />
------------rest <br />
----------------middleware <br />
----------------response <br />
----------------v1 <br />
--------------------mobile_app <br />
------------------------handlers <br />
------------------------requests <br />
------------------------routes <br />
------------------------transformers <br />

# Makefile For Development and Deployment
Here include Makefile to run some command for development and deployment you can see the command bellow:<br /><br />
`make bin` ==> this command to create folder bin <br />
`make setup-tools` ==> this command to install hot reload bin and migrate bin <br />
`make migration-create` ==> this command to create folder migrations and create file migration with name <br />
`make migration-up` ==> this command to execute all sql query from migration file <br />
`make migration-down` ==> this command to remove all sql query from migration file <br />
`make run-dev` ==> this command to run hot reload for development <br />
`make run` ==>  this command to run go file<br />
`make build` ==>  this command to build golang become binary file or executable file<br />


# Quick Setup
Clone the repository: `git clone git@github.com:anggriawanrilda88/ubersnap-test.git` <br /><br />
Setup Tools use Makefile: `make setup-tools` this will help to install hot reload, and migrate tools for golang<br /><br />
Create file .env: `all variable needed can see on file .env.example`<br /><br />
Migrate database: `make migration-up` this command will migrate all table used, ensure the connection database on .env file correct<br /><br />
Run development with hot reload: `make run-dev` this command will run hot reload golang api<br /><br />


# API Example
There are three api example in this app for testing reasong:<br />
`POST   /ubersnap-test/api/v1/users`          ==> api for registration new user dating app.<br />
`POST   /ubersnap-test/api/v1/users/login`    ==> api for login user and get token access.<br />
`GET    /ubersnap-test/api/v1/users/:id`      ==> api to get detail user by id, validate with middleware authentication with jwt token access needed.<br />

# POSTMAN TEST
Folder postman_test contains test case files that can run on postman test.