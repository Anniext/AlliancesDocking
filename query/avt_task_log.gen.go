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

func newAvtTaskLog(db *gorm.DB, opts ...gen.DOOption) avtTaskLog {
	_avtTaskLog := avtTaskLog{}

	_avtTaskLog.avtTaskLogDo.UseDB(db, opts...)
	_avtTaskLog.avtTaskLogDo.UseModel(&model.AvtTaskLog{})

	tableName := _avtTaskLog.avtTaskLogDo.TableName()
	_avtTaskLog.ALL = field.NewAsterisk(tableName)
	_avtTaskLog.ID = field.NewInt64(tableName, "id")
	_avtTaskLog.TaskID = field.NewInt64(tableName, "task_id")
	_avtTaskLog.JoinNum = field.NewInt64(tableName, "join_num")
	_avtTaskLog.Value = field.NewString(tableName, "value")
	_avtTaskLog.Status = field.NewInt64(tableName, "status")
	_avtTaskLog.CreatedTime = field.NewTime(tableName, "created_time")

	_avtTaskLog.fillFieldMap()

	return _avtTaskLog
}

type avtTaskLog struct {
	avtTaskLogDo

	ALL    field.Asterisk
	ID     field.Int64
	TaskID field.Int64 // 任务ID
	/*
		joinNum

	*/
	JoinNum field.Int64
	/*
		value

	*/
	Value       field.String
	Status      field.Int64 // 状态
	CreatedTime field.Time  // 执行时间

	fieldMap map[string]field.Expr
}

func (a avtTaskLog) Table(newTableName string) *avtTaskLog {
	a.avtTaskLogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtTaskLog) As(alias string) *avtTaskLog {
	a.avtTaskLogDo.DO = *(a.avtTaskLogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtTaskLog) updateTableName(table string) *avtTaskLog {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.TaskID = field.NewInt64(table, "task_id")
	a.JoinNum = field.NewInt64(table, "join_num")
	a.Value = field.NewString(table, "value")
	a.Status = field.NewInt64(table, "status")
	a.CreatedTime = field.NewTime(table, "created_time")

	a.fillFieldMap()

	return a
}

func (a *avtTaskLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtTaskLog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 6)
	a.fieldMap["id"] = a.ID
	a.fieldMap["task_id"] = a.TaskID
	a.fieldMap["join_num"] = a.JoinNum
	a.fieldMap["value"] = a.Value
	a.fieldMap["status"] = a.Status
	a.fieldMap["created_time"] = a.CreatedTime
}

func (a avtTaskLog) clone(db *gorm.DB) avtTaskLog {
	a.avtTaskLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtTaskLog) replaceDB(db *gorm.DB) avtTaskLog {
	a.avtTaskLogDo.ReplaceDB(db)
	return a
}

type avtTaskLogDo struct{ gen.DO }

