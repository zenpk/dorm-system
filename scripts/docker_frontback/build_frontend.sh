cd ./web/dorm-system-frontend || return
sudo docker build -t dorm-system-frontend .
sudo docker run -dp 3000:3000 --name dorm-system-frontend dorm-system-frontend
