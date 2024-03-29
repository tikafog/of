// Code generated by ent, DO NOT EDIT.

package toplist

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/tikafog/of/dbc/media/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedUnix applies equality check predicate on the "created_unix" field. It's identical to CreatedUnixEQ.
func CreatedUnix(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedUnix), v))
	})
}

// UpdatedUnix applies equality check predicate on the "updated_unix" field. It's identical to UpdatedUnixEQ.
func UpdatedUnix(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedUnix), v))
	})
}

// DeletedUnix applies equality check predicate on the "deleted_unix" field. It's identical to DeletedUnixEQ.
func DeletedUnix(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedUnix), v))
	})
}

// InformationID applies equality check predicate on the "information_id" field. It's identical to InformationIDEQ.
func InformationID(v uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInformationID), v))
	})
}

// PageID applies equality check predicate on the "page_id" field. It's identical to PageIDEQ.
func PageID(v uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPageID), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Intro applies equality check predicate on the "intro" field. It's identical to IntroEQ.
func Intro(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntro), v))
	})
}

// CreatedUnixEQ applies the EQ predicate on the "created_unix" field.
func CreatedUnixEQ(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixNEQ applies the NEQ predicate on the "created_unix" field.
func CreatedUnixNEQ(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixIn applies the In predicate on the "created_unix" field.
func CreatedUnixIn(vs ...int64) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func CreatedUnixNotIn(vs ...int64) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func CreatedUnixGT(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixGTE applies the GTE predicate on the "created_unix" field.
func CreatedUnixGTE(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixLT applies the LT predicate on the "created_unix" field.
func CreatedUnixLT(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedUnix), v))
	})
}

// CreatedUnixLTE applies the LTE predicate on the "created_unix" field.
func CreatedUnixLTE(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedUnix), v))
	})
}

// UpdatedUnixEQ applies the EQ predicate on the "updated_unix" field.
func UpdatedUnixEQ(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixNEQ applies the NEQ predicate on the "updated_unix" field.
func UpdatedUnixNEQ(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixIn applies the In predicate on the "updated_unix" field.
func UpdatedUnixIn(vs ...int64) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func UpdatedUnixNotIn(vs ...int64) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func UpdatedUnixGT(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixGTE applies the GTE predicate on the "updated_unix" field.
func UpdatedUnixGTE(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixLT applies the LT predicate on the "updated_unix" field.
func UpdatedUnixLT(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedUnix), v))
	})
}

// UpdatedUnixLTE applies the LTE predicate on the "updated_unix" field.
func UpdatedUnixLTE(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedUnix), v))
	})
}

// DeletedUnixEQ applies the EQ predicate on the "deleted_unix" field.
func DeletedUnixEQ(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixNEQ applies the NEQ predicate on the "deleted_unix" field.
func DeletedUnixNEQ(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixIn applies the In predicate on the "deleted_unix" field.
func DeletedUnixIn(vs ...int64) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func DeletedUnixNotIn(vs ...int64) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func DeletedUnixGT(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixGTE applies the GTE predicate on the "deleted_unix" field.
func DeletedUnixGTE(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixLT applies the LT predicate on the "deleted_unix" field.
func DeletedUnixLT(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixLTE applies the LTE predicate on the "deleted_unix" field.
func DeletedUnixLTE(v int64) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedUnix), v))
	})
}

// DeletedUnixIsNil applies the IsNil predicate on the "deleted_unix" field.
func DeletedUnixIsNil() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDeletedUnix)))
	})
}

// DeletedUnixNotNil applies the NotNil predicate on the "deleted_unix" field.
func DeletedUnixNotNil() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDeletedUnix)))
	})
}

// InformationIDEQ applies the EQ predicate on the "information_id" field.
func InformationIDEQ(v uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInformationID), v))
	})
}

// InformationIDNEQ applies the NEQ predicate on the "information_id" field.
func InformationIDNEQ(v uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInformationID), v))
	})
}

// InformationIDIn applies the In predicate on the "information_id" field.
func InformationIDIn(vs ...uuid.UUID) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInformationID), v...))
	})
}

// InformationIDNotIn applies the NotIn predicate on the "information_id" field.
func InformationIDNotIn(vs ...uuid.UUID) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInformationID), v...))
	})
}

// InformationIDIsNil applies the IsNil predicate on the "information_id" field.
func InformationIDIsNil() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInformationID)))
	})
}

// InformationIDNotNil applies the NotNil predicate on the "information_id" field.
func InformationIDNotNil() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInformationID)))
	})
}

// PageIDEQ applies the EQ predicate on the "page_id" field.
func PageIDEQ(v uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPageID), v))
	})
}

// PageIDNEQ applies the NEQ predicate on the "page_id" field.
func PageIDNEQ(v uuid.UUID) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPageID), v))
	})
}

