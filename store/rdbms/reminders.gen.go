package rdbms

// This file is an auto-generated file
//
// Template:    pkg/codegen/assets/store_rdbms.gen.go.tpl
// Definitions: store/reminders.yaml
//
// Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
)

var _ = errors.Is

// SearchReminders returns all matching rows
//
// This function calls convertReminderFilter with the given
// types.ReminderFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchReminders(ctx context.Context, f types.ReminderFilter) (types.ReminderSet, types.ReminderFilter, error) {
	var (
		err error
		set []*types.Reminder
		q   squirrel.SelectBuilder
	)
	q, err = s.convertReminderFilter(f)
	if err != nil {
		return nil, f, err
	}

	// Cleanup anything we've accidentally received...
	f.PrevPage, f.NextPage = nil, nil

	// When cursor for a previous page is used it's marked as reversed
	// This tells us to flip the descending flag on all used sort keys
	reversedCursor := f.PageCursor != nil && f.PageCursor.Reverse

	// If paging with reverse cursor, change the sorting
	// direction for all columns we're sorting by
	curSort := f.Sort.Clone()
	if reversedCursor {
		curSort.Reverse()
	}

	return set, f, s.config.ErrorHandler(func() error {
		set, err = s.fetchFullPageOfReminders(ctx, q, curSort, f.PageCursor, f.Limit, f.Check)

		if err != nil {
			return err
		}

		if f.Limit > 0 && len(set) > 0 {
			if f.PageCursor != nil && (!f.PageCursor.Reverse || uint(len(set)) == f.Limit) {
				f.PrevPage = s.collectReminderCursorValues(set[0], curSort.Columns()...)
				f.PrevPage.Reverse = true
			}

			// Less items fetched then requested by page-limit
			// not very likely there's another page
			f.NextPage = s.collectReminderCursorValues(set[len(set)-1], curSort.Columns()...)
		}

		f.PageCursor = nil
		return nil
	}())
}

// fetchFullPageOfReminders collects all requested results.
//
// Function applies:
//  - cursor conditions (where ...)
//  - sorting rules (order by ...)
//  - limit
//
// Main responsibility of this function is to perform additional sequential queries in case when not enough results
// are collected due to failed check on a specific row (by check fn). Function then moves cursor to the last item fetched
func (s Store) fetchFullPageOfReminders(
	ctx context.Context,
	q squirrel.SelectBuilder,
	sort filter.SortExprSet,
	cursor *filter.PagingCursor,
	limit uint,
	check func(*types.Reminder) (bool, error),
) ([]*types.Reminder, error) {
	var (
		set  = make([]*types.Reminder, 0, DefaultSliceCapacity)
		aux  []*types.Reminder
		last *types.Reminder

		// When cursor for a previous page is used it's marked as reversed
		// This tells us to flip the descending flag on all used sort keys
		reversedCursor = cursor != nil && cursor.Reverse

		// copy of the select builder
		tryQuery squirrel.SelectBuilder

		fetched uint
		err     error
	)

	// Make sure we always end our sort by primary keys
	if sort.Get("id") == nil {
		sort = append(sort, &filter.SortExpr{Column: "id"})
	}

	// Apply sorting expr from filter to query
	if q, err = setOrderBy(q, sort, s.sortableReminderColumns()...); err != nil {
		return nil, err
	}

	for try := 0; try < MaxRefetches; try++ {
		tryQuery = setCursorCond(q, cursor)
		if limit > 0 {
			tryQuery = tryQuery.Limit(uint64(limit))
		}

		if aux, fetched, last, err = s.QueryReminders(ctx, tryQuery, check); err != nil {
			return nil, err
		}

		if limit > 0 && uint(len(aux)) >= limit {
			// we should use only as much as requested
			set = append(set, aux[0:limit]...)
			break
		} else {
			set = append(set, aux...)
		}

		// if limit is not set or we've already collected enough items
		// we can break the loop right away
		if limit == 0 || fetched == 0 || fetched < limit {
			break
		}

		// In case limit is set very low and we've missed records in the first fetch,
		// make sure next fetch limit is a bit higher
		if limit < MinEnsureFetchLimit {
			limit = MinEnsureFetchLimit
		}

		// @todo improve strategy for collecting next page with lower limit

		// Point cursor to the last fetched element
		if cursor = s.collectReminderCursorValues(last, sort.Columns()...); cursor == nil {
			break
		}
	}

	if reversedCursor {
		// Cursor for previous page was used
		// Fetched set needs to be reverseCursor because we've forced a descending order to
		// get the previous page
		for i, j := 0, len(set)-1; i < j; i, j = i+1, j-1 {
			set[i], set[j] = set[j], set[i]
		}
	}

	return set, nil
}

