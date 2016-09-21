#!/usr/bin/env bash
set -e
if [[ $# -ne 1 ]];then
   echo "Please specify one version number (e.g. 1.0.0)"
   exit 1
fi
gom install
gom build -o $GOPATH/bin/aptlify
build_dir=/tmp/fpm/aptlify
[[ -d ${build_dir} ]] && rm -rf ${build_dir}
mkdir -p ${build_dir}/usr/bin
cp ${GOPATH}/bin/aptlify ${build_dir}/usr/bin
fpm -s dir -t deb -n aptlify -v ${1} -C ${build_dir}