// PageIDIn applies the In predicate on the "page_id" field.
func PageIDIn(vs ...uuid.UUID) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPageID), v...))
	})
}

// PageIDNotIn applies the NotIn predicate on the "page_id" field.
func PageIDNotIn(vs ...uuid.UUID) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPageID), v...))
	})
}

// PageIDIsNil applies the IsNil predicate on the "page_id" field.
func PageIDIsNil() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPageID)))
	})
}

// PageIDNotNil applies the NotNil predicate on the "page_id" field.
func PageIDNotNil() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPageID)))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func TitleNotIn(vs ...string) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
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
func TitleGT(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// IntroEQ applies the EQ predicate on the "intro" field.
func IntroEQ(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldIntro), v))
	})
}

// IntroNEQ applies the NEQ predicate on the "intro" field.
func IntroNEQ(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldIntro), v))
	})
}

// IntroIn applies the In predicate on the "intro" field.
func IntroIn(vs ...string) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldIntro), v...))
	})
}

// IntroNotIn applies the NotIn predicate on the "intro" field.
func IntroNotIn(vs ...string) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldIntro), v...))
	})
}

// IntroGT applies the GT predicate on the "intro" field.
func IntroGT(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldIntro), v))
	})
}

// IntroGTE applies the GTE predicate on the "intro" field.
func IntroGTE(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldIntro), v))
	})
}

// IntroLT applies the LT predicate on the "intro" field.
func IntroLT(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldIntro), v))
	})
}

// IntroLTE applies the LTE predicate on the "intro" field.
func IntroLTE(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldIntro), v))
	})
}

// IntroContains applies the Contains predicate on the "intro" field.
func IntroContains(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldIntro), v))
	})
}

// IntroHasPrefix applies the HasPrefix predicate on the "intro" field.
func IntroHasPrefix(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldIntro), v))
	})
}

// IntroHasSuffix applies the HasSuffix predicate on the "intro" field.
func IntroHasSuffix(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldIntro), v))
	})
}

// IntroEqualFold applies the EqualFold predicate on the "intro" field.
func IntroEqualFold(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldIntro), v))
	})
}

// IntroContainsFold applies the ContainsFold predicate on the "intro" field.
func IntroContainsFold(v string) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldIntro), v))
	})
}

// CategoryEQ applies the EQ predicate on the "category" field.
func CategoryEQ(v Category) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCategory), v))
	})
}

// CategoryNEQ applies the NEQ predicate on the "category" field.
func CategoryNEQ(v Category) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCategory), v))
	})
}

// CategoryIn applies the In predicate on the "category" field.
func CategoryIn(vs ...Category) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCategory), v...))
	})
}

// CategoryNotIn applies the NotIn predicate on the "category" field.
func CategoryNotIn(vs ...Category) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCategory), v...))
	})
}

// LowerBannerEQ applies the EQ predicate on the "lower_banner" field.
func LowerBannerEQ(v LowerBanner) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLowerBanner), v))
	})
}

// LowerBannerNEQ applies the NEQ predicate on the "lower_banner" field.
func LowerBannerNEQ(v LowerBanner) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLowerBanner), v))
	})
}

// LowerBannerIn applies the In predicate on the "lower_banner" field.
func LowerBannerIn(vs ...LowerBanner) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLowerBanner), v...))
	})
}

// LowerBannerNotIn applies the NotIn predicate on the "lower_banner" field.
func LowerBannerNotIn(vs ...LowerBanner) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLowerBanner), v...))
	})
}

// TopRightEQ applies the EQ predicate on the "top_right" field.
func TopRightEQ(v TopRight) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTopRight), v))
	})
}

// TopRightNEQ applies the NEQ predicate on the "top_right" field.
func TopRightNEQ(v TopRight) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTopRight), v))
	})
}

// TopRightIn applies the In predicate on the "top_right" field.
func TopRightIn(vs ...TopRight) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTopRight), v...))
	})
}

// TopRightNotIn applies the NotIn predicate on the "top_right" field.
func TopRightNotIn(vs ...TopRight) predicate.TopList {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.TopList(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTopRight), v...))
	})
}

// HasInformation applies the HasEdge predicate on the "information" edge.
func HasInformation() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InformationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InformationTable, InformationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInformationWith applies the HasEdge predicate on the "information" edge with a given conditions (other predicates).
func HasInformationWith(preds ...predicate.InformationV1) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InformationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InformationTable, InformationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPage applies the HasEdge predicate on the "page" edge.
func HasPage() predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PageTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PageTable, PageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPageWith applies the HasEdge predicate on the "page" edge with a given conditions (other predicates).
func HasPageWith(preds ...predicate.Page) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PageInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PageTable, PageColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TopList) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TopList) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
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
func Not(p predicate.TopList) predicate.TopList {
	return predicate.TopList(func(s *sql.Selector) {
		p(s.Not())
	})
}
