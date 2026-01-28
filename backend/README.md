# Build the image
docker build -t partiburo-builder .

# Create a temporary container to pull the binary out
docker create --name temp-container partiburo-builder
docker cp temp-container:/app/partiburo ./partiburo
docker rm temp-container

rm ~/projects/infra/scripts/backend/partiburo
cp partiburo ~/projects/infra/scripts/backend