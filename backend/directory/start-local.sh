DB_NAME=diplo-db
APP_NAME=diplo-directory
docker rm -f $APP_NAME

docker run -d --name $APP_NAME \
  -p 1901:1901 --network diplo-net \
  -e DIPLO_DIRECTORY_DB_NAME=directory \
  -e DIPLO_DIRECTORY_DB_USER=directory \
  -e DIPLO_DIRECTORY_DB_HOST=$DB_NAME \
  -e DIPLO_DIRECTORY_DB_PASSWORD=password \
  -e DIPLO_DIRECTORY_DB_PORT=5432 \
  -e DIPLO_DIRECTORY_SERVER_PORT=1901 \
  -e JWT_SECRET=secrettoken \
  diplo/directory:latest
