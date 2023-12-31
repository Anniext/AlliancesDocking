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

func newAvtTask(db *gorm.DB, opts ...gen.DOOption) avtTask {
	_avtTask := avtTask{}

	_avtTask.avtTaskDo.UseDB(db, opts...)
	_avtTask.avtTaskDo.UseModel(&model.AvtTask{})

	tableName := _avtTask.avtTaskDo.TableName()
	_avtTask.ALL = field.NewAsterisk(tableName)
	_avtTask.ID = field.NewInt64(tableName, "id")
	_avtTask.TaskName = field.NewString(tableName, "task_name")
	_avtTask.TaskType = field.NewString(tableName, "task_type")
	_avtTask.TaskCycle = field.NewString(tableName, "task_cycle")
	_avtTask.Description = field.NewString(tableName, "description")
	_avtTask.CronSpec = field.NewString(tableName, "cron_spec")
	_avtTask.Concurrent = field.NewInt64(tableName, "concurrent")
	_avtTask.Command = field.NewString(tableName, "command")
	_avtTask.Arge = field.NewString(tableName, "arge")
	_avtTask.Status = field.NewInt64(tableName, "status")
	_avtTask.Timeout = field.NewInt64(tableName, "timeout")
	_avtTask.VulnNum = field.NewInt64(tableName, "vuln_num")
	_avtTask.ExecuteTimes = field.NewInt64(tableName, "execute_times")
	_avtTask.PrevTime = field.NewInt64(tableName, "prev_time")
	_avtTask.CreatedTime = field.NewTime(tableName, "created_time")
	_avtTask.UpdatedTime = field.NewTime(tableName, "updated_time")
	_avtTask.CronConfig = field.NewString(tableName, "cron_config")
	_avtTask.TaskHour = field.NewString(tableName, "task_hour")
	_avtTask.TaskDay = field.NewInt64(tableName, "task_day")
	_avtTask.Target = field.NewString(tableName, "target")
	_avtTask.TaskGenType = field.NewInt64(tableName, "task_gen_type")

	_avtTask.fillFieldMap()

	return _avtTask
}

type avtTask struct {
	avtTaskDo

	ALL          field.Asterisk
	ID           field.Int64
	TaskName     field.String // 任务名称
	TaskType     field.String // 任务类型
	TaskCycle    field.String // 任务周期 now/执行一次 day每天xx点执行
	Description  field.String // 任务描述
	CronSpec     field.String // 时间表达式
	Concurrent   field.Int64  // 是否只允许一个实例
	Command      field.String // 命令详情
	Arge         field.String // url参数
	Status       field.Int64  // 0停用 1启用
	Timeout      field.Int64  // 超时设置
	VulnNum      field.Int64  // 漏洞数量
	ExecuteTimes field.Int64  // 累计执行次数
	PrevTime     field.Int64  // 上次执行时间
	CreatedTime  field.Time
	UpdatedTime  field.Time
	CronConfig   field.String // cron配置
	TaskHour     field.String
	TaskDay      field.Int64
	Target       field.String
	TaskGenType  field.Int64

	fieldMap map[string]field.Expr
}

func (a avtTask) Table(newTableName string) *avtTask {
	a.avtTaskDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a avtTask) As(alias string) *avtTask {
	a.avtTaskDo.DO = *(a.avtTaskDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *avtTask) updateTableName(table string) *avtTask {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.TaskName = field.NewString(table, "task_name")
	a.TaskType = field.NewString(table, "task_type")
	a.TaskCycle = field.NewString(table, "task_cycle")
	a.Description = field.NewString(table, "description")
	a.CronSpec = field.NewString(table, "cron_spec")
	a.Concurrent = field.NewInt64(table, "concurrent")
	a.Command = field.NewString(table, "command")
	a.Arge = field.NewString(table, "arge")
	a.Status = field.NewInt64(table, "status")
	a.Timeout = field.NewInt64(table, "timeout")
	a.VulnNum = field.NewInt64(table, "vuln_num")
	a.ExecuteTimes = field.NewInt64(table, "execute_times")
	a.PrevTime = field.NewInt64(table, "prev_time")
	a.CreatedTime = field.NewTime(table, "created_time")
	a.UpdatedTime = field.NewTime(table, "updated_time")
	a.CronConfig = field.NewString(table, "cron_config")
	a.TaskHour = field.NewString(table, "task_hour")
	a.TaskDay = field.NewInt64(table, "task_day")
	a.Target = field.NewString(table, "target")
	a.TaskGenType = field.NewInt64(table, "task_gen_type")

	a.fillFieldMap()

	return a
}

func (a *avtTask) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *avtTask) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 21)
	a.fieldMap["id"] = a.ID
	a.fieldMap["task_name"] = a.TaskName
	a.fieldMap["task_type"] = a.TaskType
	a.fieldMap["task_cycle"] = a.TaskCycle
	a.fieldMap["description"] = a.Description
	a.fieldMap["cron_spec"] = a.CronSpec
	a.fieldMap["concurrent"] = a.Concurrent
	a.fieldMap["command"] = a.Command
	a.fieldMap["arge"] = a.Arge
	a.fieldMap["status"] = a.Status
	a.fieldMap["timeout"] = a.Timeout
	a.fieldMap["vuln_num"] = a.VulnNum
	a.fieldMap["execute_times"] = a.ExecuteTimes
	a.fieldMap["prev_time"] = a.PrevTime
	a.fieldMap["created_time"] = a.CreatedTime
	a.fieldMap["updated_time"] = a.UpdatedTime
	a.fieldMap["cron_config"] = a.CronConfig
	a.fieldMap["task_hour"] = a.TaskHour
	a.fieldMap["task_day"] = a.TaskDay
	a.fieldMap["target"] = a.Target
	a.fieldMap["task_gen_type"] = a.TaskGenType
}

