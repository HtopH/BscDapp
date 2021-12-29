package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/3yXZ1ATXJ/FQxGCBEGKRDrSpDcBpQkGkBB6Cx0pgsADSCIgGnp9KAJBgoAU6R1RpEtvRgGpoSORaqSGGkp2nN1332dm39nz4d65c+89c76cmd/fGEZFzQoAAoAAo9AcM8A/xAqgAyADndzdHyOk/2eX8kT6+liYXwFQlER/c/ptb2/GAiN1HJsVnZCl5yu0JpADC/e8cj7alDZI9M0mVEtW5mxGhX/TbygtJPiYREtzOQiGxbcBb3dp+7HTPU/JEOzMgI7o15xftZXndtZVztx0RXeCpSU2NjeDSJtLREJpXkjLKS5ag49pY4wnTJe6oCZXAzcJy2iNqDc37faEG4VN1cGoLbo0Oi2X8+2Vrco4qJqMVOwbOXuLfKDXl++37FAwxITdLjDWTguYBPayRddbF8X7UdhIkj/hy+2pep1NoAAr0/TtEAOLFrGKC0mrGSmaVA5vjg73xG+2gZIeftY0G++K732o18bEIgagFyYVCSZsmpIlYhwTb4I/Vdp9RjX/grRGABUooLKVrVMih3WMbojuwbaJJQlZdWuUAQspH+Sp6I07C6XusWwl4yS73xfpTgXW+1izOR2MJRfE1dj2wD/9rbmWhh0bF1oK4QRZPkyuDScqegvG7vfQ/Xw6Htt5CmUJoeT4iz7glsbRvQDMGElnrLGISJ+y7O8lOBfJEwcJFhvLdxxNIzdpCl9DJeTet3bBKa7wnQD4ZKMmrYRQzhVrtOcb8q39+Cb10259w/e7LFTAQSaI+cW4Qkokb7/EVwKaFSS5U09jg9XAGIXIWvJMvtCk+LAQANJ6VP0CveBgpb2vEEbB+TDy93H2mLyaQrpRyKij9+rDkvkDo7TLmedkuy78W2QfN0OBbUPfUUUdhqhtnjQtsGyUJXNDbmXXggh80KNreovL8nVonYFwzGbEUBKbtRe8zIMrxrxQOCI0GQaMt1ejSMR/F8pXFl43x3og5DSopiS4k0lTnwXOkr+Yzp7vvk6UOsF+HudrkabbHWMwsPBhUx5FAETFMNQ0AagHs25ikBK3yodCs9oLDRoUFeetVt0m05AYpUGiX/BjwQ5lsoKKZh6b3hffHldiJ8MXxs5Y7gjp9jUqKm2ZWx7T17zOxrWQM7iV07bLM+XT6r+/SvuwFDvvkaz8GAp4tevWnAddOF07lkDpG2kEN2flDEOi55UumVJTu1nTWB8u2L1NHZYUjqQYV0k8PNR0pE34J4G0XbbUHe1OdZ8pth5srkUehg9kk76F7wWN6yduthaufbcjukkKSS39TKhqTi6fiVzJTeXMKbErb3Z5qaDH5c1YSWVMF/7EvBiDt6Xh5eT3zU3tE1BTSFN5McXvN2CK0/hwjvN5x3iVip8wYNyBjd/IlQWzWTIRoB9U12s+m8pUCxLJmXxmycizNJJhTGJttmPuL/GZhLUpFhe3HP1kjs1jJXMovwudjdbSCYQ24FUx6n2CBtdVmG4PPygcdXQRG/MoY44X+3jxZoIKaUXZuyzIgZpMY+9OqPt4OFgG9+ItYqc8fUUzWouOmeviCkCzDlnuHkc9NtxOLbwXfJ6uYDRTTs8+1MOAfrP1TXe4+LzpgQwKTY2UMt1t68Rwrwt4fdRZq27OPGCr0lPy7nm21NDXbKHt6gQFh91lTQf7JHW85ZY7tzOqysq2NTDI/xFIo+xjODDkNIF0fuXw9enc2sYjalZJ0OixBtfackwzsR3CVU2wjr0Lrg3S4boy7s3uP82v1sBOEJ/Nlz470gq7Uhxwt1oRNTziJMjIp0XV3U7DrPZ2Qk7msBK9T5+jMnyXanSwQ71Zp2Sx2av2Gbzr2ghq6BV62QRJgPqQWe6FPBB280A+ow2fetfK4PjRhf0w01xKPnT96JPDvnPEtrQ8RHkL8aIMdZq6fjq5+kNNFAE1fSrTcee1yhsldd7Ug5dKl90LZ3wc63tgz4nfv1Ua5fVuUDGGWm7aAOG/hhzX/tpGYbn/Wl2fSvp0yKyy9cZ9QJX+cPhtUG5e23KMRNoaAy73CX37PCpev0Mkb+3LGAn324qo33jZZpNlwjmxbjm48UCq9QHR5SCdCnlet3cSJr58JPXja9MWWK46cA10KriIq7TZbhkZyswS8w/x2K2VWVBc4K3p8cQdZC4of1i29c78fWxfFrVE+w4ytac5TW5X+RXvxXSh3R18i5KxnR4/VlqS9KVaz7do92ASkzutykb5nrAOz4CSgs2CmUPMSgpuzt3cX7MeSWdisEzpUyKd2Ox6xlUN/ZoTd/XGJPrShgLZ5GkGTxrEWRpXPXOiG1Mo6SE2OZSRzRdaIR6Db6xxD4919zs0VeieSmbYi0ZVM7LvV36tqZ/mx47X5ogX1IcqN794r/XKDB+Ybm8L1ywQ0bDYOYRQZIgO3BoDa1lVyIGPLOKAwGn1HLh/E1Ynct/VILR0vrJi4mpCzpYxc8btCfEDh9dNwvt0zp+/54MnnzJL0EI2LZ7x/NREC3xzS2DDHvpKG3/Bp8oqmvpnCP64Da+7waQ9B9Q4cCsv0tdxTrLtEBNV4mfO7z6SoJS4m1F0aTY7vGU0Zsyr5sjcgjhFjCShMdwtVnfrV/UeHYZ3R8sH1gHc39e9zQSb994fEGfZSVDA0t4/S03U3+SEXkxi0vhEiGUvoU8ezXyScd2RH5IrFFgXZBwBmzjgo9nSI7HmnSY/9Ix2yhvaWE25w7AKt3Xl4lowshlBD494J+0y70E2Kc8a9RZJN8+L6PrORXV+mbrMC/EBZar0L6Usy+gT0DiG/k4Ij0DhaHLW6V3/ZaF+unVTvpeitApSPEmsZTPIpF4vBv/m5elc0qGX4VLtKxDOEkna/450kld8Kd9wksCGFUF/gwWfVqs3ZufNlTsWVssshvCB5s0tjkGwKTeiwMjhFltpyB4fW7lf/ykJ/xNx/fdX/5z2kWIahgepMIb+WOrTO1t1HHdAT9v6xT7axezZB2+YbOglB7U9ERZBnOVFwF8OLrrHq4t0NR1EWoAKcOln7/XCxDs5SnoyXCXeMT5+UtQH5eHRnsHPO/MmiVk6x7sYbqeOJv/p++ko7G+8bJ3JYnPnZTVao2MhjyAbJFZ1pwQ/Hq+N7D75sSFMPCHIhp3Nhy46FGS+FEWsACWl+JkIR034nwj3/mthUxfBOEgI4kbdzcIQRELJnDo7HpUrjeMq3UYszHGurjuhLe6+1fyFwTj0I/7UvPTSoljMJ+oCqxo5Cu5+s64cb2Q0SwMM94Rf+EPB2El+T7gfEtP4wuG2wPCXojJdBkQZPddREJVvdn3pgM7ss+SsbssQs+dQy59chHs7OKNXUStVujqVTkVYu+tyLVe21JZeGCKtVn+xGyjttjv2HTCTalScytAJ2ZGu04ssXFAVurVwhhys9fFsjAR7psgNzUBVj4mTWwssj4vLqIcxBtphIN/JupiinQm13moZMeir+ZW9hM2la+8W68xGnlfpfWAPVTUu+Kp+VUgE+pyvl3vjY+PJ0E9tdlmnmneaTAOp3YXMckeygU0aV8SmwZPoUsX+mzplcbwpqXdmeK5ePrqwEqcoUK+e5OQ95NE91LqemPO6ObOo8fVzUZiXyycQBY/u6O2R/JRy1s+2kQUzUEw/++2cVd8ETAdlREEvrHY7LmLgVlsedGQYVFIxeh/h2SyoNgQzYQIh9fC35GscoDXIfPJKVPjAnbhHb2j5ownbYx618jRMoKV+VS9RGPfiQ8h1UN65LBj3gbVkULiLGsJHTGSythsa6fDuBu6dyUZ/XySU9jE4ioQ3sboEyb9t0I/Fnml6rvrorg7BeVSAHyqeVnUcFI/arX2HZqoqWiVn61XXaiqlI4HT7GyKJWcf8W/580FNNx7koTRLWAzz1gWyA24GcoM3Hl88Hup/9P2SIYo5Tp5z4y9e1NkAbRa6dQAVm7VPrUmGPETzZyy1oIVmROVkCT7cXyEzkjp+m9pW7T175d7q5MimK4NOPEN7YKUIsSljg6AUf7cu/eKKJYqNgjmQRTwAfv+zeR1toCrkxhKrqaedgjTEoJKuHNGSmaV4mRMQWBOp4rFOeRG6G0oBAJDJxjBaYO/O92w/egDgczQA8C92B/wfdqf9N7v/L67/+f3PN8YwCkpWqn+z/z+d/7D/v9QQ/mf9fyeBf1v95yj/LSYAWbONHvAfgl2h+XNPCaAE/A0AAEj0f07/FQAA//+5MSDpmQwAAA=="); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
