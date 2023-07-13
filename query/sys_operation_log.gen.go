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

func newSysOperationLog(db *gorm.DB, opts ...gen.DOOption) sysOperationLog {
	_sysOperationLog := sysOperationLog{}

	_sysOperationLog.sysOperationLogDo.UseDB(db, opts...)
	_sysOperationLog.sysOperationLogDo.UseModel(&model.SysOperationLog{})

	tableName := _sysOperationLog.sysOperationLogDo.TableName()
	_sysOperationLog.ALL = field.NewAsterisk(tableName)
	_sysOperationLog.ID = field.NewInt64(tableName, "id")
	_sysOperationLog.UserID = field.NewInt64(tableName, "user_id")
	_sysOperationLog.Path = field.NewString(tableName, "path")
	_sysOperationLog.Method = field.NewString(tableName, "method")
	_sysOperationLog.IP = field.NewString(tableName, "ip")
	_sysOperationLog.Input = field.NewString(tableName, "input")
	_sysOperationLog.CreatedAt = field.NewTime(tableName, "created_at")
	_sysOperationLog.UpdatedAt = field.NewTime(tableName, "updated_at")

	_sysOperationLog.fillFieldMap()

	return _sysOperationLog
}

type sysOperationLog struct {
	sysOperationLogDo

	ALL       field.Asterisk
	ID        field.Int64
	UserID    field.Int64
	Path      field.String
	Method    field.String
	IP        field.String
	Input     field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (s sysOperationLog) Table(newTableName string) *sysOperationLog {
	s.sysOperationLogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysOperationLog) As(alias string) *sysOperationLog {
	s.sysOperationLogDo.DO = *(s.sysOperationLogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysOperationLog) updateTableName(table string) *sysOperationLog {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt64(table, "id")
	s.UserID = field.NewInt64(table, "user_id")
	s.Path = field.NewString(table, "path")
	s.Method = field.NewString(table, "method")
	s.IP = field.NewString(table, "ip")
	s.Input = field.NewString(table, "input")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")

	s.fillFieldMap()

	return s
}

func (s *sysOperationLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysOperationLog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 8)
	s.fieldMap["id"] = s.ID
	s.fieldMap["user_id"] = s.UserID
	s.fieldMap["path"] = s.Path
	s.fieldMap["method"] = s.Method
	s.fieldMap["ip"] = s.IP
	s.fieldMap["input"] = s.Input
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
}

func (s sysOperationLog) clone(db *gorm.DB) sysOperationLog {
	s.sysOperationLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysOperationLog) replaceDB(db *gorm.DB) sysOperationLog {
	s.sysOperationLogDo.ReplaceDB(db)
	return s
}

type sysOperationLogDo struct{ gen.DO }

type ISysOperationLogDo interface {
	gen.SubQuery
	Debug() ISysOperationLogDo
	WithContext(ctx context.Context) ISysOperationLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISysOperationLogDo
	WriteDB() ISysOperationLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISysOperationLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISysOperationLogDo
	Not(conds ...gen.Condition) ISysOperationLogDo
	Or(conds ...gen.Condition) ISysOperationLogDo
	Select(conds ...field.Expr) ISysOperationLogDo
	Where(conds ...gen.Condition) ISysOperationLogDo
	Order(conds ...field.Expr) ISysOperationLogDo
	Distinct(cols ...field.Expr) ISysOperationLogDo
	Omit(cols ...field.Expr) ISysOperationLogDo
	Join(table schema.Tabler, on ...field.Expr) ISysOperationLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISysOperationLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISysOperationLogDo
	Group(cols ...field.Expr) ISysOperationLogDo
	Having(conds ...gen.Condition) ISysOperationLogDo
	Limit(limit int) ISysOperationLogDo
	Offset(offset int) ISysOperationLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISysOperationLogDo
	Unscoped() ISysOperationLogDo
	Create(values ...*model.SysOperationLog) error
	CreateInBatches(values []*model.SysOperationLog, batchSize int) error
	Save(values ...*model.SysOperationLog) error
	First() (*model.SysOperationLog, error)
	Take() (*model.SysOperationLog, error)
	Last() (*model.SysOperationLog, error)
	Find() ([]*model.SysOperationLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperationLog, err error)
	FindInBatches(result *[]*model.SysOperationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SysOperationLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISysOperationLogDo
	Assign(attrs ...field.AssignExpr) ISysOperationLogDo
	Joins(fields ...field.RelationField) ISysOperationLogDo
	Preload(fields ...field.RelationField) ISysOperationLogDo
	FirstOrInit() (*model.SysOperationLog, error)
	FirstOrCreate() (*model.SysOperationLog, error)
	FindByPage(offset int, limit int) (result []*model.SysOperationLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISysOperationLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sysOperationLogDo) Debug() ISysOperationLogDo {
	return s.withDO(s.DO.Debug())
}

func (s sysOperationLogDo) WithContext(ctx context.Context) ISysOperationLogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysOperationLogDo) ReadDB() ISysOperationLogDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysOperationLogDo) WriteDB() ISysOperationLogDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysOperationLogDo) Session(config *gorm.Session) ISysOperationLogDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysOperationLogDo) Clauses(conds ...clause.Expression) ISysOperationLogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysOperationLogDo) Returning(value interface{}, columns ...string) ISysOperationLogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysOperationLogDo) Not(conds ...gen.Condition) ISysOperationLogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysOperationLogDo) Or(conds ...gen.Condition) ISysOperationLogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysOperationLogDo) Select(conds ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysOperationLogDo) Where(conds ...gen.Condition) ISysOperationLogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysOperationLogDo) Order(conds ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysOperationLogDo) Distinct(cols ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysOperationLogDo) Omit(cols ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysOperationLogDo) Join(table schema.Tabler, on ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysOperationLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysOperationLogDo) RightJoin(table schema.Tabler, on ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysOperationLogDo) Group(cols ...field.Expr) ISysOperationLogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysOperationLogDo) Having(conds ...gen.Condition) ISysOperationLogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysOperationLogDo) Limit(limit int) ISysOperationLogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysOperationLogDo) Offset(offset int) ISysOperationLogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysOperationLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISysOperationLogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysOperationLogDo) Unscoped() ISysOperationLogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysOperationLogDo) Create(values ...*model.SysOperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysOperationLogDo) CreateInBatches(values []*model.SysOperationLog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysOperationLogDo) Save(values ...*model.SysOperationLog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysOperationLogDo) First() (*model.SysOperationLog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationLog), nil
	}
}

