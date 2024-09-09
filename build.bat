@echo off
setlocal

REM 设置代码页为65001（UTF-8）
chcp 65001

REM 设置临时目录和打包文件名
set TEMP_DIR=temp_deploy
set TAR_FILE=deployment_tally_gva.tar

REM 设置前端和后端代码目录
set FRONTEND_DIR=web
set BACKEND_DIR=server

REM 设置远程服务器信息
set REMOTE_USER=root
set REMOTE_HOST=47.106.160.34
set REMOTE_PASS=Qq4956225
set REMOTE_PATH=/opt/docker/admin-server
set REMOTE_SHELL_FILE=start.sh



REM 创建临时目录
echo 正在创建临时目录...
mkdir %TEMP_DIR%

REM 前端和后端打包
cd /d %~dp0%FRONTEND_DIR%
call pnpm install
call pnpm build 
cd /d %~dp0

cd /d %~dp0%BACKEND_DIR%
call go env -w GO111MODULE=on 
call go env -w GOPROXY=https://goproxy.cn,direct
call go env -w CGO_ENABLED=0 
call go env 
call go mod tidy
call go build -o %BACKEND_DIR% .
cd /d %~dp0

REM 复制前端和后端代码到临时目录
echo 正在复制前端代码...
xcopy /E /I %FRONTEND_DIR%\dist %TEMP_DIR%\%FRONTEND_DIR%

echo 正在复制后端代码...
copy %BACKEND_DIR%\%BACKEND_DIR% %TEMP_DIR%\%BACKEND_DIR%
copy %BACKEND_DIR%\%REMOTE_SHELL_FILE% %TEMP_DIR%\%BACKEND_DIR%
xcopy /E /I %BACKEND_DIR%\resource %TEMP_DIR%\%BACKEND_DIR%\resource

REM 打包前端和后端代码
echo 正在打包前端和后端代码...
tar -cvf %TAR_FILE% -C %TEMP_DIR% .

REM 发送打包文件到远程服务器
echo 正在将打包文件发送到远程服务器...
scp %TAR_FILE% %REMOTE_USER%@%REMOTE_HOST%:%REMOTE_PATH%
@REM sshpass -p %REMOTE_PASS% scp -o StrictHostKeyChecking=no %TAR_FILE% %REMOTE_USER%@%REMOTE_HOST%:%REMOTE_PATH%

REM 登录到云服务器，加载镜像并启动服务

REM 删除本地的临时目录和打包文件
echo 正在删除本地的临时目录和打包文件...
rmdir /S /Q %TEMP_DIR%
del %TAR_FILE%


echo 部署完成！
endlocal
