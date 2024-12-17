package korm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/kamalshkeir/kmap"
	"github.com/kamalshkeir/kstrct"
	"github.com/kamalshkeir/lg"
)

var (
	cachesOneM  = kmap.New[dbCache, map[string]any]()
	cacheAllM   = kmap.New[dbCache, []map[string]any](cacheMaxMemoryMb)
	setReplacer = strings.NewReplacer("=", "", "?", "")
)

// BuilderM is query builder map string any
type BuilderM struct {
	nocache    bool
	debug      bool
	limit      int
	page       int
	tableName  string
	selected   string
	orderBys   string
	whereQuery string
	offset     string
	statement  string
	db         *DatabaseEntity
	args       []any
	order      []string
	ctx        context.Context
	trace      bool
}

// Table is a starter for BuiderM
func Table(tableName string) *BuilderM {
	return &BuilderM{
		tableName: tableName,
		db:        &databases[0],
	}
}

func BuilderMap() *BuilderM {
	return &BuilderM{
		db: &databases[0],
	}
}

// Database allow to choose database to execute query on
func (b *BuilderM) Database(dbName string) *BuilderM {
	if b == nil || b.tableName == "" {
		lg.Error("korm.Table(tableName) first", "db", dbName)
		return b
	}
	db, err := GetMemoryDatabase(dbName)
	if lg.CheckError(err) {
		db = &databases[0]
	}
	b.db = db
	return b
}

// Select select table columns to return
func (b *BuilderM) Select(columns ...string) *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	for i := range columns {
		if !strings.HasPrefix(columns[i], "`") && !strings.HasPrefix(columns[i], "'") {
			columns[i] = "`" + columns[i] + "`"
		}
	}
	b.selected = strings.Join(columns, ",")
	b.order = append(b.order, "select")
	return b
}

// Where can be like: Where("id > ?",1) or Where("id = ?",1)
func (b *BuilderM) Where(query string, args ...any) *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	query = adaptConcatAndLen(query, b.db.Dialect)

	// Handle IN clauses
	var expandedArgs []any
	split := strings.Split(query, "?")
	var result strings.Builder
	argIndex := 0

	for i := range split {
		result.WriteString(split[i])
		if i < len(split)-1 && argIndex < len(args) {
			// Check if this placeholder is part of an IN clause
			beforePlaceholder := strings.TrimSpace(strings.ToUpper(split[i]))
			if strings.HasSuffix(beforePlaceholder, "IN") || strings.HasSuffix(beforePlaceholder, "IN (") {
				// Handle slice for IN clause
				switch v := args[argIndex].(type) {
				case []int:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					for _, val := range v {
						expandedArgs = append(expandedArgs, val)
					}
				case []int64:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					for _, val := range v {
						expandedArgs = append(expandedArgs, val)
					}
				case []uint:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					for _, val := range v {
						expandedArgs = append(expandedArgs, val)
					}
				case []uint8:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					for _, val := range v {
						expandedArgs = append(expandedArgs, val)
					}
				case []float32:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					for _, val := range v {
						expandedArgs = append(expandedArgs, val)
					}
				case []float64:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					for _, val := range v {
						expandedArgs = append(expandedArgs, val)
					}
				case []string:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					for _, val := range v {
						expandedArgs = append(expandedArgs, val)
					}
				case []any:
					result.WriteString(strings.Repeat("?,", len(v)-1) + "?")
					expandedArgs = append(expandedArgs, v...)
				default:
					// Not a slice, treat as normal arg
					result.WriteString("?")
					expandedArgs = append(expandedArgs, args[argIndex])
				}
			} else {
				// Normal argument
				result.WriteString("?")
				expandedArgs = append(expandedArgs, args[argIndex])
			}
			argIndex++
		}
	}

	adaptTimeToUnixArgs(&expandedArgs)
	b.whereQuery = result.String()
	b.args = append(b.args, expandedArgs...)
	b.order = append(b.order, "where")
	return b
}

// WhereNamed can be like : Where("email = :email",map[string]any{"email":"abc@mail.com"})
func (b *BuilderM) WhereNamed(query string, args map[string]any) *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	query = adaptConcatAndLen(query, b.db.Dialect)
	q, newargs, err := AdaptNamedParams(b.db.Dialect, query, args)
	if err != nil {
		b.whereQuery = query
		for _, v := range args {
			b.args = append(b.args, v)
		}
	} else {
		b.whereQuery = q
		b.args = newargs
	}
	b.order = append(b.order, "where")
	return b
}

// Limit set limit
func (b *BuilderM) Limit(limit int) *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	b.limit = limit
	b.order = append(b.order, "limit")
	return b
}

// Page return paginated elements using Limit for specific page
func (b *BuilderM) Page(pageNumber int) *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	b.page = pageNumber
	b.order = append(b.order, "page")
	return b
}

// OrderBy can be used like: OrderBy("-id","-email") OrderBy("id","-email") OrderBy("+id","email")
func (b *BuilderM) OrderBy(fields ...string) *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	b.orderBys = "ORDER BY "
	orders := []string{}
	for _, f := range fields {
		addTableName := false
		if b.tableName != "" {
			if !strings.Contains(f, b.tableName) {
				addTableName = true
			}
		}
		if strings.HasPrefix(f, "+") {
			if addTableName {
				orders = append(orders, b.tableName+"."+f[1:]+" ASC")
			} else {
				orders = append(orders, f[1:]+" ASC")
			}
		} else if strings.HasPrefix(f, "-") {
			if addTableName {
				orders = append(orders, b.tableName+"."+f[1:]+" DESC")
			} else {
				orders = append(orders, f[1:]+" DESC")
			}
		} else {
			if addTableName {
				orders = append(orders, b.tableName+"."+f+" ASC")
			} else {
				orders = append(orders, f+" ASC")
			}
		}
	}
	b.orderBys += strings.Join(orders, ",")
	b.order = append(b.order, "order_by")
	return b
}

