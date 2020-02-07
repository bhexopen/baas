#!/bin/bash -eu

error() {
  echo "$1"
  exit 1
}

git diff --exit-code >/dev/null || error "Dirty work tree, commit and try again"
./build.sh
git checkout gh-pages
rm -rf $TMPDIR/apidoc-build
mv build $TMPDIR/apidoc-build
git rm -r .
mv $TMPDIR/apidoc-build/* .
git add .
git diff --staged --exit-code >/dev/null || git commit -am "Build `git log --pretty=format:"%h %s" master | head -n 1`"
git pull --rebase origin gh-pages
git push origin gh-pages
git checkout master
