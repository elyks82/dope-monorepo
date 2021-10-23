// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import './utils/HustlerSetup.sol';
import './utils/SwapMeetSetup.sol';

struct Attribute {
    string traitType;
    string value;
}

struct Data {
    string name;
    string description;
    string image;
    Attribute[] attributes;
}

contract RLES is HustlerTest {
    function testCanAddBodies() public {
        bytes[] memory bodies = new bytes[](2);
        bodies[0] = bodyRle;
        bodies[1] = bodyRle;

        owner.addRles(RleParts.BODY, bodies);
        assertEq(string(hustler.getBody(0)), string(bodyRle));
        assertEq(string(hustler.getBody(1)), string(bodyRle));
    }

    function testCanAddBeards() public {
        bytes[] memory beards = new bytes[](2);
        beards[0] = beardRle;
        beards[1] = beardRle;
        owner.addRles(RleParts.BEARD, beards);

        assertEq(string(hustler.getBeard(0)), string(beardRle));
        assertEq(string(hustler.getBeard(1)), string(beardRle));
    }

    function testCanaddHairs() public {
        bytes[] memory hairs = new bytes[](2);
        hairs[0] = hairRle;
        hairs[1] = hairRle;
        owner.addRles(RleParts.MALE_HAIR, hairs);
        owner.addRles(RleParts.FEMALE_HAIR, hairs);
        assertEq(string(hustler.getHair(RleParts.MALE_HAIR, 0)), string(hairRle));
        assertEq(string(hustler.getHair(RleParts.MALE_HAIR, 1)), string(hairRle));
        assertEq(string(hustler.getHair(RleParts.FEMALE_HAIR, 0)), string(hairRle));
        assertEq(string(hustler.getHair(RleParts.FEMALE_HAIR, 1)), string(hairRle));
    }
}

