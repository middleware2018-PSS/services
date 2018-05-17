// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"gopkg.in/volatiletech/null.v6"
)

// Admin is an object representing the database table.
type Admin struct {
	ID       int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Password string      `boil:"password" json:"password" toml:"password" yaml:"password"`
	Info     null.String `boil:"info" json:"info,omitempty" toml:"info" yaml:"info,omitempty"`
	Name     null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Surname  null.String `boil:"surname" json:"surname,omitempty" toml:"surname" yaml:"surname,omitempty"`

	R *adminR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L adminL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AdminColumns = struct {
	ID       string
	Password string
	Info     string
	Name     string
	Surname  string
}{
	ID:       "id",
	Password: "password",
	Info:     "info",
	Name:     "name",
	Surname:  "surname",
}

// adminR is where relationships are stored.
type adminR struct {
}

// adminL is where Load methods for each relationship are stored.
type adminL struct{}

var (
	adminColumns               = []string{"id", "password", "info", "name", "surname"}
	adminColumnsWithoutDefault = []string{}
	adminColumnsWithDefault    = []string{"id", "password", "info", "name", "surname"}
	adminPrimaryKeyColumns     = []string{"id"}
)

type (
	// AdminSlice is an alias for a slice of pointers to Admin.
	// This should generally be used opposed to []Admin.
	AdminSlice []*Admin
	// AdminHook is the signature for custom Admin hook methods
	AdminHook func(boil.Executor, *Admin) error

	adminQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	adminType                 = reflect.TypeOf(&Admin{})
	adminMapping              = queries.MakeStructMapping(adminType)
	adminPrimaryKeyMapping, _ = queries.BindMapping(adminType, adminMapping, adminPrimaryKeyColumns)
	adminInsertCacheMut       sync.RWMutex
	adminInsertCache          = make(map[string]insertCache)
	adminUpdateCacheMut       sync.RWMutex
	adminUpdateCache          = make(map[string]updateCache)
	adminUpsertCacheMut       sync.RWMutex
	adminUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var adminBeforeInsertHooks []AdminHook
var adminBeforeUpdateHooks []AdminHook
var adminBeforeDeleteHooks []AdminHook
var adminBeforeUpsertHooks []AdminHook

var adminAfterInsertHooks []AdminHook
var adminAfterSelectHooks []AdminHook
var adminAfterUpdateHooks []AdminHook
var adminAfterDeleteHooks []AdminHook
var adminAfterUpsertHooks []AdminHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Admin) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Admin) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Admin) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Admin) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Admin) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Admin) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range adminAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Admin) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range adminAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Admin) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range adminAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Admin) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range adminAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAdminHook registers your hook function for all future operations.
func AddAdminHook(hookPoint boil.HookPoint, adminHook AdminHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		adminBeforeInsertHooks = append(adminBeforeInsertHooks, adminHook)
	case boil.BeforeUpdateHook:
		adminBeforeUpdateHooks = append(adminBeforeUpdateHooks, adminHook)
	case boil.BeforeDeleteHook:
		adminBeforeDeleteHooks = append(adminBeforeDeleteHooks, adminHook)
	case boil.BeforeUpsertHook:
		adminBeforeUpsertHooks = append(adminBeforeUpsertHooks, adminHook)
	case boil.AfterInsertHook:
		adminAfterInsertHooks = append(adminAfterInsertHooks, adminHook)
	case boil.AfterSelectHook:
		adminAfterSelectHooks = append(adminAfterSelectHooks, adminHook)
	case boil.AfterUpdateHook:
		adminAfterUpdateHooks = append(adminAfterUpdateHooks, adminHook)
	case boil.AfterDeleteHook:
		adminAfterDeleteHooks = append(adminAfterDeleteHooks, adminHook)
	case boil.AfterUpsertHook:
		adminAfterUpsertHooks = append(adminAfterUpsertHooks, adminHook)
	}
}

