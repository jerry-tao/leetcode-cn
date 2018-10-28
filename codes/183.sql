SELECT Customers.Name as Customers from Customers left join Orders on Orders.CustomerID=Customers.ID WHERE Orders.ID is NULL
