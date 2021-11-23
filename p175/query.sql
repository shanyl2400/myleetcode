# Write your MySQL query statement below
select 
p.FirstName as firstName, p.LastName as lastName, a.City as city, a.State as state 
from  Person p
left join
Address a
on p.PersonId = a.PersonId;