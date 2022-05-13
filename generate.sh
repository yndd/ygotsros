#!/bin/bash
# usage: ./generator.sh sros_22.2

go install github.com/openconfig/ygot/generator@v0.20.0
go install golang.org/x/tools/cmd/goimports@latest
git clone -b $1 --single-branch https://github.com/nokia/7x50_YangModels.git nokia > /dev/null 2>&1

#FILES=`find nokia/srlinux-yang-models/srl_nokia \( -iname "*.yang" ! -name "*radio*" \)`
FILES=`find nokia/YANG/nokia-combined/ \( -iname "nokia-conf.yang" \)`

generator -output_dir=. \
    -logtostderr \
    -path=nokia/YANG \
    -package_name=ygotsros -generate_fakeroot -fakeroot_name=device -compress_paths=false \
    -structs_split_files_count 32 \
    -shorten_enum_leaf_names \
    -typedef_enum_with_defmod \
    -enum_suffix_for_simple_union_enums \
    -generate_rename \
    -generate_append \
    -generate_getters \
    -generate_delete \
    -generate_simple_unions \
    -generate_populate_defaults \
    -include_schema \
    -exclude_state \
    -yangpresence \
    -ignore_circdeps \
    -include_model_data \
    -generate_leaf_getters \
    -validate_fn_name SROSValidate \
    $FILES

goimports -w enum.go
goimports -w enum_map.go
goimports -w schema.go
goimports -w structs-*.go
goimports -w union.go
rm -rf nokia/
