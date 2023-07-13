// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"AlliancesDocking/model"
)

func newAvtWtuTimetable(db *gorm.DB, opts ...gen.DOOption) avtWtuTimetable {
	_avtWtuTimetable := avtWtuTimetable{}

	_avtWtuTimetable.avtWtuTimetableDo.UseDB(db, opts...)
	_avtWtuTimetable.avtWtuTimetableDo.UseModel(&model.AvtWtuTimetable{})

	tableName := _avtWtuTimetable.avtWtuTimetableDo.TableName()
	_avtWtuTimetable.ALL = field.NewAsterisk(tableName)
	_avtWtuTimetable.ID = field.NewInt64(tableName, "id")
	_avtWtuTimetable.Wid = field.NewString(tableName, "wid")
	_avtWtuTimetable.Kch = field.NewString(tableName, "kch")
	_avtWtuTimetable.Bh = field.NewString(tableName, "bh")
	_avtWtuTimetable.Xm = field.NewString(tableName, "xm")
	_avtWtuTimetable.Xn = field.NewString(tableName, "xn")
	_avtWtuTimetable.Xq = field.NewString(tableName, "xq")
	_avtWtuTimetable.Skzc = field.NewString(tableName, "skzc")
	_avtWtuTimetable.Skjc = field.NewString(tableName, "skjc")
	_avtWtuTimetable.Skxq = field.NewInt64(tableName, "skxq")
	_avtWtuTimetable.Kcm = field.NewString(tableName, "kcm")
	_avtWtuTimetable.Ywkcm = field.NewString(tableName, "ywkcm")
	_avtWtuTimetable.Jasmc = field.NewString(tableName, "jasmc")
	_avtWtuTimetable.Jxlmc = field.NewString(tableName, "jxlmc")

	_avtWtuTimetable.fillFieldMap()

	return _avtWtuTimetable
}

type avtWtuTimetable struct {
	avtWtuTimetableDo

	ALL   field.Asterisk
	ID    field.Int64
	Wid   field.String
	Kch   field.String
	Bh    field.String
	Xm    field.String
	Xn    field.String
	Xq    field.String
	Skzc  field.String
	Skjc  field.String
	Skxq  field.Int64
	Kcm   field.String
	Ywkcm field.String
	Jasmc field.String
	Jxlmc field.String

	fieldMap map[string]field.Expr
}

func (a avtWtuTimetable) Table(newTableName string) *avtWtuTimetable {
	a.avtWtuTimetableDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtWtuTimetable) As(alias string) *avtWtuTimetable {
	a.avtWtuTimetableDo.DO = *(a.avtWtuTimetableDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtWtuTimetable) updateTableName(table string) *avtWtuTimetable {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Wid = field.NewString(table, "wid")
	a.Kch = field.NewString(table, "kch")
	a.Bh = field.NewString(table, "bh")
	a.Xm = field.NewString(table, "xm")
	a.Xn = field.NewString(table, "xn")
	a.Xq = field.NewString(table, "xq")
	a.Skzc = field.NewString(table, "skzc")
	a.Skjc = field.NewString(table, "skjc")
	a.Skxq = field.NewInt64(table, "skxq")
	a.Kcm = field.NewString(table, "kcm")
	a.Ywkcm = field.NewString(table, "ywkcm")
	a.Jasmc = field.NewString(table, "jasmc")
	a.Jxlmc = field.NewString(table, "jxlmc")

	a.fillFieldMap()

	return a
}

func (a *avtWtuTimetable) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtWtuTimetable) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 14)
	a.fieldMap["id"] = a.ID
	a.fieldMap["wid"] = a.Wid
	a.fieldMap["kch"] = a.Kch
	a.fieldMap["bh"] = a.Bh
	a.fieldMap["xm"] = a.Xm
	a.fieldMap["xn"] = a.Xn
	a.fieldMap["xq"] = a.Xq
	a.fieldMap["skzc"] = a.Skzc
	a.fieldMap["skjc"] = a.Skjc
	a.fieldMap["skxq"] = a.Skxq
	a.fieldMap["kcm"] = a.Kcm
	a.fieldMap["ywkcm"] = a.Ywkcm
	a.fieldMap["jasmc"] = a.Jasmc
	a.fieldMap["jxlmc"] = a.Jxlmc
}

func (a avtWtuTimetable) clone(db *gorm.DB) avtWtuTimetable {
	a.avtWtuTimetableDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtWtuTimetable) replaceDB(db *gorm.DB) avtWtuTimetable {
	a.avtWtuTimetableDo.ReplaceDB(db)
	return a
}

type avtWtuTimetableDo struct{ gen.DO }

