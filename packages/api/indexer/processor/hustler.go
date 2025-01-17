package processor

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/dopedao/dope-monorepo/packages/api/internal/contracts/bindings"
	"github.com/dopedao/dope-monorepo/packages/api/internal/dbprovider"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/bodypart"
	"github.com/dopedao/dope-monorepo/packages/api/internal/ent/hustler"
	"github.com/dopedao/dope-monorepo/packages/api/internal/logger"
	svgrender "github.com/dopedao/dope-monorepo/packages/api/internal/svg-render"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/rs/zerolog"
)

const (
	MaleBody         uint8  = 0
	FemaleBody       uint8  = 1
	MaleHair         uint8  = 2
	FemaleHair       uint8  = 3
	Beard            uint8  = 4
	OldestHustlerAge uint64 = 1637443660
)

var (
	componentsAddr = common.HexToAddress("0xe03C4eb2a0a797766a5DB708172e04f6A970DC7f")
	hustlerAddr    = common.HexToAddress("0xDbfEaAe58B6dA8901a8a40ba0712bEB2EE18368E")

	maskSlot      = big.NewInt(0)
	viewboxSlot   = big.NewInt(1)
	bodySlot      = big.NewInt(2)
	orderSlot     = big.NewInt(3)
	weaponSlot    = big.NewInt(5)
	clothesSlot   = big.NewInt(6)
	vehicleSlot   = big.NewInt(7)
	waistSlot     = big.NewInt(8)
	footSlot      = big.NewInt(9)
	handSlot      = big.NewInt(10)
	drugSlot      = big.NewInt(11)
	neckSlot      = big.NewInt(12)
	ringSlot      = big.NewInt(13)
	accessorySlot = big.NewInt(14)
)

var dbClient = dbprovider.Ent()

// HACK: Update addresses for testnet contracts. Stop gap until processors support dependency injection.
func init() {
	if os.Getenv("NETWORK") == "testnet" {
		componentsAddr = common.HexToAddress("0xA9d6Ef18457c4d87Ba77d5c22569c93a41b8f326")
	}

	if os.Getenv("NETWORK") == "testnet" {
		hustlerAddr = common.HexToAddress("0x5701ff301d67174d63B271cf321e3886d518370d")
	}
}

type HustlerProcessor struct {
	bindings.BaseHustlerProcessor
	components *bindings.Components
}

func (p *HustlerProcessor) Setup(address common.Address, eth interface {
	ethereum.ChainReader
	ethereum.ChainStateReader
	ethereum.TransactionReader
	bind.ContractBackend
}) error {
	if err := p.BaseHustlerProcessor.Setup(address, eth); err != nil {
		return err
	}

	var err error
	p.components, err = bindings.NewComponents(componentsAddr, eth)
	if err != nil {
		return err
	}

	return nil
}

func (p *HustlerProcessor) ProcessAddRles(ctx context.Context, e bindings.HustlerAddRles) (func(tx *ent.Tx) error, error) {
	// Tagged logger
	ctx, log := logger.LogFor(
		ctx,
		func(zctx *zerolog.Context) zerolog.Context {
			return zctx.Str("HustlerProcessor", "ProcessAddRles")
		},
	)

	var builders []*ent.BodyPartCreate

	var part bodypart.Type
	var sex bodypart.Sex

	switch e.Part {
	case MaleBody:
		part = bodypart.TypeBODY
		sex = bodypart.SexMALE
	case FemaleBody:
		part = bodypart.TypeBODY
		sex = bodypart.SexFEMALE
	case MaleHair:
		part = bodypart.TypeHAIR
		sex = bodypart.SexMALE
	case FemaleHair:
		part = bodypart.TypeHAIR
		sex = bodypart.SexFEMALE
	case Beard:
		part = bodypart.TypeBEARD
		sex = bodypart.SexMALE
	}

	log.Debug().
		Str("Body Part", part.String()).
		Str("Sex", sex.String()).
		Msg("Processing body part")

	return func(tx *ent.Tx) error {
		n, err := tx.BodyPart.
			Query().
			Where(bodypart.
				And(bodypart.TypeEQ(part), bodypart.SexEQ(sex))).
			Count(ctx)

		if err != nil {
			return fmt.Errorf("hustler: getting body count: %w", err)
		}

		for i := 0; i < int(e.Len.Int64()); i++ {
			id := int64(n + i)
			rle, err := p.Contract.BodyRle(nil, e.Part, big.NewInt(id))
			if err != nil {
				return fmt.Errorf("hustler: getting body rle part %d, id: %d: %w", e.Part, id, err)
			}
			builders = append(builders, tx.BodyPart.Create().
				SetID(fmt.Sprintf("%s-%s-%d", sex, part, id)).
				SetRle(hex.EncodeToString(rle)).
				SetType(part).
				SetSex(sex),
			)
		}

		if err := tx.BodyPart.CreateBulk(builders...).Exec(ctx); err != nil {
			return fmt.Errorf("hustler: creating bodyparts: %w", err)
		}

		return nil
	}, nil
}

