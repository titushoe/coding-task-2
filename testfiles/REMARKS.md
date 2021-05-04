# Remarks to the implemented code

At the first step the code ist designed to solve the exercise
and get an result.

## To anwere the questions

    1) How acts the run-time and how ist the storage-usage
       The run-time and the storage-usage wil expand if we had more data in the input.
       To change the code for a better manner of this we should test this code with
       bigger data sets and check the bottlenecks with added go-packages like "benchmarks"
       or "perf"

    2) How could we get a better robustness eg. of a bigger data-input
       A better way to produce bigger data for the input is read them from files.
       We can made excel-sheets with interval-datas and export them to .csv-files.
       to read from .csv-files is actually realized in the coding!
       Another way could be to get the datas from databases and store the results in them.
