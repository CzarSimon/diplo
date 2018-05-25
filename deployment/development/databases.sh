DB_NAME=diplo-db
docker rm -f $DB_NAME

docker run -d --name $DB_NAME \
  -p 5432:5432 --network diplo-net \
  -e POSTGRES_PASSWORD=password \
  postgres:10.2-alpine

echo "Installing databases"
sleep 3
docker exec -i $DB_NAME psql -U postgres < databases.sql

echo "Installing schema"
sleep 1
docker exec -i $DB_NAME psql -U chat chat < ../../backend/chat/database/1_baseline.sql
