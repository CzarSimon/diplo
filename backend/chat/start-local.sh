DB_NAME=diplo-db
APP_NAME=diplo-chat
docker rm -f $APP_NAME

docker run -d --name $APP_NAME \
  -p 1902:1902 --network diplo-net \
  -e DIPLO_CHAT_DB_NAME=chat \
  -e DIPLO_CHAT_DB_USER=chat \
  -e DIPLO_CHAT_DB_HOST=$DB_NAME \
  -e DIPLO_CHAT_DB_PASSWORD=password \
  -e DIPLO_CHAT_DB_PORT=5432 \
  -e DIPLO_CHAT_SERVER_PORT=1901 \
  -e JWT_SECRET=secrettoken \
  diplo/chat:latest