// QueryReminders queries the database, converts and checks each row and
// returns collected set
//
// Fn also returns total number of fetched items and last fetched item so that the caller can construct cursor
// for next page of results
func (s Store) QueryReminders(
	ctx context.Context,
	q squirrel.Sqlizer,
	check func(*types.Reminder) (bool, error),
) ([]*types.Reminder, uint, *types.Reminder, error) {
	var (
		set = make([]*types.Reminder, 0, DefaultSliceCapacity)
		res *types.Reminder

		// Query rows with
		rows, err = s.Query(ctx, q)

		fetched uint
	)

	if err != nil {
		return nil, 0, nil, err
	}

	defer rows.Close()
	for rows.Next() {
		fetched++
		if err = rows.Err(); err == nil {
			res, err = s.internalReminderRowScanner(rows)
		}

		if err != nil {
			return nil, 0, nil, err
		}

		// If check function is set, call it and act accordingly
		if check != nil {
			if chk, err := check(res); err != nil {
				return nil, 0, nil, err
			} else if !chk {
				// did not pass the check
				// go with the next row
				continue
			}
		}

		set = append(set, res)
	}

	return set, fetched, res, rows.Err()
}

// LookupReminderByID searches for reminder by its ID
//
// It returns reminder even if deleted or suspended
func (s Store) LookupReminderByID(ctx context.Context, id uint64) (*types.Reminder, error) {
	return s.execLookupReminder(ctx, squirrel.Eq{
		s.preprocessColumn("rmd.id", ""): s.preprocessValue(id, ""),
	})
}

// CreateReminder creates one or more rows in reminders table
func (s Store) CreateReminder(ctx context.Context, rr ...*types.Reminder) (err error) {
	for _, res := range rr {
		err = s.checkReminderConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execCreateReminders(ctx, s.internalReminderEncoder(res))
		if err != nil {
			return err
		}
	}

	return
}

// UpdateReminder updates one or more existing rows in reminders
func (s Store) UpdateReminder(ctx context.Context, rr ...*types.Reminder) error {
	return s.config.ErrorHandler(s.PartialReminderUpdate(ctx, nil, rr...))
}