func (a avtTask) clone(db *gorm.DB) avtTask {
	a.avtTaskDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a avtTask) replaceDB(db *gorm.DB) avtTask {
	a.avtTaskDo.ReplaceDB(db)
	return a
}

type avtTaskDo struct{ gen.DO }

type IAvtTaskDo interface {
	gen.SubQuery
	Debug() IAvtTaskDo
	WithContext(ctx context.Context) IAvtTaskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAvtTaskDo
	WriteDB() IAvtTaskDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAvtTaskDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAvtTaskDo
	Not(conds ...gen.Condition) IAvtTaskDo
	Or(conds ...gen.Condition) IAvtTaskDo
	Select(conds ...field.Expr) IAvtTaskDo
	Where(conds ...gen.Condition) IAvtTaskDo
	Order(conds ...field.Expr) IAvtTaskDo
	Distinct(cols ...field.Expr) IAvtTaskDo
	Omit(cols ...field.Expr) IAvtTaskDo
	Join(table schema.Tabler, on ...field.Expr) IAvtTaskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAvtTaskDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAvtTaskDo
	Group(cols ...field.Expr) IAvtTaskDo
	Having(conds ...gen.Condition) IAvtTaskDo
	Limit(limit int) IAvtTaskDo
	Offset(offset int) IAvtTaskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtTaskDo
	Unscoped() IAvtTaskDo
	Create(values ...*model.AvtTask) error
	CreateInBatches(values []*model.AvtTask, batchSize int) error
	Save(values ...*model.AvtTask) error
	First() (*model.AvtTask, error)
	Take() (*model.AvtTask, error)
	Last() (*model.AvtTask, error)
	Find() ([]*model.AvtTask, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtTask, err error)
	FindInBatches(result *[]*model.AvtTask, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.AvtTask) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAvtTaskDo
	Assign(attrs ...field.AssignExpr) IAvtTaskDo
	Joins(fields ...field.RelationField) IAvtTaskDo
	Preload(fields ...field.RelationField) IAvtTaskDo
	FirstOrInit() (*model.AvtTask, error)
	FirstOrCreate() (*model.AvtTask, error)
	FindByPage(offset int, limit int) (result []*model.AvtTask, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAvtTaskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a avtTaskDo) Debug() IAvtTaskDo {
	return a.withDO(a.DO.Debug())
}

func (a avtTaskDo) WithContext(ctx context.Context) IAvtTaskDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a avtTaskDo) ReadDB() IAvtTaskDo {
	return a.Clauses(dbresolver.Read)
}

func (a avtTaskDo) WriteDB() IAvtTaskDo {
	return a.Clauses(dbresolver.Write)
}

func (a avtTaskDo) Session(config *gorm.Session) IAvtTaskDo {
	return a.withDO(a.DO.Session(config))
}

func (a avtTaskDo) Clauses(conds ...clause.Expression) IAvtTaskDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a avtTaskDo) Returning(value interface{}, columns ...string) IAvtTaskDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a avtTaskDo) Not(conds ...gen.Condition) IAvtTaskDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a avtTaskDo) Or(conds ...gen.Condition) IAvtTaskDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a avtTaskDo) Select(conds ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a avtTaskDo) Where(conds ...gen.Condition) IAvtTaskDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a avtTaskDo) Order(conds ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a avtTaskDo) Distinct(cols ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a avtTaskDo) Omit(cols ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a avtTaskDo) Join(table schema.Tabler, on ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a avtTaskDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a avtTaskDo) RightJoin(table schema.Tabler, on ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a avtTaskDo) Group(cols ...field.Expr) IAvtTaskDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a avtTaskDo) Having(conds ...gen.Condition) IAvtTaskDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a avtTaskDo) Limit(limit int) IAvtTaskDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a avtTaskDo) Offset(offset int) IAvtTaskDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a avtTaskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAvtTaskDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a avtTaskDo) Unscoped() IAvtTaskDo {
	return a.withDO(a.DO.Unscoped())
}

func (a avtTaskDo) Create(values ...*model.AvtTask) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a avtTaskDo) CreateInBatches(values []*model.AvtTask, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a avtTaskDo) Save(values ...*model.AvtTask) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a avtTaskDo) First() (*model.AvtTask, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTask), nil
	}
}

func (a avtTaskDo) Take() (*model.AvtTask, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTask), nil
	}
}

func (a avtTaskDo) Last() (*model.AvtTask, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTask), nil
	}
}

func (a avtTaskDo) Find() ([]*model.AvtTask, error) {
	result, err := a.DO.Find()
	return result.([]*model.AvtTask), err
}

func (a avtTaskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AvtTask, err error) {
	buf := make([]*model.AvtTask, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a avtTaskDo) FindInBatches(result *[]*model.AvtTask, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a avtTaskDo) Attrs(attrs ...field.AssignExpr) IAvtTaskDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a avtTaskDo) Assign(attrs ...field.AssignExpr) IAvtTaskDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a avtTaskDo) Joins(fields ...field.RelationField) IAvtTaskDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a avtTaskDo) Preload(fields ...field.RelationField) IAvtTaskDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a avtTaskDo) FirstOrInit() (*model.AvtTask, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTask), nil
	}
}

func (a avtTaskDo) FirstOrCreate() (*model.AvtTask, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AvtTask), nil
	}
}

func (a avtTaskDo) FindByPage(offset int, limit int) (result []*model.AvtTask, count int64, err error) {
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

func (a avtTaskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a avtTaskDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a avtTaskDo) Delete(models ...*model.AvtTask) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *avtTaskDo) withDO(do gen.Dao) *avtTaskDo {
	a.DO = *do.(*gen.DO)
	return a
}
