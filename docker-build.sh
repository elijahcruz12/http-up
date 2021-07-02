docker run --name="tasker_builder" --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:latest ./build.sh
