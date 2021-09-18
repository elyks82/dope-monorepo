/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import {
  ethers,
  EventFilter,
  Signer,
  BigNumber,
  BigNumberish,
  PopulatedTransaction,
  BaseContract,
  ContractTransaction,
  Overrides,
  CallOverrides,
} from "ethers";
import { BytesLike } from "@ethersproject/bytes";
import { Listener, Provider } from "@ethersproject/providers";
import { FunctionFragment, EventFragment, Result } from "@ethersproject/abi";
import { TypedEventFilter, TypedEvent, TypedListener } from "./commons";

interface StockpileInterface extends ethers.utils.Interface {
  functions: {
    "attributes(uint256)": FunctionFragment;
    "balanceOf(address,uint256)": FunctionFragment;
    "balanceOfBatch(address[],uint256[])": FunctionFragment;
    "clothesId(uint256)": FunctionFragment;
    "componentsToString(uint256[5],uint256)": FunctionFragment;
    "contractURI()": FunctionFragment;
    "drugsId(uint256)": FunctionFragment;
    "footId(uint256)": FunctionFragment;
    "handId(uint256)": FunctionFragment;
    "ids(uint256)": FunctionFragment;
    "idsMany(uint256[])": FunctionFragment;
    "isApprovedForAll(address,address)": FunctionFragment;
    "itemName(uint256,uint256)": FunctionFragment;
    "name()": FunctionFragment;
    "names(uint256)": FunctionFragment;
    "namesMany(uint256[])": FunctionFragment;
    "neckId(uint256)": FunctionFragment;
    "open(uint256)": FunctionFragment;
    "ownedBatchRLE(uint256[])": FunctionFragment;
    "ringId(uint256)": FunctionFragment;
    "rleOf(uint256)": FunctionFragment;
    "rleOfBatch(uint256[])": FunctionFragment;
    "safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)": FunctionFragment;
    "safeTransferFrom(address,address,uint256,uint256,bytes)": FunctionFragment;
    "setApprovalForAll(address,bool)": FunctionFragment;
    "supportsInterface(bytes4)": FunctionFragment;
    "symbol()": FunctionFragment;
    "tokenName(uint256)": FunctionFragment;
    "tokenURI(uint256)": FunctionFragment;
    "uri(uint256)": FunctionFragment;
    "vehicleId(uint256)": FunctionFragment;
    "waistId(uint256)": FunctionFragment;
    "weaponId(uint256)": FunctionFragment;
  };

