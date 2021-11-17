pragma solidity 0.6.4;

contract Purchase {
   using SafeMath for uint128;
   uint128 public totalTickets =18000000000000000000000000;
   uint128 public spendTickets =0;
   uint128 public baseNum =100000000000000000000000;
   uint64  private WEI = 1000000000000000000;
   uint32 private percentWei = 1000000;

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

    function getNo() public view returns(uint32 n){
        if (spendTickets>0){
            n=uint32((spendTickets.sub(1)).div(baseNum).add(1));
        }else{
            n=1;
        }
        return n;
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