// Context allow to query or execute using ctx
func (b *BuilderM) Context(ctx context.Context) *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	b.ctx = ctx
	return b
}

// Debug print prepared statement and values for this operation
func (b *BuilderM) Debug() *BuilderM {
	if b == nil || b.tableName == "" {
		return nil
	}
	b.debug = true
	return b
}

func (b *BuilderM) NoCache() *BuilderM {
	b.nocache = true
	return b
}

// All get all data
func (b *BuilderM) All() ([]map[string]any, error) {
	if b.trace {
		trace := TraceData{
			Query:     b.statement,
			Args:      b.args,
			Database:  b.db.Name,
			StartTime: time.Now(),
		}
		defer func() {
			trace.Duration = time.Since(trace.StartTime)
			defaultTracer.addTrace(trace)
		}()
	}

	if b == nil || b.tableName == "" {
		return nil, ErrTableNotFound
	}

	c := dbCache{
		database:   b.db.Name,
		table:      b.tableName,
		selected:   b.selected,
		statement:  b.statement,
		orderBys:   b.orderBys,
		whereQuery: b.whereQuery,
		offset:     b.offset,
		limit:      b.limit,
		page:       b.page,
		args:       fmt.Sprint(b.args...),
	}
	if useCache && !b.nocache {
		if v, ok := cacheAllM.Get(c); ok {
			return v, nil
		}
	}

	if b.selected != "" {
		b.statement = "select " + b.selected + " from " + b.tableName
	} else {
		b.statement = "select * from " + b.tableName
	}

	if b.whereQuery != "" {
		b.statement += " WHERE " + b.whereQuery
	}

	if b.orderBys != "" {
		b.statement += " " + b.orderBys
	}

	if b.limit > 0 {
		i := strconv.Itoa(b.limit)
		b.statement += " LIMIT " + i
		if b.page > 0 {
			o := strconv.Itoa((b.page - 1) * b.limit)
			b.statement += " OFFSET " + o
		}
	}

	if b.debug {
		lg.InfoC("debug", "statement", b.statement, "args", b.args)
	}
	models, err := b.QueryM(b.statement, b.args...)
	if err != nil {
		return nil, err
	}
	if useCache && !b.nocache {
		_ = cacheAllM.Set(c, models)
	}
	return models, nil
}

// One get single row
func (b *BuilderM) One() (map[string]any, error) {
	if b.trace {
		if b.ctx == nil {
			b.ctx = context.Background()
		}
		b.ctx = context.WithValue(b.ctx, traceEnabledKey, true)
	}

	if b == nil || b.tableName == "" {
		return nil, ErrTableNotFound
	}

	c := dbCache{
		database:   b.db.Name,
		table:      b.tableName,
		selected:   b.selected,
		statement:  b.statement,
		orderBys:   b.orderBys,
		whereQuery: b.whereQuery,
		offset:     b.offset,
		limit:      b.limit,
		page:       b.page,
		args:       fmt.Sprint(b.args...),
	}
	if useCache && !b.nocache {
		if v, ok := cachesOneM.Get(c); ok {
			return v, nil
		}
	}

	if b.selected != "" && b.selected != "*" {
		b.statement = "select " + b.selected + " from " + b.tableName
	} else {
		b.statement = "select * from " + b.tableName
	}

	if b.whereQuery != "" {
		b.statement += " WHERE " + b.whereQuery
	}

	if b.orderBys != "" {
		b.statement += " " + b.orderBys
	}

	if b.limit > 0 {
		i := strconv.Itoa(b.limit)
		b.statement += " LIMIT " + i
	} else {
		b.statement += " LIMIT 1"
	}

	if b.debug {
		lg.InfoC("debug", "statement", b.statement, "args", b.args)
	}

	models, err := b.Database(b.db.Name).QueryM(b.statement, b.args...)
	if err != nil {
		return nil, err
	}

	if len(models) == 0 {
		return nil, ErrNoData
	}
	if useCache && !b.nocache {
		_ = cachesOneM.Set(c, models[0])
	}

	return models[0], nil
}

