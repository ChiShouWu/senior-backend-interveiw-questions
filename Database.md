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
5. Why do databases treat null as a so special case? For example, why does SELECT * FROM table WHERE field = null not match records with null field in SQL?
6. Lazy Loading
7. How would you find the most expensive queries in an application?
8. In your opinion, is it always needed to use database normalization? When is it advisable to use demoralized databases?
9.  What do you understand by Data Redundancy?
    <details><summary> Ans </summary>      
      Duplication of data in the database is known as data redundancy. As a result of data redundancy, duplicated data is present at multiple locations, hence it leads to wastage of the storage space and the integrity of the database is destroyed.
    </details>
10. What do you understand by Data Independence? What are its two types?
    <details><summary> Ans </summary>
      Data Independence refers to the ability to modify the schema definition in one level in such a way that it does not affect the schema definition in the next higher level.

      The 2 types of Data Independence are:

      Physical Data Independence: It modifies the schema at the physical level without affecting the schema at the conceptual level.
      Logical Data Independence: It modifies the schema at the conceptual level without affecting or causing changes in the schema at the view level.
    </details>

11. Define the relationship between ‘View’ and ‘Data Independence’.
    <details><summary> Ans </summary>
      Answer: View is a virtual table that does not have its data on its own rather the data is defined from one or more underlying base tables.
      Views account for logical data independence as the growth and restructuring of base tables are not reflected in views. 
    </details>

12. What are the advantages and disadvantages of views in the database?
13. What is transaction?
14. Define Database Lock and its types.
15. Define Phantom deadlock.
16. What do you understand by B-Trees?
17. How would you find the most expensive queries in an application?
18. What are some common issues with ORMs?
19. What is phantom read?
20. 讀未提交(Read Uncommitted) 讀已提交(Read Committed) 可重複讀(Repeatable Read) 可序列化(Serializable)
21. Write Skew? Postgres的可重複讀是能解決更新丟失的，但同樣無法解決寫入偏斜。
22. Race condition
    <details><summary>Answer</summary>
    - Atomic update
    - Transaction lock
    - Version control
    </details>
23. Lost Updates in Mysql, Postgres

## RDBMS
1. Explain count(*), count(1), count(column_name) in SQL?
2. How to improve count(*)?
3. Why MySQL use B+ tree?
    <details><summary>Answer</summary>
    B+Tree only store data in leaf node, B-Tree store data in all node. So there is less data in B+Tree, so it is faster to find data cuz less IO read.

    The time complexity of B+Tree is O(log`d`n), d is the number of children of a node, n is the number of data in the tree. In actual situation d is larger than 100, so even data rows comes to 10million, the height of the tree is only 3.

    Compare with hash table, the time complexity of hash table is O(1), but the hash table is not sorted, so it is not suitable for range query.
    </details>
4. What is B tree and B+ tree?
5. 如果user可以 follow的人/訂閱的topic/favorite的商品數, 是無上限/1萬筆以上的時候, db schema會怎麼設計?
6. 怎麼維持不同table之間的data consistency
7. How to prevent same position be bought by two or more users at the same time?
8. How database process a command?
9. What could happen if multiple users create transactions at the same time?
10. What kind of isolation level do you know?
11. What is MVCC?
12. What is buffer pool?
13. Describe the different of `drop` `delete` `truncate`?
14. Primary key vs Unique key, what is the difference? which one is faster?
15. for update lock

### Index
1. What do you understand by Index hunting?
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
 

## NoSql
1. 你的MongoDB怎麼做text search？
2. MongoDB每個document有32MB上限，如果你的文章超過32MB要怎麼辦？
3. What is eventual consistency?

5. [What are virtual property in mongoose?
](https://github.com/Gauthamjm007/Backend-NodeJS-Golang-Interview_QA#what-are-virtual-property-in-mongoose)
6. [What is Datamasking?](https://github.com/Gauthamjm007/Backend-NodeJS-Golang-Interview_QA#what-is-datamasking)


## Reference