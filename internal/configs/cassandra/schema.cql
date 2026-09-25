CREATE KEYSPACE IF NOT EXISTS vortex WITH replication = 
{'class': 'SimpleStrategy', 'replication_factor': 3};

CREATE TABLE IF NOT EXISTS vortex.tasks (
    taskid uuid, 
    priority int,
    status text,
    retry_count int,
    execution_time duration,
    last_executed time,

    PRIMARY KEY (taskid)

) WITH comment='vortex tasks';

INSERT INTO vortex.tasks (taskid, priority, status, retry_count, execution_time, last_executed) VALUES (6e1fc47d-48e8-4b50-adbc-b825856768a5, 1, 'PENDING', 3, 0m0s, '00:00:00');
INSERT INTO vortex.tasks (taskid, priority, status, retry_count, execution_time, last_executed) VALUES (7a20314f-c2f1-413d-94da-b5ad28ce849a, 2, 'PENDING', 2, 0m0s, '00:00:00');
INSERT INTO vortex.tasks (taskid, priority, status, retry_count, execution_time, last_executed) VALUES (4a9adf9f-c2a6-4e97-b3fd-278f633d2c74, 7, 'PENDING', 6, 0m0s, '00:00:00');

SELECT * FROM tasks
