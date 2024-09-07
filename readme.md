# Building a sample go app based on the fridge project for Panenco

## 1. Initial setup

I started by following [this guide](https://soheilnik.medium.com/building-a-golang-project-with-firebase-authentication-part-1-146c91969d26).

I also added modules using `go get xx`.

The current version uses a live version of firebase only for the creation of tokens. The database is very simply migrated. There is no real interaction with the database.

Next steps will be:

1. Integrate with local emulated firebase authentication
2. Use firebase to store and retrieve tokens
3. Add a sample authenticated route.