func (p *HustlerProcessor) ProcessMetadataUpdate(
	ctx context.Context,
	e bindings.HustlerMetadataUpdate,
) (func(tx *ent.Tx) error, error) {

	ctx, log := logger.LogFor(
		ctx,
		func(zctx *zerolog.Context) zerolog.Context {
			return zctx.Str("HustlerProcessor", "ProcessMetadataUpdate")
		},
		func(zctx *zerolog.Context) zerolog.Context {
			return zctx.Str("hustler_id", e.Id.String())
		},
	)

	meta, err := p.Contract.Metadata(nil, e.Id)
	if err != nil {
		return nil, fmt.Errorf("hustler: getting metadata: %w", err)
	}

	hustlerSvg, err := GetHustlerSvg(p.Contract, e.Id, &log)
	if err != nil {
		return nil, fmt.Errorf("hustler: getting svg from tokenURI: %w", err)
	}

	metadataKey := new(big.Int).SetBytes(solsha3.SoliditySHA3(
		// types
		[]string{"uint256", "uint256"},

		// values
		[]interface{}{
			e.Id.String(),
			"19",
		},
	))

	viewbox, err := p.Eth.StorageAt(
		ctx,
		p.Address,
		common.BytesToHash(
			new(big.Int).Add(metadataKey, viewboxSlot).Bytes(),
		),
		new(big.Int).SetUint64(e.Raw.BlockNumber))
	if err != nil {
		return nil, fmt.Errorf("getting viewbox from storage: %w", err)
	}
	// log.Debug().Msgf("viewbox: %s", viewbox)

	order, err := p.Eth.StorageAt(
		ctx,
		p.Address,
		common.BytesToHash(
			new(big.Int).Add(metadataKey, orderSlot).Bytes(),
		),
		new(big.Int).SetUint64(e.Raw.BlockNumber))
	if err != nil {
		return nil, fmt.Errorf("getting order from storage: %w", err)
	}
	// log.Debug().Msgf("order: %s", order)

	bodyParts, err := p.Eth.StorageAt(
		ctx,
		p.Address,
		common.BytesToHash(
			new(big.Int).Add(metadataKey, bodySlot).Bytes(),
		),
		new(big.Int).SetUint64(e.Raw.BlockNumber))
	if err != nil {
		return nil, fmt.Errorf("getting body from storage: %w", err)
	}
	// log.Debug().Msgf("body: %s", bodyParts)

	// BEARD_ID has to exist in our database or indexer can fail
	// It's a pointer to string because it can be nil.
	var beardID *string
	sex := hustler.DefaultSex
	if new(big.Int).SetBytes(bodyParts[31:32]).Uint64() == 1 {
		sex = hustler.SexFEMALE
	} else {
		sex = hustler.SexMALE
		beardID_ := fmt.Sprintf("%s-%s-%d", sex, bodypart.TypeBEARD, new(big.Int).SetBytes(bodyParts[28:29]).Uint64())
		beardID = &beardID_
		// beard check
		_, err = dbClient.BodyPart.Query().Where(bodypart.IDEQ(*beardID)).Only(ctx)
		if err != nil {
			*beardID = ""
			log.Warn().Str("BeardID", *beardID).Msg("BAD BEARD")
		}

	}

	// BODY_ID has to exist in our database or indexer can fail
	bodyID := fmt.Sprintf(
		"%s-%s-%d", sex, bodypart.TypeBODY,
		new(big.Int).SetBytes(bodyParts[30:31]).Uint64(),
	)
	_, err = dbClient.BodyPart.Query().Where(bodypart.IDEQ(bodyID)).Only(ctx)
	if err != nil {
		log.Warn().Str("BodyID", bodyID).Msg("BAD BODY")
		bodyID = fmt.Sprintf(
			"%s-%s-%d", sex, bodypart.TypeBODY,
			// set to default if people are using random, non-existent parts
			0,
		)
	}

	// HAIR_ID has to exist in our database or indexer can fail
	hairID := fmt.Sprintf("%s-%s-%d", sex, bodypart.TypeHAIR, new(big.Int).SetBytes(bodyParts[29:30]).Uint64())
	_, err = dbClient.BodyPart.Query().Where(bodypart.IDEQ(hairID)).Only(ctx)
	if err != nil {
		log.Warn().Str("HairID", hairID).Msg("BAD HAIR")
		hairID = fmt.Sprintf(
			"%s-%s-%d", sex, bodypart.TypeHAIR,
			// set to default if people are using random, non-existent parts
			0,
		)
	}

	// We have to use a pointer here because title can be nil
	// and Go is draconian.
	var title *string
	typ := hustler.TypeREGULAR
	if e.Id.Cmp(big.NewInt(500)) == -1 {
		title_, err := p.components.Title(nil, e.Id)
		if err != nil {
			return nil, fmt.Errorf("getting hustler title %s metadata: %w", e.Id.String(), err)
		}
		title = &title_
		typ = hustler.TypeORIGINAL_GANGSTA
	}

	block, err := p.Eth.BlockByNumber(ctx, new(big.Int).SetUint64(e.Raw.BlockNumber))
	if err != nil {
		return nil, fmt.Errorf("Eth.ByBlockNumber ProcessMetadataUpdate updating hustler %s metadata: %w", e.Id.String(), err)
	}

	safeName := CleanJsonString(meta.Name)

	// LOG BEFORE UPDATING …and deal with stupid nil pointers
	var beardIDStr string
	if (beardID) != nil {
		beardIDStr = *beardID
	}
	var titleStr string
	if (title) != nil {
		titleStr = *title
	}
	log.Debug().
		Str("bodyParts", string(bodyParts)).
		Str("viewBox", string(viewbox)).
		Str("order", string(order)).
		Str("body_id", bodyID).
		Str("hair_id", hairID).
		Str("beard_id", beardIDStr).
		Str("title", titleStr).
		Str("name", safeName).
		Msg("Upsert hustler metadata")

	return func(tx *ent.Tx) error {
		if err := tx.Hustler.Create().
			SetID(e.Id.String()).
			SetType(typ).
			SetAge(block.Time()).
			SetName(safeName).
			SetBackground(hex.EncodeToString(meta.Background[:])).
			SetColor(hex.EncodeToString(meta.Color[:])).
			SetSex(sex).
			SetBodyID(bodyID).
			SetHairID(hairID).
			SetNillableBeardID(beardID).
			SetNillableTitle(title).
			SetSvg(hustlerSvg).
			SetViewbox([]int{
				int(new(big.Int).SetBytes(viewbox[31:32]).Int64()),
				int(new(big.Int).SetBytes(viewbox[30:31]).Int64()),
				int(new(big.Int).SetBytes(viewbox[29:30]).Int64()),
				int(new(big.Int).SetBytes(viewbox[28:29]).Int64()),
			}).
			SetOrder([]int{
				int(new(big.Int).SetBytes(order[31:32]).Int64()),
				int(new(big.Int).SetBytes(order[30:31]).Int64()),
				int(new(big.Int).SetBytes(order[29:30]).Int64()),
				int(new(big.Int).SetBytes(order[28:29]).Int64()),
				int(new(big.Int).SetBytes(order[27:28]).Int64()),
				int(new(big.Int).SetBytes(order[26:27]).Int64()),
				int(new(big.Int).SetBytes(order[25:26]).Int64()),
				int(new(big.Int).SetBytes(order[24:25]).Int64()),
				int(new(big.Int).SetBytes(order[23:24]).Int64()),
				int(new(big.Int).SetBytes(order[22:23]).Int64()),
			}).
			OnConflictColumns(hustler.FieldID).
			UpdateNewValues().
			Exec(ctx); err != nil {
			log.Debug().Str("indexer", "HUSTLER").
				Str("name", safeName).
				Msg("Failed Saving Hustler")
			return fmt.Errorf("ProcessMetadataUpdate ent tx updating hustler %s metadata: %w", e.Id.String(), err)
		}

		return nil
	}, nil
}

