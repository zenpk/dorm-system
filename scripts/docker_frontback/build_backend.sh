docker build -t dorm-system-backend .
docker run -dp 8080:8080 --name dorm-system-backend --add-host=host.docker.internal:host-gateway dorm-system-backend
