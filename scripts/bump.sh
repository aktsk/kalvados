#!/bin/sh

set -e

echo current version: $(gobump show -r ./version)

read -p "input next version: " next_version

gobump set $next_version -w ./version
ghch -w -N v$next_version

git commit -am "Checking in changes prior to tagging of version v$next_version"
git tag v$next_version
git push && git push --tags