// Insert add row to a table using input map, and return PK of the inserted row
func (b *BuilderM) Insert(rowData map[string]any) (int, error) {
	if b.trace {
		trace := TraceData{
			Query:     b.statement,
			Args:      b.args,
			Database:  b.db.Name,
			StartTime: time.Now(),
		}
		defer func() {
			trace.Duration = time.Since(trace.StartTime)
			defaultTracer.addTrace(trace)
		}()
	}

	if len(rowData) == 0 {
		return 0, fmt.Errorf("cannot insert empty map, rowData:%v", rowData)
	}
	if b == nil || b.tableName == "" {
		return 0, ErrTableNotFound
	}
	if onInsert != nil {
		err := onInsert(b.db.Name, b.tableName, rowData)
		if err != nil {
			return 0, err
		}

	}
	pk := ""
	var tbmem TableEntity
	for _, t := range b.db.Tables {
		if t.Name == b.tableName {
			pk = t.Pk
			tbmem = t
		}
	}

	placeholdersSlice := []string{}
	keys := []string{}
	values := []any{}
	count := 0
	for k, v := range rowData {
		switch b.db.Dialect {
		case POSTGRES, SQLITE:
			placeholdersSlice = append(placeholdersSlice, "$"+strconv.Itoa(count+1))
		case MYSQL, MARIA:
			placeholdersSlice = append(placeholdersSlice, "?")
		default:
			return 0, errors.New("database is neither sqlite3, postgres or mysql")
		}
		if !strings.HasPrefix(k, "`") && !strings.HasPrefix(k, "'") {
			keys = append(keys, "`"+k+"`")
		} else {
			keys = append(keys, k)
		}
		if v == true {
			v = 1
		} else if v == false {
			v = 0
		}

		if vvv, ok := tbmem.ModelTypes[k]; ok && strings.HasSuffix(vvv, "Time") {
			switch tyV := v.(type) {
			case time.Time:
				v = tyV.Unix()
			case string:
				v = strings.ReplaceAll(tyV, "T", " ")
			}
		}

		values = append(values, v)
		count++
	}
	placeholders := strings.Join(placeholdersSlice, ",")
	stat := strings.Builder{}
	stat.WriteString("INSERT INTO `" + b.tableName + "` (")
	stat.WriteString(strings.Join(keys, ","))
	stat.WriteString(") VALUES (")
	stat.WriteString(placeholders)
	stat.WriteString(")")
	statement := stat.String()
	var id int
	if b.db.Dialect != POSTGRES {
		if b.debug {
			lg.InfoC("debug", "statement", b.statement, "args", values)
		}
		var res sql.Result
		var err error
		if b.ctx != nil {
			res, err = b.db.Conn.ExecContext(b.ctx, statement, values...)
		} else {
			res, err = b.db.Conn.Exec(statement, values...)
		}
		if err != nil {
			return 0, err
		}
		rows, err := res.LastInsertId()
		if err != nil {
			id = -1
		} else {
			id = int(rows)
		}
	} else {
		if b.debug {
			lg.InfoC("debug", "statement", b.statement+" RETURNING "+pk, "args", values)
		}
		var err error
		if b.ctx != nil {
			err = b.db.Conn.QueryRowContext(b.ctx, statement+" RETURNING "+pk, values...).Scan(&id)
		} else {
			err = b.db.Conn.QueryRow(statement+" RETURNING "+pk, values...).Scan(&id)
		}
		if err != nil {
			id = -1
			return id, err
		}
	}
	return id, nil
}

// InsertR add row to a table using input map, and return the inserted row
func (b *BuilderM) InsertR(rowData map[string]any) (map[string]any, error) {
	if b.trace {
		trace := TraceData{
			Query:     b.statement,
			Args:      b.args,
			Database:  b.db.Name,
			StartTime: time.Now(),
		}
		defer func() {
			trace.Duration = time.Since(trace.StartTime)
			defaultTracer.addTrace(trace)
		}()
	}

	if b == nil || b.tableName == "" {
		return nil, ErrTableNotFound
	}
	if onInsert != nil {
		err := onInsert(b.db.Name, b.tableName, rowData)
		if err != nil {
			return nil, err
		}

	}
	pk := ""
	var tbmem TableEntity
	for _, t := range b.db.Tables {
		if t.Name == b.tableName {
			pk = t.Pk
			tbmem = t
		}
	}

	placeholdersSlice := []string{}
	keys := []string{}
	values := []any{}
	count := 0
	for k, v := range rowData {
		switch b.db.Dialect {
		case POSTGRES, SQLITE:
			placeholdersSlice = append(placeholdersSlice, "$"+strconv.Itoa(count+1))
		case MYSQL, MARIA:
			placeholdersSlice = append(placeholdersSlice, "?")
		default:
			return nil, errors.New("database is neither sqlite3, postgres or mysql")
		}
		if !strings.HasPrefix(k, "`") && !strings.HasPrefix(k, "'") {
			keys = append(keys, "`"+k+"`")
		} else {
			keys = append(keys, k)
		}

		if v == true {
			v = 1
		} else if v == false {
			v = 0
		}

		if vvv, ok := tbmem.ModelTypes[k]; ok && strings.HasSuffix(vvv, "Time") {
			switch tyV := v.(type) {
			case time.Time:
				v = tyV.Unix()
			case string:
				v = strings.ReplaceAll(tyV, "T", " ")
			}
		}

		values = append(values, v)
		count++
	}
	placeholders := strings.Join(placeholdersSlice, ",")
	stat := strings.Builder{}
	stat.WriteString("INSERT INTO `" + b.tableName + "` (")
	stat.WriteString(strings.Join(keys, ","))
	stat.WriteString(") VALUES (")
	stat.WriteString(placeholders)
	stat.WriteString(")")
	statement := stat.String()
	if b.debug {
		lg.InfoC("debug", "statement", statement, "args", values)
	}
	var id int
	if b.db.Dialect != POSTGRES {
		var res sql.Result
		var err error
		if b.ctx != nil {
			res, err = b.db.Conn.ExecContext(b.ctx, statement, values...)
		} else {
			res, err = b.db.Conn.Exec(statement, values...)
		}
		if err != nil {
			return nil, err
		}
		rows, err := res.LastInsertId()
		if err != nil {
			id = -1
		} else {
			id = int(rows)
		}
	} else {
		var err error
		if b.ctx != nil {
			err = b.db.Conn.QueryRowContext(b.ctx, statement+" RETURNING "+pk, values...).Scan(&id)
		} else {
			err = b.db.Conn.QueryRow(statement+" RETURNING "+pk, values...).Scan(&id)
		}
		if err != nil {
			return nil, err
		}
	}
	m, err := Table(b.tableName).Where(pk+"= ?", id).One()
	if err != nil {
		return nil, err
	}
	return m, nil
}

