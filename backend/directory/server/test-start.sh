export DIPLO_DIRECTORY_DB_NAME=directory
export DIPLO_DIRECTORY_DB_USER=directory
export DIPLO_DIRECTORY_DB_HOST=localhost
export DIPLO_DIRECTORY_DB_PASSWORD=password
export DIPLO_DIRECTORY_DB_PORT=5432

export DIPLO_DIRECTORY_SERVER_PORT=1901

export JWT_SECRET=secrettoken

echo "Building DIPLO_DIRECTORY_SERVER"
go build
./server
