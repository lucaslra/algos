
var r = new Rocket("Apollo");

Task.Delay(TimeSpan.FromSeconds(5)).ContinueWith(_ => { r.Launch(); }).Wait();

record Rocket(string name)
{
    public void Launch()
    {
        Console.WriteLine($"{name} is launched!");
    }
}