func (s sysOperationLogDo) Take() (*model.SysOperationLog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationLog), nil
	}
}

func (s sysOperationLogDo) Last() (*model.SysOperationLog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationLog), nil
	}
}

func (s sysOperationLogDo) Find() ([]*model.SysOperationLog, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysOperationLog), err
}

func (s sysOperationLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysOperationLog, err error) {
	buf := make([]*model.SysOperationLog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysOperationLogDo) FindInBatches(result *[]*model.SysOperationLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysOperationLogDo) Attrs(attrs ...field.AssignExpr) ISysOperationLogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysOperationLogDo) Assign(attrs ...field.AssignExpr) ISysOperationLogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysOperationLogDo) Joins(fields ...field.RelationField) ISysOperationLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysOperationLogDo) Preload(fields ...field.RelationField) ISysOperationLogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysOperationLogDo) FirstOrInit() (*model.SysOperationLog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationLog), nil
	}
}

func (s sysOperationLogDo) FirstOrCreate() (*model.SysOperationLog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysOperationLog), nil
	}
}

func (s sysOperationLogDo) FindByPage(offset int, limit int) (result []*model.SysOperationLog, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysOperationLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysOperationLogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysOperationLogDo) Delete(models ...*model.SysOperationLog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysOperationLogDo) withDO(do gen.Dao) *sysOperationLogDo {
	s.DO = *do.(*gen.DO)
	return s
}
