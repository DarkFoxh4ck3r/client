#!/usr/bin/env bash

set -e -u -o pipefail # Fail on error

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $dir

# Original sources
client_go_dir="$GOPATH/src/github.com/keybase/client/go"
kbfs_dir="$GOPATH/src/github.com/keybase/kbfs"

# Our custom gopath for iOS build
GOPATH="/tmp/goios"
echo "Using temp GOPATH: $GOPATH"

# Clear source
rm -rf "$GOPATH/src/github.com/keybase"

# Copy source
go_client_dir="$GOPATH/src/github.com/keybase/client/go"
mkdir -p "$go_client_dir"
go_kbfs_dir="$GOPATH/src/github.com/keybase/kbfs"
mkdir -p "$go_kbfs_dir"

echo "Copying client..."
cp -R "$client_go_dir"/* "$go_client_dir"
echo "Copying KBFS..."
cp -R "$kbfs_dir"/* "$go_kbfs_dir"

# Move all vendoring up a directory to github.com/keybase/vendor
echo "Re-vendoring..."
mkdir -p "$GOPATH/src/github.com/keybase/vendor"
# Remove client vendored in kbfs
rm -rf "$go_kbfs_dir/vendor/github.com/keybase/client/go"
# Vendoring client over kbfs (ignore time)
rsync -pr --ignore-times "$go_kbfs_dir/vendor" "$GOPATH/src/github.com/keybase"
rsync -pr --ignore-times "$go_client_dir/vendor" "$GOPATH/src/github.com/keybase"
# Remove their vendoring
rm -rf "$go_kbfs_dir/vendor"
rm -rf "$go_client_dir/vendor"


if [ ! -f "$GOPATH/bin/gomobile" ]; then
  echo "Build gomobile..."
  go get -u golang.org/x/mobile/cmd/gomobile
  gomobile init
fi

package="github.com/keybase/client/go/loopback"

if [ "$1" = "ios" ]; then
  ios_dest="$dir/ios/keybase.framework"
  echo "Building for iOS ($ios_dest)..."
  "$GOPATH/bin/gomobile" bind -target=ios -tags="ios" -o "$ios_dest" "$package"
elif [ "$1" = "android" ]; then
  android_dest="$dir/android/keybaselib/keybaselib.aar"
  echo "Building for Android ($android_dest)..."
  "$GOPATH/bin/gomobile" bind -target=android -tags="android" -o "$android_dest" "$package"
else
  echo "Nothing to build, you need to specify 'ios' or 'android'"
fi