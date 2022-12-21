sudo docker build -t dorm-system-backend .
# Not working
#sudo docker run -dp 8080:8080 --name dorm-system-backend --add-host=host.docker.internal:host-gateway dorm-system-backend
sudo docker run -dp 8080:8080 --name dorm-system-backend dorm-system-backend
