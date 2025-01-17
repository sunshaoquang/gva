#!/bin/bash

# 设置临时目录和打包文件名
TEMP_DIR=temp_deploy
TAR_FILE=deployment_tally_gva.tar

# 设置前端和后端代码目录
FRONTEND_DIR=web
BACKEND_DIR=server

# 设置远程服务器信息
REMOTE_USER=root
REMOTE_HOST=47.106.160.34
REMOTE_PASS=Qq4956225
REMOTE_PATH=/opt/docker/admin-server
REMOTE_SHELL_FILE=start.sh

# 创建临时目录
echo "正在创建临时目录..."
mkdir -p $TEMP_DIR

# 前端和后端打包
echo "正在打包前端代码..."
cd $FRONTEND_DIR
pnpm install
pnpm build
cd ..

echo "正在打包后端代码..."
cd $BACKEND_DIR
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go env -w CGO_ENABLED=0
go env -w GOARCH=amd64
go mod tidy
go build -o $BACKEND_DIR .
cd ..

# 复制前端和后端代码到临时目录
echo "正在复制前端和后端代码到临时目录..."
mkdir -p $TEMP_DIR/$FRONTEND_DIR
cp -r $FRONTEND_DIR/dist/* $TEMP_DIR/$FRONTEND_DIR

mkdir -p $TEMP_DIR/$BACKEND_DIR

cp -r $BACKEND_DIR/$BACKEND_DIR $TEMP_DIR/$BACKEND_DIR
cp -r $BACKEND_DIR/config.docker.yaml $TEMP_DIR/$BACKEND_DIR
cp $BACKEND_DIR/$REMOTE_SHELL_FILE $TEMP_DIR/$BACKEND_DIR
cp -r $BACKEND_DIR/resource $TEMP_DIR/$BACKEND_DIR

# 打包前端和后端代码
echo "正在打包前端和后端代码..."
tar -cvf $TAR_FILE -C $TEMP_DIR .

# 发送打包文件到远程服务器
echo "正在将打包文件发送到远程服务器..."
# scp $TAR_FILE $REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH
# sshpass -p "$REMOTE_PASS" rsync -avz -e "ssh -o StrictHostKeyChecking=no" -r `pwd` "$REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH"
tar cf - "$TAR_FILE" "$COMPOSE_FILE" | pv | sshpass -p "$REMOTE_PASS" rsync -avz -e "ssh -o StrictHostKeyChecking=no" "$TAR_FILE" "$REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH"

echo "正在登录到远程服务器并启动服务..."
sshpass -p "$REMOTE_PASS" ssh -o StrictHostKeyChecking=no "$REMOTE_USER@$REMOTE_HOST" "
cd $REMOTE_PATH && tar -xvf $TAR_FILE -C . --overwrite && sh $REMOTE_PATH/$BACKEND_DIR/$REMOTE_SHELL_FILE && sh $REMOTE_PATH/reload.sh
"


# 删除本地的临时目录和打包文件
echo "正在删除本地的临时目录和打包文件..."
rm -rf $TEMP_DIR
rm $TAR_FILE

echo "部署完成！"
