sudo docker build -t dorm-system-backend .
# Windows and Mac
#sudo docker run -dp 8080:8080 --name dorm-system-backend --add-host=host.docker.internal:host-gateway dorm-system-backend
# Linux
sudo docker run -dp 8080:8080 --name dorm-system-backend --add-host=host.docker.internal:172.17.0.1 dorm-system-backend
