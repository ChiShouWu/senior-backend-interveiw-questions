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
5. Is redis single thread?
    <details><summary>Answer</summary>
    Yes, but it is multi-process. It is single-threaded per process.
    </details>
6. How redis be consistent?
7. How to implement redis cluster?
8. [How redis rehash](https://github.com/doocs/advanced-java/blob/main/docs/high-concurrency/redis-rehash.md)
9.  [What will happen if wrong usage of redis?](https://github.com/doocs/advanced-java/blob/main/docs/high-concurrency/why-cache.md)
10. Why sentinels? How it work?
11. Why redis cluster? How it work?
12. What is redis Cache Avalanche, Hotspot Invalid, Cache Penetration? 
    <details><summary>Answer</summary>
    Cache avalanche: When lot of cache expires. This will cause a large number of requests to the database in same time.

      - We can give each key a random expiration time, so that the cache expires at different times.
    
    Hotspot invalid: When a key is frequently accessed and it is expired, this will cause a large number of requests to the database in same time. 
    
      1. We can use a distributed lock to prevent multiple requests from accessing the database at the same time. But this will cause slow response.
      2. Set the key to non-expired.

    Cache penetration: When a data is not in cache, the request will directly go through to the database.
    
      - Filter the request if we know this data should not exist in the database.
      - Insert an empty value into the cache for the key .
    </details>
13. What is the disadvantage of redis?
14. How to resolve redis shutdown?
15. How redis recycle progress work?
16. How redis implement queuing delay? When to use?
17. How to know nodes in cluster are working?
18. Redis is sync replica or async replica?
19. If main slave data is different, how to resolve?
20. If main node is down, what will happen?
21. Further, if main node is back to normal, what will happen?
22. What is brain split?
23. 

## Reference
- [小林redis](https://xiaolincoding.com/redis/)