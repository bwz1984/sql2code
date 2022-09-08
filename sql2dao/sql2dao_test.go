package sql2dao

import (
	"fmt"
	"strings"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func TestCrudTmpFile(t *testing.T) {
	tpl, err := template.ParseFiles(gDaoTplPath)
	assert.Nil(t, err)

	builder := strings.Builder{}
	crudf := DaoFile{
		PackageName:  "student_service",
		ModelPackage: "student_name",
		ModelName:    "Student",
		DBConect:     "db",
	}
	err = tpl.Execute(&builder, crudf)
	assert.Nil(t, err)
	fmt.Println(builder.String())
}

func TestSQL2Dao(t *testing.T) {
	sql := "CREATE TABLE `student` (" +
		"`id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID'," +
		"`age` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '年龄'," +
		"`height` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '身高'," +
		"PRIMARY KEY (`id`)," +
		"KEY `age` (`age`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学生表';"

	str, err := SQL2Dao(sql, "stu_db")
	assert.Nil(t, err)
	fmt.Println(str)
}
