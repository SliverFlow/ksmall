# 使用一个轻量级的基础镜像，如alpine
FROM alpine:latest

# 设置环境变量
ENV APP_NAME app
ENV APP_DIR /data

# 创建应用目录
RUN mkdir -p ${APP_DIR}
RUN mkdir -p ${APP_DIR}/etc

# 将编译好的二进制文件复制到镜像中
COPY ${APP_NAME} ${APP_DIR}/${APP_NAME}
COPY etc ${APP_DIR}/etc

# 给予执行权限（如果需要）
RUN chmod +x ${APP_DIR}/${APP_NAME}

# 设置工作目录
WORKDIR ${APP_DIR}

# 启动应用程序
CMD ["./app -env=dev"]

