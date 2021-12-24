// Code generated by entc, DO NOT EDIT.

package asset

import (
	"entgo.io/ent/dialect/sql"
	"github.com/dopedao/dope-monorepo/packages/api/ent/predicate"
	"github.com/dopedao/dope-monorepo/packages/api/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Symbol applies equality check predicate on the "symbol" field. It's identical to SymbolEQ.
func Symbol(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AssetID applies equality check predicate on the "asset_id" field. It's identical to AssetIDEQ.
func AssetID(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetID), v))
	})
}

// Decimals applies equality check predicate on the "decimals" field. It's identical to DecimalsEQ.
func Decimals(v int) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDecimals), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v Type) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v Type) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...Type) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...Type) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// SymbolEQ applies the EQ predicate on the "symbol" field.
func SymbolEQ(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymbol), v))
	})
}

// SymbolNEQ applies the NEQ predicate on the "symbol" field.
func SymbolNEQ(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymbol), v))
	})
}

// SymbolIn applies the In predicate on the "symbol" field.
func SymbolIn(vs ...string) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSymbol), v...))
	})
}

// SymbolNotIn applies the NotIn predicate on the "symbol" field.
func SymbolNotIn(vs ...string) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSymbol), v...))
	})
}

// SymbolGT applies the GT predicate on the "symbol" field.
func SymbolGT(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymbol), v))
	})
}

// SymbolGTE applies the GTE predicate on the "symbol" field.
func SymbolGTE(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymbol), v))
	})
}

// SymbolLT applies the LT predicate on the "symbol" field.
func SymbolLT(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymbol), v))
	})
}

// SymbolLTE applies the LTE predicate on the "symbol" field.
func SymbolLTE(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymbol), v))
	})
}

// SymbolContains applies the Contains predicate on the "symbol" field.
func SymbolContains(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymbol), v))
	})
}

// SymbolHasPrefix applies the HasPrefix predicate on the "symbol" field.
func SymbolHasPrefix(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymbol), v))
	})
}

// SymbolHasSuffix applies the HasSuffix predicate on the "symbol" field.
func SymbolHasSuffix(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymbol), v))
	})
}

// SymbolEqualFold applies the EqualFold predicate on the "symbol" field.
func SymbolEqualFold(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymbol), v))
	})
}

// SymbolContainsFold applies the ContainsFold predicate on the "symbol" field.
func SymbolContainsFold(v string) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymbol), v))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...schema.BigInt) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...schema.BigInt) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AssetIDEQ applies the EQ predicate on the "asset_id" field.
func AssetIDEQ(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAssetID), v))
	})
}

// AssetIDNEQ applies the NEQ predicate on the "asset_id" field.
func AssetIDNEQ(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAssetID), v))
	})
}

// AssetIDIn applies the In predicate on the "asset_id" field.
func AssetIDIn(vs ...schema.BigInt) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAssetID), v...))
	})
}

// AssetIDNotIn applies the NotIn predicate on the "asset_id" field.
func AssetIDNotIn(vs ...schema.BigInt) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAssetID), v...))
	})
}

// AssetIDGT applies the GT predicate on the "asset_id" field.
func AssetIDGT(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAssetID), v))
	})
}

// AssetIDGTE applies the GTE predicate on the "asset_id" field.
func AssetIDGTE(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAssetID), v))
	})
}

// AssetIDLT applies the LT predicate on the "asset_id" field.
func AssetIDLT(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAssetID), v))
	})
}

// AssetIDLTE applies the LTE predicate on the "asset_id" field.
func AssetIDLTE(v schema.BigInt) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAssetID), v))
	})
}

// DecimalsEQ applies the EQ predicate on the "decimals" field.
func DecimalsEQ(v int) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDecimals), v))
	})
}

// DecimalsNEQ applies the NEQ predicate on the "decimals" field.
func DecimalsNEQ(v int) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDecimals), v))
	})
}

// DecimalsIn applies the In predicate on the "decimals" field.
func DecimalsIn(vs ...int) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDecimals), v...))
	})
}

// DecimalsNotIn applies the NotIn predicate on the "decimals" field.
func DecimalsNotIn(vs ...int) predicate.Asset {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Asset(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDecimals), v...))
	})
}

// DecimalsGT applies the GT predicate on the "decimals" field.
func DecimalsGT(v int) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDecimals), v))
	})
}

// DecimalsGTE applies the GTE predicate on the "decimals" field.
func DecimalsGTE(v int) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDecimals), v))
	})
}

// DecimalsLT applies the LT predicate on the "decimals" field.
func DecimalsLT(v int) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDecimals), v))
	})
}

// DecimalsLTE applies the LTE predicate on the "decimals" field.
func DecimalsLTE(v int) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDecimals), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
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
func Not(p predicate.Asset) predicate.Asset {
	return predicate.Asset(func(s *sql.Selector) {
		p(s.Not())
	})
}