// BulkInsert insert many row at the same time in one query
func (b *BuilderM) BulkInsert(rowsData ...map[string]any) ([]int, error) {
	if b.trace {
		trace := TraceData{
			Query:     b.statement,
			Args:      b.args,
			Database:  b.db.Name,
			StartTime: time.Now(),
		}
		defer func() {
			trace.Duration = time.Since(trace.StartTime)
			defaultTracer.addTrace(trace)
		}()
	}

	if b == nil || b.tableName == "" {
		return nil, ErrTableNotFound
	}

	tx, err := b.db.Conn.BeginTx(context.Background(), &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	ids := []int{}
	pk := ""
	var tbmem TableEntity
	for _, t := range b.db.Tables {
		if t.Name == b.tableName {
			pk = t.Pk
			tbmem = t
		}
	}
	for ii := range rowsData {
		if onInsert != nil {
			err := onInsert(b.db.Name, b.tableName, rowsData[ii])
			if err != nil {
				return nil, err
			}
		}
		placeholdersSlice := []string{}
		keys := []string{}
		values := []any{}
		count := 0
		for k, v := range rowsData[ii] {
			switch b.db.Dialect {
			case POSTGRES, SQLITE:
				placeholdersSlice = append(placeholdersSlice, "$"+strconv.Itoa(count+1))
			case MYSQL, MARIA:
				placeholdersSlice = append(placeholdersSlice, "?")
			default:
				return nil, errors.New("database is neither sqlite3, postgres or mysql")
			}
			keys = append(keys, k)
			if v == true {
				v = 1
			} else if v == false {
				v = 0
			}
			if vvv, ok := tbmem.ModelTypes[k]; ok && strings.HasSuffix(vvv, "Time") {
				switch tyV := v.(type) {
				case time.Time:
					v = tyV.Unix()
				case string:
					v = strings.ReplaceAll(tyV, "T", " ")
				}
			}
			values = append(values, v)
			count++
		}
		placeholders := strings.Join(placeholdersSlice, ",")

		stat := strings.Builder{}
		stat.WriteString("INSERT INTO " + b.tableName + " (")
		stat.WriteString(strings.Join(keys, ","))
		stat.WriteString(") VALUES (")
		stat.WriteString(placeholders)
		stat.WriteString(")")
		statement := stat.String()
		if b.debug {
			lg.InfoC("debug", "statement", statement, "args", values)
		}
		if b.db.Dialect != POSTGRES {
			var res sql.Result
			var err error
			if b.ctx != nil {
				res, err = b.db.Conn.ExecContext(b.ctx, statement, values...)
			} else {
				res, err = b.db.Conn.Exec(statement, values...)
			}
			if err != nil {
				errRoll := tx.Rollback()
				if errRoll != nil {
					return nil, errRoll
				}
				return nil, err
			}
			idInserted, err := res.LastInsertId()
			if err != nil {
				return ids, err
			}
			ids = append(ids, int(idInserted))
		} else {
			var idInserted int
			if b.ctx != nil {
				err = b.db.Conn.QueryRowContext(b.ctx, statement+" RETURNING "+pk, values...).Scan(&idInserted)
			} else {
				err = b.db.Conn.QueryRow(statement+" RETURNING "+pk, values...).Scan(&idInserted)
			}
			if err != nil {
				return ids, err
			}
			ids = append(ids, idInserted)
		}
	}
	err = tx.Commit()
	if err != nil {
		return ids, err
	}
	return ids, nil
}

// Set used to update, Set("email,is_admin","example@mail.com",true) or Set("email = ? AND is_admin = ?","example@mail.com",true)
func (b *BuilderM) Set(query string, args ...any) (int, error) {
	if b.trace {
		trace := TraceData{
			Query:     b.statement,
			Args:      b.args,
			Database:  b.db.Name,
			StartTime: time.Now(),
		}
		defer func() {
			trace.Duration = time.Since(trace.StartTime)
			defaultTracer.addTrace(trace)
		}()
	}

	if b == nil || b.tableName == "" {
		return 0, ErrTableNotFound
	}
	if onSet != nil {
		mToSet := map[string]any{}
		sp := strings.Split(query, ",")
		if strings.Contains(query, "?") {
			for i := range sp {
				sp[i] = setReplacer.Replace(sp[i])
				mToSet[strings.TrimSpace(sp[i])] = args[i]
			}
		} else {
			for i := range sp {
				mToSet[strings.TrimSpace(sp[i])] = args[i]
			}
		}
		err := onSet(b.db.Name, b.tableName, mToSet)
		if err != nil {
			return 0, err
		}
	}
	if b.whereQuery == "" {
		return 0, errors.New("you should use Where before Update")
	}
	adaptSetQuery(&query)
	b.statement = "UPDATE " + b.tableName + " SET " + query + " WHERE " + b.whereQuery
	adaptTimeToUnixArgs(&args)
	AdaptPlaceholdersToDialect(&b.statement, b.db.Dialect)
	args = append(args, b.args...)
	if b.debug {
		lg.InfoC("debug", "statement", b.statement, "args", args)
	}

	var res sql.Result
	var err error
	if b.ctx != nil {
		res, err = b.db.Conn.ExecContext(b.ctx, b.statement, args...)
	} else {
		res, err = b.db.Conn.Exec(b.statement, args...)
	}
	if err != nil {
		return 0, err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(aff), nil
}

// Delete data from database, can be multiple, depending on the where, return affected rows(Not every database or database driver may support affected rows)
func (b *BuilderM) Delete() (int, error) {
	if b.trace {
		trace := TraceData{
			Query:     b.statement,
			Args:      b.args,
			Database:  b.db.Name,
			StartTime: time.Now(),
		}
		defer func() {
			trace.Duration = time.Since(trace.StartTime)
			defaultTracer.addTrace(trace)
		}()
	}

	if b == nil || b.tableName == "" {
		return 0, ErrTableNotFound
	}
	if onDelete != nil {
		err := onDelete(b.db.Name, b.tableName, b.whereQuery, b.args...)
		if err != nil {
			return 0, err
		}
	}
	b.statement = "DELETE FROM " + b.tableName
	if b.whereQuery != "" {
		b.statement += " WHERE " + b.whereQuery
	} else {
		return 0, errors.New("no Where was given for this query:" + b.whereQuery)
	}
	AdaptPlaceholdersToDialect(&b.statement, b.db.Dialect)
	if b.debug {
		lg.InfoC("debug", "statement", b.statement, "args", b.args)
	}

	var res sql.Result
	var err error
	if b.ctx != nil {
		res, err = b.db.Conn.ExecContext(b.ctx, b.statement, b.args...)
	} else {
		res, err = b.db.Conn.Exec(b.statement, b.args...)
	}
	if err != nil {
		return 0, err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return int(affectedRows), err
	}
	return int(affectedRows), nil
}

// Drop drop table from db
func (b *BuilderM) Drop() (int, error) {
	if b == nil || b.tableName == "" {
		return 0, ErrTableNotFound
	}
	if onDrop != nil {
		err := onDrop(b.db.Name, b.tableName)
		if err != nil {
			return 0, err
		}
	}
	b.statement = "DROP TABLE IF EXISTS " + b.tableName
	var res sql.Result
	var err error
	if b.ctx != nil {
		res, err = b.db.Conn.ExecContext(b.ctx, b.statement)
	} else {
		res, err = b.db.Conn.Exec(b.statement)
	}
	if err != nil {
		return 0, err
	}
	aff, err := res.RowsAffected()
	if err != nil {
		return int(aff), err
	}
	return int(aff), err
}

// AddRelated used for many to many, and after korm.ManyToMany, to add a class to a student or a student to a class, class or student should exist in the database before adding them
func (b *BuilderM) AddRelated(relatedTable string, whereRelatedTable string, whereRelatedArgs ...any) (int, error) {
	if b == nil || b.tableName == "" {
		return 0, errors.New("unable to find model, try korm.AutoMigrate before")
	}

	relationTableName := "m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable
	if _, ok := relationsMap.Get("m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable); !ok {
		relationTableName = "m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName
		if _, ok2 := relationsMap.Get("m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName); !ok2 {
			return 0, fmt.Errorf("no relations many to many between theses 2 tables: %s, %s", b.tableName, relatedTable)
		}
	}

	cols := ""
	wherecols := ""
	inOrder := false
	if strings.HasPrefix(relationTableName, "m2m_"+b.tableName) {
		inOrder = true
		relationTableName = "m2m_" + b.tableName + "_" + relatedTable
		cols = b.tableName + "_id," + relatedTable + "_id"
		wherecols = b.tableName + "_id = ? and " + relatedTable + "_id = ?"
	} else if strings.HasPrefix(relationTableName, "m2m_"+relatedTable) {
		relationTableName = "m2m_" + relatedTable + "_" + b.tableName
		cols = relatedTable + "_id," + b.tableName + "_id"
		wherecols = relatedTable + "_id = ? and " + b.tableName + "_id = ?"
	}

	memoryRelatedTable, err := GetMemoryTable(relatedTable)
	if err != nil {
		return 0, fmt.Errorf("memory table not found:" + relatedTable)
	}
	memoryTypedTable, err := GetMemoryTable(b.tableName)
	if err != nil {
		return 0, fmt.Errorf("memory table not found:" + relatedTable)
	}
	ids := make([]any, 4)
	adaptTimeToUnixArgs(&whereRelatedArgs)
	whereRelatedTable = adaptConcatAndLen(whereRelatedTable, b.db.Dialect)
	data, err := Table(relatedTable).Where(whereRelatedTable, whereRelatedArgs...).One()
	if err != nil {
		return 0, err
	}
	if v, ok := data[memoryRelatedTable.Pk]; ok {
		if inOrder {
			ids[1] = v
			ids[3] = v
		} else {
			ids[0] = v
			ids[2] = v
		}
	}
	// get the typed model
	if b.whereQuery == "" {
		return 0, fmt.Errorf("you must specify a where for the typed struct")
	}
	typedModel, err := Table(b.tableName).Where(b.whereQuery, b.args...).One()
	if err != nil {
		return 0, err
	}
	if v, ok := typedModel[memoryTypedTable.Pk]; ok {
		if inOrder {
			ids[0] = v
			ids[2] = v
		} else {
			ids[1] = v
			ids[3] = v
		}
	}
	if onInsert != nil {
		var mInsert map[string]any
		if inOrder {
			mInsert = map[string]any{
				b.tableName + "_id":  ids[0],
				relatedTable + "_id": ids[1],
			}
		} else {
			mInsert = map[string]any{
				b.tableName + "_id":  ids[1],
				relatedTable + "_id": ids[0],
			}
		}
		err := onInsert(b.db.Name, relationTableName, mInsert)
		if err != nil {
			return 0, err
		}
	}
	stat := "INSERT INTO " + relationTableName + "(" + cols + ") select ?,? WHERE NOT EXISTS (select * FROM " + relationTableName + " WHERE " + wherecols + ");"
	AdaptPlaceholdersToDialect(&stat, b.db.Dialect)
	if b.debug {
		lg.InfoC("debug", "statement", stat, "args", ids)
	}
	err = Exec(b.db.Name, stat, ids...)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// GetRelated used for many to many to get related classes to a student or related students to a class
func (b *BuilderM) GetRelated(relatedTable string, dest *[]map[string]any) error {
	if b == nil || b.tableName == "" {
		return errors.New("unable to find model, try db.Table before")
	}
	relationTableName := "m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable
	if _, ok := relationsMap.Get("m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable); !ok {
		relationTableName = "m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName
		if _, ok2 := relationsMap.Get("m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName); !ok2 {
			return fmt.Errorf("no relations many to many between theses 2 tables: %s, %s", b.tableName, relatedTable)
		}
	}

	if strings.HasPrefix(relationTableName, "m2m_"+b.tableName) {
		relationTableName = "m2m_" + b.tableName + "_" + relatedTable
	} else if strings.HasPrefix(relationTableName, "m2m_"+relatedTable) {
		relationTableName = "m2m_" + relatedTable + "_" + b.tableName
	}

	// get the typed model
	if b.whereQuery == "" {
		return fmt.Errorf("you must specify a where query like 'email = ?'")
	}
	b.whereQuery = strings.TrimSpace(b.whereQuery)
	if b.selected != "" {
		if !strings.Contains(b.selected, b.tableName) && !strings.Contains(b.selected, relatedTable) {
			if strings.Contains(b.selected, ",") {
				sp := strings.Split(b.selected, ",")
				for i := range sp {
					sp[i] = b.tableName + "." + sp[i]
				}
				b.selected = strings.Join(sp, ",")
			} else {
				b.selected = b.tableName + "." + b.selected
			}
		}
		b.statement = "select " + b.selected + " FROM " + relatedTable
	} else {
		b.statement = "select " + relatedTable + ".* FROM " + relatedTable
	}

	b.statement += " JOIN " + relationTableName + " ON " + relatedTable + ".id = " + relationTableName + "." + relatedTable + "_id"
	b.statement += " JOIN " + b.tableName + " ON " + b.tableName + ".id = " + relationTableName + "." + b.tableName + "_id"
	if !strings.Contains(b.whereQuery, b.tableName) {
		return fmt.Errorf("you should specify table name like : %s.id = ? , instead of %s", b.tableName, b.whereQuery)
	}
	b.statement += " WHERE " + b.whereQuery
	if b.orderBys != "" {
		b.statement += " " + b.orderBys
	}
	if b.limit > 0 {
		i := strconv.Itoa(b.limit)
		b.statement += " LIMIT " + i
		if b.page > 0 {
			o := strconv.Itoa((b.page - 1) * b.limit)
			b.statement += " OFFSET " + o
		}
	}
	if b.debug {
		lg.InfoC("debug", "statement", b.statement, "args", b.args)
	}
	var err error
	*dest, err = Table(relationTableName).Database(b.db.Name).QueryM(b.statement, b.args...)
	if err != nil {
		return err
	}

	return nil
}

// JoinRelated same as get, but it join data
func (b *BuilderM) JoinRelated(relatedTable string, dest *[]map[string]any) error {
	if b == nil || b.tableName == "" {
		return errors.New("unable to find model, try db.Table before")
	}
	relationTableName := "m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable
	if _, ok := relationsMap.Get("m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable); !ok {
		relationTableName = "m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName
		if _, ok2 := relationsMap.Get("m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName); !ok2 {
			return fmt.Errorf("no relations many to many between theses 2 tables: %s, %s", b.tableName, relatedTable)
		}
	}

	if strings.HasPrefix(relationTableName, "m2m_"+b.tableName) {
		relationTableName = "m2m_" + b.tableName + "_" + relatedTable
	} else if strings.HasPrefix(relationTableName, "m2m_"+relatedTable) {
		relationTableName = "m2m_" + relatedTable + "_" + b.tableName
	}

	// get the typed model
	if b.whereQuery == "" {
		return fmt.Errorf("you must specify a where query like 'email = ? and username like ...' for structs")
	}
	b.whereQuery = strings.TrimSpace(b.whereQuery)
	if b.selected != "" {
		if !strings.Contains(b.selected, b.tableName) && !strings.Contains(b.selected, relatedTable) {
			if strings.Contains(b.selected, ",") {
				sp := strings.Split(b.selected, ",")
				for i := range sp {
					sp[i] = b.tableName + "." + sp[i]
				}
				b.selected = strings.Join(sp, ",")
			} else {
				b.selected = b.tableName + "." + b.selected
			}
		}
		b.statement = "select " + b.selected + " FROM " + relatedTable
	} else {
		b.statement = "select " + relatedTable + ".*," + b.tableName + ".* FROM " + relatedTable
	}
	b.statement += " JOIN " + relationTableName + " ON " + relatedTable + ".id = " + relationTableName + "." + relatedTable + "_id"
	b.statement += " JOIN " + b.tableName + " ON " + b.tableName + ".id = " + relationTableName + "." + b.tableName + "_id"
	if !strings.Contains(b.whereQuery, b.tableName) {
		return fmt.Errorf("you should specify table name like : %s.id = ? , instead of %s", b.tableName, b.whereQuery)
	}
	b.statement += " WHERE " + b.whereQuery
	if b.orderBys != "" {
		b.statement += " " + b.orderBys
	}
	if b.limit > 0 {
		i := strconv.Itoa(b.limit)
		b.statement += " LIMIT " + i
		if b.page > 0 {
			o := strconv.Itoa((b.page - 1) * b.limit)
			b.statement += " OFFSET " + o
		}
	}
	if b.debug {
		lg.InfoC("debug", "statement", b.statement, "args", b.args)
	}
	var err error
	*dest, err = Table(relationTableName).QueryM(b.statement, b.args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteRelated delete a relations many to many
func (b *BuilderM) DeleteRelated(relatedTable string, whereRelatedTable string, whereRelatedArgs ...any) (int, error) {
	if b == nil || b.tableName == "" {
		return 0, errors.New("unable to find model, try db.Table before")
	}
	relationTableName := "m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable
	if _, ok := relationsMap.Get("m2m_" + b.tableName + "-" + b.db.Name + "-" + relatedTable); !ok {
		relationTableName = "m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName
		if _, ok2 := relationsMap.Get("m2m_" + relatedTable + "-" + b.db.Name + "-" + b.tableName); !ok2 {
			return 0, fmt.Errorf("no relations many to many between theses 2 tables: %s, %s", b.tableName, relatedTable)
		}
	}

	wherecols := ""
	inOrder := false
	if strings.HasPrefix(relationTableName, "m2m_"+b.tableName) {
		inOrder = true
		relationTableName = "m2m_" + b.tableName + "_" + relatedTable
		wherecols = b.tableName + "_id = ? and " + relatedTable + "_id = ?"
	} else if strings.HasPrefix(relationTableName, "m2m_"+relatedTable) {
		relationTableName = "m2m_" + relatedTable + "_" + b.tableName
		wherecols = relatedTable + "_id = ? and " + b.tableName + "_id = ?"
	}
	memoryRelatedTable, err := GetMemoryTable(relatedTable)
	if err != nil {
		return 0, fmt.Errorf("memory table not found:" + relatedTable)
	}
	memoryTypedTable, err := GetMemoryTable(b.tableName)
	if err != nil {
		return 0, fmt.Errorf("memory table not found:" + relatedTable)
	}
	ids := make([]any, 2)
	adaptTimeToUnixArgs(&whereRelatedArgs)
	whereRelatedTable = adaptConcatAndLen(whereRelatedTable, b.db.Dialect)

	data, err := Table(relatedTable).Where(whereRelatedTable, whereRelatedArgs...).One()
	if err != nil {
		return 0, err
	}
	if v, ok := data[memoryRelatedTable.Pk]; ok {
		if inOrder {
			ids[1] = v
		} else {
			ids[0] = v
		}
	}
	// get the typed model
	if b.whereQuery == "" {
		return 0, fmt.Errorf("you must specify a where for the typed struct")
	}
	typedModel, err := Table(b.tableName).Where(b.whereQuery, b.args...).One()
	if err != nil {
		return 0, err
	}
	if v, ok := typedModel[memoryTypedTable.Pk]; ok {
		if inOrder {
			ids[0] = v
		} else {
			ids[1] = v
		}
	}
	n, err := Table(relationTableName).Where(wherecols, ids...).Delete()
	if err != nil {
		return 0, err
	}
	return n, nil
}

// QueryM query sql and return result as slice maps
func (b *BuilderM) QueryM(statement string, args ...any) ([]map[string]any, error) {
	if b.trace {
		if b.ctx == nil {
			b.ctx = context.Background()
		}
		b.ctx = context.WithValue(b.ctx, traceEnabledKey, true)
	}

	if b.db.Conn == nil {
		return nil, errors.New("no connection")
	}
	c := dbCache{
		database:  b.db.Name,
		statement: statement,
		args:      fmt.Sprint(args...),
	}
	if useCache && !b.nocache {
		if v, ok := cacheQueryM.Get(c); ok {
			return v.([]map[string]any), nil
		}
	}
	AdaptPlaceholdersToDialect(&statement, b.db.Dialect)
	adaptTimeToUnixArgs(&args)
	var rows *sql.Rows
	var err error
	if b.ctx != nil {
		rows, err = b.db.Conn.QueryContext(b.ctx, statement, args...)
	} else {
		rows, err = b.db.Conn.Query(statement, args...)
	}
	if err == sql.ErrNoRows {
		return nil, ErrNoData
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()
	var columns []string
	if b.selected != "" && b.selected != "*" {
		columns = strings.Split(b.selected, ",")
	} else {
		columns, err = rows.Columns()
		if err != nil {
			return nil, err
		}
	}

	models := make([]any, len(columns))
	modelsPtrs := make([]any, len(columns))

	listMap := make([]map[string]any, 0)

	for rows.Next() {
		for i := range models {
			models[i] = &modelsPtrs[i]
		}

		err := rows.Scan(models...)
		if err != nil {
			return nil, err
		}

		m := map[string]any{}
		for i := range columns {
			if b.db.Dialect == MYSQL || b.db.Dialect == MARIA {
				if v, ok := modelsPtrs[i].([]byte); ok {
					modelsPtrs[i] = string(v)
				}
			}
			m[columns[i]] = modelsPtrs[i]
		}
		listMap = append(listMap, m)
	}
	if len(listMap) == 0 {
		return nil, ErrNoData
	}
	if useCache && !b.nocache {
		_ = cacheQueryM.Set(c, listMap)
	}
	return listMap, nil
}

// QueryMNamed query sql and return result as slice maps
//
// Example:
//
//		QueryMNamed("select * from users where email = :email",map[string]any{
//			"email":"email@mail.com",
//	    })
func (b *BuilderM) QueryMNamed(statement string, args map[string]any, unsafe ...bool) ([]map[string]any, error) {
	if b.trace {
		if b.ctx == nil {
			b.ctx = context.Background()
		}
		b.ctx = context.WithValue(b.ctx, traceEnabledKey, true)
	}

	if b.db.Conn == nil {
		return nil, errors.New("no connection")
	}
	rgs := ""
	for _, v := range args {
		rgs += fmt.Sprint(v)
	}
	c := dbCache{
		database:  b.db.Name,
		statement: statement,
		args:      rgs,
	}
	if useCache && !b.nocache {
		if v, ok := cacheQueryM.Get(c); ok {
			return v.([]map[string]any), nil
		}
	}
	var query string
	var newargs []any
	if len(unsafe) > 0 && unsafe[0] {
		var err error
		query, err = UnsafeNamedQuery(statement, args)
		if err != nil {
			return nil, err
		}
	} else {
		var err error
		query, newargs, err = AdaptNamedParams(b.db.Dialect, statement, args)
		if err != nil {
			return nil, err
		}
	}
	var rows *sql.Rows
	var err error
	if b.ctx != nil {
		rows, err = b.db.Conn.QueryContext(b.ctx, query, newargs...)
	} else {
		rows, err = b.db.Conn.Query(query, newargs...)
	}
	if err == sql.ErrNoRows {
		return nil, ErrNoData
	} else if err != nil {
		return nil, err
	}
	defer rows.Close()
	var columns []string
	if b.selected != "" && b.selected != "*" {
		columns = strings.Split(b.selected, ",")
	} else {
		columns, err = rows.Columns()
		if err != nil {
			return nil, err
		}
	}

	models := make([]any, len(columns))
	modelsPtrs := make([]any, len(columns))

	listMap := make([]map[string]any, 0)

	for rows.Next() {
		for i := range models {
			models[i] = &modelsPtrs[i]
		}

		err := rows.Scan(models...)
		if err != nil {
			return nil, err
		}

		m := map[string]any{}
		for i := range columns {
			if b.db.Dialect == MYSQL || b.db.Dialect == MARIA {
				if v, ok := modelsPtrs[i].([]byte); ok {
					modelsPtrs[i] = string(v)
				}
			}
			m[columns[i]] = modelsPtrs[i]
		}
		listMap = append(listMap, m)
	}
	if len(listMap) == 0 {
		return nil, ErrNoData
	}
	if useCache && !b.nocache {
		_ = cacheQueryM.Set(c, listMap)
	}
	return listMap, nil
}

func (b *BuilderM) queryS(ptrStrctSlice any, statement string, args ...any) error {
	if b == nil || b.tableName == "" {
		return ErrTableNotFound
	}
	AdaptPlaceholdersToDialect(&statement, b.db.Dialect)
	adaptTimeToUnixArgs(&args)
	if b.db.Conn == nil {
		return errors.New("no connection")
	}
	var rows *sql.Rows
	var err error
	if b.ctx != nil {
		rows, err = b.db.Conn.QueryContext(b.ctx, statement, args...)
	} else {
		rows, err = b.db.Conn.Query(statement, args...)
	}
	if err == sql.ErrNoRows {
		return ErrNoData
	} else if err != nil {
		return err
	}
	defer rows.Close()
	var columns []string
	if b.selected != "" && b.selected != "*" {
		columns = strings.Split(b.selected, ",")
	} else {
		columns, err = rows.Columns()
		if err != nil {
			return err
		}
	}
	models := make([]any, len(columns))
	modelsPtrs := make([]any, len(columns))

	var value = reflect.ValueOf(ptrStrctSlice)
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	} else {
		return errors.New("expected destination struct to be a pointer")
	}

	if value.Kind() != reflect.Slice {
		return fmt.Errorf("expected strct to be a ptr slice")
	}

	for rows.Next() {
		for i := range models {
			models[i] = &modelsPtrs[i]
		}

		err := rows.Scan(models...)
		if err != nil {
			return err
		}

		m := map[string]any{}
		if b.selected != "" && b.selected != "*" {
			for i, key := range strings.Split(b.selected, ",") {
				if b.db.Dialect == MYSQL || b.db.Dialect == MARIA {
					if v, ok := modelsPtrs[i].([]byte); ok {
						modelsPtrs[i] = string(v)
					}
				}
				m[key] = modelsPtrs[i]
			}
		} else {
			for i, key := range columns {
				if b.db.Dialect == MYSQL || b.db.Dialect == MARIA {
					if v, ok := modelsPtrs[i].([]byte); ok {
						modelsPtrs[i] = string(v)
					}
				}
				m[key] = modelsPtrs[i]
			}
		}
		ptr := reflect.New(value.Type().Elem()).Interface()
		err = kstrct.FillM(ptr, m)
		if err != nil {
			return err
		}
		if value.CanAddr() && value.CanSet() {
			value.Set(reflect.Append(value, reflect.ValueOf(ptr).Elem()))
		}
	}
	return nil
}

func (b *BuilderM) Trace() *BuilderM {
	if b == nil {
		return nil
	}
	b.trace = true
	if b.ctx == nil {
		b.ctx = context.Background()
	}
	b.ctx = context.WithValue(b.ctx, traceEnabledKey, true)
	return b
}
