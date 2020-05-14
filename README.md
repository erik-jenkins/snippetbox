# Snippetbox

This repo contains my code for following along with Alex Edwards' book, [Let's Go](https://lets-go.alexedwards.net/).

I am not following along exactly with the book. Here are some of the changes:

- The project is dockerized, including the MySQL insteance.
- There are scripts for running the project and launching chrome if available.
- The `.nvimrc` file enables convenient shortcuts for development if you are using neovim (AsyncRun plugin is required!)
- I am using the gorilla mux router instead of pat for RESTful routing.

## Debugging

The code contains some support for debugging with delve. As there isn't a convenient GUI for delve, a `launch.json` for debugging through VS code is included. Make sure to launch the project in debug mode from the command line (`startdev.sh` followed by `debuggo.sh`), set breakpoints in VS Code, and launch the included configuration.
