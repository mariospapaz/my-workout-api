VERSION=1.0
NAME="my-workout"

echo -e "#####################################\n\n"

echo -e "Starting building process....\n\n"

echo -e "#####################################"

docker buildx create --name ${NAME}

docker buildx use ${NAME}

docker buildx build --platform linux/amd64,linux/arm64 -t mariospapaz/${NAME}:${VERSION} --push .