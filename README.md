# Snippetbox

The tutorial project for the "Let's Go!" book by Alex Edwards.

This repo has some changes from the book in that the mysql database
is run in a container that is reset every time it is started. If you
have Neovim with the asyncrun plugin (and set excr in your init.vim),
you can start the dev server and docker-compose with `<leader>pr`. Otherwise,
use the `startdev.sh` script in the project root.
