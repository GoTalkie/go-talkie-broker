# go-talkie-broker
Broker is responsible for broking information on server side

# creation of build images
docker buildx build -t gotalkie/go-talkie-broker --push  --platform=linux/amd64,linux/arm64,linux/arm/v7,linux/arm/v6 .