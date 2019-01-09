# Gophin' Off With Linked Lists

I produced this code as a result of creating my Golang tutorial series, Gophin' Off with Linked Lists, on my blog at [http://saharsh.org](http://saharsh.org). All the code written throught that blog series can be found here. Check out the series, starting with the first part [here](http://saharsh.org/2019/01/08/gophin-off-with-linked-lists-part-6/), for more information.

## Clone this repository

        go get github.com/saharshsingh/gophinoff

## Run all tests

Following will run all the unit tests in this repository.

        ./testall.sh

To generate an HTML coverage report after running all the unit tests, run the above command with the `--show-html` option.

        ./testall.sh --show-html

NOTE: The `testall.sh` script is pretty general purpose and can be dropped in the top level directory of any Go project to achieve the same functionality for that project without any code changes.

## Task Master web service

The `Task Master` web service from the last part of the blog series can be built into an executable by running the following command.

        go build github.com/saharshsingh/gophinoff/taskmaster

The resulting executable will be created in the current working directory (customizable using the `-o` option in `go build`).
