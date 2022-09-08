package sql2code_tpl

import (
	"os"
	"path"
	"runtime"
)

// 模版文件放在同级目录
func Sql2ModelTplPathGet() (string, error) {
	_, file, _, _ := runtime.Caller(0)
	tpl := path.Join(file, "../sql2model_tpl.txt")
	_, err := os.Stat(tpl)
	return tpl, err
}

// 模版文件放在同级目录
func Sql2DaoTplPathGet() (string, error) {
	_, file, _, _ := runtime.Caller(0)
	tpl := path.Join(file, "../sql2dao_tpl.txt")
	_, err := os.Stat(tpl)
	return tpl, err
}
