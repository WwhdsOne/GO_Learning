#!/bin/bash

# 获取当前目录下的所有 .proto 文件
proto_files=$(find . -maxdepth 1 -name "*.proto")

# 遍历每个 .proto 文件
for proto_file in $proto_files; do
    # 获取文件名（不包括扩展名）
    prefix=$(basename "$proto_file" .proto)

    # 设置输出目录
    output_dir="../gen/$prefix"

    # 创建输出目录（如果不存在）
    mkdir -p "$output_dir"

    # 使用 protoc 生成 Go 代码
    protoc --go_out="$output_dir" --go_opt=paths=source_relative \
           --go-grpc_out="$output_dir" --go-grpc_opt=paths=source_relative \
           "$proto_file"

    echo "Generated Go code for $proto_file in $output_dir"
done