// OneP returns a single admin record from the query, and panics on error.
func (q adminQuery) OneP() *Admin {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single admin record from the query.
func (q adminQuery) One() (*Admin, error) {
	o := &Admin{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for admins")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Admin records from the query, and panics on error.
func (q adminQuery) AllP() AdminSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Admin records from the query.
func (q adminQuery) All() (AdminSlice, error) {
	var o []*Admin

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Admin slice")
	}

	if len(adminAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Admin records in the query, and panics on error.
func (q adminQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Admin records in the query.
func (q adminQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count admins rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q adminQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q adminQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if admins exists")
	}

	return count > 0, nil
}

// AdminsG retrieves all records.
func AdminsG(mods ...qm.QueryMod) adminQuery {
	return Admins(boil.GetDB(), mods...)
}

// Admins retrieves all the records using an executor.
func Admins(exec boil.Executor, mods ...qm.QueryMod) adminQuery {
	mods = append(mods, qm.From("\"back2school\".\"admins\""))
	return adminQuery{NewQuery(exec, mods...)}
}

// FindAdminG retrieves a single record by ID.
func FindAdminG(id int, selectCols ...string) (*Admin, error) {
	return FindAdmin(boil.GetDB(), id, selectCols...)
}

// FindAdminGP retrieves a single record by ID, and panics on error.
func FindAdminGP(id int, selectCols ...string) *Admin {
	retobj, err := FindAdmin(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindAdmin retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAdmin(exec boil.Executor, id int, selectCols ...string) (*Admin, error) {
	adminObj := &Admin{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"back2school\".\"admins\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(adminObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from admins")
	}

	return adminObj, nil
}

// FindAdminP retrieves a single record by ID with an executor, and panics on error.
func FindAdminP(exec boil.Executor, id int, selectCols ...string) *Admin {
	retobj, err := FindAdmin(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Admin) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Admin) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Admin) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Admin) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no admins provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	adminInsertCacheMut.RLock()
	cache, cached := adminInsertCache[key]
	adminInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			adminColumns,
			adminColumnsWithDefault,
			adminColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(adminType, adminMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(adminType, adminMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"back2school\".\"admins\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"back2school\".\"admins\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into admins")
	}

	if !cached {
		adminInsertCacheMut.Lock()
		adminInsertCache[key] = cache
		adminInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Admin record. See Update for
// whitelist behavior description.
func (o *Admin) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Admin record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Admin) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Admin, and panics on error.
// See Update for whitelist behavior description.
func (o *Admin) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Admin.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Admin) Update(exec boil.Executor, whitelist ...string) error {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	adminUpdateCacheMut.RLock()
	cache, cached := adminUpdateCache[key]
	adminUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			adminColumns,
			adminPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update admins, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"back2school\".\"admins\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, adminPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(adminType, adminMapping, append(wl, adminPrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update admins row")
	}

	if !cached {
		adminUpdateCacheMut.Lock()
		adminUpdateCache[key] = cache
		adminUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q adminQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q adminQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for admins")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o AdminSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o AdminSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o AdminSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AdminSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"back2school\".\"admins\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, adminPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in admin slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Admin) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Admin) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Admin) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Admin) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no admins provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(adminColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()

	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	adminUpsertCacheMut.RLock()
	cache, cached := adminUpsertCache[key]
	adminUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			adminColumns,
			adminColumnsWithDefault,
			adminColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			adminColumns,
			adminPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert admins, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(adminPrimaryKeyColumns))
			copy(conflict, adminPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"back2school\".\"admins\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(adminType, adminMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(adminType, adminMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert admins")
	}

	if !cached {
		adminUpsertCacheMut.Lock()
		adminUpsertCache[key] = cache
		adminUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Admin record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Admin) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Admin record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Admin) DeleteG() error {
	if o == nil {
		return errors.New("models: no Admin provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Admin record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Admin) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Admin record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Admin) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Admin provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), adminPrimaryKeyMapping)
	sql := "DELETE FROM \"back2school\".\"admins\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from admins")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q adminQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q adminQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no adminQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from admins")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o AdminSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o AdminSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Admin slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o AdminSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AdminSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Admin slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(adminBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"back2school\".\"admins\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from admin slice")
	}

	if len(adminAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Admin) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Admin) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Admin) ReloadG() error {
	if o == nil {
		return errors.New("models: no Admin provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Admin) Reload(exec boil.Executor) error {
	ret, err := FindAdmin(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AdminSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *AdminSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty AdminSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AdminSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	admins := AdminSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), adminPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"back2school\".\"admins\".* FROM \"back2school\".\"admins\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, adminPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&admins)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AdminSlice")
	}

	*o = admins

	return nil
}

// AdminExists checks if the Admin row exists.
func AdminExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"back2school\".\"admins\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if admins exists")
	}

	return exists, nil
}

// AdminExistsG checks if the Admin row exists.
func AdminExistsG(id int) (bool, error) {
	return AdminExists(boil.GetDB(), id)
}

// AdminExistsGP checks if the Admin row exists. Panics on error.
func AdminExistsGP(id int) bool {
	e, err := AdminExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// AdminExistsP checks if the Admin row exists. Panics on error.
func AdminExistsP(exec boil.Executor, id int) bool {
	e, err := AdminExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