// PartialReminderUpdate updates one or more existing rows in reminders
func (s Store) PartialReminderUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Reminder) (err error) {
	for _, res := range rr {
		err = s.checkReminderConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execUpdateReminders(
			ctx,
			squirrel.Eq{
				s.preprocessColumn("rmd.id", ""): s.preprocessValue(res.ID, ""),
			},
			s.internalReminderEncoder(res).Skip("id").Only(onlyColumns...))
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return
}

// UpsertReminder updates one or more existing rows in reminders
func (s Store) UpsertReminder(ctx context.Context, rr ...*types.Reminder) (err error) {
	for _, res := range rr {
		err = s.checkReminderConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.config.ErrorHandler(s.execUpsertReminders(ctx, s.internalReminderEncoder(res)))
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteReminder Deletes one or more rows from reminders table
func (s Store) DeleteReminder(ctx context.Context, rr ...*types.Reminder) (err error) {
	for _, res := range rr {

		err = s.execDeleteReminders(ctx, squirrel.Eq{
			s.preprocessColumn("rmd.id", ""): s.preprocessValue(res.ID, ""),
		})
		if err != nil {
			return s.config.ErrorHandler(err)
		}
	}

	return nil
}

// DeleteReminderByID Deletes row from the reminders table
func (s Store) DeleteReminderByID(ctx context.Context, ID uint64) error {
	return s.execDeleteReminders(ctx, squirrel.Eq{
		s.preprocessColumn("rmd.id", ""): s.preprocessValue(ID, ""),
	})
}

// TruncateReminders Deletes all rows from the reminders table
func (s Store) TruncateReminders(ctx context.Context) error {
	return s.config.ErrorHandler(s.Truncate(ctx, s.reminderTable()))
}

// execLookupReminder prepares Reminder query and executes it,
// returning types.Reminder (or error)
func (s Store) execLookupReminder(ctx context.Context, cnd squirrel.Sqlizer) (res *types.Reminder, err error) {
	var (
		row rowScanner
	)

	row, err = s.QueryRow(ctx, s.remindersSelectBuilder().Where(cnd))
	if err != nil {
		return
	}

	res, err = s.internalReminderRowScanner(row)
	if err != nil {
		return
	}

	return res, nil
}

// execCreateReminders updates all matched (by cnd) rows in reminders with given data
func (s Store) execCreateReminders(ctx context.Context, payload store.Payload) error {
	return s.config.ErrorHandler(s.Exec(ctx, s.InsertBuilder(s.reminderTable()).SetMap(payload)))
}

// execUpdateReminders updates all matched (by cnd) rows in reminders with given data
func (s Store) execUpdateReminders(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return s.config.ErrorHandler(s.Exec(ctx, s.UpdateBuilder(s.reminderTable("rmd")).Where(cnd).SetMap(set)))
}

// execUpsertReminders inserts new or updates matching (by-primary-key) rows in reminders with given data
func (s Store) execUpsertReminders(ctx context.Context, set store.Payload) error {
	upsert, err := s.config.UpsertBuilder(
		s.config,
		s.reminderTable(),
		set,
		"id",
	)

	if err != nil {
		return err
	}

	return s.config.ErrorHandler(s.Exec(ctx, upsert))
}

// execDeleteReminders Deletes all matched (by cnd) rows in reminders with given data
func (s Store) execDeleteReminders(ctx context.Context, cnd squirrel.Sqlizer) error {
	return s.config.ErrorHandler(s.Exec(ctx, s.DeleteBuilder(s.reminderTable("rmd")).Where(cnd)))
}

func (s Store) internalReminderRowScanner(row rowScanner) (res *types.Reminder, err error) {
	res = &types.Reminder{}

	if _, has := s.config.RowScanners["reminder"]; has {
		scanner := s.config.RowScanners["reminder"].(func(_ rowScanner, _ *types.Reminder) error)
		err = scanner(row, res)
	} else {
		err = row.Scan(
			&res.ID,
			&res.Resource,
			&res.Payload,
			&res.SnoozeCount,
			&res.AssignedTo,
			&res.AssignedBy,
			&res.AssignedAt,
			&res.DismissedBy,
			&res.DismissedAt,
			&res.RemindAt,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)
	}

	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("could not scan db row for Reminder: %w", err)
	} else {
		return res, nil
	}
}

// QueryReminders returns squirrel.SelectBuilder with set table and all columns
func (s Store) remindersSelectBuilder() squirrel.SelectBuilder {
	return s.SelectBuilder(s.reminderTable("rmd"), s.reminderColumns("rmd")...)
}

// reminderTable name of the db table
func (Store) reminderTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "reminders" + alias
}

// ReminderColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) reminderColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "resource",
		alias + "payload",
		alias + "snooze_count",
		alias + "assigned_to",
		alias + "assigned_by",
		alias + "assigned_at",
		alias + "dismissed_by",
		alias + "dismissed_at",
		alias + "remind_at",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// {true true true true true}

// sortableReminderColumns returns all Reminder columns flagged as sortable
//
// With optional string arg, all columns are returned aliased
func (Store) sortableReminderColumns() []string {
	return []string{
		"id",
		"remind_at",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

// internalReminderEncoder encodes fields from types.Reminder to store.Payload (map)
//
// Encoding is done by using generic approach or by calling encodeReminder
// func when rdbms.customEncoder=true
func (s Store) internalReminderEncoder(res *types.Reminder) store.Payload {
	return store.Payload{
		"id":           res.ID,
		"resource":     res.Resource,
		"payload":      res.Payload,
		"snooze_count": res.SnoozeCount,
		"assigned_to":  res.AssignedTo,
		"assigned_by":  res.AssignedBy,
		"assigned_at":  res.AssignedAt,
		"dismissed_by": res.DismissedBy,
		"dismissed_at": res.DismissedAt,
		"remind_at":    res.RemindAt,
		"created_at":   res.CreatedAt,
		"updated_at":   res.UpdatedAt,
		"deleted_at":   res.DeletedAt,
	}
}

// collectReminderCursorValues collects values from the given resource that and sets them to the cursor
// to be used for pagination
//
// Values that are collected must come from sortable, unique or primary columns/fields
// At least one of the collected columns must be flagged as unique, otherwise fn appends primary keys at the end
//
// Known issue:
//   when collecting cursor values for query that sorts by unique column with partial index (ie: unique handle on
//   undeleted items)
func (s Store) collectReminderCursorValues(res *types.Reminder, cc ...string) *filter.PagingCursor {
	var (
		cursor = &filter.PagingCursor{}

		hasUnique bool

		// All known primary key columns

		pkId bool

		collect = func(cc ...string) {
			for _, c := range cc {
				switch c {
				case "id":
					cursor.Set(c, res.ID, false)

					pkId = true
				case "remind_at":
					cursor.Set(c, res.RemindAt, false)

				case "created_at":
					cursor.Set(c, res.CreatedAt, false)

				case "updated_at":
					cursor.Set(c, res.UpdatedAt, false)

				case "deleted_at":
					cursor.Set(c, res.DeletedAt, false)

				}
			}
		}
	)

	collect(cc...)
	if !hasUnique || !(pkId && true) {
		collect("id")
	}

	return cursor
}

func (s *Store) checkReminderConstraints(ctx context.Context, res *types.Reminder) error {

	return nil
}