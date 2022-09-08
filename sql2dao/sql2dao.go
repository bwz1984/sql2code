package sql2dao

import (
	"log"
	"sql2code/sql2code_tpl"
	"sql2code/sql2model"
	"sql2code/util/util_strings"
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

func SQL2Dao(sql string, dbCon string) (string, error) {
	tblOriName, err := sql2model.ModelOriNameGet(sql)
	if err != nil {
		return "", err
	}
	modelPackage := tblOriName + "_model"
	modelName := util_strings.ToCamel(tblOriName)
	packageName := tblOriName + "_dao"
	df := DaoFile{
		PackageName:  packageName,
		ModelPackage: modelPackage,
		ModelName:    modelName,
		DBConect:     dbCon,
	}
	return daoFileGen(df)
}
