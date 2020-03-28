package crud

func CrudTemplate() string {

	return `
// Code generated by 'freedom new-crud'
package object
{{.Import}}
{{.Content}}

// TakeChanges .
func (obj *{{.Name}})TakeChanges() map[string]interface{} {
	if obj.changes == nil {
		return nil
	}
	result := make(map[string]interface{})
	for k, v := range obj.changes {
		result[k] = v
	}
	obj.changes = nil
	return result
}

{{range .Fields}}
// Set{{.Value}} .
func (obj *{{.StructName}}) Set{{.Value}} ({{.Arg}} {{.Type}}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.{{.Value}} = {{.Arg}} 
	obj.changes["{{.Column}}"] = {{.Arg}}
}
{{ end }}

{{range .NumberFields}}
// Add{{.Value}} .
func (obj *{{.StructName}}) Add{{.Value}} ({{.Arg}} {{.Type}}) {
	if obj.changes == nil {
		obj.changes = make(map[string]interface{})
	}
	obj.{{.Value}} += {{.Arg}} 
	obj.changes["{{.Column}}"] = gorm.Expr("{{.Column}} + ?", {{.Arg}})
}
{{ end }}
`
}

func FunTemplatePackage() string {
	return `
	// Code generated by 'freedom new-crud'
	package repository
	import (
		"github.com/8treenet/freedom"
		"github.com/jinzhu/gorm"
		"time"
		"{{.PackagePath}}"
	)

	func errorLog(repo freedom.GORMRepository, model, method string, e error, expression ...interface{}) {
		if e == nil || e == gorm.ErrRecordNotFound {
			return
		}
		repo.GetRuntime().Logger().Errorf("Orm error, model: %s, method: %s, expression :%v, reason for error:%v", model, method, expression, e)
	}
`
}
func FunTemplate() string {
	return `
	// find{{.Name}}ByPrimary .
	func find{{.Name}}ByPrimary(repo freedom.GORMRepository, result interface{}, primary interface{}) (e error) {
		now := time.Now()
		e = repo.DB().Find(result, primary).Error
		freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ByPrimary", e, now)
		errorLog(repo, "{{.Name}}", "find{{.Name}}ByPrimary", e, primary)
		return
	}
	
	// find{{.Name}}sByPrimarys .
	func find{{.Name}}sByPrimarys(repo freedom.GORMRepository, results interface{}, primarys ...interface{}) (e error) {
		now := time.Now()
		e = repo.DB().Find(results, primarys).Error
		freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}sByPrimarys", e, now)
		errorLog(repo, "{{.Name}}", "find{{.Name}}sByPrimarys", e, primarys)
		return
	}
	
	// find{{.Name}} .
	func find{{.Name}}(repo freedom.GORMRepository, query object.{{.Name}}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}", e, now)
			errorLog(repo, "{{.Name}}", "find{{.Name}}", e, query)
		}()
		db := repo.DB().Where(query)
		if len(builders) == 0 {
			e = db.Last(result).Error
			return
		}

		e = builders[0].Execute(db.Limit(1), result)
		return
	}
	
	// find{{.Name}}ByWhere .
	func find{{.Name}}ByWhere(repo freedom.GORMRepository, query string, args []interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ByWhere", e, now)
			errorLog(repo, "{{.Name}}", "find{{.Name}}ByWhere", e, query, args)
		}()
		db := repo.DB()
		if query != "" {
			db = db.Where(query, args...)
		}
		if len(builders) == 0 {
			e = db.Last(result).Error
			return
		}
	
		e = builders[0].Execute(db.Limit(1), result)
		return
	}
	
	// find{{.Name}}ByMap .
	func find{{.Name}}ByMap(repo freedom.GORMRepository, query map[string]interface{}, result interface{}, builders ...freedom.QueryBuilder) (e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}ByMap", e, now)
			errorLog(repo, "{{.Name}}", "find{{.Name}}ByMap", e, query)
		}()

		db := repo.DB().Where(query)
		if len(builders) == 0 {
			e = db.Last(result).Error
			return
		}
	
		e = builders[0].Execute(db.Limit(1), result)
		return
	}
	
	// find{{.Name}}s .
	func find{{.Name}}s(repo freedom.GORMRepository, query object.{{.Name}}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}s", e, now)
			errorLog(repo, "{{.Name}}", "find{{.Name}}s", e, query)
		}()
		db := repo.DB().Where(query)
	
		if len(builders) == 0 {
			e = db.Find(results).Error
			return
		}
		e = builders[0].Execute(db, results)
		return
	}
	
	// find{{.Name}}sByWhere .
	func find{{.Name}}sByWhere(repo freedom.GORMRepository, query string, args []interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}sByWhere", e, now)
			errorLog(repo, "{{.Name}}", "find{{.Name}}sByWhere", e, query, args)
		}()
		db := repo.DB()
		if query != "" {
			db = db.Where(query, args...)
		}
	
		if len(builders) == 0 {
			e = db.Find(results).Error
			return
		}
		e = builders[0].Execute(db, results)
		return
	}
	
	// find{{.Name}}sByMap .
	func find{{.Name}}sByMap(repo freedom.GORMRepository, query map[string]interface{}, results interface{}, builders ...freedom.QueryBuilder) (e error) {
		now := time.Now()
		defer func() {
			freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "find{{.Name}}sByMap", e, now)
			errorLog(repo, "{{.Name}}", "find{{.Name}}sByMap", e, query)
		}()

		db := repo.DB().Where(query)
	
		if len(builders) == 0 {
			e = db.Find(results).Error
			return
		}
		e = builders[0].Execute(db, results)
		return
	}
	
	// create{{.Name}} .
	func create{{.Name}}(repo freedom.GORMRepository, object *object.{{.Name}}) (rowsAffected int64, e error) {
		now := time.Now()
		db := repo.DB().Create(object)
		rowsAffected = db.RowsAffected
		e = db.Error
		freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "create{{.Name}}", e, now)
		errorLog(repo, "{{.Name}}", "create{{.Name}}", e, *object)
		return
	}

	// save{{.Name}} .
	func save{{.Name}}(repo freedom.GORMRepository, object *object.{{.Name}}) (affected int64, e error) {
		now := time.Now()
		db := repo.DB().Model(object).Updates(object.TakeChanges())
		e = db.Error
		affected = db.RowsAffected
		freedom.Prometheus().OrmWithLabelValues("{{.Name}}", "save{{.Name}}", e, now)
		errorLog(repo, "{{.Name}}", "save{{.Name}}", e, *object)
		return
	}
`
}
