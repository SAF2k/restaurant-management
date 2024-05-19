# restaurant-management

# to login mariadb through docker
docker exec -it <container_id> mysql -u root -p123456 

# to start the db
docker start <container_id>  

# to check if there is anything runing in dockers backgroud 
docker ps -a  

# to check if anything is running in the docker container 
docker ps 

# to create a database in docker 
# it will create a maria 10.6 database in the docker container with the root user 
docker run -d --name mariadbresturent -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 mariadb:10.6


docker exec -it <container_id> bash

CREATE USER 'new_user'@'%' IDENTIFIED BY 'password';


GRANT ALL PRIVILEGES ON *.* TO 'new_user'@'%';


FLUSH PRIVILEGES;



docker exec -it <container_id> mysql -u<user> -p<password>

