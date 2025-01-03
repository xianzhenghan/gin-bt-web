// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"gin-api-mono/internal/repository/mysql/model"
)

func newCode(db *gorm.DB, opts ...gen.DOOption) code {
	_code := code{}

	_code.codeDo.UseDB(db, opts...)
	_code.codeDo.UseModel(&model.Code{})

	tableName := _code.codeDo.TableName()
	_code.ALL = field.NewAsterisk(tableName)
	_code.ID = field.NewInt64(tableName, "id")
	_code.UserID = field.NewInt64(tableName, "user_id")
	_code.Username = field.NewString(tableName, "username")
	_code.Title = field.NewString(tableName, "title")
	_code.Description = field.NewString(tableName, "description")
	_code.CodeContent = field.NewString(tableName, "code_content")
	_code.Language = field.NewString(tableName, "language")
	_code.CreatedAt = field.NewTime(tableName, "created_at")
	_code.UpdatedAt = field.NewTime(tableName, "updated_at")
	_code.IsPublic = field.NewBool(tableName, "is_public")
	_code.Tags = field.NewString(tableName, "tags")
	_code.Version = field.NewString(tableName, "version")
	_code.FileName = field.NewString(tableName, "file_name")
	_code.LineNumbers = field.NewBool(tableName, "line_numbers")

	_code.fillFieldMap()

	return _code
}

type code struct {
	codeDo

	ALL         field.Asterisk
	ID          field.Int64
	UserID      field.Int64
	Username    field.String // 用户名
	Title       field.String
	Description field.String
	CodeContent field.String
	Language    field.String
	CreatedAt   field.Time
	UpdatedAt   field.Time
	IsPublic    field.Bool
	Tags        field.String
	Version     field.String
	FileName    field.String
	LineNumbers field.Bool

	fieldMap map[string]field.Expr
}

func (c code) Table(newTableName string) *code {
	c.codeDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c code) As(alias string) *code {
	c.codeDo.DO = *(c.codeDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *code) updateTableName(table string) *code {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.UserID = field.NewInt64(table, "user_id")
	c.Username = field.NewString(table, "username")
	c.Title = field.NewString(table, "title")
	c.Description = field.NewString(table, "description")
	c.CodeContent = field.NewString(table, "code_content")
	c.Language = field.NewString(table, "language")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.IsPublic = field.NewBool(table, "is_public")
	c.Tags = field.NewString(table, "tags")
	c.Version = field.NewString(table, "version")
	c.FileName = field.NewString(table, "file_name")
	c.LineNumbers = field.NewBool(table, "line_numbers")

	c.fillFieldMap()

	return c
}

func (c *code) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *code) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 14)
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["username"] = c.Username
	c.fieldMap["title"] = c.Title
	c.fieldMap["description"] = c.Description
	c.fieldMap["code_content"] = c.CodeContent
	c.fieldMap["language"] = c.Language
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["is_public"] = c.IsPublic
	c.fieldMap["tags"] = c.Tags
	c.fieldMap["version"] = c.Version
	c.fieldMap["file_name"] = c.FileName
	c.fieldMap["line_numbers"] = c.LineNumbers
}

func (c code) clone(db *gorm.DB) code {
	c.codeDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c code) replaceDB(db *gorm.DB) code {
	c.codeDo.ReplaceDB(db)
	return c
}

type codeDo struct{ gen.DO }

func (c codeDo) Debug() *codeDo {
	return c.withDO(c.DO.Debug())
}

func (c codeDo) WithContext(ctx context.Context) *codeDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c codeDo) ReadDB() *codeDo {
	return c.Clauses(dbresolver.Read)
}

func (c codeDo) WriteDB() *codeDo {
	return c.Clauses(dbresolver.Write)
}

func (c codeDo) Session(config *gorm.Session) *codeDo {
	return c.withDO(c.DO.Session(config))
}

func (c codeDo) Clauses(conds ...clause.Expression) *codeDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c codeDo) Returning(value interface{}, columns ...string) *codeDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c codeDo) Not(conds ...gen.Condition) *codeDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c codeDo) Or(conds ...gen.Condition) *codeDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c codeDo) Select(conds ...field.Expr) *codeDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c codeDo) Where(conds ...gen.Condition) *codeDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c codeDo) Order(conds ...field.Expr) *codeDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c codeDo) Distinct(cols ...field.Expr) *codeDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c codeDo) Omit(cols ...field.Expr) *codeDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c codeDo) Join(table schema.Tabler, on ...field.Expr) *codeDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c codeDo) LeftJoin(table schema.Tabler, on ...field.Expr) *codeDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c codeDo) RightJoin(table schema.Tabler, on ...field.Expr) *codeDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c codeDo) Group(cols ...field.Expr) *codeDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c codeDo) Having(conds ...gen.Condition) *codeDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c codeDo) Limit(limit int) *codeDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c codeDo) Offset(offset int) *codeDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c codeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *codeDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c codeDo) Unscoped() *codeDo {
	return c.withDO(c.DO.Unscoped())
}

func (c codeDo) Create(values ...*model.Code) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c codeDo) CreateInBatches(values []*model.Code, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c codeDo) Save(values ...*model.Code) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c codeDo) First() (*model.Code, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Code), nil
	}
}

func (c codeDo) Take() (*model.Code, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Code), nil
	}
}

func (c codeDo) Last() (*model.Code, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Code), nil
	}
}

func (c codeDo) Find() ([]*model.Code, error) {
	result, err := c.DO.Find()
	return result.([]*model.Code), err
}

func (c codeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Code, err error) {
	buf := make([]*model.Code, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c codeDo) FindInBatches(result *[]*model.Code, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c codeDo) Attrs(attrs ...field.AssignExpr) *codeDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c codeDo) Assign(attrs ...field.AssignExpr) *codeDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c codeDo) Joins(fields ...field.RelationField) *codeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c codeDo) Preload(fields ...field.RelationField) *codeDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c codeDo) FirstOrInit() (*model.Code, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Code), nil
	}
}

func (c codeDo) FirstOrCreate() (*model.Code, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Code), nil
	}
}

func (c codeDo) FindByPage(offset int, limit int) (result []*model.Code, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c codeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c codeDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c codeDo) Delete(models ...*model.Code) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *codeDo) withDO(do gen.Dao) *codeDo {
	c.DO = *do.(*gen.DO)
	return c
}
