# Signposting

## About

This is a very simple Go program to explore signposting.  It is not in a useful state right now.
It has been made by inspecting the behavior of https://github.com/stain/signposting/
by @stain and @kinow.  If you need a package to work with signposting links you should visit that repo.

If you really want a Go based solution, feel free to check this out and send in any
requests.  I am happy to do what I can and any PRs are welcome!  Note, this really is
just a personal testing repo for now.  If this works out I might incorporate this into
the main Gleaner code based for use if any collaborators wish to use it.

However, it does seem to be working with the HTTP and HTML signposting approaches and runs
relatively well.  There are some edge cases, like warning on mis-matches between HTTP and HTML
link values for example.  I will try and work my way through all the tests now that the plunbing is
working and resolve things. 

## Running

Running with no flags will provide the options

```bash
❯ go run cmd/signposting/main.go
no flags provided
Usage:
-all
    Run all known test URLs
-position int
    Position of the element to keep in the array (default 1000)
-testset
    Show the set of test URLs and exit (overrides other flags)
-url string
    URL to test
exit status 2
```

The -all flag will run all the test in the [testURLs.md](./docs/testURLs.md) file.

If you use -testset these will be printed out with an index number to the left.  You can
can use the -position and the index number to test that URL

Otherwise you can test any use with the -url flag as in

```bash
❯ go run cmd/signposting/main.go -url=https://w3id.org/a2a-fair-metrics/06-http-citeas-describedby-item/

-----------------------------------------------------
HTTP HEAD Results for https://w3id.org/a2a-fair-metrics/06-http-citeas-describedby-item/

cite-as          to https://w3id.org/a2a-fair-metrics/06-http-citeas-describedby-item/                           with type 
describedby      to https://s11.no/2022/a2a-fair-metrics/06-http-citeas-describedby-item/index.ttl               with type text/turtle
item             to https://s11.no/2022/a2a-fair-metrics/06-http-citeas-describedby-item/test-apple-data.csv     with type text/csv

HTML HEAD Results for https://w3id.org/a2a-fair-metrics/06-http-citeas-describedby-item/
```



## Testing Set

I am using the test URLs at  https://s11.no/2022/a2a-fair-metrics/
to test this code with.  You can also see this in [testURLs.md](./docs/testURLs.md).


