// The Employee table holds all employees. Every employee has an Id, and there is also a column for the department Id.

// +----+-------+--------+--------------+
// | Id | Name  | Salary | DepartmentId |
// +----+-------+--------+--------------+
// | 1  | Joe   | 70000  | 1            |
// | 2  | Henry | 80000  | 2            |
// | 3  | Sam   | 60000  | 2            |
// | 4  | Max   | 90000  | 1            |
// | 5  | Janet | 69000  | 1            |
// | 6  | Randy | 85000  | 1            |
// +----+-------+--------+--------------+
// The Department table holds all departments of the company.

// +----+----------+
// | Id | Name     |
// +----+----------+
// | 1  | IT       |
// | 2  | Sales    |
// +----+----------+
// Write a SQL query to find employees who earn the top three salaries in each of the department. For the above tables, your SQL query should return the following rows.

// +------------+----------+--------+
// | Department | Employee | Salary |
// +------------+----------+--------+
// | IT         | Max      | 90000  |
// | IT         | Randy    | 85000  |
// | IT         | Joe      | 70000  |
// | Sales      | Henry    | 80000  |
// | Sales      | Sam      | 60000  |
// +------------+----------+--------+
// https://leetcode.com/problems/department-top-three-salaries/

select count(*) , emp.name, emp.salary ,dep.name from Employee as emp,department as dep
where emp.DepartmentId = dep.Id 
group by emp.Id
having count(*) < 3
order by dep.name, Salary desc

SELECT Email, COUNT(*)
FROM user_log
GROUP BY Email
HAVING COUNT(*) > 1
ORDER BY UpdateDate DESC


SELECT    
d.name  as "Department"                     
,e1.Name as "Employee"
, e1.Salary as "Salary"

FROM 
Employee e1 
JOIN Employee e2  JOIN Department d

WHERE 
e1.DepartmentId = e2.DepartmentId 
AND e1.Salary <= e2.Salary  AND d.id = e2.DepartmentId


GROUP BY d.name,e1.id
HAVING COUNT(DISTINCT(e2.Salary)) <= 3
ORDER BY d.name , salary DESC 


SELECT
    d.Name AS 'Department', e1.Name AS 'Employee', e1.Salary
FROM
    Employee e1
        JOIN
    Department d ON e1.DepartmentId = d.Id
WHERE
    3 > (SELECT
            COUNT(DISTINCT e2.Salary)
        FROM
            Employee e2
        WHERE
            e2.Salary > e1.Salary
                AND e1.DepartmentId = e2.DepartmentId
        )
;

