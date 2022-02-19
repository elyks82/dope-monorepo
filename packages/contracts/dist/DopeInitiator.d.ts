import { BaseContract, BigNumber, BigNumberish, BytesLike, CallOverrides, ContractTransaction, PayableOverrides, PopulatedTransaction, Signer, utils } from "ethers";
import { FunctionFragment, Result } from "@ethersproject/abi";
import { Listener, Provider } from "@ethersproject/providers";
import { TypedEventFilter, TypedEvent, TypedListener, OnEvent } from "./common";
export declare type OrderStruct = {
    undefined: string;
    undefined: BigNumberish;
    undefined: [BytesLike, BytesLike];
    undefined: BigNumberish;
    undefined: BigNumberish;
    undefined: BigNumberish;
    undefined: BigNumberish;
    undefined: BigNumberish;
    undefined: BytesLike;
    undefined: BytesLike;
};
export declare type OrderStructOutput = [
    string,
    number,
    [
        string,
        string
    ],
    BigNumber,
    BigNumber,
    BigNumber,
    BigNumber,
    BigNumber,
    string,
    string
];
export declare type SetMetadataStruct = {
    undefined: BytesLike;
    undefined: BytesLike;
    undefined: BytesLike;
    undefined: [BigNumberish, BigNumberish, BigNumberish, BigNumberish];
    undefined: [BigNumberish, BigNumberish, BigNumberish, BigNumberish];
    undefined: BigNumberish[];
    undefined: BytesLike;
    undefined: string;
};
export declare type SetMetadataStructOutput = [
    string,
    string,
    string,
    [
        number,
        number,
        number,
        number
    ],
    [
        number,
        number,
        number,
        number
    ],
    number[],
    string,
    string
];
export interface DopeInitiatorInterface extends utils.Interface {
    functions: {
        "estimate(uint256)": FunctionFragment;
        "initiate((address,uint8,bytes32[2],uint256,uint256,uint256,uint256,uint256,bytes,bytes),uint256,(bytes4,bytes4,bytes2,uint8[4],uint8[4],uint8[10],bytes2,string),address,uint256,uint256,uint256,uint256)": FunctionFragment;
    };
    encodeFunctionData(functionFragment: "estimate", values: [BigNumberish]): string;
    encodeFunctionData(functionFragment: "initiate", values: [
        OrderStruct,
        BigNumberish,
        SetMetadataStruct,
        string,
        BigNumberish,
        BigNumberish,
        BigNumberish,
        BigNumberish
    ]): string;
    decodeFunctionResult(functionFragment: "estimate", data: BytesLike): Result;
    decodeFunctionResult(functionFragment: "initiate", data: BytesLike): Result;
    events: {};
}
export interface DopeInitiator extends BaseContract {
    connect(signerOrProvider: Signer | Provider | string): this;
    attach(addressOrName: string): this;
    deployed(): Promise<this>;
    interface: DopeInitiatorInterface;
    queryFilter<TEvent extends TypedEvent>(event: TypedEventFilter<TEvent>, fromBlockOrBlockhash?: string | number | undefined, toBlock?: string | number | undefined): Promise<Array<TEvent>>;
    listeners<TEvent extends TypedEvent>(eventFilter?: TypedEventFilter<TEvent>): Array<TypedListener<TEvent>>;
    listeners(eventName?: string): Array<Listener>;
    removeAllListeners<TEvent extends TypedEvent>(eventFilter: TypedEventFilter<TEvent>): this;
    removeAllListeners(eventName?: string): this;
    off: OnEvent<this>;
    on: OnEvent<this>;
    once: OnEvent<this>;
    removeListener: OnEvent<this>;
    functions: {
        estimate(paper: BigNumberish, overrides?: PayableOverrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
        initiate(order: OrderStruct, id: BigNumberish, meta: SetMetadataStruct, to: string, openseaEth: BigNumberish, paperEth: BigNumberish, paperOut: BigNumberish, deadline: BigNumberish, overrides?: PayableOverrides & {
            from?: string | Promise<string>;
        }): Promise<ContractTransaction>;
    };
    estimate(paper: BigNumberish, overrides?: PayableOverrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    initiate(order: OrderStruct, id: BigNumberish, meta: SetMetadataStruct, to: string, openseaEth: BigNumberish, paperEth: BigNumberish, paperOut: BigNumberish, deadline: BigNumberish, overrides?: PayableOverrides & {
        from?: string | Promise<string>;
    }): Promise<ContractTransaction>;
    callStatic: {
        estimate(paper: BigNumberish, overrides?: CallOverrides): Promise<BigNumber>;
        initiate(order: OrderStruct, id: BigNumberish, meta: SetMetadataStruct, to: string, openseaEth: BigNumberish, paperEth: BigNumberish, paperOut: BigNumberish, deadline: BigNumberish, overrides?: CallOverrides): Promise<void>;
    };
    filters: {};
    estimateGas: {
        estimate(paper: BigNumberish, overrides?: PayableOverrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
        initiate(order: OrderStruct, id: BigNumberish, meta: SetMetadataStruct, to: string, openseaEth: BigNumberish, paperEth: BigNumberish, paperOut: BigNumberish, deadline: BigNumberish, overrides?: PayableOverrides & {
            from?: string | Promise<string>;
        }): Promise<BigNumber>;
    };
    populateTransaction: {
        estimate(paper: BigNumberish, overrides?: PayableOverrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
        initiate(order: OrderStruct, id: BigNumberish, meta: SetMetadataStruct, to: string, openseaEth: BigNumberish, paperEth: BigNumberish, paperOut: BigNumberish, deadline: BigNumberish, overrides?: PayableOverrides & {
            from?: string | Promise<string>;
        }): Promise<PopulatedTransaction>;
    };
}
