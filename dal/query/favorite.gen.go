// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"ByteDance/dal/model"
)

func newFavorite(db *gorm.DB) favorite {
	_favorite := favorite{}

	_favorite.favoriteDo.UseDB(db)
	_favorite.favoriteDo.UseModel(&model.Favorite{})

	tableName := _favorite.favoriteDo.TableName()
	_favorite.ALL = field.NewField(tableName, "*")
	_favorite.ID = field.NewInt32(tableName, "id")
	_favorite.VideoID = field.NewInt32(tableName, "video_id")
	_favorite.UserID = field.NewInt32(tableName, "user_id")
	_favorite.Removed = field.NewInt32(tableName, "removed")
	_favorite.Deleted = field.NewInt32(tableName, "deleted")

	_favorite.fillFieldMap()

	return _favorite
}

type favorite struct {
	favoriteDo

	ALL     field.Field
	ID      field.Int32
	VideoID field.Int32
	UserID  field.Int32
	Removed field.Int32
	Deleted field.Int32

	fieldMap map[string]field.Expr
}

func (f favorite) Table(newTableName string) *favorite {
	f.favoriteDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f favorite) As(alias string) *favorite {
	f.favoriteDo.DO = *(f.favoriteDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *favorite) updateTableName(table string) *favorite {
	f.ALL = field.NewField(table, "*")
	f.ID = field.NewInt32(table, "id")
	f.VideoID = field.NewInt32(table, "video_id")
	f.UserID = field.NewInt32(table, "user_id")
	f.Removed = field.NewInt32(table, "removed")
	f.Deleted = field.NewInt32(table, "deleted")

	f.fillFieldMap()

	return f
}

func (f *favorite) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *favorite) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 5)
	f.fieldMap["id"] = f.ID
	f.fieldMap["video_id"] = f.VideoID
	f.fieldMap["user_id"] = f.UserID
	f.fieldMap["removed"] = f.Removed
	f.fieldMap["deleted"] = f.Deleted
}

func (f favorite) clone(db *gorm.DB) favorite {
	f.favoriteDo.ReplaceDB(db)
	return f
}

type favoriteDo struct{ gen.DO }

//查询视频点赞数目
//
//select count(1) from favorite where video_id = @videoID and removed = 0 and deleted = 0
func (f favoriteDo) QueryFavoriteCount(videoID int32) (result int64) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["videoID"] = videoID
	generateSQL.WriteString("select count(1) from favorite where video_id = @videoID and removed = 0 and deleted = 0 ")

	if len(params) > 0 {
		_ = f.UnderlyingDB().Raw(generateSQL.String(), params).Take(&result)
	} else {
		_ = f.UnderlyingDB().Raw(generateSQL.String()).Take(&result)
	}
	return
}

//removed置反
//
//update favorite set removed = -removed where user_id = @userID and video_id = @videoID and deleted = 0
func (f favoriteDo) UpdateFavoriteRemoved(userID int32, videoID int32) (rowsAffected int64, err error) {
	params := make(map[string]interface{}, 0)

	var generateSQL strings.Builder
	params["userID"] = userID
	params["videoID"] = videoID
	generateSQL.WriteString("update favorite set removed = -removed where user_id = @userID and video_id = @videoID and deleted = 0 ")

	var executeSQL *gorm.DB
	if len(params) > 0 {
		executeSQL = f.UnderlyingDB().Exec(generateSQL.String(), params)
	} else {
		executeSQL = f.UnderlyingDB().Exec(generateSQL.String())
	}
	rowsAffected = executeSQL.RowsAffected
	err = executeSQL.Error
	return
}

func (f favoriteDo) Debug() *favoriteDo {
	return f.withDO(f.DO.Debug())
}

func (f favoriteDo) WithContext(ctx context.Context) *favoriteDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f favoriteDo) Clauses(conds ...clause.Expression) *favoriteDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f favoriteDo) Returning(value interface{}, columns ...string) *favoriteDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f favoriteDo) Not(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f favoriteDo) Or(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f favoriteDo) Select(conds ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f favoriteDo) Where(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f favoriteDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *favoriteDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f favoriteDo) Order(conds ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f favoriteDo) Distinct(cols ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f favoriteDo) Omit(cols ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f favoriteDo) Join(table schema.Tabler, on ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f favoriteDo) LeftJoin(table schema.Tabler, on ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f favoriteDo) RightJoin(table schema.Tabler, on ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f favoriteDo) Group(cols ...field.Expr) *favoriteDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f favoriteDo) Having(conds ...gen.Condition) *favoriteDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f favoriteDo) Limit(limit int) *favoriteDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f favoriteDo) Offset(offset int) *favoriteDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f favoriteDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *favoriteDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f favoriteDo) Unscoped() *favoriteDo {
	return f.withDO(f.DO.Unscoped())
}

func (f favoriteDo) Create(values ...*model.Favorite) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f favoriteDo) CreateInBatches(values []*model.Favorite, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f favoriteDo) Save(values ...*model.Favorite) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f favoriteDo) First() (*model.Favorite, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) Take() (*model.Favorite, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) Last() (*model.Favorite, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) Find() ([]*model.Favorite, error) {
	result, err := f.DO.Find()
	return result.([]*model.Favorite), err
}

func (f favoriteDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Favorite, err error) {
	buf := make([]*model.Favorite, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f favoriteDo) FindInBatches(result *[]*model.Favorite, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f favoriteDo) Attrs(attrs ...field.AssignExpr) *favoriteDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f favoriteDo) Assign(attrs ...field.AssignExpr) *favoriteDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f favoriteDo) Joins(fields ...field.RelationField) *favoriteDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f favoriteDo) Preload(fields ...field.RelationField) *favoriteDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f favoriteDo) FirstOrInit() (*model.Favorite, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) FirstOrCreate() (*model.Favorite, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Favorite), nil
	}
}

func (f favoriteDo) FindByPage(offset int, limit int) (result []*model.Favorite, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f favoriteDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f *favoriteDo) withDO(do gen.Dao) *favoriteDo {
	f.DO = *do.(*gen.DO)
	return f
}
