// Code generated by entc, DO NOT EDIT.

package discovery

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedUnix applies equality check predicate on the "created_unix" field. It's identical to CreatedUnixEQ.
func CreatedUnix(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedUnix), v))
	})
}

// UpdatedUnix applies equality check predicate on the "updated_unix" field. It's identical to UpdatedUnixEQ.
func UpdatedUnix(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedUnix), v))
	})
}

// DeletedUnix applies equality check predicate on the "deleted_unix" field. It's identical to DeletedUnixEQ.
func DeletedUnix(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedUnix), v))
	})
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// Rid applies equality check predicate on the "rid" field. It's identical to RidEQ.
func Rid(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRid), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Detail applies equality check predicate on the "detail" field. It's identical to DetailEQ.
func Detail(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// CreatedUnixEQ applies the EQ predicate on the "created_unix" field.
func CreatedUnixEQ(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixNEQ applies the NEQ predicate on the "created_unix" field.
func CreatedUnixNEQ(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixIn applies the In predicate on the "created_unix" field.
func CreatedUnixIn(vs ...int64) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedUnix), v...))
	})
}

// CreatedUnixNotIn applies the NotIn predicate on the "created_unix" field.
func CreatedUnixNotIn(vs ...int64) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedUnix), v...))
	})
}

// CreatedUnixGT applies the GT predicate on the "created_unix" field.
func CreatedUnixGT(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixGTE applies the GTE predicate on the "created_unix" field.
func CreatedUnixGTE(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixLT applies the LT predicate on the "created_unix" field.
func CreatedUnixLT(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixLTE applies the LTE predicate on the "created_unix" field.
func CreatedUnixLTE(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedUnix), v))
	})
}

// UpdatedUnixEQ applies the EQ predicate on the "updated_unix" field.
func UpdatedUnixEQ(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixNEQ applies the NEQ predicate on the "updated_unix" field.
func UpdatedUnixNEQ(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixIn applies the In predicate on the "updated_unix" field.
func UpdatedUnixIn(vs ...int64) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedUnix), v...))
	})
}

// UpdatedUnixNotIn applies the NotIn predicate on the "updated_unix" field.
func UpdatedUnixNotIn(vs ...int64) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedUnix), v...))
	})
}

// UpdatedUnixGT applies the GT predicate on the "updated_unix" field.
func UpdatedUnixGT(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixGTE applies the GTE predicate on the "updated_unix" field.
func UpdatedUnixGTE(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixLT applies the LT predicate on the "updated_unix" field.
func UpdatedUnixLT(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixLTE applies the LTE predicate on the "updated_unix" field.
func UpdatedUnixLTE(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedUnix), v))
	})
}

// DeletedUnixEQ applies the EQ predicate on the "deleted_unix" field.
func DeletedUnixEQ(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixNEQ applies the NEQ predicate on the "deleted_unix" field.
func DeletedUnixNEQ(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixIn applies the In predicate on the "deleted_unix" field.
func DeletedUnixIn(vs ...int64) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedUnix), v...))
	})
}

// DeletedUnixNotIn applies the NotIn predicate on the "deleted_unix" field.
func DeletedUnixNotIn(vs ...int64) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedUnix), v...))
	})
}

// DeletedUnixGT applies the GT predicate on the "deleted_unix" field.
func DeletedUnixGT(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixGTE applies the GTE predicate on the "deleted_unix" field.
func DeletedUnixGTE(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixLT applies the LT predicate on the "deleted_unix" field.
func DeletedUnixLT(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixLTE applies the LTE predicate on the "deleted_unix" field.
func DeletedUnixLTE(v int64) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixIsNil applies the IsNil predicate on the "deleted_unix" field.
func DeletedUnixIsNil() predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedUnix)))
	})
}

// DeletedUnixNotNil applies the NotNil predicate on the "deleted_unix" field.
func DeletedUnixNotNil() predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedUnix)))
	})
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// DateContains applies the Contains predicate on the "date" field.
func DateContains(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDate), v))
	})
}

// DateHasPrefix applies the HasPrefix predicate on the "date" field.
func DateHasPrefix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDate), v))
	})
}