type IAvtWtuTimetableDo interface {
	gen.SubQuery
	Debug() IAvtWtuTimetableDo
	WithContext(ctx context.Context) IAvtWtuTimetableDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtWtuTimetableDo
	WriteDB() IAvtWtuTimetableDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtWtuTimetableDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtWtuTimetableDo
	Not(conds ...gen.Condition) IAvtWtuTimetableDo
	Or(conds ...gen.Condition) IAvtWtuTimetableDo
	Select(conds ...field.Expr) IAvtWtuTimetableDo
	Where(conds ...gen.Condition) IAvtWtuTimetableDo
	Order(conds ...field.Expr) IAvtWtuTimetableDo
	Distinct(cols ...field.Expr) IAvtWtuTimetableDo
	Omit(cols ...field.Expr) IAvtWtuTimetableDo
	Join(table schema.Tabler, on ...field.Expr) IAvtWtuTimetableDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtWtuTimetableDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtWtuTimetableDo
	Group(cols ...field.Expr) IAvtWtuTimetableDo
	Having(conds ...gen.Condition) IAvtWtuTimetableDo
	Limit(limit int) IAvtWtuTimetableDo
	Offset(offset int) IAvtWtuTimetableDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtWtuTimetableDo
	Unscoped() IAvtWtuTimetableDo
	Create(values ...*model.AvtWtuTimetable) error
	CreateInBatches(values []*model.AvtWtuTimetable, batchSize int) error
	Save(values ...*model.AvtWtuTimetable) error
	First() (*model.AvtWtuTimetable, error)
	Take() (*model.AvtWtuTimetable, error)
	Last() (*model.AvtWtuTimetable, error)
	Find() ([]*model.AvtWtuTimetable, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtWtuTimetable, err error)
	FindInBatches(result *[]*model.AvtWtuTimetable, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtWtuTimetable) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtWtuTimetableDo
	Assign(attrs ...field.AssignExpr) IAvtWtuTimetableDo
	Joins(fields ...field.RelationField) IAvtWtuTimetableDo
	Preload(fields ...field.RelationField) IAvtWtuTimetableDo
	FirstOrInit() (*model.AvtWtuTimetable, error)
	FirstOrCreate() (*model.AvtWtuTimetable, error)
	FindByPage(offset int, limit int) (result []*model.AvtWtuTimetable, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtWtuTimetableDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtWtuTimetableDo) Debug() IAvtWtuTimetableDo {
	return a.withDO(a.DO.Debug())
}

func (a avtWtuTimetableDo) WithContext(ctx context.Context) IAvtWtuTimetableDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtWtuTimetableDo) ReadDB() IAvtWtuTimetableDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtWtuTimetableDo) WriteDB() IAvtWtuTimetableDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtWtuTimetableDo) Session(config *gorm.Session) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtWtuTimetableDo) Clauses(conds ...clause.Expression) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtWtuTimetableDo) Returning(value interface{}, columns ...string) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtWtuTimetableDo) Not(conds ...gen.Condition) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtWtuTimetableDo) Or(conds ...gen.Condition) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtWtuTimetableDo) Select(conds ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtWtuTimetableDo) Where(conds ...gen.Condition) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtWtuTimetableDo) Order(conds ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtWtuTimetableDo) Distinct(cols ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtWtuTimetableDo) Omit(cols ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtWtuTimetableDo) Join(table schema.Tabler, on ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtWtuTimetableDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtWtuTimetableDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtWtuTimetableDo) Group(cols ...field.Expr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtWtuTimetableDo) Having(conds ...gen.Condition) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtWtuTimetableDo) Limit(limit int) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtWtuTimetableDo) Offset(offset int) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtWtuTimetableDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtWtuTimetableDo) Unscoped() IAvtWtuTimetableDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtWtuTimetableDo) Create(values ...*model.AvtWtuTimetable) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtWtuTimetableDo) CreateInBatches(values []*model.AvtWtuTimetable, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtWtuTimetableDo) Save(values ...*model.AvtWtuTimetable) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtWtuTimetableDo) First() (*model.AvtWtuTimetable, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtWtuTimetable), nil
	}
}

func (a avtWtuTimetableDo) Take() (*model.AvtWtuTimetable, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtWtuTimetable), nil
	}
}

func (a avtWtuTimetableDo) Last() (*model.AvtWtuTimetable, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtWtuTimetable), nil
	}
}

func (a avtWtuTimetableDo) Find() ([]*model.AvtWtuTimetable, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtWtuTimetable), err
}

func (a avtWtuTimetableDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtWtuTimetable, err error) {
	buf := make([]*model.AvtWtuTimetable, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtWtuTimetableDo) FindInBatches(result *[]*model.AvtWtuTimetable, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtWtuTimetableDo) Attrs(attrs ...field.AssignExpr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtWtuTimetableDo) Assign(attrs ...field.AssignExpr) IAvtWtuTimetableDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtWtuTimetableDo) Joins(fields ...field.RelationField) IAvtWtuTimetableDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtWtuTimetableDo) Preload(fields ...field.RelationField) IAvtWtuTimetableDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtWtuTimetableDo) FirstOrInit() (*model.AvtWtuTimetable, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtWtuTimetable), nil
	}
}

func (a avtWtuTimetableDo) FirstOrCreate() (*model.AvtWtuTimetable, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtWtuTimetable), nil
	}
}

func (a avtWtuTimetableDo) FindByPage(offset int, limit int) (result []*model.AvtWtuTimetable, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a avtWtuTimetableDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtWtuTimetableDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtWtuTimetableDo) Delete(models ...*model.AvtWtuTimetable) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtWtuTimetableDo) withDO(do gen.Dao) *avtWtuTimetableDo {
	a.DO = *do.(*gen.DO)
	return a
}
