#!/bin/zsh

hugo -d ./s3
#for f in $(ls s3/*.html s3/**/*.html)
#do
#    mv $f "${f%%.*}"
#done

s3cmd sync --guess-mime-type --delete-removed s3/ s3://blog.owulveryck.info
