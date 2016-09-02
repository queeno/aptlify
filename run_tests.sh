#!/usr/bin/env bash

repo_present() {
    return `aptly repo list|grep ${1} >/dev/null`
}

create_repos() {
    for testing_repo in aptlify_testing_1 aptlify_testing_2;do
        if ! repo_present ${testing_repo};then
            aptly repo create ${testing_repo}
        fi
    done
}

drop_repos() {
    for testing_repo in aptlify_testing_1 aptlify_testing_2;do
        if repo_present ${testing_repo};then
            aptly repo drop ${testing_repo}
        fi
    done
}


create_repos
aptlify dump
drop_repos

go test -v github.com/queeno/aptlify/utils -check.v
