// Code generated by ent, DO NOT EDIT.

package version

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tikafog/of/dbc/kernel/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Current applies equality check predicate on the "Current" field. It's identical to CurrentEQ.
func Current(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrent), v))
	})
}

// Last applies equality check predicate on the "Last" field. It's identical to LastEQ.
func Last(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLast), v))
	})
}

// CurrentEQ applies the EQ predicate on the "Current" field.
func CurrentEQ(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCurrent), v))
	})
}

// CurrentNEQ applies the NEQ predicate on the "Current" field.
func CurrentNEQ(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCurrent), v))
	})
}

// CurrentIn applies the In predicate on the "Current" field.
func CurrentIn(vs ...int) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCurrent), v...))
	})
}

// CurrentNotIn applies the NotIn predicate on the "Current" field.
func CurrentNotIn(vs ...int) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCurrent), v...))
	})
}

// CurrentGT applies the GT predicate on the "Current" field.
func CurrentGT(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCurrent), v))
	})
}

// CurrentGTE applies the GTE predicate on the "Current" field.
func CurrentGTE(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCurrent), v))
	})
}

// CurrentLT applies the LT predicate on the "Current" field.
func CurrentLT(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCurrent), v))
	})
}

// CurrentLTE applies the LTE predicate on the "Current" field.
func CurrentLTE(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCurrent), v))
	})
}

// LastEQ applies the EQ predicate on the "Last" field.
func LastEQ(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLast), v))
	})
}

// LastNEQ applies the NEQ predicate on the "Last" field.
func LastNEQ(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLast), v))
	})
}

// LastIn applies the In predicate on the "Last" field.
func LastIn(vs ...int) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLast), v...))
	})
}

// LastNotIn applies the NotIn predicate on the "Last" field.
func LastNotIn(vs ...int) predicate.Version {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Version(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLast), v...))
	})
}

// LastGT applies the GT predicate on the "Last" field.
func LastGT(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLast), v))
	})
}

// LastGTE applies the GTE predicate on the "Last" field.
func LastGTE(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLast), v))
	})
}

// LastLT applies the LT predicate on the "Last" field.
func LastLT(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLast), v))
	})
}

// LastLTE applies the LTE predicate on the "Last" field.
func LastLTE(v int) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLast), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Version) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Version) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
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
func Not(p predicate.Version) predicate.Version {
	return predicate.Version(func(s *sql.Selector) {
		p(s.Not())
	})
}
