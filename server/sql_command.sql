CREATE TABLE IF NOT EXISTS employee (
    id INTEGER NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    birth_date TEXT NOT NULL,
    start_date TEXT NOT NULL,
    salary NUMERIC NOT NULL,
    pvd_fund_rate NUMERIC NOT NULL,
    PRIMARY KEY(id AUTOINCREMENT)
);

INSERT or REPLACE INTO employee (id, first_name,last_name, birth_date, start_date, salary, pvd_fund_rate)
VALUES
(1, 'John', 'Wick', '1964-09-02', '2004-03-01', 46000, 10),
(2, 'Beyoncé', 'Knowles', '1981-09-04', '2008-09-16', 25000, 3),
(3, 'Justin', 'Bieber', '1994-03-01', '2008-12-01', 21000, 3),
(4, 'Justin', 'Timberlake', '1981-01-31', '2014-01-20', 35000, 5),
(5, 'Taylor', 'Swift', '1989-12-13', '2018-05-01', 23000, 4),
(6, 'Justin', 'Timberlake', '1981-01-31', '2021-06-01', 20000, 5),
(7, 'Taylor1', 'Swift', '1989-12-13', '2021-10-01', 20000, 3),
(8, 'Taylor2', 'Swift', '1989-12-13', '2022-01-01', 20000, 3),
(9, 'Taylor3', 'Swift', '1989-12-13', '2022-05-01', 20000, 3);
