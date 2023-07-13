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

func newAvtLockLog(db *gorm.DB, opts ...gen.DOOption) avtLockLog {
	_avtLockLog := avtLockLog{}

	_avtLockLog.avtLockLogDo.UseDB(db, opts...)
	_avtLockLog.avtLockLogDo.UseModel(&model.AvtLockLog{})

	tableName := _avtLockLog.avtLockLogDo.TableName()
	_avtLockLog.ALL = field.NewAsterisk(tableName)
	_avtLockLog.ID = field.NewInt64(tableName, "id")
	_avtLockLog.UnitNo = field.NewString(tableName, "unit_no")
	_avtLockLog.RoomName = field.NewString(tableName, "room_name")
	_avtLockLog.LockID = field.NewString(tableName, "lock_id")
	_avtLockLog.PeopleID = field.NewString(tableName, "people_id")
	_avtLockLog.PeopleName = field.NewString(tableName, "people_name")
	_avtLockLog.CardID = field.NewString(tableName, "card_id")
	_avtLockLog.Date = field.NewString(tableName, "date")
	_avtLockLog.OpenType = field.NewInt64(tableName, "open_type")
	_avtLockLog.AppID = field.NewString(tableName, "app_id")
	_avtLockLog.Time = field.NewInt64(tableName, "time")

	_avtLockLog.fillFieldMap()

	return _avtLockLog
}

type avtLockLog struct {
	avtLockLogDo

	ALL        field.Asterisk
	ID         field.Int64
	UnitNo     field.String
	RoomName   field.String
	LockID     field.String
	PeopleID   field.String
	PeopleName field.String
	CardID     field.String
	Date       field.String
	OpenType   field.Int64
	AppID      field.String
	Time       field.Int64

	fieldMap map[string]field.Expr
}

func (a avtLockLog) Table(newTableName string) *avtLockLog {
	a.avtLockLogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtLockLog) As(alias string) *avtLockLog {
	a.avtLockLogDo.DO = *(a.avtLockLogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtLockLog) updateTableName(table string) *avtLockLog {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UnitNo = field.NewString(table, "unit_no")
	a.RoomName = field.NewString(table, "room_name")
	a.LockID = field.NewString(table, "lock_id")
	a.PeopleID = field.NewString(table, "people_id")
	a.PeopleName = field.NewString(table, "people_name")
	a.CardID = field.NewString(table, "card_id")
	a.Date = field.NewString(table, "date")
	a.OpenType = field.NewInt64(table, "open_type")
	a.AppID = field.NewString(table, "app_id")
	a.Time = field.NewInt64(table, "time")

	a.fillFieldMap()

	return a
}

func (a *avtLockLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtLockLog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 11)
	a.fieldMap["id"] = a.ID
	a.fieldMap["unit_no"] = a.UnitNo
	a.fieldMap["room_name"] = a.RoomName
	a.fieldMap["lock_id"] = a.LockID
	a.fieldMap["people_id"] = a.PeopleID
	a.fieldMap["people_name"] = a.PeopleName
	a.fieldMap["card_id"] = a.CardID
	a.fieldMap["date"] = a.Date
	a.fieldMap["open_type"] = a.OpenType
	a.fieldMap["app_id"] = a.AppID
	a.fieldMap["time"] = a.Time
}

func (a avtLockLog) clone(db *gorm.DB) avtLockLog {
	a.avtLockLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtLockLog) replaceDB(db *gorm.DB) avtLockLog {
	a.avtLockLogDo.ReplaceDB(db)
	return a
}

type avtLockLogDo struct{ gen.DO }

