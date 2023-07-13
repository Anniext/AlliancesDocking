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

func newAvtFloor(db *gorm.DB, opts ...gen.DOOption) avtFloor {
	_avtFloor := avtFloor{}

	_avtFloor.avtFloorDo.UseDB(db, opts...)
	_avtFloor.avtFloorDo.UseModel(&model.AvtFloor{})

	tableName := _avtFloor.avtFloorDo.TableName()
	_avtFloor.ALL = field.NewAsterisk(tableName)
	_avtFloor.ID = field.NewInt64(tableName, "id")
	_avtFloor.Floor = field.NewString(tableName, "floor")
	_avtFloor.BuildingID = field.NewInt64(tableName, "building_id")
	_avtFloor.IsDelete = field.NewInt64(tableName, "is_delete")
	_avtFloor.Sort = field.NewInt64(tableName, "sort")
	_avtFloor.FloorCode = field.NewString(tableName, "floor_code")
	_avtFloor.RoomNum = field.NewInt64(tableName, "room_num")

	_avtFloor.fillFieldMap()

	return _avtFloor
}

type avtFloor struct {
	avtFloorDo

	ALL        field.Asterisk
	ID         field.Int64
	Floor      field.String
	BuildingID field.Int64
	IsDelete   field.Int64  // 0-未删除；1-已删除
	Sort       field.Int64  // 越大越靠前
	FloorCode  field.String // 楼层编号
	RoomNum    field.Int64  // 教室数量

	fieldMap map[string]field.Expr
}

func (a avtFloor) Table(newTableName string) *avtFloor {
	a.avtFloorDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtFloor) As(alias string) *avtFloor {
	a.avtFloorDo.DO = *(a.avtFloorDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtFloor) updateTableName(table string) *avtFloor {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Floor = field.NewString(table, "floor")
	a.BuildingID = field.NewInt64(table, "building_id")
	a.IsDelete = field.NewInt64(table, "is_delete")
	a.Sort = field.NewInt64(table, "sort")
	a.FloorCode = field.NewString(table, "floor_code")
	a.RoomNum = field.NewInt64(table, "room_num")

	a.fillFieldMap()

	return a
}

func (a *avtFloor) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtFloor) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 7)
	a.fieldMap["id"] = a.ID
	a.fieldMap["floor"] = a.Floor
	a.fieldMap["building_id"] = a.BuildingID
	a.fieldMap["is_delete"] = a.IsDelete
	a.fieldMap["sort"] = a.Sort
	a.fieldMap["floor_code"] = a.FloorCode
	a.fieldMap["room_num"] = a.RoomNum
}

func (a avtFloor) clone(db *gorm.DB) avtFloor {
	a.avtFloorDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtFloor) replaceDB(db *gorm.DB) avtFloor {
	a.avtFloorDo.ReplaceDB(db)
	return a
}

type avtFloorDo struct{ gen.DO }

