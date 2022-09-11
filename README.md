# sql2code
## Introduce
generate the golang code from table creation statement
* model code
* dao code of basic CRUD
## Download
```
git@github.com:bwz1984/sql2code.git
```
### Usage
```
Usage of this program:
  -dbcon string
        db connect name
  -if string
        File path of the SQL statement that creates the table
  -op int
        1:gen model code 2:gen dao code 3:both (default 1)
  -pp string
        package prefix add for go file
  -sql string
        SQL statement to create table
  -tp string
        table prefix of table name to cut
```
### Use
Use the terminal to enter the ```code sql2code``` directory
```
go run main.go -if=xxx.sql -dbC
```