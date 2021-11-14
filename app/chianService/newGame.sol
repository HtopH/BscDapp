pragma solidity 0.6.4;

interface IBEP20 {

    function totalSupply() external view returns (uint256);
    function decimals() external view returns (uint8);
    function symbol() external view returns (string memory);
    function name() external view returns (string memory);
    function getOwner() external view returns (address);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address recipient, uint256 amount) external returns (bool);
    function allowance(address _owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address sender, address recipient, uint256 amount) external returns (bool);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

}

contract NewGame {
    using SafeMath for uint128;

    IBEP20 private SYS_TOKEN;
    IBEP20 private TUSD_TOKEN;


    address constant private ADMIN_ADDR = 0x125a0daEE26BD73B37A3c2a86c84426c68743750;
    address private op_addr = 0x125a0daEE26BD73B37A3c2a86c84426c68743750;
    address private owner = 0x125a0daEE26BD73B37A3c2a86c84426c68743750;

    address public SYS_ADDR = 0x843c95480D6BC6E23C33adb0ec5D2246CBA8E1E9;
    address public TUSD_ADDR=0xD47636fb8bdBAF9AF3c90bc913800e3bf72E7cB4;

    event join(uint64 id, address indexed addr);

    event paylog(uint64 id,uint128 _value,uint8 _type);

    struct Player {
        address ADDR;
        uint128 SYS;
        uint128 TUSD;
    }

    Player[] private players;
    mapping (address => uint64) public addrToId;
    mapping (uint64 => address) public idToAddr;

     modifier onlyAdmin() {
        require(msg.sender == ADMIN_ADDR);
        _;
    }

    modifier onlyOperator() {
        require(msg.sender == op_addr);
        _;
    }

    constructor() public {
        SYS_TOKEN = IBEP20(SYS_ADDR);
        TUSD_TOKEN = IBEP20(TUSD_ADDR);

        Player memory _player = Player({
            ADDR:owner,
            SYS : 0,
            TUSD : 0
        });
        players.push(_player);
        players.push(_player);
        uint64 playerId = uint64(players.length - 1);
        addrToId[owner] = playerId;
        idToAddr[playerId] =owner;
    }


    function register() public{
        require(addrToId[msg.sender]==0,"accout have register");
        Player memory _player = Player({
            ADDR:msg.sender,
            SYS : 0,
            TUSD : 0
        });
        players.push(_player);
        uint64 playerId = uint64(players.length - 1);
        addrToId[msg.sender] = playerId;
        idToAddr[playerId] = msg.sender;
        emit join(playerId,msg.sender);
    }

    function pay (uint128 _value,uint8 _type) public {
        uint64 playerId=addrToId[msg.sender];
        require(idToAddr[playerId]!=address(0),'There is no user');
        Player storage this_player=players[playerId];

        if(_type==1){

            require(SYS_TOKEN.balanceOf(msg.sender)>_value,"You have not enough mount");
            SYS_TOKEN.transferFrom(msg.sender,address(this),uint256(_value));
            players[playerId].SYS = this_player.SYS.add(_value);
        }else if(_type==2){

            require(TUSD_TOKEN.balanceOf(msg.sender)>_value,"You have not enough mount");
            TUSD_TOKEN.transferFrom(msg.sender,address(this),uint256(_value));
            players[playerId].TUSD = this_player.TUSD.add(_value);

        }
        emit paylog(playerId,_value,_type);
    }

    function get (uint8 _type) public {
        require(_type <3 && _type >0, "type error");
        uint64 playerId=addrToId[msg.sender];
        require(idToAddr[playerId]!=address(0),'There is no user');
        Player storage this_player=players[playerId];
        if(_type==1){
            require(this_player.SYS>0,"There is no Amount");
            players[playerId].SYS = 0;
           if (!SYS_TOKEN.transfer(msg.sender,uint256(this_player.SYS))){
               players[playerId].SYS =this_player.SYS;
           }
        }else if(_type==2){
            require(this_player.TUSD>0,"There is no Amount");
            players[playerId].TUSD = 0;
            TUSD_TOKEN.transfer(msg.sender,uint256(this_player.TUSD));
            if (!TUSD_TOKEN.transfer(msg.sender,uint256(this_player.TUSD))){
               players[playerId].TUSD =this_player.TUSD;
           }

        }

    }


    function withdrawAdmin(uint256 val, uint8 _token_type) public onlyAdmin {
        require(_token_type <3 && _token_type >0, "Token type error");
        if(_token_type == 1){
            require(val <= SYS_TOKEN.balanceOf(address(this)), "Not enough balance.");
            SYS_TOKEN.transfer(address(uint160(ADMIN_ADDR)),val);
        }else if(_token_type == 2){
            require(val <= TUSD_TOKEN.balanceOf(address(this)), "Not enough balance.");
            TUSD_TOKEN.transfer(address(uint160(ADMIN_ADDR)),val);
        }
    }

    function set_op_addr(address opAddr) external onlyAdmin{
        require(opAddr!=address(0),"Address error");
        op_addr = opAddr;
    }

    function get_balance_info() external view returns (uint256 SYS_balance,uint256 TUSD_balance){
        SYS_balance = SYS_TOKEN.balanceOf(address(this));
        TUSD_balance = TUSD_TOKEN.balanceOf(address(this));
        return(SYS_balance,TUSD_balance);
    }

    function get_player_count() external view
    returns(uint64){
        return uint64(players.length - 1);
    }

}

library SafeMath {
    function add(uint128 x, uint128 y) internal pure returns (uint128 z) {
        require((z = x + y) >= x, 'ds-math-add-overflow');
    }

    function sub(uint128 x, uint128 y) internal pure returns (uint128 z) {
        require((z = x - y) <= x, 'ds-math-sub-underflow');
    }

    function mul(uint128 x, uint128 y) internal pure returns (uint128 z) {
        require(y == 0 || (z = x * y) / y == x, 'ds-math-mul-overflow');
    }
    function div(uint128 a, uint128 b) internal pure returns (uint128 c) {
        require(b > 0,'ds-math-div-overflow');
        c = a / b;
    }
}