contract Hustlers is HustlerTest {
    function setUp() public override {
        super.setUp();

        bytes[] memory heads = new bytes[](2);
        heads[0] = hairRle;
        heads[1] = hairRle;
        owner.addRles(RleParts.MALE_HAIR, heads);
        owner.addRles(RleParts.FEMALE_HAIR, heads);

        bytes[] memory beards = new bytes[](2);
        beards[0] = beardRle;
        beards[1] = beardRle;
        owner.addRles(RleParts.BEARD, beards);

        bytes[] memory bodies = new bytes[](2);
        bodies[0] = bodyRle;
        bodies[1] = bodyRle;
        owner.addRles(RleParts.BODY, bodies);
    }

    function testCanMintFromDope() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000ff';
        bytes4 color = hex'fafafaff';

        uint256 hustlerId = 500;
        alice.mintFromDope(OTHER_BAG, name, background, color);
        ItemIds memory ids = swapmeet.ids(OTHER_BAG);
        checkOwns1155s(ids, address(hustler));
        checkIsEquipped(ids, hustlerId);

        assertEq(hustler.getMetadata(hustlerId).background, background);
        assertEq(hustler.getMetadata(hustlerId).color, color);
        assertEq(hustler.getMetadata(hustlerId).name, name);

        hustler.tokenURI(hustlerId);
    }

    function testCanMintOGFromDope() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000ff';
        bytes4 color = hex'fafafaff';

        uint256 hustlerId = 0;
        alice.mintOGFromDope{ value: 250000000000000000 }(OTHER_BAG, name, background, color, '');
        ItemIds memory ids = swapmeet.ids(OTHER_BAG);
        checkOwns1155s(ids, address(hustler));
        checkIsEquipped(ids, hustlerId);

        assertEq(hustler.getMetadata(hustlerId).background, background);
        assertEq(hustler.getMetadata(hustlerId).color, color);
        assertEq(hustler.getMetadata(hustlerId).name, name);

        hustler.tokenURI(hustlerId);
    }

    function testFailMintOGFromDopeWithLessEth() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000ff';
        bytes4 color = hex'fafafaff';

        alice.mintOGFromDope{ value: 250000000000000000 - 1 }(OTHER_BAG, name, background, color, '');
    }

    function testFailMintOGFromDopeWithMoreEth() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000ff';
        bytes4 color = hex'fafafaff';

        alice.mintOGFromDope{ value: 250000000000000000 + 1 }(OTHER_BAG, name, background, color, '');
    }

    function testCanMintOG() public {
        uint256 hustlerId = 0;
        alice.mintOG{ value: 250000000000000000 }('');
        hustler.tokenURI(hustlerId);
    }

    function testFailMintOGWithLessEth() public {
        alice.mintOG{ value: 250000000000000000 - 1 }('');
    }

    function testFailMintOGWithMoreEth() public {
        alice.mintOG{ value: 250000000000000000 + 1 }('');
    }

    function testCanSetAccessory() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000ff';
        bytes4 color = hex'fafafaff';

        uint256 id = 500;
        alice.mintFromDope(OTHER_BAG, name, background, color);
        ItemIds memory ids = swapmeet.ids(OTHER_BAG);
        checkOwns1155s(ids, address(hustler));
        checkIsEquipped(ids, id);

        alice.safeTransferFrom(address(alice), address(hustler), ACCESSORY, 1, abi.encode(equip, id));
        hustler.tokenURI(id);
    }

    function testCanMintThenTransferHustler() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000';
        bytes4 color = hex'fafafa';

        uint256 hustlerId = 500;
        alice.mintFromDope(OTHER_BAG, name, background, color);
        ItemIds memory ids = swapmeet.ids(OTHER_BAG);
        checkOwns1155s(ids, address(hustler));
        checkIsEquipped(ids, hustlerId);

        alice.safeTransferHustlerFrom(address(alice), address(bob), hustlerId, 1, '');

        assertEq(hustler.balanceOf(address(bob), hustlerId), 1);
    }

    function testCanUnequipHustler() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000';
        bytes4 color = hex'fafafa';

        uint256 hustlerId = 500;
        alice.mintFromDope(OTHER_BAG, name, background, color);
        ItemIds memory ids = swapmeet.ids(OTHER_BAG);
        checkOwns1155s(ids, address(hustler));
        checkIsEquipped(ids, hustlerId);

        // 0000 0001 1111 1111
        assertEq(hustler.getMetadata(hustlerId).mask, bytes2(hex'01ff'));
        uint8[] memory slots = new uint8[](1);
        slots[0] = 0;
        alice.unequip(hustlerId, slots);
        // 0000 0001 1111 1110
        assertEq(swapmeet.balanceOf(address(alice), ids.weapon), 1);
        assertEq(bytes32(hustler.getMetadata(hustlerId).mask), bytes32(bytes2(hex'01fe')));

        uint8[] memory slots2 = new uint8[](5);
        slots2[0] = 1;
        slots2[1] = 2;
        slots2[2] = 3;
        slots2[3] = 4;
        slots2[4] = 5;
        alice.unequip(hustlerId, slots2);

        // 0000 0001 1100 0000
        assertEq(bytes32(hustler.getMetadata(hustlerId).mask), bytes32(bytes2(hex'01c0')));
        assertEq(swapmeet.balanceOf(address(alice), ids.clothes), 1);
        assertEq(swapmeet.balanceOf(address(alice), ids.vehicle), 1);
        assertEq(swapmeet.balanceOf(address(alice), ids.waist), 1);
        assertEq(swapmeet.balanceOf(address(alice), ids.foot), 1);
        assertEq(swapmeet.balanceOf(address(alice), ids.hand), 1);

        alice.safeTransferFrom(address(alice), address(hustler), ids.weapon, 1, abi.encode(equip, hustlerId));
        assertEq(bytes32(hustler.getMetadata(hustlerId).mask), bytes32(bytes2(hex'01c1')));
        assertEq(swapmeet.balanceOf(address(alice), ids.weapon), 0);
        assertEq(swapmeet.balanceOf(address(hustler), ids.weapon), 1);
    }

    function testFailMintFromDopeWithoutApproval() public {
        alice.mintFromDope(OTHER_BAG, '', '', '');
    }

    function testCanSetName() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body;
        uint8[4] memory viewbox;

        alice.setMetadata(id, 'hustler', '', '', '', viewbox, body, hex'0001');

        assertEq(hustler.getMetadata(id).name, 'hustler');
    }

    function testSettingNameWithoutMaskIsNoop() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body;
        uint8[4] memory viewbox;

        alice.setMetadata(id, 'hustler', '', '', '', viewbox, body, hex'0000');

        assertEq(hustler.getMetadata(id).name, '');
    }

    function testFailCantSetLongName() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body;
        uint8[4] memory viewbox;

        alice.setMetadata(id, 'hustler name', '', '', '', viewbox, body, hex'0001');
    }

    function testCanSetBackground() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body;
        uint8[4] memory viewbox;

        alice.setMetadata(id, '', '', hex'123456', '', viewbox, body, hex'0004');

        assertEq(hustler.getMetadata(id).background, hex'123456');
    }

    function testSettingBackgroundWithoutMaskIsNoop() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body;
        uint8[4] memory viewbox;

        alice.setMetadata(id, '', '', hex'123456', '', viewbox, body, hex'0000');

        assertEq(hustler.getMetadata(id).background, '');
    }

    function testCanSetColor() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body;
        uint8[4] memory viewbox;

        alice.setMetadata(id, '', hex'123456', '', '', viewbox, body, hex'0002');

        assertEq(hustler.getMetadata(id).color, hex'123456');
    }

    function testSettingColorWithoutMaskIsNoop() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body;
        uint8[4] memory viewbox;

        alice.setMetadata(id, '', hex'123456', '', '', viewbox, body, hex'0000');

        assertEq(hustler.getMetadata(id).color, '');
    }

    function testCanSetBodyPartial() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body = [1, 0, 12, 0];
        uint8[4] memory viewbox;

        alice.setMetadata(id, '', '', '', '', viewbox, body, hex'0050');

        assertEq(hustler.getMetadata(id).body[0], 1);
        assertEq(hustler.getMetadata(id).body[1], 0);
        assertEq(hustler.getMetadata(id).body[2], 12);
        assertEq(hustler.getMetadata(id).body[3], 0);
    }

    function testCanSetBodyFull() public {
        uint256 id = 500;
        alice.mint();
        uint8[4] memory body = [2, 3, 4, 5];
        uint8[4] memory viewbox;

        alice.setMetadata(id, '', '', '', '', viewbox, body, hex'00f0');

        assertEq(hustler.getMetadata(id).body[0], 2);
        assertEq(hustler.getMetadata(id).body[1], 3);
        assertEq(hustler.getMetadata(id).body[2], 4);
        assertEq(hustler.getMetadata(id).body[3], 5);
    }

    function testFailCantSetMetadataOfUnowned() public {
        uint256 id = 500;
        hustler.mint('');
        uint8[4] memory body = [2, 3, 4, 5];
        uint8[4] memory viewbox;

        alice.setMetadata(id, '', '', '', '', viewbox, body, hex'00f0');
    }

    function testCanSetSlotsPartial() public {
        uint256 id = 500;
        alice.mint();
        ItemIds memory ids = swapmeet.ids(BAG);
        uint256[] memory items = new uint256[](2);
        items[0] = ids.weapon;
        items[1] = ids.vehicle;
        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 1;
        amounts[1] = 1;

        alice.safeBatchTransferFrom(address(alice), address(hustler), items, amounts, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], ids.weapon);
        assertEq(hustler.getMetadata(id).slots[1], 0);
        assertEq(hustler.getMetadata(id).slots[2], ids.vehicle);
        assertEq(hustler.getMetadata(id).slots[3], 0);
        assertEq(hustler.getMetadata(id).slots[4], 0);
        assertEq(hustler.getMetadata(id).slots[5], 0);
        assertEq(hustler.getMetadata(id).slots[6], 0);
        assertEq(hustler.getMetadata(id).slots[7], 0);
        assertEq(hustler.getMetadata(id).slots[8], 0);
        assertEq(hustler.getMetadata(id).slots[9], 0);
    }

    function testCanSetSlotsFull() public {
        uint256 id = 500;
        alice.mint();
        ItemIds memory ids = swapmeet.ids(BAG);

        uint256[] memory items = new uint256[](9);
        items[0] = ids.weapon;
        items[1] = ids.clothes;
        items[2] = ids.vehicle;
        items[3] = ids.waist;
        items[4] = ids.foot;
        items[5] = ids.hand;
        items[6] = ids.drugs;
        items[7] = ids.neck;
        items[8] = ids.ring;

        uint256[] memory amounts = new uint256[](9);
        amounts[0] = 1;
        amounts[1] = 1;
        amounts[2] = 1;
        amounts[3] = 1;
        amounts[4] = 1;
        amounts[5] = 1;
        amounts[6] = 1;
        amounts[7] = 1;
        amounts[8] = 1;

        alice.setApprovalForAll(address(hustler), true);
        alice.safeBatchTransferFrom(address(alice), address(hustler), items, amounts, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], ids.weapon);
        assertEq(hustler.getMetadata(id).slots[1], ids.clothes);
        assertEq(hustler.getMetadata(id).slots[2], ids.vehicle);
        assertEq(hustler.getMetadata(id).slots[3], ids.waist);
        assertEq(hustler.getMetadata(id).slots[4], ids.foot);
        assertEq(hustler.getMetadata(id).slots[5], ids.hand);
        assertEq(hustler.getMetadata(id).slots[6], ids.drugs);
        assertEq(hustler.getMetadata(id).slots[7], ids.neck);
        assertEq(hustler.getMetadata(id).slots[8], ids.ring);
        assertEq(hustler.getMetadata(id).slots[9], 0);
        checkOwns1155s(ids, address(hustler));
    }

    function testSettingSlotsFullReturnsExisting() public {
        uint256 id = 500;
        alice.mint();
        ItemIds memory ids = swapmeet.ids(BAG);

        uint256[] memory items = new uint256[](9);
        items[0] = ids.weapon;
        items[1] = ids.clothes;
        items[2] = ids.vehicle;
        items[3] = ids.waist;
        items[4] = ids.foot;
        items[5] = ids.hand;
        items[6] = ids.drugs;
        items[7] = ids.neck;
        items[8] = ids.ring;

        uint256[] memory amounts = new uint256[](9);
        amounts[0] = 1;
        amounts[1] = 1;
        amounts[2] = 1;
        amounts[3] = 1;
        amounts[4] = 1;
        amounts[5] = 1;
        amounts[6] = 1;
        amounts[7] = 1;
        amounts[8] = 1;

        alice.setApprovalForAll(address(hustler), true);
        alice.safeBatchTransferFrom(address(alice), address(hustler), items, amounts, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], ids.weapon);
        assertEq(hustler.getMetadata(id).slots[1], ids.clothes);
        assertEq(hustler.getMetadata(id).slots[2], ids.vehicle);
        assertEq(hustler.getMetadata(id).slots[3], ids.waist);
        assertEq(hustler.getMetadata(id).slots[4], ids.foot);
        assertEq(hustler.getMetadata(id).slots[5], ids.hand);
        assertEq(hustler.getMetadata(id).slots[6], ids.drugs);
        assertEq(hustler.getMetadata(id).slots[7], ids.neck);
        assertEq(hustler.getMetadata(id).slots[8], ids.ring);
        assertEq(hustler.getMetadata(id).slots[9], 0);
        checkOwns1155s(ids, address(hustler));

        alice.open(OTHER_BAG);
        ItemIds memory otherIds = swapmeet.ids(OTHER_BAG);
        uint256[] memory otherItems = new uint256[](9);
        otherItems[0] = otherIds.weapon;
        otherItems[1] = otherIds.clothes;
        otherItems[2] = otherIds.vehicle;
        otherItems[3] = otherIds.waist;
        otherItems[4] = otherIds.foot;
        otherItems[5] = otherIds.hand;
        otherItems[6] = otherIds.drugs;
        otherItems[7] = otherIds.neck;
        otherItems[8] = otherIds.ring;
        alice.safeBatchTransferFrom(address(alice), address(hustler), otherItems, amounts, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], otherIds.weapon);
        assertEq(hustler.getMetadata(id).slots[1], otherIds.clothes);
        assertEq(hustler.getMetadata(id).slots[2], otherIds.vehicle);
        assertEq(hustler.getMetadata(id).slots[3], otherIds.waist);
        assertEq(hustler.getMetadata(id).slots[4], otherIds.foot);
        assertEq(hustler.getMetadata(id).slots[5], otherIds.hand);
        assertEq(hustler.getMetadata(id).slots[6], otherIds.drugs);
        assertEq(hustler.getMetadata(id).slots[7], otherIds.neck);
        assertEq(hustler.getMetadata(id).slots[8], otherIds.ring);
        assertEq(hustler.getMetadata(id).slots[9], 0);
        checkOwns1155s(ids, address(alice));
        checkOwns1155s(otherIds, address(hustler));
    }

    function testBatchSettingSlotsSameSlotReturnsFirst() public {
        uint256 id = 500;
        alice.mint();
        alice.open(OTHER_BAG);
        ItemIds memory ids = swapmeet.ids(BAG);
        ItemIds memory otherIds = swapmeet.ids(OTHER_BAG);

        uint256[] memory items = new uint256[](2);
        items[0] = ids.weapon;
        items[1] = otherIds.weapon;

        uint256[] memory amounts = new uint256[](2);
        amounts[0] = 1;
        amounts[1] = 1;

        alice.setApprovalForAll(address(hustler), true);
        alice.safeBatchTransferFrom(address(alice), address(hustler), items, amounts, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], otherIds.weapon);
        assertEq(swapmeet.balanceOf(address(hustler), otherIds.weapon), 1);
        assertEq(swapmeet.balanceOf(address(alice), ids.weapon), 1);
    }

    function testFailBatchSettingSlotsOfUnowned() public {
        uint256 id = 500;
        hustler.mint('');
        alice.open(OTHER_BAG);
        ItemIds memory ids = swapmeet.ids(BAG);

        uint256[] memory items = new uint256[](1);
        items[0] = ids.weapon;

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 1;

        alice.setApprovalForAll(address(hustler), true);
        alice.safeBatchTransferFrom(address(alice), address(hustler), items, amounts, abi.encode(equip, id));
    }

    function testFailBatchSettingSlotWithDuplicate() public {
        uint256 id = 500;
        alice.mint();
        alice.claim(UZI_BAG_1);
        alice.open(UZI_BAG_1);
        alice.claim(UZI_BAG_2);
        alice.open(UZI_BAG_2);

        ItemIds memory ids = swapmeet.ids(UZI_BAG_1);

        uint256[] memory items = new uint256[](1);
        items[0] = ids.weapon;

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 2;

        alice.setApprovalForAll(address(hustler), true);
        alice.safeBatchTransferFrom(address(alice), address(hustler), items, amounts, abi.encode(equip, id));
        assertEq(hustler.getMetadata(id).slots[0], 0);
        assertEq(swapmeet.balanceOf(address(alice), ids.weapon), 2);
    }

    function testCanSetSingleSlot() public {
        uint256 id = 500;
        alice.mint();
        ItemIds memory ids = swapmeet.ids(BAG);

        alice.safeTransferFrom(address(alice), address(hustler), ids.weapon, 1, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], ids.weapon);
        assertEq(swapmeet.balanceOf(address(hustler), ids.weapon), 1);
    }

    function testFailSettingSlotWithDuplicate() public {
        uint256 id = 500;
        alice.mint();
        alice.claim(UZI_BAG_1);
        alice.open(UZI_BAG_1);
        alice.claim(UZI_BAG_2);
        alice.open(UZI_BAG_2);

        ItemIds memory ids = swapmeet.ids(UZI_BAG_1);

        alice.setApprovalForAll(address(hustler), true);
        alice.safeTransferFrom(address(alice), address(hustler), ids.weapon, 2, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], 0);
        assertEq(swapmeet.balanceOf(address(alice), ids.weapon), 2);
    }

    function tesFailtCantSetSlotOfUnownedHustler() public {
        uint256 id = 500;
        hustler.mint('');
        ItemIds memory ids = swapmeet.ids(BAG);
        alice.safeTransferFrom(address(alice), address(hustler), ids.weapon, 1, abi.encode(equip, id));
    }

    function testSettingSlotReturnsExistingItem() public {
        uint256 id = 500;
        alice.mint();
        ItemIds memory ids = swapmeet.ids(BAG);

        alice.safeTransferFrom(address(alice), address(hustler), ids.weapon, 1, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], ids.weapon);
        assertEq(swapmeet.balanceOf(address(hustler), ids.weapon), 1);

        alice.open(OTHER_BAG);
        ItemIds memory otherIds = swapmeet.ids(OTHER_BAG);

        alice.safeTransferFrom(address(alice), address(hustler), otherIds.weapon, 1, abi.encode(equip, id));

        assertEq(hustler.getMetadata(id).slots[0], otherIds.weapon);
        assertEq(swapmeet.balanceOf(address(hustler), otherIds.weapon), 1);
        assertEq(swapmeet.balanceOf(address(alice), ids.weapon), 1);
    }

    function testCanBuildTokenURIWithNoRLEs() public {
        bytes[] memory bodies = new bytes[](1);
        bodies[0] = bodyRle;
        bytes[] memory hairs = new bytes[](1);
        hairs[0] = hairRle;
        bytes[] memory beards = new bytes[](1);
        beards[0] = beardRle;

        owner.addRles(RleParts.BODY, bodies);
        owner.addRles(RleParts.MALE_HAIR, hairs);
        owner.addRles(RleParts.BEARD, beards);

        uint256 id = 500;
        alice.mint();
        ItemIds memory ids = swapmeet.ids(BAG);

        uint256[] memory items = new uint256[](9);
        items[0] = ids.weapon;
        items[1] = ids.clothes;
        items[2] = ids.vehicle;
        items[3] = ids.waist;
        items[4] = ids.foot;
        items[5] = ids.hand;
        items[6] = ids.drugs;
        items[7] = ids.neck;
        items[8] = ids.ring;

        alice.setApprovalForAll(address(hustler), true);

        hustler.tokenURI(id);
    }

    function testCanRenderCarMode() public {
        alice.setDopeApprovalForAll(address(hustler), true);

        string memory name = 'gangsta';
        bytes4 background = hex'000000ff';
        bytes4 color = hex'fafafaff';

        uint256 hustlerId = 500;
        alice.mintFromDope(OTHER_BAG, name, background, color);

        uint8[4] memory body;
        uint8[4] memory viewbox;
        alice.setMetadata(hustlerId, '', '', '', hex'0001', viewbox, body, hex'0000');
        assertEq(hustler.getMetadata(hustlerId).options, hex'0001');

        hustler.tokenURI(hustlerId);
    }

    function checkOwns1155s(ItemIds memory ids, address who) private {
        assertEq(swapmeet.balanceOf(who, ids.weapon), 1);
        assertEq(swapmeet.balanceOf(who, ids.clothes), 1);
        assertEq(swapmeet.balanceOf(who, ids.vehicle), 1);
        assertEq(swapmeet.balanceOf(who, ids.waist), 1);
        assertEq(swapmeet.balanceOf(who, ids.foot), 1);
        assertEq(swapmeet.balanceOf(who, ids.hand), 1);
        assertEq(swapmeet.balanceOf(who, ids.drugs), 1);
        assertEq(swapmeet.balanceOf(who, ids.neck), 1);
        assertEq(swapmeet.balanceOf(who, ids.ring), 1);
    }

    function checkIsEquipped(ItemIds memory ids, uint256 hustlerId) private {
        assertEq(hustler.getMetadata(hustlerId).slots[0], ids.weapon);
        assertEq(hustler.getMetadata(hustlerId).slots[1], ids.clothes);
        assertEq(hustler.getMetadata(hustlerId).slots[2], ids.vehicle);
        assertEq(hustler.getMetadata(hustlerId).slots[3], ids.waist);
        assertEq(hustler.getMetadata(hustlerId).slots[4], ids.foot);
        assertEq(hustler.getMetadata(hustlerId).slots[5], ids.hand);
        assertEq(hustler.getMetadata(hustlerId).slots[6], ids.drugs);
        assertEq(hustler.getMetadata(hustlerId).slots[7], ids.neck);
        assertEq(hustler.getMetadata(hustlerId).slots[8], ids.ring);
        assertEq(hustler.getMetadata(hustlerId).slots[9], 0);
    }

    function assertMetadata(
        uint256 hustlerId,
        Attribute[] memory attributes,
        string memory name
    ) private {
        string memory meta = hustler.uri(hustlerId);
        string[] memory inputs = new string[](3);
        inputs[0] = 'node';
        inputs[1] = 'scripts/metadata.js';
        inputs[2] = meta;
        bytes memory res = hevm.ffi(inputs);
        Data memory data = abi.decode(res, (Data));
        assertEq(data.name, name);
        // for (uint256 i = 0; i < attributes.length; i++) {
        //     assertEq(data.attributes[i], attributes[i]);
        // }
    }
}

contract Benchmark is HustlerTest {
    function setUp() public override {
        super.setUp();
        alice.setDopeApprovalForAll(address(hustler), true);
    }

    function testMintFromDope() public {
        alice.mintFromDope(OTHER_BAG, 'gangsta', hex'000000ff', hex'fafafaff');
    }
}