type IAvtFloorDo interface {
	gen.SubQuery
	Debug() IAvtFloorDo
	WithContext(ctx context.Context) IAvtFloorDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtFloorDo
	WriteDB() IAvtFloorDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtFloorDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtFloorDo
	Not(conds ...gen.Condition) IAvtFloorDo
	Or(conds ...gen.Condition) IAvtFloorDo
	Select(conds ...field.Expr) IAvtFloorDo
	Where(conds ...gen.Condition) IAvtFloorDo
	Order(conds ...field.Expr) IAvtFloorDo
	Distinct(cols ...field.Expr) IAvtFloorDo
	Omit(cols ...field.Expr) IAvtFloorDo
	Join(table schema.Tabler, on ...field.Expr) IAvtFloorDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtFloorDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtFloorDo
	Group(cols ...field.Expr) IAvtFloorDo
	Having(conds ...gen.Condition) IAvtFloorDo
	Limit(limit int) IAvtFloorDo
	Offset(offset int) IAvtFloorDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtFloorDo
	Unscoped() IAvtFloorDo
	Create(values ...*model.AvtFloor) error
	CreateInBatches(values []*model.AvtFloor, batchSize int) error
	Save(values ...*model.AvtFloor) error
	First() (*model.AvtFloor, error)
	Take() (*model.AvtFloor, error)
	Last() (*model.AvtFloor, error)
	Find() ([]*model.AvtFloor, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtFloor, err error)
	FindInBatches(result *[]*model.AvtFloor, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtFloor) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtFloorDo
	Assign(attrs ...field.AssignExpr) IAvtFloorDo
	Joins(fields ...field.RelationField) IAvtFloorDo
	Preload(fields ...field.RelationField) IAvtFloorDo
	FirstOrInit() (*model.AvtFloor, error)
	FirstOrCreate() (*model.AvtFloor, error)
	FindByPage(offset int, limit int) (result []*model.AvtFloor, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtFloorDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtFloorDo) Debug() IAvtFloorDo {
	return a.withDO(a.DO.Debug())
}

func (a avtFloorDo) WithContext(ctx context.Context) IAvtFloorDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtFloorDo) ReadDB() IAvtFloorDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtFloorDo) WriteDB() IAvtFloorDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtFloorDo) Session(config *gorm.Session) IAvtFloorDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtFloorDo) Clauses(conds ...clause.Expression) IAvtFloorDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtFloorDo) Returning(value interface{}, columns ...string) IAvtFloorDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtFloorDo) Not(conds ...gen.Condition) IAvtFloorDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtFloorDo) Or(conds ...gen.Condition) IAvtFloorDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtFloorDo) Select(conds ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtFloorDo) Where(conds ...gen.Condition) IAvtFloorDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtFloorDo) Order(conds ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtFloorDo) Distinct(cols ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtFloorDo) Omit(cols ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtFloorDo) Join(table schema.Tabler, on ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtFloorDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtFloorDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtFloorDo) Group(cols ...field.Expr) IAvtFloorDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtFloorDo) Having(conds ...gen.Condition) IAvtFloorDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtFloorDo) Limit(limit int) IAvtFloorDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtFloorDo) Offset(offset int) IAvtFloorDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtFloorDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtFloorDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtFloorDo) Unscoped() IAvtFloorDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtFloorDo) Create(values ...*model.AvtFloor) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtFloorDo) CreateInBatches(values []*model.AvtFloor, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtFloorDo) Save(values ...*model.AvtFloor) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtFloorDo) First() (*model.AvtFloor, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtFloor), nil
	}
}

func (a avtFloorDo) Take() (*model.AvtFloor, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtFloor), nil
	}
}

func (a avtFloorDo) Last() (*model.AvtFloor, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtFloor), nil
	}
}

func (a avtFloorDo) Find() ([]*model.AvtFloor, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtFloor), err
}

func (a avtFloorDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtFloor, err error) {
	buf := make([]*model.AvtFloor, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtFloorDo) FindInBatches(result *[]*model.AvtFloor, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtFloorDo) Attrs(attrs ...field.AssignExpr) IAvtFloorDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtFloorDo) Assign(attrs ...field.AssignExpr) IAvtFloorDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtFloorDo) Joins(fields ...field.RelationField) IAvtFloorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtFloorDo) Preload(fields ...field.RelationField) IAvtFloorDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtFloorDo) FirstOrInit() (*model.AvtFloor, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtFloor), nil
	}
}

func (a avtFloorDo) FirstOrCreate() (*model.AvtFloor, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtFloor), nil
	}
}

func (a avtFloorDo) FindByPage(offset int, limit int) (result []*model.AvtFloor, count int64, err error) {
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

func (a avtFloorDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtFloorDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtFloorDo) Delete(models ...*model.AvtFloor) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtFloorDo) withDO(do gen.Dao) *avtFloorDo {
	a.DO = *do.(*gen.DO)
	return a
}