// Nasty Hustlers put bad unicode and null chars in things.
// We have to clean it up or it breaks our system.
func CleanJsonString(json string) string {
	return strings.ReplaceAll(strings.ToValidUTF8(json, ""), "\x00", "")
}

func (p *HustlerProcessor) ProcessTransferBatch(
	ctx context.Context,
	e bindings.HustlerTransferBatch,
) (func(tx *ent.Tx) error, error) {
	ctx, log := logger.LogFor(ctx)
	log.Debug().Str("indexer", "HUSTLER").Msgf("ProcessTransferBatch %v", e.Ids)

	return func(tx *ent.Tx) error {
		if err := ensureWalletExists(ctx, tx, e.To); err != nil {
			return fmt.Errorf("hustler: %w", err)
		}

		var ids []string
		for _, id := range e.Ids {
			ids = append(ids, id.String())
		}

		// TODO: reset age for non-og
		if err := tx.Hustler.Update().
			Where(hustler.IDIn(ids...)).
			SetWalletID(e.To.Hex()).
			Exec(ctx); err != nil {
			return fmt.Errorf("hustler: upsert to wallet: %w", err)
		}

		return nil
	}, nil
}

func (p *HustlerProcessor) ProcessTransferSingle(
	ctx context.Context,
	e bindings.HustlerTransferSingle,
) (func(tx *ent.Tx) error, error) {
	ctx, log := logger.LogFor(ctx)
	log.Debug().Str("indexer", "HUSTLER").Msgf("ProcessTransferSingle %s", e.Id.String())

	block, blockErr := p.Eth.BlockByNumber(
		ctx,
		new(big.Int).SetUint64(e.Raw.BlockNumber))
	if blockErr != nil {
		fmtErr := fmt.Errorf(
			"ProcessTransferSingle updating hustler %s metadata: %w",
			e.Id.String(),
			blockErr)
		return nil, fmtErr
	}

	return func(tx *ent.Tx) error {
		if err := ensureWalletExists(ctx, tx, e.To); err != nil {
			return fmt.Errorf("hustler: %w", err)
		}

		if e.From == (common.Address{}) {
			typ := hustler.TypeREGULAR
			if e.Id.Cmp(big.NewInt(500)) == -1 {
				typ = hustler.TypeORIGINAL_GANGSTA
			}

			if err := tx.Hustler.
				Create().
				SetID(e.Id.String()).
				SetWalletID(e.To.Hex()).
				SetType(typ).
				SetAge(block.Time()).
				OnConflictColumns(hustler.FieldID).
				UpdateNewValues().
				Exec(ctx); err != nil {
				return fmt.Errorf("hustler: create hustler: %w", err)
			}

			if err := RefreshHustlerEquipment(ctx, p.Eth, tx, e.Id.String(), hustlerAddr, new(big.Int).SetUint64(e.Raw.BlockNumber)); err != nil {
				return err
			}
		}

		if e.To != (common.Address{}) {
			chain := tx.Hustler.
				Create().
				SetID(e.Id.String()).
				SetWalletID(e.To.Hex())

			typ := hustler.TypeREGULAR
			if e.Id.Cmp(big.NewInt(500)) == -1 {
				typ = hustler.TypeORIGINAL_GANGSTA
				chain = chain.SetAge(OldestHustlerAge) // oldest OG age
			} else {
				// Reset age for non-ogs
				chain = chain.SetAge(block.Time())
			}

			if err := chain.
				SetType(typ).
				OnConflictColumns(hustler.FieldID).
				UpdateNewValues().
				Exec(ctx); err != nil {
				return fmt.Errorf("hustler: update hustler owner: %w", err)
			}
		}

		return nil
	}, nil
}

