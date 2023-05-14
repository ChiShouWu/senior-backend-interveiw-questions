# Redis

1. What is redis?
    <details><summary>Answer</summary>
    A store that is in-memory. Can be used as a cache or queue or pub/sub.
    All execution is atomic and single-threaded.
    It also support transactions, consistent, Lua scripting, and clusters.
    </details>
2. What is difference between redis and memcached?
    <details><summary>Answer</summary>

    - Redis supports more data types, and is more flexible.

    - Redis supports consistent, can store data on disk.
  
    - Redis has original clustering.

    - Redis has Lua scripting, pub/sub, transactions.
    </details>
3. When to use redis?
    <details><summary>Answer</summary>
    - When you need a high performance cache.
    - When you are facing concurrency issues. Redis can easily handle over 10k QPS when 1k QPS is hard to MySql.
    </details>
4. Is redis single thread?
    <details><summary>Answer</summary>
    Yes, but it is multi-process. It is single-threaded per process.
    </details>
5. How redis be consistent?
    <details><summary>Answer</summary>
    Redis have two ways to implement consistency. One is AOF, it store command record into hard disk. Second is RDB, it store data snapshot into hard disk.
    </details>
6. [How redis rehash](https://github.com/doocs/advanced-java/blob/main/docs/high-concurrency/redis-rehash.md)
7.  [What will happen if wrong usage of redis?](https://github.com/doocs/advanced-java/blob/main/docs/high-concurrency/why-cache.md)
8.  Why sentinels? How it work?
    <details><summary>Answer</summary>
    - When we using redis cluster as master-slave mode. If master is down, there will be no master to response the query and we must manually set a slave to master node. So we need a sentinel to monitor the master. Sentinel will promote a slave to master.
    <details>
9.  Why redis cluster? How it work?
    <details><summary>Answer</summary>
    - The cluster provide a way to scale out. We can add more nodes to the cluster. The cluster will automatically distribute the data to the nodes.
    <details>
10. What is redis Cache Avalanche, Hotspot Invalid, Cache Penetration? 
    <details><summary>Answer</summary>
    Cache avalanche: When lot of cache expires. This will cause a large number of requests to the database in same time.

      - We can give each key a random expiration time, so that the cache expires at different times.
      - Use redis cluster to prevent node is down.
      - Use Hystrix like rate limiter to prevent too many requests stream to database.
    
    Hotspot invalid: When a key is frequently accessed and it is expired, this will cause a large number of requests to the database in same time. 
    
      1. We can use a distributed lock to prevent multiple requests from accessing the database at the same time. But this will cause slow response.
      2. Set the key to non-expired.

    Cache penetration: When a data is not in cache, the request will directly go through to the database.
    
      - Filter the request if we know this data should not exist in the database.
      - Insert an empty value into the cache for the key .
    </details>
11. What is the disadvantage of redis?
    <details><summary>Answer</summary>
    
    - It store data in memory, so it is cost.

    - Security is not good.

    - Redis is single-threaded, if there is a large key to be processed, it will block other requests.
    <details>
12. How redis implement queuing delay? When to use?
    <details><summary>Answer</summary>
    We can use `zset` to store data with score. The score is the time when the data should be executed. A thread will check the timestamp of the first data in the zset. If the timestamp is smaller than current time, it will delete the data and remove it from the zset. We can use this to implement queuing delay.
    </details>

13. How to know nodes in cluster are working?
    <details><summary>Answer</summary>
    Redis use  ping-pong to check. If half node to ping one node and the node is not response, the node will be considered down.
    </details>
14. Redis is sync replica or async replica?
15. If main slave data is different, how to resolve?
16. If main node is down, what will happen?
17. Further, if main node is back to normal, what will happen?
18. What is brain split?
19. What kinds of key sharding in redis cluster?
    <details><summary>Answer</summary>
     

## Reference
- [小林redis](https://xiaolincoding.com/redis/)