# Things to test

## Client

* A missing JSON file
* An empty JSON file
* A JSON file with only one entry
* A JSON file with duplicated entries
* A large JSON file with a small `ulimit`
* Malformed data is rejected (port missing coordinates?)
* API calls that return counts (simplifies testing)
* Reloading reflects new file changes
* Reloading on the same file doesn't change anything

## Server

* Check there are no shortcodes before client starts up
* Check there are N shortcodes after client starts up
* Check there are N shortcodes after a client reload
* Check that absent shortcodes don't return a port
* Check that known shortcodes return the correct port
* Check that duplicated codes end up with the last set of data