func RefreshHustlerEquipment(
	ctx context.Context,
	eth interface {
		bind.ContractBackend
		ethereum.ChainStateReader
		ethereum.TransactionReader
	},
	tx *ent.Tx,
	id string,
	address common.Address,
	blockNumber *big.Int,
) error {
	ctx, log := logger.LogFor(ctx)
	log.Debug().Str("indexer", "HUSTLER").Msgf("RefreshEquipment for %s", id)
	slots, err := equipmentSlots(ctx, eth, id, address, blockNumber)
	if err != nil {
		return err
	}

	h, err := bindings.NewHustler(address, eth)
	if err != nil {
		return fmt.Errorf("initialize hustler contract: %w", err)
	}

	big, ok := new(big.Int).SetString(id, 10)
	if !ok {
		return fmt.Errorf("casting id to int: %s", id)
	}

	u := tx.Hustler.Update().
		Where(hustler.IDEQ(id))

	hustlerSvg, err := GetHustlerSvg(h, big, &log)
	if err == nil {
		u.SetSvg(hustlerSvg)
	} else {
		// we dont want to cancel updating the equipment
		// just because we cant get the image
		log.Err(err).Msgf("couldnt get svg of hustler: %v", id)
	}

	if slots.Weapon != nil {
		u = u.SetWeaponID(slots.Weapon.String())
	} else {
		u = u.ClearWeapon()
	}

	if slots.Clothes != nil {
		u = u.SetClothesID(slots.Clothes.String())
	} else {
		u = u.ClearClothes()
	}

	if slots.Vehicle != nil {
		u = u.SetVehicleID(slots.Vehicle.String())
	} else {
		u = u.ClearVehicle()
	}

	if slots.Waist != nil {
		u = u.SetWaistID(slots.Waist.String())
	} else {
		u = u.ClearWaist()
	}

	if slots.Foot != nil {
		u = u.SetFootID(slots.Foot.String())
	} else {
		u = u.ClearFoot()
	}

	if slots.Hand != nil {
		u = u.SetHandID(slots.Hand.String())
	} else {
		u = u.ClearHand()
	}

	if slots.Drug != nil {
		u = u.SetDrugID(slots.Drug.String())
	} else {
		u = u.ClearDrug()
	}

	if slots.Neck != nil {
		u = u.SetNeckID(slots.Neck.String())
	} else {
		u = u.ClearNeck()
	}

	if slots.Ring != nil {
		u = u.SetRingID(slots.Ring.String())
	} else {
		u = u.ClearRing()
	}

	if slots.Accessory != nil {
		u = u.SetAccessoryID(slots.Accessory.String())
	} else {
		u = u.ClearAccessory()
	}

	if err := u.Exec(ctx); err != nil {
		return fmt.Errorf("updating equipment: %w", err)
	}

	return nil
}

