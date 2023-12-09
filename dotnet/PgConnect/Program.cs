// See https://aka.ms/new-console-template for more information
using Npgsql;

Console.WriteLine("User is ?");
var userName = Console.ReadLine();
Console.WriteLine("Password is ?");
var password = Console.ReadLine();
Console.WriteLine($"User is {userName}");
Console.WriteLine($"Password is {password}");



var connectionString = "Host=localhost;Username=user;Password=userpass;Database=sdl";
await using var dataSource = NpgsqlDataSource.Create(connectionString);



// Retrieve all rows
await using (var cmd = dataSource.CreateCommand("SELECT VERSION();"))
await using (var reader = await cmd.ExecuteReaderAsync())
{
    while (await reader.ReadAsync())
    {
        Console.WriteLine(reader.GetString(0));
    }
}
