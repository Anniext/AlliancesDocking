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

func newAvtCard(db *gorm.DB, opts ...gen.DOOption) avtCard {
	_avtCard := avtCard{}

	_avtCard.avtCardDo.UseDB(db, opts...)
	_avtCard.avtCardDo.UseModel(&model.AvtCard{})

	tableName := _avtCard.avtCardDo.TableName()
	_avtCard.ALL = field.NewAsterisk(tableName)
	_avtCard.ID = field.NewInt64(tableName, "id")
	_avtCard.CardNo = field.NewString(tableName, "card_no")
	_avtCard.Name = field.NewString(tableName, "name")
	_avtCard.College = field.NewString(tableName, "college")
	_avtCard.Status = field.NewInt64(tableName, "status")
	_avtCard.CreateTime = field.NewTime(tableName, "create_time")
	_avtCard.UpdateTime = field.NewTime(tableName, "update_time")
	_avtCard.JobNumber = field.NewString(tableName, "job_number")

	_avtCard.fillFieldMap()

	return _avtCard
}

type avtCard struct {
	avtCardDo

	ALL        field.Asterisk
	ID         field.Int64
	CardNo     field.String // 卡号
	Name       field.String // 老师名称
	College    field.String // 学院
	Status     field.Int64  // 1-正常；2-禁用
	CreateTime field.Time
	UpdateTime field.Time
	JobNumber  field.String // 工号

	fieldMap map[string]field.Expr
}

func (a avtCard) Table(newTableName string) *avtCard {
	a.avtCardDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtCard) As(alias string) *avtCard {
	a.avtCardDo.DO = *(a.avtCardDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtCard) updateTableName(table string) *avtCard {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.CardNo = field.NewString(table, "card_no")
	a.Name = field.NewString(table, "name")
	a.College = field.NewString(table, "college")
	a.Status = field.NewInt64(table, "status")
	a.CreateTime = field.NewTime(table, "create_time")
	a.UpdateTime = field.NewTime(table, "update_time")
	a.JobNumber = field.NewString(table, "job_number")

	a.fillFieldMap()

	return a
}

func (a *avtCard) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtCard) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["card_no"] = a.CardNo
	a.fieldMap["name"] = a.Name
	a.fieldMap["college"] = a.College
	a.fieldMap["status"] = a.Status
	a.fieldMap["create_time"] = a.CreateTime
	a.fieldMap["update_time"] = a.UpdateTime
	a.fieldMap["job_number"] = a.JobNumber
}

func (a avtCard) clone(db *gorm.DB) avtCard {
	a.avtCardDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtCard) replaceDB(db *gorm.DB) avtCard {
	a.avtCardDo.ReplaceDB(db)
	return a
}

type avtCardDo struct{ gen.DO }

type IAvtCardDo interface {
	gen.SubQuery
	Debug() IAvtCardDo
	WithContext(ctx context.Context) IAvtCardDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtCardDo
	WriteDB() IAvtCardDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtCardDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtCardDo
	Not(conds ...gen.Condition) IAvtCardDo
	Or(conds ...gen.Condition) IAvtCardDo
	Select(conds ...field.Expr) IAvtCardDo
	Where(conds ...gen.Condition) IAvtCardDo
	Order(conds ...field.Expr) IAvtCardDo
	Distinct(cols ...field.Expr) IAvtCardDo
	Omit(cols ...field.Expr) IAvtCardDo
	Join(table schema.Tabler, on ...field.Expr) IAvtCardDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtCardDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtCardDo
	Group(cols ...field.Expr) IAvtCardDo
	Having(conds ...gen.Condition) IAvtCardDo
	Limit(limit int) IAvtCardDo
	Offset(offset int) IAvtCardDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtCardDo
	Unscoped() IAvtCardDo
	Create(values ...*model.AvtCard) error
	CreateInBatches(values []*model.AvtCard, batchSize int) error
	Save(values ...*model.AvtCard) error
	First() (*model.AvtCard, error)
	Take() (*model.AvtCard, error)
	Last() (*model.AvtCard, error)
	Find() ([]*model.AvtCard, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtCard, err error)
	FindInBatches(result *[]*model.AvtCard, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtCard) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtCardDo
	Assign(attrs ...field.AssignExpr) IAvtCardDo
	Joins(fields ...field.RelationField) IAvtCardDo
	Preload(fields ...field.RelationField) IAvtCardDo
	FirstOrInit() (*model.AvtCard, error)
	FirstOrCreate() (*model.AvtCard, error)
	FindByPage(offset int, limit int) (result []*model.AvtCard, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtCardDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtCardDo) Debug() IAvtCardDo {
	return a.withDO(a.DO.Debug())
}

func (a avtCardDo) WithContext(ctx context.Context) IAvtCardDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtCardDo) ReadDB() IAvtCardDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtCardDo) WriteDB() IAvtCardDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtCardDo) Session(config *gorm.Session) IAvtCardDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtCardDo) Clauses(conds ...clause.Expression) IAvtCardDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtCardDo) Returning(value interface{}, columns ...string) IAvtCardDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtCardDo) Not(conds ...gen.Condition) IAvtCardDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtCardDo) Or(conds ...gen.Condition) IAvtCardDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtCardDo) Select(conds ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtCardDo) Where(conds ...gen.Condition) IAvtCardDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtCardDo) Order(conds ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtCardDo) Distinct(cols ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtCardDo) Omit(cols ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtCardDo) Join(table schema.Tabler, on ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtCardDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtCardDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtCardDo) Group(cols ...field.Expr) IAvtCardDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtCardDo) Having(conds ...gen.Condition) IAvtCardDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtCardDo) Limit(limit int) IAvtCardDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtCardDo) Offset(offset int) IAvtCardDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtCardDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtCardDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtCardDo) Unscoped() IAvtCardDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtCardDo) Create(values ...*model.AvtCard) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtCardDo) CreateInBatches(values []*model.AvtCard, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtCardDo) Save(values ...*model.AvtCard) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtCardDo) First() (*model.AvtCard, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtCard), nil
	}
}

func (a avtCardDo) Take() (*model.AvtCard, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtCard), nil
	}
}

func (a avtCardDo) Last() (*model.AvtCard, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtCard), nil
	}
}

func (a avtCardDo) Find() ([]*model.AvtCard, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtCard), err
}

func (a avtCardDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtCard, err error) {
	buf := make([]*model.AvtCard, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtCardDo) FindInBatches(result *[]*model.AvtCard, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtCardDo) Attrs(attrs ...field.AssignExpr) IAvtCardDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtCardDo) Assign(attrs ...field.AssignExpr) IAvtCardDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtCardDo) Joins(fields ...field.RelationField) IAvtCardDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtCardDo) Preload(fields ...field.RelationField) IAvtCardDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtCardDo) FirstOrInit() (*model.AvtCard, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtCard), nil
	}
}

func (a avtCardDo) FirstOrCreate() (*model.AvtCard, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtCard), nil
	}
}

func (a avtCardDo) FindByPage(offset int, limit int) (result []*model.AvtCard, count int64, err error) {
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

func (a avtCardDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtCardDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtCardDo) Delete(models ...*model.AvtCard) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtCardDo) withDO(do gen.Dao) *avtCardDo {
	a.DO = *do.(*gen.DO)
	return a
}
