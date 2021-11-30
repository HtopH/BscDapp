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

    uint128 public  spendTickets;
    uint128 public  joinBase = 10000000000000000000;
    uint128 private baseNum =100000000000000000000000;
    uint64  private WEI = 1000000000000000000;
    uint32  private percentWei = 1000000;
    uint32  public  round =0;

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
    mapping (uint32=>mapping(address=>uint128)) roundInfo;

    event registerLog(uint8 doType,uint64 id,address indexed addr,uint64 refId);
    event buyTicketLog(uint8 doType,uint64 id,uint128 _value,uint128 getTicket,uint32 percent);
    event joinLog(uint8 doType,uint64 id,uint128 _value,uint32 _round);
    event userGetLog(uint8 doType,uint64 id,uint128 _value);
    event transferLog(uint8 doType,address fromAddr,address toAddr,uint128 _value);

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
        require(ref_addr!=msg.sender,"ref_addr can not be self!");
        require(addrToId[msg.sender]==0,"user is exist!");
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
        emit registerLog(1,playerId,msg.sender,refId);
    }

    function payForTickets(uint128 _value) external {
        require(addrToId[msg.sender]!=0,"please register first");
        uint64 playerId=addrToId[msg.sender];
        Player storage this_player=players[playerId];
        USD_TOKEN.transferFrom(msg.sender,address(this),_value);
        players[playerId].useU=this_player.useU+_value;

        uint32 percent=getPercent();
        uint128 ticketNum= _value.mul(percent).div(percentWei);
        Ticket_TOKEN.transfer(msg.sender,ticketNum);
        players[playerId].getTickets=this_player.getTickets+ticketNum;

        spendTickets+=ticketNum;
        emit buyTicketLog(2,playerId,_value,ticketNum,percent);
    }

    function joinGame (uint128 _value) external {
        require(addrToId[msg.sender]!=0,"please register first");
        require(roundInfo[round][msg.sender]==0,"you have joined!");
        require(_value>=joinBase,"amount is too low!");
        require(USD_TOKEN.balanceOf(msg.sender)>=_value,"balance is not enough!");

        uint128 pay_ticket=_value.mul(10).div(100);
        require(Ticket_TOKEN.balanceOf(msg.sender)>=pay_ticket,"ticket is not enough!");

        Ticket_TOKEN.transferFrom(msg.sender,address(this),pay_ticket);
        USD_TOKEN.transferFrom(msg.sender,address(this),_value);

        uint64 playerId=addrToId[msg.sender];
        Player storage this_player=players[playerId];
        players[playerId].useTickets=this_player.useTickets+pay_ticket;
        players[playerId].useU=this_player.useU+_value;
        roundInfo[round][msg.sender]=_value;

        emit joinLog(3,playerId,_value,round);

    }

    function userWithdraw(uint128 _value) external {
        require(addrToId[msg.sender]!=0,"user is not exist!");
        require(_value <= USD_TOKEN.balanceOf(address(this)), "Contract have not enough balance!");
        emit userGetLog(4,addrToId[msg.sender],_value);
    }

    function transferToUser(address _addr,uint128 _value) external{
        require(Ticket_TOKEN.balanceOf(msg.sender)>=_value,"ticket is not enough!");
        Ticket_TOKEN.transferFrom(msg.sender,_addr,_value);
        emit transferLog(5,msg.sender,_addr, _value);
    }

    function pay(uint64 _id,uint128 _value) external onlyOperator{
        require(idToAddr[_id]!=address(0),"User is not exist!");
        require(_value <= USD_TOKEN.balanceOf(address(this)), "Contract have not enough balance.");
        USD_TOKEN.transfer(idToAddr[_id],_value);
        Player storage this_player=players[_id];
        players[_id].balanceU=this_player.balanceU.add(_value);
    }

    function adminWithdraw(uint128 val, uint8 _token_type) external onlyAdmin returns (bool) {
        require(_token_type <3 && _token_type >0, "Token type error");
        if(_token_type == 1){
            require(val <= Ticket_TOKEN.balanceOf(address(this)), "Not enough balance.");
            Ticket_TOKEN.transfer(address(uint160(ADMIN_ADDR)),val);
        }else if(_token_type == 2){
            require(val <= USD_TOKEN.balanceOf(address(this)), "Not enough balance.");
            USD_TOKEN.transfer(address(uint160(ADMIN_ADDR)),val);
        }
        return true;
    }

    function getUserInfo(uint64 _id) external view onlyOperator returns (uint128 _getTickets,
            uint128 _useTickets,
            uint128 _balanceU,
            uint128 _useU,
            uint128 _refId){
            require(idToAddr[_id]!=address(0),"user is no exit!");
            Player storage this_player=players[_id];
            _getTickets=this_player.getTickets;
            _useTickets=this_player.useTickets;
            _balanceU=this_player.balanceU;
            _useU=this_player.useU;
            _refId=this_player.refId;
    }

    function userOut(uint64 userId,uint32 _round) external onlyOperator returns(bool){
        require(idToAddr[userId]!=address(0));
        roundInfo[_round][idToAddr[userId]]=0;
    }

    function setSpend (uint128 _value) external onlyOperator returns (bool){
        spendTickets=_value;
        return true;
    }


    function setRound() external onlyOperator returns (bool) {
        round=round+1;
        return true;
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
        }else if(n>1&&n<11){
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

    function getBase(uint32 n) public pure returns(uint32){
        uint128 base=10;
        if (n>10){
            uint128 b=uint128((n-11)/9+2);
            base=base.power(b);
        }
        return uint32(base);
    }

    function getNo() public view returns(uint32 n){
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
    function power(uint128 a,uint128 b)internal pure returns (uint128 c){
        if (a==0) return 0;
        if (b==0) return 1;
        c = 1;
        for (uint128 i=0;i<b;i++){
            c=mul(c,a);
        }
    }
}

