# Wego
A go client for the [Wekan REST API](https://wekan.github.io/api/v6.97/#wekan-rest-api)  

## Features
- Automatic login and token renewal
- 100% of the official API implemented
- Implements Wekan REST API v6.97

## Sample
```go
// Create client.
// The client will automatically login and renew its token regularly.
c, err := wego.NewClient(wego.Options{
    RemoteAddr: "https://your.wekanboard.com",
    Username:   "user",
    Password:   "secure-password",
})
if err != nil {
    log.Fatal(err)
}

boards, err := c.GetPublicBoards(context.Background())
if err != nil {
    log.Fatal(err)
} 
fmt.Printf("Public boards: %+v\n", boards)

self, err := c.GetCurrentUser(context.Background())
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Self: %+v\n", self)

other, err := c.GetUser(context.Background(), "user-id-of-somebody")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("Other: %+v\n", other)
```

## Known problems
The current state of the Wekan API is slightly brittle.  
Some API funcs are implemented according to spec, but do currently not work on my testing instance.  
I need to create issues in the Wekan repository for them.

## Issues
When you find issues or bugs, please create an issue in this repository and/or submit a PR.

## License
MIT, see `LICENSE` file of this repository.