// DateHasSuffix applies the HasSuffix predicate on the "date" field.
func DateHasSuffix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDate), v))
	})
}

// DateEqualFold applies the EqualFold predicate on the "date" field.
func DateEqualFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDate), v))
	})
}

// DateContainsFold applies the ContainsFold predicate on the "date" field.
func DateContainsFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDate), v))
	})
}

// RidEQ applies the EQ predicate on the "rid" field.
func RidEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRid), v))
	})
}

// RidNEQ applies the NEQ predicate on the "rid" field.
func RidNEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRid), v))
	})
}

// RidIn applies the In predicate on the "rid" field.
func RidIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRid), v...))
	})
}

// RidNotIn applies the NotIn predicate on the "rid" field.
func RidNotIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRid), v...))
	})
}

// RidGT applies the GT predicate on the "rid" field.
func RidGT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRid), v))
	})
}

// RidGTE applies the GTE predicate on the "rid" field.
func RidGTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRid), v))
	})
}

// RidLT applies the LT predicate on the "rid" field.
func RidLT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRid), v))
	})
}

// RidLTE applies the LTE predicate on the "rid" field.
func RidLTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRid), v))
	})
}

// RidContains applies the Contains predicate on the "rid" field.
func RidContains(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRid), v))
	})
}

// RidHasPrefix applies the HasPrefix predicate on the "rid" field.
func RidHasPrefix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRid), v))
	})
}

// RidHasSuffix applies the HasSuffix predicate on the "rid" field.
func RidHasSuffix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRid), v))
	})
}

// RidEqualFold applies the EqualFold predicate on the "rid" field.
func RidEqualFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRid), v))
	})
}

// RidContainsFold applies the ContainsFold predicate on the "rid" field.
func RidContainsFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRid), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// DetailEQ applies the EQ predicate on the "detail" field.
func DetailEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetail), v))
	})
}

// DetailNEQ applies the NEQ predicate on the "detail" field.
func DetailNEQ(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetail), v))
	})
}

// DetailIn applies the In predicate on the "detail" field.
func DetailIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetail), v...))
	})
}

// DetailNotIn applies the NotIn predicate on the "detail" field.
func DetailNotIn(vs ...string) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetail), v...))
	})
}

// DetailGT applies the GT predicate on the "detail" field.
func DetailGT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetail), v))
	})
}

// DetailGTE applies the GTE predicate on the "detail" field.
func DetailGTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetail), v))
	})
}

// DetailLT applies the LT predicate on the "detail" field.
func DetailLT(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetail), v))
	})
}

// DetailLTE applies the LTE predicate on the "detail" field.
func DetailLTE(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetail), v))
	})
}

// DetailContains applies the Contains predicate on the "detail" field.
func DetailContains(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetail), v))
	})
}

// DetailHasPrefix applies the HasPrefix predicate on the "detail" field.
func DetailHasPrefix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetail), v))
	})
}

// DetailHasSuffix applies the HasSuffix predicate on the "detail" field.
func DetailHasSuffix(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetail), v))
	})
}

// DetailEqualFold applies the EqualFold predicate on the "detail" field.
func DetailEqualFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetail), v))
	})
}

// DetailContainsFold applies the ContainsFold predicate on the "detail" field.
func DetailContainsFold(v string) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetail), v))
	})
}

// MtypeEQ applies the EQ predicate on the "mtype" field.
func MtypeEQ(v Mtype) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMtype), v))
	})
}

// MtypeNEQ applies the NEQ predicate on the "mtype" field.
func MtypeNEQ(v Mtype) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMtype), v))
	})
}

// MtypeIn applies the In predicate on the "mtype" field.
func MtypeIn(vs ...Mtype) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMtype), v...))
	})
}

// MtypeNotIn applies the NotIn predicate on the "mtype" field.
func MtypeNotIn(vs ...Mtype) predicate.Discovery {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Discovery(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMtype), v...))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Discovery) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Discovery) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Discovery) predicate.Discovery {
	return predicate.Discovery(func(s *sql.Selector) {
		p(s.Not())
	})
}