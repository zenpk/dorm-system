cd ./web/dorm-system-frontend || return
docker build -t dorm-system-frontend .
docker run -dp 3000:3000 --name dorm-system-frontend dorm-system-frontend
