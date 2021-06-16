#!/usr/bin/env bash

package="main.go"

package_name="up"

if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64" "linux/386" "linux/arm64" "linux/arm")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })

    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}

    output_name=$package_name'-'$GOOS'-'$GOARCH

    output_folder="build/"

    output_total=$output_folder$output_name

    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    echo "Building for ${GOOS} with arch ${GOARCH}"

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_total $package

    echo "Completed building ${output_name}"

done

echo "All builds complete"
