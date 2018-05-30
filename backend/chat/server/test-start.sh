export DIPLO_CHAT_DB_NAME=chat
export DIPLO_CHAT_DB_USER=chat
export DIPLO_CHAT_DB_HOST=localhost
export DIPLO_CHAT_DB_PASSWORD=password
export DIPLO_CHAT_DB_PORT=5432

export DIPLO_CHAT_SERVER_PORT=1902

export JWT_SECRET=secrettoken

echo "Building DIPLO_CHAT_SERVER"
go build
./server
