package sql2model

import (
	"fmt"
	"html/template"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSQL2Model(t *testing.T) {
	sql := "CREATE TABLE `t_student` (" +
		"`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID'," +
		"`age` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '年龄'," +
		"`height` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '身高'," +
		"PRIMARY KEY (`id`)," +
		"KEY `age` (`age`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学生表';"

	str, err := SQL2Model(sql, "t_", "college")
	assert.Nil(t, err)
	fmt.Println(str)
}

func TestTmp(t *testing.T) {
	tpl, err := template.ParseFiles("../template/sql2model.txt")
	assert.Nil(t, err)
	builder := strings.Builder{}
	mf := ModelFile{
		PackageName: "aaa",
		ModelTable: &ModelTable{
			ModelName:     "bb",
			OriginTblName: "t_bb",
			Comment:       "tmp test",
			Rows: []ModelRow{
				{Name: "col", Comment: "col comment", GoType: "uint64", Tags: "gorm:column:col json:col"},
			},
		},
	}
	err = tpl.Execute(&builder, mf)
	assert.Nil(t, err)
	fmt.Println(builder.String())
}
