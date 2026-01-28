# Build the image
docker build -t mail-builder .

# Create a temporary container to pull the binary out
docker create --name temp-container mail-builder
docker cp temp-container:/app/mail ./mail
docker rm temp-container

rm ~/projects/infra/scripts/backend/service/mail
cp mail ~/projects/infra/scripts/backend/servicew