package sql2dao

import (
	"log"
	"sql2code/sql2code_tpl"
	"sql2code/sql2model"
	"strings"
	"text/template"
)

var (
	gDaoTplPath = ""
)

func init() {
	var err error
	gDaoTplPath, err = sql2code_tpl.Sql2DaoTplPathGet()
	if err != nil {
		panic(err)
	}
}

type DaoFile struct {
	PackageName  string
	ModelPackage string
	ModelName    string
	DBConect     string
}

func daoFileGen(df DaoFile) (string, error) {
	tpl, err := template.ParseFiles(gDaoTplPath)
	if err != nil {
		log.Printf("template.ParseFiles fail,err:%v path:%v", err, gDaoTplPath)
		return "", err
	}
	builder := strings.Builder{}
	err = tpl.Execute(&builder, df)
	return builder.String(), err
}

func DaoPackageNameGet(tableName, tablePrefix, packagePrefix string) string {
	var packageName string
	if len(packagePrefix) > 0 {
		packageName = packagePrefix + "_" + tableName + "_dao"
	} else {
		packageName = tableName + "_dao"
	}
	return packageName
}

func DaoFileNameGetFromSQL(sql, tablePrefix, packagePrefix string) (string, error) {
	tblName, err := sql2model.TableNameGetFromSQL(sql, tablePrefix)
	if err != nil {
		return "", err
	}
	var fileName string
	if len(packagePrefix) > 0 {
		fileName = packagePrefix + "_" + tblName + "_service.go"
	} else {
		fileName = tblName + ".go"
	}
	return fileName, nil
}

func SQL2Dao(sql string, tablePrefix, packagePrefix, dbCon string) (string, error) {
	tblName, err := sql2model.TableNameGetFromSQL(sql, tablePrefix)
	if err != nil {
		return "", err
	}

	modelPackage := sql2model.ModelPackageGet(tblName, tablePrefix, packagePrefix)
	packageName := DaoPackageNameGet(tblName, tablePrefix, packagePrefix)

	df := DaoFile{
		PackageName:  packageName,
		ModelPackage: modelPackage,
		ModelName:    sql2model.TableName2ModelName(tblName),
		DBConect:     dbCon,
	}
	return daoFileGen(df)
}
