# Database Questions
- [Database Questions](#database-questions)
  - [Concept](#concept)
  - [RDBMS](#rdbms)
    - [Index](#index)
  - [NoSql](#nosql)
  - [Reference](#reference)
  
---
## Concept
1. What is RDBMS?
    <details><summary>Answer</summary>
    Many data table with relation ship for store data.
    </details>
2. What is Isolation Levels
    <details><summary>Answer</summary>
      Database isolation refers to the ability of a database to allow a transaction to execute as if there are no other concurrently running transactions (even though in reality there can be a large number of concurrently running transactions). The overarching goal is to prevent reads and writes of temporary, aborted, or otherwise incorrect data written by concurrent transactions.
    
      The four levels of isolation are:
      - Read Uncommitted: A transaction can read data that has been modified by other transactions but not yet committed.
      - Read Committed: A transaction can only read data that has been committed by other transactions.
      - Repeatable Read: A transaction can only read data that has been committed by other transactions. It also prevents other transactions from modifying the data that has already been read by the current transaction.
      - Serializable: A transaction can only read data that has been committed by other transactions. It also prevents other transactions from modifying the data that has already been read by the current transaction. It also prevents other transactions from inserting new rows that match the WHERE clause of a query that has already been executed by the current transaction.
    </details>

3. Use mongodb or mysql? Why?
    <details><summary>Answer</summary>
      Actually we can change this question to compare with NoSQL and RDBMS.
      RDBMS is good at ACID if your business need transaction, and you want it to have high performance, you can use RDBMS.
      NoSQL is good at scalability if your business model is not structured will growth fast and become really dynamic. Choose NoSQL is good for you.

    </details>
4. Why do databases treat null as a so special case? For example, why does `SELECT * FROM table WHERE field = null` not match records with null field in SQL?
    <details><summary>Answer</summary>
      NULL means no value, not `zero` or `empty string`. And it have no type in SQL.
      Not `VARCHAR` or `DATE` or others. It will not to equal anything neither itself.
      If you want to check if a field is null, you should use `IS NULL` or `IS NOT NULL` instead of `=` or `<>`. 
    </details>
5. How would you find the most expensive queries in an application?
    <details><summary>Answer</summary>
      If you use SQL server , You can query `dm_exec_query_stats` to get the most expensive query.
      If you use Postgres, you can use an extension module `pg_stat_statements` to get the most expensive query.
      If you use MySQL, you need to capture this information from a log file, and not via a query. `See slow query log`

      Than with `EXPLAIN` you can see the query plan, and you can see the cost of each step. The cost is the number of rows that the step will process. So you can find the most expensive query by the cost of each step. 
    </details>
6. In your opinion, is it always needed to use database normalization? When is it advisable to use demoralized databases?
    <details><summary>Answer</summary>
      Normalization is a process of organizing data in a database. It is used to minimize data redundancy and improve data integrity. It is a technique to reduce the size of the database and increase the speed of the database.
      It is advisable to use demoralized databases when you need to improve the performance of the database cause by lots of join operation.
      `Normalize until it hurts, denormalize until it works.`
    </details>
7.  What do you understand by Data Redundancy?
    <details><summary> Ans </summary>      
      Duplication of data in the database is known as data redundancy. As a result of data redundancy, duplicated data is present at multiple locations, hence it leads to wastage of the storage space and the integrity of the database is destroyed.
    </details>
8.  What do you understand by Data Independence? What are its two types?
    <details><summary> Ans </summary>
      Data Independence refers to the ability to modify the schema definition in one level in such a way that it does not affect the schema definition in the next higher level.

      The 2 types of Data Independence are:

      - Physical Data Independence: It modifies the schema at the physical level without affecting the schema at the conceptual level. User application don't know how the data is stored in the disk, it is own by the database system.
      - Logical Data Independence: It modifies the schema at the conceptual level without affecting or causing changes in the schema at the view level. User application won't be affect by database schema change.
    </details>

9.  Define the relationship between `View’`and `Data Independence`.
    <details><summary> Ans </summary>
      View is a virtual table that does not have its data on its own rather the data is defined from one or more underlying base tables.
      Views account for logical data independence as the growth and restructuring of base tables are not reflected in views. 
    </details>
10. What are the advantages and disadvantages of views in the database?
    <details><summary> Ans </summary>
      Advantages:
        
        1. Views don't store data in a physical location.

        2. The view can be used to hide some of the columns from the table.

        3. Views can provide Access Restriction, since data insertion, update and deletion is not possible with the view.

      Disadvantages:
        
        4. When a table is dropped, associated view become irrelevant.
        
        5. Since the view is created when a query requesting data from view is triggered, its a bit slow.
        
        6. When views are created for large tables, it occupies more memory.
    </summary>
11. What is transaction?
    <details><summary> Ans</summary>
    Transaction is a logic unit in database, means a group of data read & write actions. It will be two results committed(all success) or rollback(all cancel).

    It is common to see in transfer money, you will transfer money from A to B, if A's balance is not enough, you will cancel the transaction, and the money will not transfer to B. To prevent data inconsistency, we need to use transaction.
    </details>
12. Define Database Lock and its types.
    <details><summary> Ans</summary>
    Database lock is a mechanism to prevent data inconsistency. It is used to prevent other transaction from accessing the data that is being accessed by the current transaction.
    - Exclusive Lock: Only one transaction can hold an exclusive lock on a row at a time. Others can't read or write, Until the transaction releases the lock.
  
    - Shared Lock: Multiple transactions can hold a shared lock on a row at a time. Others can read, but can't write, Until the transaction releases the lock. And the same time exclusive lock can't be acquired.

    - Range Lock: It is used to lock a range of rows. Also for InnoDB, it is related to `isolation level`. When isolation level is serializable, it will lock W/R to the range of rows. If isolation level is repeatable read, it will lock write to the range of rows.
    </details>
13. Define Phantom deadlock.
    <details><summary> Ans</summary>
    Phantom deadlock is a deadlock in a distributed DBMS. It is caused by the following conditions:
     - One process is waiting for resource which is being held by another process.
     - When second process release the resource and here comes a delay so no one knows resource is released.
    </details>
14. What is lost update?
    <details><summary> Ans</summary>
    Lost update is a phenomenon that two transactions update the same data, and the last transaction will overwrite the first transaction's update.
    </details>
15. What do you understand by B-Trees?
    <details><summary> Ans</summary>
      B-Trees are a type of self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. B-Trees are a generalization of a binary search tree in that a node can have more than two children.
    </details>
16. What are some common issues with ORMs?
    <details><summary> Ans</summary>
      Props: 
        - Make query sentence more readable. Easy to maintain
        - Prevent SQL injection
        - Models use OOP, which means you an extend and inherit from Models.
      Cons:
        - When you need to do complex query, you need to write raw SQL. ORM may have performance issue.
    </details>

17. What is phantom read?
    <details><summary> Ans</summary>
      Phantom read is a phenomenon two transactions read same data but get different rows.
    </details>

18. What is dirty read?
    <details><summary> Ans</summary>
      Dirty read is a phenomenon that occurs when a transaction reads a data item that has been modified by another uncommitted transaction.

      For example: Transaction A reads a data item, Transaction B modifies the data item, Transaction A reads the data item again, Transaction B update the data item, Transaction A reads the data item again. Transaction A will get the data item that is modified by Transaction B, but not committed.
    </details>

19. What is non-repeatable read?
    <details><summary> Ans</summary>
      Non-repeatable read is a phenomenon that occurs when a transaction re-reads a data item that it has previously read, and finds that the data item's value has been modified by another transaction.

      For example: Transaction A reads a data item, Transaction B modifies the data item, Transaction A reads the data item again. Transaction A will get the data item that is modified by Transaction B, but not committed.
    </details>
20. What is write Skew?
    <details><summary> Ans</summary>
      Write Skew is a phenomenon that occurs when two transactions that each have to make sure conditions are met before they can commit. And both pass the condition check, but when both commit, it will make wrong result.

      For example: Our customers want to buy products which are in stock 20. We have a table to store the stock of products. We have two transactions, Transaction A and Transaction B. Transaction A wants to buy 10 products, Transaction B wants to buy 5 products. Both transactions check the stock of products, and both pass the check. But when both commit, the stock of products will be minus 15, which is wrong.
    </details>

21. Race condition
    <details><summary>Answer</summary>
    - Atomic update
    - Transaction lock
    - Version control
    </details>
22. What is Lost Updates in Mysql, Postgres, how to prevent it?
    <details><summary>Answer</summary>
    - It is a phenomenon that two transactions update the same data, and the last transaction will overwrite the first transaction's update.
    - Use serializable isolation level or repeatable read isolation level
    </details>
23. When many people want to buy a commodity and the stock is only one, how to solve this problem?
    <details><summary>Answer</summary>
    - Use serializable isolation level
    </details>
24. What is optimistic lock?
    <details><summary>Answer</summary>
    Optimistic lock is a mechanism to prevent data inconsistency. When accessing data, it will not lock. So other transaction can access the data. But when commit, it will check if the data is changed. If the data is changed, it will throw an exception.

    - The data I/O is more then pessimistic lock, but the lock time is less.
    - It is more suitable for less resource competition.
    </details>
25. What is pessimistic lock?
    <details><summary>Answer</summary>
    Pessimistic lock is a mechanism to prevent data inconsistency. It is used to prevent other transaction from accessing the data that is being accessed by the current transaction.

    - The data I/O is less then optimistic lock, but the lock time is more.
    - It is more suitable for more resource competition.
    </details>

## RDBMS
1. Explain count(*), count(1), count(column_name) in SQL?
    <details><summary> Ans</summary>
      - count(*) will count all rows, including null.
      - count(1) will count all rows, including null.
      - count(column_name) will count all rows that column_name is not null.
    </details>
2. How to improve count(*)?
    <details><summary> Ans</summary>
      - Add non-clustered index.
    </details>
3. Why MySQL use B+tree?
    <details><summary>Answer</summary>
    B+Tree only store data in leaf node, B-Tree store data in all node. So there is less data in B+Tree, so it is faster to find data cuz less IO read.

    The time complexity of B+Tree is O(log`d`n), d is the number of children of a node, n is the number of data in the tree. In actual situation d is larger than 100, so even data rows comes to 10 millions, the height of the tree is only 3.

    Compare with hash table, the time complexity of hash table is O(1), but the hash table is not sorted, so it is not suitable for range query.
    </details>
4. What is B tree and B+ tree?
    <details><summary>Answer</summary>
    B tree is a self-balancing tree data structure that keeps data sorted and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree is a generalization of a binary search tree in that a node can have more than two children. Unlike self-balancing binary search trees, the B-tree is optimized for systems that read and write large blocks of data. It is commonly used in databases and file-systems.

    B+ tree is a B tree with an additional level at the bottom with linked leaves. It is used to store data in a database or file system. It is a B tree with the following properties:
    - All leaves are on the same level.
    - Every non-leaf node has a pointer to the leftmost leaf node.
    - Only the leaf nodes contain data.
    </details>
5. if users can follow/subscribe topics/favorite commodities, is unlimited / larger than 10 thousands rows, how to design data schema?
    <details><summary>Answer</summary>
    The interviewer want to know is the relation between user and topic/commodity is one-to-one or one-to-many. If it is one-to-one, we can use one table to store the relation. If it is one-to-many, we can use two tables to store the relation. If the relation is one-to-many, we can use one table to store the relation, and use a column to store the number of relation.

    And you will use RDBMS or NoSQL?
    </details>
   
6. How to keep the data consistency in distributed system?
    <details><summary>Answer</summary>
    - Use transaction
    - Use isolation lock, which one is better? why?
    </details>
7. How database process a command?
    <details><summary>Answer</summary>
    There are two layers of DB, the first is server layer. The second is storage engine layer. The server layer is responsible for parsing SQL, generating execution plan, and executing SQL. The storage engine layer is responsible for storing data and index.
    - Parse SQL
    - Generate execution plan
    - Execute SQL
    - Return result
    </details>
8.  What could happen if multiple users create transactions at the same time?
    <details><summary>Answer</summary>
    First, all these users access the same data? Second, what are there transactions doing?
    - If all these users access the same data, and all these transactions are read-only, then there is no problem.
    - If all these users access the same data, and all these transactions are write, then there is a problem. The data will be inconsistent. And we need to use isolation lock to solve this problem.
    </details>
9.  What is MVCC?
    <details><summary>Answer</summary>
    Multi-Version Concurrency Control (MVCC) is a method of concurrency control in which multiple versions of a record can coexist in the database. It is used to solve the problem of dirty read and non-repeatable read. It is used in InnoDB and Postgres.

    It provide a snapshot of the data at a particular time. So the data is consistent. When updated a row of data, it will mark the old row is obsolete, and add a new row with new version. There will be many versions of the same row of data. But only the latest version is visible to the user.
    </details>
10. What is buffer pool?
    <details><summary>Answer</summary>
    Buffer pool is a memory area used to cache data. It is used to improve the performance of database. When a query is executed, the data will be read from buffer pool. If the data is not in buffer pool, it will be read from disk. And the data will be put into buffer pool. So the next time the data is read, it will be read from buffer pool. So it is faster.
    </details>
11. Describe the different of `drop`, `delete`, `truncate`?
    <details><summary>Answer</summary>

        - `drop` will delete the table and the data. It is a DDL command. It will lock the table. It will not release the space of the table.

        - `delete` will delete the data. It is a DML command. It will not lock the table but lock row each execute. It will delete the data row by row and record it. So it is slow.
        
        - `truncate` will delete the data in table. It is a DDL command. It will lock the table. It is faster than `delete` cuz it will not record the deleted data.
    </details>

12. Primary key vs Unique key, what is the difference? which one is faster?
    <details><summary>Answer</summary>
    - Primary key is a unique key, but not allow null.
    - Primary key is faster than unique key.
    - Only one primary key in a table.
    - Both should have same performance.
    </details>

13. Master-slave consistency, how to solve it?
    <details><summary>Answer</summary>
    - semi-sync
    - 
    </detail>


### Index
1. What do you understand by `Index hunting`?
    <details><summary>Answer</summary>
    Index hunting is the process of boosting the collection of indexes which help in improving the query performance as well as the speed of the database.
    </details>
2. What is `index` in database? And how it works?
    <details><summary>Answer</summary>
    Index is a data structure that improves the speed of data retrieval operations on a database table at the cost of additional writes and storage space to maintain the index data structure. Indexes can be created using one or more columns, or expressions. 
    </details>
3. What kind of index do you know?
    <details><summary>Answer</summary>
       - Data structure :B-tree, Hash, R-tree, Bitmap, Full-text, Spatial
       - Physical storage: Clustered Index, Non-clustered Index, secondary index
       - Characteristics: Primary key, Unique key, Non-unique key, Full-text index, Spatial index
  
    Most usually to see are B+Tree, Hash, Full-Text. 
    </details>
4. When to create `index`? When not to use?
    <details><summary>Answer</summary>
    Database will create index defaultly when you create a table. But you can also create index by yourself.

    - If there is primary key, database will create a clustered index on it.

    - If not, database will create a clustered index on non-null and unique column.
    
    - Last, InnoDB will create a clustered index on explicit id row automatically.

    - Disadvantage:
      - Index as B+Tree need to be updated when data is changed, so it will slow down the write speed.
      - Need more physical space to store index.
      - Create/maintain index will slow down the write speed and time spent grow up with data.

    - When to use:
      - Column is unique.
      - Columns are frequently used in `where` clause. You can use union index to combine multiple columns.
      - Columns are frequently used in `order by` and `orderBy` clause. Thus when you use `order by` clause, the data already sorted by index, so it will be faster.

    - When not to use:
      - Columns not frequently used in `where`, `groupBy` or `orderBy` clause.
      - Data is duplicate or null.
      - Data is less.
      - Data frequently changed. This will make index recreate.
    </details> 
5. What is leftmost prefix matching?
    <details><summary>Answer</summary>
    When you create a index on multiple columns, the index will use the leftmost column to sort the data. So if you want to use the index, you must use the leftmost column to query. If you use the other column to query, the index will not be used.
    </details>
6. Differentiate between ‘Cluster’ and ‘Non-cluster’ index.
    <details><summary>Answer</summary>
    
    - Cluster index store real data in the index.
  
    - Non-cluster index store the pointer of data, need to find the data by primary key.
    </details>
7. How to prevent index failure?
    <details><summary>Answer</summary>
    - Using left or right fuzzy query.
    - Use count or functions will make index not work.
    - Union index with leftmost prefix matching.
    - In where clause, use index before `or`, and after `or` is not index.
    </details>
8. Is that speed up the query if we add more index on a table?
    <details><summary>Answer</summary>
    When we read data with index, it will faster. But if we write data, it will slower. Because we need to update the index.
    </details>
9. How full-text index work?
   <details><summary>Answer</summary>
    Full-text index is a special index that is used to search text. It is a kind of inverted index. It will store the word and the document id. When you search the word, it will return the document id. And you can get the document by the document id.
    </details>

## NoSql
1. How MongoDB execute text search？
    <details><summary>Answer</summary>
    MongoDB uses a text index to store the text search data. When you execute a text search, it will use the text index to search the data.
    </details>
2. MongoDB document have 32MB limit, if your article larger then that, what would you do？
    <details><summary>Answer</summary>
    Actually it is 16mb now.
    You can split the article into multiple documents. And use a reference to link them.
    Or use GridFS to store the article.
    </details>
3. What is eventual consistency?
    <details><summary>Answer</summary>
    Eventual consistency is a consistency model in which all replicas will eventually show the same data. It is a consistency model that is weaker than strong consistency. It is used in distributed systems and will enhance the system availability and performance.
    </details>
4. [What are virtual property in mongoose?](https://github.com/Gauthamjm007/Backend-NodeJS-Golang-Interview_QA#what-are-virtual-property-in-mongoose)
   <details><summary>Answer</summary>
   A virtual property is not persisted to the database. We can add it to our schema as a helper to get and set values.but it wont store in database
   </details>
5. [What is Datamasking?](https://github.com/Gauthamjm007/Backend-NodeJS-Golang-Interview_QA#what-is-datamasking)
    <details><summary>Answer</summary>
    It protects the sensitive data from unauthorized access. It is a technique to hide the sensitive data from the user. 
    </details>

## Reference