// Tries to get hustler svg from chain,
// if it fails, it tries to get it from offchain
func GetHustlerSvg(h *bindings.Hustler, big *big.Int, log *zerolog.Logger) (string, error) {

	// On-chain SVG
	metadata, err := h.TokenURI(nil, big)
	if err != nil {
		log.
			Err(err).
			Str("indexer", "HUSTLER").
			Str("hustler_id", big.String()).
			Msg("COULD NOT GET HUSTLER IMAGE FROM TOKENURI.")

		// Off-chain SVG
		offchainHustlerSvg, err := svgrender.GetOffchainRender(big)
		if err != nil {
			return "", err
		}

		return offchainHustlerSvg, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(metadata, "data:application/json;base64,"))
	if err != nil {
		return "", fmt.Errorf("decoding metadata: %w", err)
	}

	safeStr := CleanJsonString(string(decoded))

	var parsed Metadata
	if err := json.Unmarshal([]byte(safeStr), &parsed); err != nil {
		return "", fmt.Errorf("hustler unmarshalling metadata: %w", err)
	}

	return parsed.Image, nil
}

type Slots struct {
	Weapon    *big.Int
	Clothes   *big.Int
	Vehicle   *big.Int
	Waist     *big.Int
	Foot      *big.Int
	Hand      *big.Int
	Drug      *big.Int
	Neck      *big.Int
	Ring      *big.Int
	Accessory *big.Int
}

