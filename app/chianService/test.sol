pragma solidity 0.6.4;

contract Purchase {
   using SafeMath for uint128;
   uint128 public totalTickets =18000000000000000000000000;
   uint128 public spendTickets =0;
   uint128 public baseNum =100000000000000000000000;
   uint64  private WEI = 1000000000000000000;
   uint128 private percentWei = 1000000;

   uint128 public n;
   uint128 public base;
   uint128 public num;

    function one (uint128 _value) public {
        if (spendTickets>0){
            n=(spendTickets.sub(1)).div(baseNum).add(1);
        }else{
            n=1;
        }
        if (n>1){
            base=10**((n-1)/10+1);
        }else{
            base=0;
        }
        num=percentWei-(n-1)*percentWei/base;

    }

    function set (uint128 _value) public{
        spendTickets=_value.mul(WEI);
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
