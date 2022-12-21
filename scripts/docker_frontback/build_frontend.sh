cd ./web/dorm-system-frontend || return
sudo docker build -t dorm-system-frontend .
# Windows and Mac
#sudo docker run -dp 3000:3000 --name dorm-system-frontend --add-host host.docker.internal:host-gateway dorm-system-frontend
# Linux
sudo docker run -dp 3000:3000 --name dorm-system-frontend --network="host" dorm-system-frontend
