package sql2model

import (
	"fmt"
	"sql2code/sql2code_tpl"
	"sql2code/util/util_strings"
	"strings"
	"text/template"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/test_driver"
)

var (
	gModelTplPath = ""
)

func init() {
	var err error
	gModelTplPath, err = sql2code_tpl.Sql2ModelTplPathGet()
	if err != nil {
		panic(err)
	}
}

type ModelRow struct {
	Name    string
	Comment string
	GoType  string
	Tags    string
}

type ModelTable struct {
	TblName       string
	OriginTblName string
	Comment       string
	Rows          []ModelRow
}

type ModelFile struct {
	PackageName string
	*ModelTable
}

func parseCreateTableStmt(sql string) (*ast.CreateTableStmt, error) {
	p := parser.New()

	astNode, err := p.ParseOneStmt(sql, "", "")
	if err != nil {
		return nil, err
	}
	cts, ok := astNode.(*ast.CreateTableStmt)
	if !ok {
		return nil, fmt.Errorf("turn to ast.CreateTableStmt fail")
	}
	return cts, nil
}

func sqlType2GoType(tp byte, flag uint) string {
	switch tp {
	case TypeTiny, TypeShort, TypeInt24, TypeLong,
		TypeBit, TypeYear:
		if HasUnsignedFlag(flag) {
			return "uint32"
		} else {
			return "int32"
		}
	case TypeLonglong:
		if HasUnsignedFlag(flag) {
			return "uint64"
		} else {
			return "int64"
		}
	case TypeFloat, TypeDouble:
		return "float64"
	case TypeTimestamp, TypeDate, TypeDatetime:
		return "time.Time"
	case TypeEnum, TypeSet:
		if flag&EnumSetAsIntFlag > 0 {
			return "int32"
		}
	}
	return "string"
}

func SQLParse(sql string) (*ModelTable, error) {
	cts, err := parseCreateTableStmt(sql)
	if err != nil {
		return nil, err
	}
	mt := &ModelTable{}
	primaryKey := ""

	// table name
	mt.TblName = util_strings.ToCamel(cts.Table.Name.L)
	mt.OriginTblName = cts.Table.Name.L

	// primary
	for _, ctt := range cts.Constraints {
		// only contain one primary key
		if ctt.Tp == ast.ConstraintPrimaryKey {
			if len(ctt.Keys) >= 0 {
				primaryKey = ctt.Keys[0].Column.Name.L
			}
			break
		}
	}

	// comment
	for _, op := range cts.Options {
		if op.Tp == ast.TableOptionComment {
			mt.Comment = op.StrValue
		}
	}

	modelRows := make([]ModelRow, 0, len(cts.Cols))
	for _, col := range cts.Cols {

		nameLow := col.Name.Name.L

		modelRow := ModelRow{
			Name: util_strings.ToCamel(col.Name.Name.L), // 需要去除下划线转驼峰
		}
		//fmt.Printf("col: %+v %+v %v %v\n", col.Name, col.Tp, HasUnsignedFlag(col.Tp.GetFlag()), col.Tp.GetType())
		modelRow.GoType = sqlType2GoType(col.Tp.GetType(), col.Tp.GetFlag())

		for _, op := range col.Options {
			if op.Tp == ast.ColumnOptionComment {
				exprVal, ok := op.Expr.(*test_driver.ValueExpr)
				if !ok {
					fmt.Println("op.Expr.(*test_driver.ValueExpr) fail.")
					continue
				}
				modelRow.Comment = exprVal.Datum.GetString()
				break
			}
		}
		if primaryKey == col.Name.Name.L {
			modelRow.Tags = fmt.Sprintf("`gorm:\"column:%v; primary_key\" json:\"%v\"`", nameLow, nameLow)
		} else {
			modelRow.Tags = fmt.Sprintf("`gorm:\"column:%v;\" json:\"%v\"`", nameLow, nameLow)
		}
		//fmt.Println(modelRow.Tags)
		modelRows = append(modelRows, modelRow)
	}
	mt.Rows = modelRows

	return mt, nil
}

func modelCodeGen(modelFile *ModelFile) (string, error) {
	tpl, err := template.ParseFiles(gModelTplPath)
	if err != nil {
		return "", err
	}
	builder := strings.Builder{}
	err = tpl.Execute(&builder, modelFile)
	return builder.String(), err
}

func ModelOriNameGet(sql string) (string, error) {
	cts, err := parseCreateTableStmt(sql)
	if err != nil {
		return "", err
	}
	return cts.Table.Name.L, err
}

func SQL2Model(sql string) (string, error) {
	modelTable, err := SQLParse(sql)
	if err != nil {
		return "", err
	}
	mf := &ModelFile{
		PackageName: modelTable.OriginTblName + "_model",
		ModelTable:  modelTable,
	}
	//fmt.Printf("%+v\n", modelTable)
	return modelCodeGen(mf)
}