type IAvtLockLogDo interface {
	gen.SubQuery
	Debug() IAvtLockLogDo
	WithContext(ctx context.Context) IAvtLockLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtLockLogDo
	WriteDB() IAvtLockLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtLockLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtLockLogDo
	Not(conds ...gen.Condition) IAvtLockLogDo
	Or(conds ...gen.Condition) IAvtLockLogDo
	Select(conds ...field.Expr) IAvtLockLogDo
	Where(conds ...gen.Condition) IAvtLockLogDo
	Order(conds ...field.Expr) IAvtLockLogDo
	Distinct(cols ...field.Expr) IAvtLockLogDo
	Omit(cols ...field.Expr) IAvtLockLogDo
	Join(table schema.Tabler, on ...field.Expr) IAvtLockLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtLockLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtLockLogDo
	Group(cols ...field.Expr) IAvtLockLogDo
	Having(conds ...gen.Condition) IAvtLockLogDo
	Limit(limit int) IAvtLockLogDo
	Offset(offset int) IAvtLockLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtLockLogDo
	Unscoped() IAvtLockLogDo
	Create(values ...*model.AvtLockLog) error
	CreateInBatches(values []*model.AvtLockLog, batchSize int) error
	Save(values ...*model.AvtLockLog) error
	First() (*model.AvtLockLog, error)
	Take() (*model.AvtLockLog, error)
	Last() (*model.AvtLockLog, error)
	Find() ([]*model.AvtLockLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtLockLog, err error)
	FindInBatches(result *[]*model.AvtLockLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtLockLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtLockLogDo
	Assign(attrs ...field.AssignExpr) IAvtLockLogDo
	Joins(fields ...field.RelationField) IAvtLockLogDo
	Preload(fields ...field.RelationField) IAvtLockLogDo
	FirstOrInit() (*model.AvtLockLog, error)
	FirstOrCreate() (*model.AvtLockLog, error)
	FindByPage(offset int, limit int) (result []*model.AvtLockLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtLockLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtLockLogDo) Debug() IAvtLockLogDo {
	return a.withDO(a.DO.Debug())
}

func (a avtLockLogDo) WithContext(ctx context.Context) IAvtLockLogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtLockLogDo) ReadDB() IAvtLockLogDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtLockLogDo) WriteDB() IAvtLockLogDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtLockLogDo) Session(config *gorm.Session) IAvtLockLogDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtLockLogDo) Clauses(conds ...clause.Expression) IAvtLockLogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtLockLogDo) Returning(value interface{}, columns ...string) IAvtLockLogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtLockLogDo) Not(conds ...gen.Condition) IAvtLockLogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtLockLogDo) Or(conds ...gen.Condition) IAvtLockLogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtLockLogDo) Select(conds ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtLockLogDo) Where(conds ...gen.Condition) IAvtLockLogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtLockLogDo) Order(conds ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtLockLogDo) Distinct(cols ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtLockLogDo) Omit(cols ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtLockLogDo) Join(table schema.Tabler, on ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtLockLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtLockLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtLockLogDo) Group(cols ...field.Expr) IAvtLockLogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtLockLogDo) Having(conds ...gen.Condition) IAvtLockLogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtLockLogDo) Limit(limit int) IAvtLockLogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtLockLogDo) Offset(offset int) IAvtLockLogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtLockLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtLockLogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtLockLogDo) Unscoped() IAvtLockLogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtLockLogDo) Create(values ...*model.AvtLockLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtLockLogDo) CreateInBatches(values []*model.AvtLockLog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtLockLogDo) Save(values ...*model.AvtLockLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtLockLogDo) First() (*model.AvtLockLog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtLockLog), nil
	}
}

func (a avtLockLogDo) Take() (*model.AvtLockLog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtLockLog), nil
	}
}

func (a avtLockLogDo) Last() (*model.AvtLockLog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtLockLog), nil
	}
}

func (a avtLockLogDo) Find() ([]*model.AvtLockLog, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtLockLog), err
}

func (a avtLockLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtLockLog, err error) {
	buf := make([]*model.AvtLockLog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtLockLogDo) FindInBatches(result *[]*model.AvtLockLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtLockLogDo) Attrs(attrs ...field.AssignExpr) IAvtLockLogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtLockLogDo) Assign(attrs ...field.AssignExpr) IAvtLockLogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtLockLogDo) Joins(fields ...field.RelationField) IAvtLockLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtLockLogDo) Preload(fields ...field.RelationField) IAvtLockLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtLockLogDo) FirstOrInit() (*model.AvtLockLog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtLockLog), nil
	}
}

func (a avtLockLogDo) FirstOrCreate() (*model.AvtLockLog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtLockLog), nil
	}
}

func (a avtLockLogDo) FindByPage(offset int, limit int) (result []*model.AvtLockLog, count int64, err error) {
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

func (a avtLockLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtLockLogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtLockLogDo) Delete(models ...*model.AvtLockLog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtLockLogDo) withDO(do gen.Dao) *avtLockLogDo {
	a.DO = *do.(*gen.DO)
	return a
}
