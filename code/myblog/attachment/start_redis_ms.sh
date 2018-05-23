cd /home/caojx/redis/redis_master/bin
./redis-server ../conf/redis.conf 

cd /home/caojx/redis/redis_slave1/bin
./redis-server ../conf/redis.conf 

cd /home/caojx/redis/redis_slave2/bin
./redis-server ../conf/redis.conf 

cd /home/caojx/redis/redis_slave3/bin
./redis-server ../conf/redis.conf 


/home/caojx/redis/redis_master/bin/redis-server /home/caojx/redis/redis_master/conf/sentinel.conf --sentinel

/home/caojx/redis/redis_slave1/bin/redis-server /home/caojx/redis/redis_slave1/conf/sentinel.conf --sentinel

/home/caojx/redis/redis_slave2/bin/redis-server /home/caojx/redis/redis_slave2/conf/sentinel.conf --sentinel

/home/caojx/redis/redis_slave3/bin/redis-server /home/caojx/redis/redis_slave3/conf/sentinel.conf --sentinel
