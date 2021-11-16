#!/usr/bin/env

docker pull busk3r/genericuniversity:latest
docker run --name genericuniversity -itd --rm -p 8000:8000 busk3r/genericuniversity && docker exec genericuniversity service mysql start && docker exec genericuniversity mysql -u root -p -e "ALTER USER 'root'@'localhost' IDENTIFIED BY 'password';"
docker exec genericuniversity php /root/Generic-University/artisan serve --host 0.0.0.0

curl http://localhost:8000
