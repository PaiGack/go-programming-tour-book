package sql2struct

import (
	"fmt"
	"go-programming-tour-book/ch1/part2/a/internal/word"
	"os"
	"text/template"
)

const structTpl = `
package model

type {{.TableName | ToCamelCase}}Model struct {
{{range .Columns}}
  {{ $typeLen := len .Type }} {{ if gt $typeLen 0 }}{{.Name | ToCamelCase}} {{.Type}} {{.Tag}}{{ else }} {{.Name}}{{ end }}{{ $length := len .Comment}} {{ if gt $length 0 }}// {{.Comment}} {{else}}// {{.Name}} {{ end }}
{{end}}
}
func (model *{{.TableName | ToCamelCase}}Model) TableName() string {
  return "{{.TableName}}"
}`

type StructTemplate struct {
	structTpl string
}

type StructColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName string
	Columns   []*StructColumn
}

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{structTpl: structTpl}
}

func (t *StructTemplate) AssemblyColumns(tbColumns []*TableColumn) []*StructColumn {
	tplColumns := make([]*StructColumn, 0, len(tbColumns))
	for _, column := range tbColumns {
		tag := fmt.Sprintf("`json:\"%s\"`", column.ColumnName)
		tplColumns = append(tplColumns, &StructColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}
	return tplColumns
}

func (t *StructTemplate) Generate(tableName string, tplColumns []*StructColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(template.FuncMap{
		"ToCamelCase": word.UnderscoreToUpperCamelCase,
	}).Parse(t.structTpl))
	tplDB := StructTemplateDB{
		TableName: tableName,
		Columns:   tplColumns,
	}
	f, err := os.Create(fmt.Sprintf("./%s.go", tableName))
	if err != nil {
		return err
	}
	defer f.Close()
	//err = tpl.Execute(os.Stdout, tplDB)
	err = tpl.Execute(f, tplDB)
	if err != nil {
		return err
	}
	return nil
}
