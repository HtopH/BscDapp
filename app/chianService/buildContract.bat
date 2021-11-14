@echo off


echo *********************begin publish cont*********************
echo 工作目录:%cd%

set srcfolder=%cd%
set abigen=D:\Code_world\Golang_code\src\github.com\ethereum\go-ethereum\cmd\abigen\abigen.exe



rem abigen --abi xx.abi --pkg pkgname --type apiname --out xx.go
rem 1. abi 文件在 remix 部署时可以得到
rem 2. Pkg 指定的是编译成的 go 文件对应的 package 名称
rem 3. type指定的是go文件的入口函数，可以认为是类名
rem 4. out 指定输出go文件名称


rem abi myContract.abi	指定abi文件来源
rem pkg chianService	指定输出文件的包名
rem type bscGame	指定合约结构体名称
rem out bscGame.go	指定合约交互文件名称


%abigen% --abi %srcfolder%\myContract.abi --pkg chianService --type bscGame --out bscGame.go

pause