SELECT Person.FirstName,Person.LastName,Address.City, Address.State from Person left join Address on Address.PersonID=Person.PersonID
