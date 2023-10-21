#!/bin/bash

# 执行每个文件的pb3编译，参数为一级目录的文件夹
function docompile() {
    for file in ${1}/*.proto
    do
        if test -f ${file}
        then
            echo "   compile start: ${file}"
            echo "   " | protoc -I=. --go_out=paths=source_relative:./ --micro_out=paths=source_relative:./ ${file}
            echo "   compile end: ${file}"
        fi
    done
}

# 编译proto文件
function build() {
    for file in ./*
    do
        if test -d ${file}
        then
            echo "---build pkg start: ${file}"
            docompile ${file}
            echo "---build pkg end: ${file}"
        fi
    done

    # 生成mock文件
    mockery --name '(.+)Service' --recursive --inpackage --keeptree --case underscore
}

build