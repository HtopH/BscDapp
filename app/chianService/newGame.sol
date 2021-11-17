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

    IBEP20 private Ticket_TOKEN;
    IBEP20 private USD_TOKEN;

    address constant private ADMIN_ADDR = 0x125a0daEE26BD73B37A3c2a86c84426c68743750;
    address private op_addr = 0x125a0daEE26BD73B37A3c2a86c84426c68743750;

    uint128 public seedPool;
    uint128 public jackPor;
    uint128 public spendTickets;
    uint128 private baseNum =100000000000000000000000;
    uint64  private WEI = 1000000000000000000;
    uint32  private percentWei = 1000000;
    uint16  public round =1;

    struct Player {
        uint128 getTickets;
        uint128 useTickets;
        uint128 balanceU;
        uint128 useU;
        uint64  refId;
    }

    Player[] private players;
    mapping (address => uint64) public addrToId;
    mapping (uint64 => address) public idToAddr;
    mapping (uint16=>mapping(address=>uint128)) roundInfo;

    event registerLog(uint64 id,address indexed addr,uint64 refId);
    event buyTicketLog(address indexed addr,uint128 _value,uint128 getTicket,uint32 percent);
    event joinLog(address indexed addr,uint128 _value,uint128 ticketUse);

     modifier onlyAdmin() {
        require(msg.sender == ADMIN_ADDR,"You are not Admin");
        _;
    }

    modifier onlyOperator() {
        require(msg.sender == op_addr,"You are not operator");
        _;
    }

    constructor() public {
        Ticket_TOKEN = IBEP20(0x843c95480D6BC6E23C33adb0ec5D2246CBA8E1E9);
        USD_TOKEN = IBEP20(0xD47636fb8bdBAF9AF3c90bc913800e3bf72E7cB4);

        Player memory _player = Player({
            getTickets:0,
            useTickets:0,
            balanceU:0,
            useU:0,
            refId:0
        });
        players.push(_player);
        players.push(_player);
        uint64 playerId = uint64(players.length - 1);
        addrToId[0x125a0daEE26BD73B37A3c2a86c84426c68743750] = playerId;
        idToAddr[playerId] =0x125a0daEE26BD73B37A3c2a86c84426c68743750;
    }

    function register(address ref_addr) external {
        require(addrToId[msg.sender]==0,"user is exit!");
        uint64 refId;

        if (addrToId[ref_addr]!=0){
            refId=addrToId[ref_addr];
        }else{
            refId=1;
        }

        Player memory _player = Player({
            getTickets:0,
            useTickets:0,
            balanceU:0,
            useU:0,
            refId:refId
        });
        players.push(_player);
        uint64 playerId = uint64(players.length - 1);
        addrToId[msg.sender] = playerId;
        idToAddr[playerId] =msg.sender;
        emit registerLog(playerId,msg.sender,addrToId[ref_addr]);
    }

    function payForTickets(uint128 _value) external {
        require(addrToId[msg.sender]!=0,"please register first");
        uint64 playerId=addrToId[msg.sender];
        Player storage this_player=players[playerId];
        USD_TOKEN.transferFrom(msg.sender,address(this),_value);
        players[playerId].useU=this_player.useU+_value;

        uint32 percent=getPercent();
        uint128 ticketNum= _value.mul(percent).div(percent);
        Ticket_TOKEN.transfer(msg.sender,ticketNum);
        players[playerId].getTickets=this_player.getTickets+ticketNum;

        spendTickets+=ticketNum;
        jackPor+=_value.mul(25).div(100);
        seedPool+=_value.mul(25).div(100);
        Player storage ref_player=players[this_player.refId];
        players[this_player.refId].balanceU=ref_player.balanceU.add(_value.mul(20).div(100));

        emit buyTicketLog(msg.sender,_value,ticketNum,percent);
    }

    // function joinGame (uint128 _value) external {

    //     uint64 playerId=addrToId[msg.sender];

    // }


    function withdrawAdmin(uint256 val, uint8 _token_type) external onlyAdmin {
        require(_token_type <3 && _token_type >0, "Token type error");
        if(_token_type == 1){
            require(val <= Ticket_TOKEN.balanceOf(address(this)), "Not enough balance.");
            Ticket_TOKEN.transfer(address(uint160(ADMIN_ADDR)),val);
        }else if(_token_type == 2){
            require(val <= USD_TOKEN.balanceOf(address(this)), "Not enough balance.");
            USD_TOKEN.transfer(address(uint160(ADMIN_ADDR)),val);
        }
    }

    function setRound() external onlyOperator {
        round=round+1;
        uint128 half= seedPool.mul(50).div(100);
        jackPor=half;
        seedPool-=half;
    }

    function set_op_addr(address opAddr) external onlyAdmin{
        require(opAddr!=address(0),"Address error");
        op_addr = opAddr;
    }

    function getPercent() public view returns (uint32 percent) {
        uint32 n=getNo();
        uint32 base=getBase(n);
        if (n==1){
            percent=percentWei;
        }else if(n<1&&n<11){
            percent=uint32(percentWei-percentWei*(n-1)/base);
        }else{
            if ((n-10)%9>0){
                percent=uint32(percentWei*10/base-percentWei*((n-10)%9)/base);
            }else{
              percent=uint32(percentWei*10/base-percentWei*9/base);
            }
        }
        return percent;
    }

    function getBase(uint32 n) public pure returns(uint32 base){
        base=10;
        if (n>10){
            base=10**((n-11)/9+2);
        }
        return base;
    }

    function getNo() internal view returns(uint32 n){
        if (spendTickets>0){
            n=uint32((spendTickets.sub(1)).div(baseNum).add(1));
        }else{
            n=1;
        }
        return n;
    }

    function get_balance_info() external view returns (uint256 ticketBalance,uint256 usdBalance){
        ticketBalance = Ticket_TOKEN.balanceOf(address(this));
        usdBalance = USD_TOKEN.balanceOf(address(this));
        return(ticketBalance,usdBalance);
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

