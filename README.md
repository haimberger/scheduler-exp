Happy Shiny Scheduler
=====================

This is a web application intended to help improve productivity by promoting a positive, focused approach to working on tasks.

More details can be found [here](https://docs.google.com/document/d/19fPz48mgMIbgPmBH03mRc4Ii2RGQ8nqGLT9ySna3WvE).

## Development

Before committing any changes, make sure to run

    make lint

It will concurrently run a bunch of linters including [go vet](https://golang.org/cmd/vet/) and [megacheck](https://github.com/dominikh/go-tools/tree/master/cmd/megacheck).

You can also run linters for individual packages:

    make clock.lint // run linters in clock package

## Testing

The following command runs all tests in all packages:

    make test

If you'd like to see the test coverage information in a more visually appealing form, you can try the following instead:

    make test-coverage

It will save the coverage information to a file, then open a browser window showing the covered (green), uncovered (red), and uninstrumented (grey) source. You can find more information under "Viewing the results" [here](https://blog.golang.org/cover).

If you only want to run tests for one package, there are commands for that as well:

    make clock.t  # run tests in clock package
    make clock.cv # show coverage information for clock package