func equipmentSlots(
	ctx context.Context,
	eth interface {
		ethereum.ChainStateReader
		ethereum.TransactionReader
	},
	id string,
	address common.Address,
	blockNumber *big.Int,
) (*Slots, error) {
	slots := &Slots{}

	metadataKey := new(big.Int).SetBytes(solsha3.SoliditySHA3(
		[]string{"uint256", "uint256"},
		[]interface{}{
			id,
			"19",
		},
	))

	mask, err := eth.StorageAt(
		ctx,
		address,
		common.BytesToHash(
			new(big.Int).Add(metadataKey, maskSlot).Bytes(),
		),
		blockNumber)
	if err != nil {
		return nil, fmt.Errorf("getting mask from storage: %w", err)
	}

	// Little endian
	if mask[31-8]&1 != 0 {
		slots.Weapon, err = equipmentSlot(ctx, eth, id, address, weaponSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting weapon from storage: %w", err)
		}
	}

	if mask[31-8]&2 != 0 {
		slots.Clothes, err = equipmentSlot(ctx, eth, id, address, clothesSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting clothes from storage: %w", err)
		}
	}

	if mask[31-8]&4 != 0 {
		slots.Vehicle, err = equipmentSlot(ctx, eth, id, address, vehicleSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting vehicle from storage: %w", err)
		}
	}

	if mask[31-8]&8 != 0 {
		slots.Waist, err = equipmentSlot(ctx, eth, id, address, waistSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting waist from storage: %w", err)
		}
	}

	if mask[31-8]&16 != 0 {
		slots.Foot, err = equipmentSlot(ctx, eth, id, address, footSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting foot from storage: %w", err)
		}
	}

	if mask[31-8]&32 != 0 {
		slots.Hand, err = equipmentSlot(ctx, eth, id, address, handSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting hand from storage: %w", err)
		}
	}

	if mask[31-8]&64 != 0 {
		slots.Drug, err = equipmentSlot(ctx, eth, id, address, drugSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting drug from storage: %w", err)
		}
	}

	if mask[31-8]&128 != 0 {
		slots.Neck, err = equipmentSlot(ctx, eth, id, address, neckSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting neck from storage: %w", err)
		}
	}

	if mask[31-9]&1 != 0 {
		slots.Ring, err = equipmentSlot(ctx, eth, id, address, ringSlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting ring from storage: %w", err)
		}
	}

	if mask[31-9]&2 != 0 {
		slots.Accessory, err = equipmentSlot(ctx, eth, id, address, accessorySlot, blockNumber)
		if err != nil {
			return nil, fmt.Errorf("getting accessory from storage: %w", err)
		}
	}

	return slots, nil
}

func equipmentSlot(ctx context.Context, eth interface {
	ethereum.ChainStateReader
	ethereum.TransactionReader
}, id string, address common.Address, slot *big.Int, blockNumber *big.Int) (*big.Int, error) {
	metadataKey := new(big.Int).SetBytes(solsha3.SoliditySHA3(
		[]string{"uint256", "uint256"},
		[]interface{}{
			id,
			"19",
		},
	))

	value, err := eth.StorageAt(
		ctx,
		address,
		common.BytesToHash(
			new(big.Int).Add(metadataKey, slot).Bytes(),
		),
		blockNumber)
	if err != nil {
		return nil, fmt.Errorf("getting equipment from storage: %w", err)
	}

	return new(big.Int).SetBytes(value), nil
}
