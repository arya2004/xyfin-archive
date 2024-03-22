# Transaction Isolation Levels

ACID Properties:

1. Atomicity: All or nothing

2. Consistency: Data is in a consistent state before and after the transaction

3. Isolation: Transactions are isolated from each other

4. Durability: Changes are permanent

Read Phenomena:

1. Dirty Read: Reading uncommitted data

2. Non-Repeatable Read: Reading same data that has been modified by another transaction are different

3. Phantom Read: Same query returns different results

4. Serialization Anomaly: Transactions are not executed in the order they were received


Isolation Levels:

1. Read Uncommitted (Lowest Level): Dirty reads, Non-repeatable reads, Phantom reads, Serialization anomalies
    - Can see data written by other transactions that have not been committed
    - Not available in SQL Server and PostgreSQL

2. Read Committed: Non-repeatable reads, Phantom reads, Serialization anomalies
    - Can only see data that has been committed

3. Repeatable Read: Phantom reads, Serialization anomalies
    - Same read will return the same data. 
    - Data is same as when transaction started to when it ends

4. Serializable: Serialization anomalies
    - Can achieve same results as if transactions were executed in serial order instead of concurrently