  encodeFunctionData(
    functionFragment: "attributes",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "balanceOf",
    values: [string, BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "balanceOfBatch",
    values: [string[], BigNumberish[]]
  ): string;
  encodeFunctionData(
    functionFragment: "clothesId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "componentsToString",
    values: [
      [BigNumberish, BigNumberish, BigNumberish, BigNumberish, BigNumberish],
      BigNumberish
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "contractURI",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "drugsId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "footId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "handId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "ids", values: [BigNumberish]): string;
  encodeFunctionData(
    functionFragment: "idsMany",
    values: [BigNumberish[]]
  ): string;
  encodeFunctionData(
    functionFragment: "isApprovedForAll",
    values: [string, string]
  ): string;
  encodeFunctionData(
    functionFragment: "itemName",
    values: [BigNumberish, BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "name", values?: undefined): string;
  encodeFunctionData(functionFragment: "names", values: [BigNumberish]): string;
  encodeFunctionData(
    functionFragment: "namesMany",
    values: [BigNumberish[]]
  ): string;
  encodeFunctionData(
    functionFragment: "neckId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "open", values: [BigNumberish]): string;
  encodeFunctionData(
    functionFragment: "ownedBatchRLE",
    values: [BigNumberish[]]
  ): string;
  encodeFunctionData(
    functionFragment: "ringId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "rleOf", values: [BigNumberish]): string;
  encodeFunctionData(
    functionFragment: "rleOfBatch",
    values: [BigNumberish[]]
  ): string;
  encodeFunctionData(
    functionFragment: "safeBatchTransferFrom",
    values: [string, string, BigNumberish[], BigNumberish[], BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "safeTransferFrom",
    values: [string, string, BigNumberish, BigNumberish, BytesLike]
  ): string;
  encodeFunctionData(
    functionFragment: "setApprovalForAll",
    values: [string, boolean]
  ): string;
  encodeFunctionData(
    functionFragment: "supportsInterface",
    values: [BytesLike]
  ): string;
  encodeFunctionData(functionFragment: "symbol", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "tokenName",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "tokenURI",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(functionFragment: "uri", values: [BigNumberish]): string;
  encodeFunctionData(
    functionFragment: "vehicleId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "waistId",
    values: [BigNumberish]
  ): string;
  encodeFunctionData(
    functionFragment: "weaponId",
    values: [BigNumberish]
  ): string;

  decodeFunctionResult(functionFragment: "attributes", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "balanceOf", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "balanceOfBatch",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "clothesId", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "componentsToString",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "contractURI",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "drugsId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "footId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "handId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "ids", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "idsMany", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "isApprovedForAll",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "itemName", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "name", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "names", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "namesMany", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "neckId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "open", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "ownedBatchRLE",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "ringId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "rleOf", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "rleOfBatch", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "safeBatchTransferFrom",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "safeTransferFrom",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "setApprovalForAll",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "supportsInterface",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "symbol", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "tokenName", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "tokenURI", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "uri", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "vehicleId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "waistId", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "weaponId", data: BytesLike): Result;

  events: {
    "ApprovalForAll(address,address,bool)": EventFragment;
    "TransferBatch(address,address,address,uint256[],uint256[])": EventFragment;
    "TransferSingle(address,address,address,uint256,uint256)": EventFragment;
    "URI(string,uint256)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "ApprovalForAll"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "TransferBatch"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "TransferSingle"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "URI"): EventFragment;
}

export type ApprovalForAllEvent = TypedEvent<
  [string, string, boolean] & {
    account: string;
    operator: string;
    approved: boolean;
  }
>;

export type TransferBatchEvent = TypedEvent<
  [string, string, string, BigNumber[], BigNumber[]] & {
    operator: string;
    from: string;
    to: string;
    ids: BigNumber[];
    values: BigNumber[];
  }
>;

export type TransferSingleEvent = TypedEvent<
  [string, string, string, BigNumber, BigNumber] & {
    operator: string;
    from: string;
    to: string;
    id: BigNumber;
    value: BigNumber;
  }
>;

export type URIEvent = TypedEvent<
  [string, BigNumber] & { value: string; id: BigNumber }
>;

export class Stockpile extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  listeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter?: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): Array<TypedListener<EventArgsArray, EventArgsObject>>;
  off<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  on<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  once<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  removeListener<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>,
    listener: TypedListener<EventArgsArray, EventArgsObject>
  ): this;
  removeAllListeners<EventArgsArray extends Array<any>, EventArgsObject>(
    eventFilter: TypedEventFilter<EventArgsArray, EventArgsObject>
  ): this;

  listeners(eventName?: string): Array<Listener>;
  off(eventName: string, listener: Listener): this;
  on(eventName: string, listener: Listener): this;
  once(eventName: string, listener: Listener): this;
  removeListener(eventName: string, listener: Listener): this;
  removeAllListeners(eventName?: string): this;

  queryFilter<EventArgsArray extends Array<any>, EventArgsObject>(
    event: TypedEventFilter<EventArgsArray, EventArgsObject>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TypedEvent<EventArgsArray & EventArgsObject>>>;

  interface: StockpileInterface;

  functions: {
    attributes(id: BigNumberish, overrides?: CallOverrides): Promise<[string]>;

    balanceOf(
      account: string,
      id: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    balanceOfBatch(
      accounts: string[],
      ids: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<[BigNumber[]]>;

    clothesId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    componentsToString(
      components: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
      ],
      itemType: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string]>;

    contractURI(overrides?: CallOverrides): Promise<[string]>;

    drugsId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    footId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    handId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    ids(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<
      [
        [
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber
        ] & {
          weapon: BigNumber;
          clothes: BigNumber;
          vehicle: BigNumber;
          waist: BigNumber;
          foot: BigNumber;
          hand: BigNumber;
          drugs: BigNumber;
          neck: BigNumber;
          ring: BigNumber;
        }
      ]
    >;

    idsMany(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<
      [
        ([
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber,
          BigNumber
        ] & {
          weapon: BigNumber;
          clothes: BigNumber;
          vehicle: BigNumber;
          waist: BigNumber;
          foot: BigNumber;
          hand: BigNumber;
          drugs: BigNumber;
          neck: BigNumber;
          ring: BigNumber;
        })[]
      ]
    >;

    isApprovedForAll(
      account: string,
      operator: string,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    itemName(
      itemType: BigNumberish,
      idx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string]>;

    name(overrides?: CallOverrides): Promise<[string]>;

    names(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<
      [
        [
          string,
          string,
          string,
          string,
          string,
          string,
          string,
          string,
          string
        ] & {
          weapon: string;
          clothes: string;
          vehicle: string;
          waist: string;
          foot: string;
          hand: string;
          drugs: string;
          neck: string;
          ring: string;
        }
      ]
    >;

    namesMany(
      tokenNames: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<
      [
        ([
          string,
          string,
          string,
          string,
          string,
          string,
          string,
          string,
          string
        ] & {
          weapon: string;
          clothes: string;
          vehicle: string;
          waist: string;
          foot: string;
          hand: string;
          drugs: string;
          neck: string;
          ring: string;
        })[]
      ]
    >;

    neckId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    open(
      tokenId: BigNumberish,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    ownedBatchRLE(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<[string[]] & { rles: string[] }>;

    ringId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    rleOf(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string] & { rle: string }>;

    rleOfBatch(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<[string[]] & { rles: string[] }>;

    safeBatchTransferFrom(
      from: string,
      to: string,
      ids: BigNumberish[],
      amounts: BigNumberish[],
      data: BytesLike,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    safeTransferFrom(
      from: string,
      to: string,
      id: BigNumberish,
      amount: BigNumberish,
      data: BytesLike,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    setApprovalForAll(
      operator: string,
      approved: boolean,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<ContractTransaction>;

    supportsInterface(
      interfaceId: BytesLike,
      overrides?: CallOverrides
    ): Promise<[boolean]>;

    symbol(overrides?: CallOverrides): Promise<[string]>;

    tokenName(id: BigNumberish, overrides?: CallOverrides): Promise<[string]>;

    tokenURI(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[string]>;

    uri(tokenId: BigNumberish, overrides?: CallOverrides): Promise<[string]>;

    vehicleId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    waistId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;

    weaponId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<[BigNumber]>;
  };

  attributes(id: BigNumberish, overrides?: CallOverrides): Promise<string>;

  balanceOf(
    account: string,
    id: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  balanceOfBatch(
    accounts: string[],
    ids: BigNumberish[],
    overrides?: CallOverrides
  ): Promise<BigNumber[]>;

  clothesId(
    tokenId: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  componentsToString(
    components: [
      BigNumberish,
      BigNumberish,
      BigNumberish,
      BigNumberish,
      BigNumberish
    ],
    itemType: BigNumberish,
    overrides?: CallOverrides
  ): Promise<string>;

  contractURI(overrides?: CallOverrides): Promise<string>;

  drugsId(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

  footId(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

  handId(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

  ids(
    tokenId: BigNumberish,
    overrides?: CallOverrides
  ): Promise<
    [
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber
    ] & {
      weapon: BigNumber;
      clothes: BigNumber;
      vehicle: BigNumber;
      waist: BigNumber;
      foot: BigNumber;
      hand: BigNumber;
      drugs: BigNumber;
      neck: BigNumber;
      ring: BigNumber;
    }
  >;

  idsMany(
    tokenIds: BigNumberish[],
    overrides?: CallOverrides
  ): Promise<
    ([
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber,
      BigNumber
    ] & {
      weapon: BigNumber;
      clothes: BigNumber;
      vehicle: BigNumber;
      waist: BigNumber;
      foot: BigNumber;
      hand: BigNumber;
      drugs: BigNumber;
      neck: BigNumber;
      ring: BigNumber;
    })[]
  >;

  isApprovedForAll(
    account: string,
    operator: string,
    overrides?: CallOverrides
  ): Promise<boolean>;

  itemName(
    itemType: BigNumberish,
    idx: BigNumberish,
    overrides?: CallOverrides
  ): Promise<string>;

  name(overrides?: CallOverrides): Promise<string>;

  names(
    tokenId: BigNumberish,
    overrides?: CallOverrides
  ): Promise<
    [string, string, string, string, string, string, string, string, string] & {
      weapon: string;
      clothes: string;
      vehicle: string;
      waist: string;
      foot: string;
      hand: string;
      drugs: string;
      neck: string;
      ring: string;
    }
  >;

  namesMany(
    tokenNames: BigNumberish[],
    overrides?: CallOverrides
  ): Promise<
    ([
      string,
      string,
      string,
      string,
      string,
      string,
      string,
      string,
      string
    ] & {
      weapon: string;
      clothes: string;
      vehicle: string;
      waist: string;
      foot: string;
      hand: string;
      drugs: string;
      neck: string;
      ring: string;
    })[]
  >;

  neckId(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

  open(
    tokenId: BigNumberish,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  ownedBatchRLE(
    tokenIds: BigNumberish[],
    overrides?: CallOverrides
  ): Promise<string[]>;

  ringId(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

  rleOf(tokenId: BigNumberish, overrides?: CallOverrides): Promise<string>;

  rleOfBatch(
    tokenIds: BigNumberish[],
    overrides?: CallOverrides
  ): Promise<string[]>;

  safeBatchTransferFrom(
    from: string,
    to: string,
    ids: BigNumberish[],
    amounts: BigNumberish[],
    data: BytesLike,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  safeTransferFrom(
    from: string,
    to: string,
    id: BigNumberish,
    amount: BigNumberish,
    data: BytesLike,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  setApprovalForAll(
    operator: string,
    approved: boolean,
    overrides?: Overrides & { from?: string | Promise<string> }
  ): Promise<ContractTransaction>;

  supportsInterface(
    interfaceId: BytesLike,
    overrides?: CallOverrides
  ): Promise<boolean>;

  symbol(overrides?: CallOverrides): Promise<string>;

  tokenName(id: BigNumberish, overrides?: CallOverrides): Promise<string>;

  tokenURI(tokenId: BigNumberish, overrides?: CallOverrides): Promise<string>;

  uri(tokenId: BigNumberish, overrides?: CallOverrides): Promise<string>;

  vehicleId(
    tokenId: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  waistId(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

  weaponId(
    tokenId: BigNumberish,
    overrides?: CallOverrides
  ): Promise<BigNumber>;

  callStatic: {
    attributes(id: BigNumberish, overrides?: CallOverrides): Promise<string>;

    balanceOf(
      account: string,
      id: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    balanceOfBatch(
      accounts: string[],
      ids: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber[]>;

    clothesId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    componentsToString(
      components: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
      ],
      itemType: BigNumberish,
      overrides?: CallOverrides
    ): Promise<string>;

    contractURI(overrides?: CallOverrides): Promise<string>;

    drugsId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    footId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    handId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    ids(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<
      [
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber
      ] & {
        weapon: BigNumber;
        clothes: BigNumber;
        vehicle: BigNumber;
        waist: BigNumber;
        foot: BigNumber;
        hand: BigNumber;
        drugs: BigNumber;
        neck: BigNumber;
        ring: BigNumber;
      }
    >;

    idsMany(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<
      ([
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber,
        BigNumber
      ] & {
        weapon: BigNumber;
        clothes: BigNumber;
        vehicle: BigNumber;
        waist: BigNumber;
        foot: BigNumber;
        hand: BigNumber;
        drugs: BigNumber;
        neck: BigNumber;
        ring: BigNumber;
      })[]
    >;

    isApprovedForAll(
      account: string,
      operator: string,
      overrides?: CallOverrides
    ): Promise<boolean>;

    itemName(
      itemType: BigNumberish,
      idx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<string>;

    name(overrides?: CallOverrides): Promise<string>;

    names(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<
      [
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string
      ] & {
        weapon: string;
        clothes: string;
        vehicle: string;
        waist: string;
        foot: string;
        hand: string;
        drugs: string;
        neck: string;
        ring: string;
      }
    >;

    namesMany(
      tokenNames: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<
      ([
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string,
        string
      ] & {
        weapon: string;
        clothes: string;
        vehicle: string;
        waist: string;
        foot: string;
        hand: string;
        drugs: string;
        neck: string;
        ring: string;
      })[]
    >;

    neckId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    open(tokenId: BigNumberish, overrides?: CallOverrides): Promise<void>;

    ownedBatchRLE(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<string[]>;

    ringId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    rleOf(tokenId: BigNumberish, overrides?: CallOverrides): Promise<string>;

    rleOfBatch(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<string[]>;

    safeBatchTransferFrom(
      from: string,
      to: string,
      ids: BigNumberish[],
      amounts: BigNumberish[],
      data: BytesLike,
      overrides?: CallOverrides
    ): Promise<void>;

    safeTransferFrom(
      from: string,
      to: string,
      id: BigNumberish,
      amount: BigNumberish,
      data: BytesLike,
      overrides?: CallOverrides
    ): Promise<void>;

    setApprovalForAll(
      operator: string,
      approved: boolean,
      overrides?: CallOverrides
    ): Promise<void>;

    supportsInterface(
      interfaceId: BytesLike,
      overrides?: CallOverrides
    ): Promise<boolean>;

    symbol(overrides?: CallOverrides): Promise<string>;

    tokenName(id: BigNumberish, overrides?: CallOverrides): Promise<string>;

    tokenURI(tokenId: BigNumberish, overrides?: CallOverrides): Promise<string>;

    uri(tokenId: BigNumberish, overrides?: CallOverrides): Promise<string>;

    vehicleId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    waistId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    weaponId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  filters: {
    "ApprovalForAll(address,address,bool)"(
      account?: string | null,
      operator?: string | null,
      approved?: null
    ): TypedEventFilter<
      [string, string, boolean],
      { account: string; operator: string; approved: boolean }
    >;

    ApprovalForAll(
      account?: string | null,
      operator?: string | null,
      approved?: null
    ): TypedEventFilter<
      [string, string, boolean],
      { account: string; operator: string; approved: boolean }
    >;

    "TransferBatch(address,address,address,uint256[],uint256[])"(
      operator?: string | null,
      from?: string | null,
      to?: string | null,
      ids?: null,
      values?: null
    ): TypedEventFilter<
      [string, string, string, BigNumber[], BigNumber[]],
      {
        operator: string;
        from: string;
        to: string;
        ids: BigNumber[];
        values: BigNumber[];
      }
    >;

    TransferBatch(
      operator?: string | null,
      from?: string | null,
      to?: string | null,
      ids?: null,
      values?: null
    ): TypedEventFilter<
      [string, string, string, BigNumber[], BigNumber[]],
      {
        operator: string;
        from: string;
        to: string;
        ids: BigNumber[];
        values: BigNumber[];
      }
    >;

    "TransferSingle(address,address,address,uint256,uint256)"(
      operator?: string | null,
      from?: string | null,
      to?: string | null,
      id?: null,
      value?: null
    ): TypedEventFilter<
      [string, string, string, BigNumber, BigNumber],
      {
        operator: string;
        from: string;
        to: string;
        id: BigNumber;
        value: BigNumber;
      }
    >;

    TransferSingle(
      operator?: string | null,
      from?: string | null,
      to?: string | null,
      id?: null,
      value?: null
    ): TypedEventFilter<
      [string, string, string, BigNumber, BigNumber],
      {
        operator: string;
        from: string;
        to: string;
        id: BigNumber;
        value: BigNumber;
      }
    >;

    "URI(string,uint256)"(
      value?: null,
      id?: BigNumberish | null
    ): TypedEventFilter<[string, BigNumber], { value: string; id: BigNumber }>;

    URI(
      value?: null,
      id?: BigNumberish | null
    ): TypedEventFilter<[string, BigNumber], { value: string; id: BigNumber }>;
  };

  estimateGas: {
    attributes(id: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    balanceOf(
      account: string,
      id: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    balanceOfBatch(
      accounts: string[],
      ids: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    clothesId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    componentsToString(
      components: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
      ],
      itemType: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    contractURI(overrides?: CallOverrides): Promise<BigNumber>;

    drugsId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    footId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    handId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    ids(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    idsMany(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    isApprovedForAll(
      account: string,
      operator: string,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    itemName(
      itemType: BigNumberish,
      idx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    name(overrides?: CallOverrides): Promise<BigNumber>;

    names(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    namesMany(
      tokenNames: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    neckId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    open(
      tokenId: BigNumberish,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    ownedBatchRLE(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    ringId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    rleOf(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    rleOfBatch(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    safeBatchTransferFrom(
      from: string,
      to: string,
      ids: BigNumberish[],
      amounts: BigNumberish[],
      data: BytesLike,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    safeTransferFrom(
      from: string,
      to: string,
      id: BigNumberish,
      amount: BigNumberish,
      data: BytesLike,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    setApprovalForAll(
      operator: string,
      approved: boolean,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<BigNumber>;

    supportsInterface(
      interfaceId: BytesLike,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    symbol(overrides?: CallOverrides): Promise<BigNumber>;

    tokenName(id: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    tokenURI(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    uri(tokenId: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;

    vehicleId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    waistId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    weaponId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    attributes(
      id: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    balanceOf(
      account: string,
      id: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    balanceOfBatch(
      accounts: string[],
      ids: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    clothesId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    componentsToString(
      components: [
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
      ],
      itemType: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    contractURI(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    drugsId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    footId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    handId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    ids(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    idsMany(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    isApprovedForAll(
      account: string,
      operator: string,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    itemName(
      itemType: BigNumberish,
      idx: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    name(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    names(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    namesMany(
      tokenNames: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    neckId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    open(
      tokenId: BigNumberish,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    ownedBatchRLE(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    ringId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    rleOf(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    rleOfBatch(
      tokenIds: BigNumberish[],
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    safeBatchTransferFrom(
      from: string,
      to: string,
      ids: BigNumberish[],
      amounts: BigNumberish[],
      data: BytesLike,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    safeTransferFrom(
      from: string,
      to: string,
      id: BigNumberish,
      amount: BigNumberish,
      data: BytesLike,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    setApprovalForAll(
      operator: string,
      approved: boolean,
      overrides?: Overrides & { from?: string | Promise<string> }
    ): Promise<PopulatedTransaction>;

    supportsInterface(
      interfaceId: BytesLike,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    symbol(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    tokenName(
      id: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    tokenURI(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    uri(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    vehicleId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    waistId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    weaponId(
      tokenId: BigNumberish,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
