if test "$#" -ne 1; then
    echo "Illegal number of parameters"
    exit
fi

VERSION=$1
DATE=$(date '+%Y-%m-%d')
IMG_NAME=us.gcr.io/mibici/back


echo "Building image v$VERSION..."
docker build -t $IMG_NAME:$VERSION \
  --build-arg GITHUB_TOKEN=8c4551b79c8ed1e182d85f8a18cb16eb4d297961 \
  --build-arg ENVIRONMENT=DEV .

echo "Pushing image to Google Cloud..."
gcloud docker -- push $IMG_NAME:$VERSION
gcloud container images add-tag $IMG_NAME:$VERSION $IMG_NAME:$DATE -q
gcloud container images untag $IMG_NAME:latest -q
gcloud container images add-tag $IMG_NAME:$VERSION $IMG_NAME:latest -q