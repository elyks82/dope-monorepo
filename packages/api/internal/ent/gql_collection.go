// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AmountQuery) CollectFields(ctx context.Context, satisfies ...string) *AmountQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		a = a.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return a
}

func (a *AmountQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *AmountQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "listing_input":
			a = a.WithListingInput(func(query *ListingQuery) {
				query.collectField(ctx, field)
			})
		case "listing_output":
			a = a.WithListingOutput(func(query *ListingQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return a
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (bp *BodyPartQuery) CollectFields(ctx context.Context, satisfies ...string) *BodyPartQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		bp = bp.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return bp
}

func (bp *BodyPartQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *BodyPartQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "hustler_beards":
			bp = bp.WithHustlerBeards(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_bodies":
			bp = bp.WithHustlerBodies(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_hairs":
			bp = bp.WithHustlerHairs(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return bp
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DopeQuery) CollectFields(ctx context.Context, satisfies ...string) *DopeQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		d = d.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return d
}

func (d *DopeQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *DopeQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "index":
			d = d.WithIndex(func(query *SearchQuery) {
				query.collectField(ctx, field)
			})
		case "items":
			d = d.WithItems(func(query *ItemQuery) {
				query.collectField(ctx, field)
			})
		case "last_sale":
			d = d.WithLastSale(func(query *ListingQuery) {
				query.collectField(ctx, field)
			})
		case "listings":
			d = d.WithListings(func(query *ListingQuery) {
				query.collectField(ctx, field)
			})
		case "wallet":
			d = d.WithWallet(func(query *WalletQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return d
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (e *EventQuery) CollectFields(ctx context.Context, satisfies ...string) *EventQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		e = e.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return e
}

func (e *EventQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *EventQuery {
	return e
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (gh *GameHustlerQuery) CollectFields(ctx context.Context, satisfies ...string) *GameHustlerQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		gh = gh.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return gh
}

func (gh *GameHustlerQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GameHustlerQuery {
	return gh
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ghi *GameHustlerItemQuery) CollectFields(ctx context.Context, satisfies ...string) *GameHustlerItemQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ghi = ghi.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ghi
}

func (ghi *GameHustlerItemQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GameHustlerItemQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "hustler":
			ghi = ghi.WithHustler(func(query *GameHustlerQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return ghi
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ghq *GameHustlerQuestQuery) CollectFields(ctx context.Context, satisfies ...string) *GameHustlerQuestQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ghq = ghq.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ghq
}

func (ghq *GameHustlerQuestQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GameHustlerQuestQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "hustler":
			ghq = ghq.WithHustler(func(query *GameHustlerQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return ghq
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ghr *GameHustlerRelationQuery) CollectFields(ctx context.Context, satisfies ...string) *GameHustlerRelationQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ghr = ghr.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ghr
}

func (ghr *GameHustlerRelationQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *GameHustlerRelationQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "hustler":
			ghr = ghr.WithHustler(func(query *GameHustlerQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return ghr
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (h *HustlerQuery) CollectFields(ctx context.Context, satisfies ...string) *HustlerQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		h = h.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return h
}

func (h *HustlerQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *HustlerQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "index":
			h = h.WithIndex(func(query *SearchQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return h
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (i *ItemQuery) CollectFields(ctx context.Context, satisfies ...string) *ItemQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		i = i.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return i
}

func (i *ItemQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ItemQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "dopes":
			i = i.WithDopes(func(query *DopeQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_accessories":
			i = i.WithHustlerAccessories(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_clothes":
			i = i.WithHustlerClothes(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_drugs":
			i = i.WithHustlerDrugs(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_feet":
			i = i.WithHustlerFeet(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_hands":
			i = i.WithHustlerHands(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_necks":
			i = i.WithHustlerNecks(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_rings":
			i = i.WithHustlerRings(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_vehicles":
			i = i.WithHustlerVehicles(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_waists":
			i = i.WithHustlerWaists(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "hustler_weapons":
			i = i.WithHustlerWeapons(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "index":
			i = i.WithIndex(func(query *SearchQuery) {
				query.collectField(ctx, field)
			})
		case "wallets":
			i = i.WithWallets(func(query *WalletItemsQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return i
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (l *ListingQuery) CollectFields(ctx context.Context, satisfies ...string) *ListingQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		l = l.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return l
}

func (l *ListingQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ListingQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "dope":
			l = l.WithDope(func(query *DopeQuery) {
				query.collectField(ctx, field)
			})
		case "dope_lastsales":
			l = l.WithDopeLastsales(func(query *DopeQuery) {
				query.collectField(ctx, field)
			})
		case "inputs":
			l = l.WithInputs(func(query *AmountQuery) {
				query.collectField(ctx, field)
			})
		case "outputs":
			l = l.WithOutputs(func(query *AmountQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return l
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (s *SearchQuery) CollectFields(ctx context.Context, satisfies ...string) *SearchQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		s = s.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return s
}

func (s *SearchQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *SearchQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "dope":
			s = s.WithDope(func(query *DopeQuery) {
				query.collectField(ctx, field)
			})
		case "hustler":
			s = s.WithHustler(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "item":
			s = s.WithItem(func(query *ItemQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return s
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ss *SyncStateQuery) CollectFields(ctx context.Context, satisfies ...string) *SyncStateQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		ss = ss.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return ss
}

func (ss *SyncStateQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *SyncStateQuery {
	return ss
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (w *WalletQuery) CollectFields(ctx context.Context, satisfies ...string) *WalletQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		w = w.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return w
}

func (w *WalletQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *WalletQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "dopes":
			w = w.WithDopes(func(query *DopeQuery) {
				query.collectField(ctx, field)
			})
		case "hustlers":
			w = w.WithHustlers(func(query *HustlerQuery) {
				query.collectField(ctx, field)
			})
		case "items":
			w = w.WithItems(func(query *WalletItemsQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return w
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (wi *WalletItemsQuery) CollectFields(ctx context.Context, satisfies ...string) *WalletItemsQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		wi = wi.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return wi
}

func (wi *WalletItemsQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *WalletItemsQuery {
	for _, field := range graphql.CollectFields(ctx, field.Selections, satisfies) {
		switch field.Name {
		case "item":
			wi = wi.WithItem(func(query *ItemQuery) {
				query.collectField(ctx, field)
			})
		case "wallet":
			wi = wi.WithWallet(func(query *WalletQuery) {
				query.collectField(ctx, field)
			})
		}
	}
	return wi
}