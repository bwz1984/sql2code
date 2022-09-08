package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sql2code/sql2dao"
	"sql2code/sql2model"

	_ "github.com/pingcap/tidb/parser/test_driver"
)

const (
	OperationGenModel = int(1 << 0)
	OperationGenDao   = int(1 << 1)
)

type sql2CodeParams struct {
	sqlStr   string
	filePath string
	dbConStr string
	op       int
}

var inputParams sql2CodeParams = sql2CodeParams{}

func init() {
	flag.StringVar(&inputParams.sqlStr, "sql", "", "SQL statement to create table")
	flag.StringVar(&inputParams.filePath, "file", "", "File path of the SQL statement that creates the table")
	flag.StringVar(&inputParams.dbConStr, "dbcon", "", "db connect name")
	flag.IntVar(&inputParams.op, "op", int(1), "1:gen model code 2:gen dao code 3:both")
}

func main() {
	flag.Parse()
	sql := inputParams.sqlStr
	if sql == "" && inputParams.filePath != "" {
		b, err := ioutil.ReadFile(inputParams.filePath)
		if err != nil {
			fmt.Println("aaaaas")
			log.Printf("ioutil.ReadFile fail, err:%v path:%v\n", err, inputParams.filePath)
			return
		}
		sql = string(b)
	}
	op := inputParams.op
	if op&OperationGenModel == OperationGenModel {
		modelCode, err := sql2model.SQL2Model(sql)
		if err != nil {
			log.Printf("sql2model.SQL2Model fail, err:%v sql:%v\n", err, sql)
			return
		}
		fmt.Printf("%+v\n", modelCode)
	}
	if op&OperationGenDao == OperationGenDao {
		dbConStr := inputParams.dbConStr
		if len(dbConStr) <= 0 {
			log.Printf("should set db connect when gen dao code")
		}
		daoCode, err := sql2dao.SQL2Dao(sql, dbConStr)
		if err != nil {
			log.Printf("sql2dao.SQL2Dao fail, err:%v sql:%v\n", err, sql)
			return
		}
		fmt.Printf("%+v\n", daoCode)
	}
}
