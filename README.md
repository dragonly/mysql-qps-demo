# MySQL QPS Visualizer

This is a demo of using mysql's PERFORMANCE_SCHEMA table as a table source and visualizing the qps.


# demo

https://user-images.githubusercontent.com/8433465/151288041-73af2da2-65e1-4ffe-be57-b2a52446f40e.mp4


# Instructions

1. start a mysql instance, probably using docker

```bash
# start mysql using docker
docker run -d -e MYSQL_ROOT_PASSWORD=pass --name mysql -p 3306:3306 mysql
```

2. run sysbench workload

```bash
# create test database
mysql -u root -ppass -h 127.0.0.1 -e "create database if not exists sbtest;"
# insert test data
sysbench oltp_point_select --mysql-host=127.0.0.1 --mysql-user=root --mysql-password=pass --tables=1 --table-size=10000 --db-ps-mode=disable prepare
# run point get query workload
sysbench oltp_point_select --mysql-host=127.0.0.1 --mysql-user=root --mysql-password=pass --tables=1 --table-size=10000 --time=0 --report-interval=3 --db-ps-mode=disable run
```

3. start demo server

```bash
make run
```

4. watch the qps visualization

visit http://localhost:8080 to see the visualization live.
