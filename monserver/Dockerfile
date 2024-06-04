# 表示依赖 alpine 最新版
FROM alpine:latest
ENV VERSION 1.0
# 在容器根目录 创建一个 apps 目录
WORKDIR /apps
# 挂载容器目录
VOLUME ["/apps/etc"]
# 拷贝当前目录下 go_docker_demo1 可以执行文件
COPY dist/go_docker_demo1_linux_amd64/go_docker_demo1 /apps/golang_app
# 拷贝配置文件到容器中
COPY etc/config-dev.yaml /apps/etc/config-dev.yaml
# 设置时区为上海
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone
# 设置编码
ENV LANG C.UTF-8
# 暴露端口
EXPOSE 9000
# 运行golang程序的命令
ENTRYPOINT ["/apps/golang_app"]