type IAvtTaskLogDo interface {
	gen.SubQuery
	Debug() IAvtTaskLogDo
	WithContext(ctx context.Context) IAvtTaskLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtTaskLogDo
	WriteDB() IAvtTaskLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtTaskLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtTaskLogDo
	Not(conds ...gen.Condition) IAvtTaskLogDo
	Or(conds ...gen.Condition) IAvtTaskLogDo
	Select(conds ...field.Expr) IAvtTaskLogDo
	Where(conds ...gen.Condition) IAvtTaskLogDo
	Order(conds ...field.Expr) IAvtTaskLogDo
	Distinct(cols ...field.Expr) IAvtTaskLogDo
	Omit(cols ...field.Expr) IAvtTaskLogDo
	Join(table schema.Tabler, on ...field.Expr) IAvtTaskLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtTaskLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtTaskLogDo
	Group(cols ...field.Expr) IAvtTaskLogDo
	Having(conds ...gen.Condition) IAvtTaskLogDo
	Limit(limit int) IAvtTaskLogDo
	Offset(offset int) IAvtTaskLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtTaskLogDo
	Unscoped() IAvtTaskLogDo
	Create(values ...*model.AvtTaskLog) error
	CreateInBatches(values []*model.AvtTaskLog, batchSize int) error
	Save(values ...*model.AvtTaskLog) error
	First() (*model.AvtTaskLog, error)
	Take() (*model.AvtTaskLog, error)
	Last() (*model.AvtTaskLog, error)
	Find() ([]*model.AvtTaskLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtTaskLog, err error)
	FindInBatches(result *[]*model.AvtTaskLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtTaskLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtTaskLogDo
	Assign(attrs ...field.AssignExpr) IAvtTaskLogDo
	Joins(fields ...field.RelationField) IAvtTaskLogDo
	Preload(fields ...field.RelationField) IAvtTaskLogDo
	FirstOrInit() (*model.AvtTaskLog, error)
	FirstOrCreate() (*model.AvtTaskLog, error)
	FindByPage(offset int, limit int) (result []*model.AvtTaskLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtTaskLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtTaskLogDo) Debug() IAvtTaskLogDo {
	return a.withDO(a.DO.Debug())
}

func (a avtTaskLogDo) WithContext(ctx context.Context) IAvtTaskLogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtTaskLogDo) ReadDB() IAvtTaskLogDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtTaskLogDo) WriteDB() IAvtTaskLogDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtTaskLogDo) Session(config *gorm.Session) IAvtTaskLogDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtTaskLogDo) Clauses(conds ...clause.Expression) IAvtTaskLogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtTaskLogDo) Returning(value interface{}, columns ...string) IAvtTaskLogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtTaskLogDo) Not(conds ...gen.Condition) IAvtTaskLogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtTaskLogDo) Or(conds ...gen.Condition) IAvtTaskLogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtTaskLogDo) Select(conds ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtTaskLogDo) Where(conds ...gen.Condition) IAvtTaskLogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtTaskLogDo) Order(conds ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtTaskLogDo) Distinct(cols ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtTaskLogDo) Omit(cols ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtTaskLogDo) Join(table schema.Tabler, on ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtTaskLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtTaskLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtTaskLogDo) Group(cols ...field.Expr) IAvtTaskLogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtTaskLogDo) Having(conds ...gen.Condition) IAvtTaskLogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtTaskLogDo) Limit(limit int) IAvtTaskLogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtTaskLogDo) Offset(offset int) IAvtTaskLogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtTaskLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtTaskLogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtTaskLogDo) Unscoped() IAvtTaskLogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtTaskLogDo) Create(values ...*model.AvtTaskLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtTaskLogDo) CreateInBatches(values []*model.AvtTaskLog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtTaskLogDo) Save(values ...*model.AvtTaskLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtTaskLogDo) First() (*model.AvtTaskLog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTaskLog), nil
	}
}

func (a avtTaskLogDo) Take() (*model.AvtTaskLog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTaskLog), nil
	}
}

func (a avtTaskLogDo) Last() (*model.AvtTaskLog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTaskLog), nil
	}
}

func (a avtTaskLogDo) Find() ([]*model.AvtTaskLog, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtTaskLog), err
}

func (a avtTaskLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtTaskLog, err error) {
	buf := make([]*model.AvtTaskLog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtTaskLogDo) FindInBatches(result *[]*model.AvtTaskLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtTaskLogDo) Attrs(attrs ...field.AssignExpr) IAvtTaskLogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtTaskLogDo) Assign(attrs ...field.AssignExpr) IAvtTaskLogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtTaskLogDo) Joins(fields ...field.RelationField) IAvtTaskLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtTaskLogDo) Preload(fields ...field.RelationField) IAvtTaskLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtTaskLogDo) FirstOrInit() (*model.AvtTaskLog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTaskLog), nil
	}
}

func (a avtTaskLogDo) FirstOrCreate() (*model.AvtTaskLog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTaskLog), nil
	}
}

func (a avtTaskLogDo) FindByPage(offset int, limit int) (result []*model.AvtTaskLog, count int64, err error) {
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

func (a avtTaskLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtTaskLogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtTaskLogDo) Delete(models ...*model.AvtTaskLog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtTaskLogDo) withDO(do gen.Dao) *avtTaskLogDo {
	a.DO = *do.(*gen.DO)
	return a
}
