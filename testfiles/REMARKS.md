# Remarks to the implemented code

At the first step the code ist designed to solve the exercise
and get an result in a short time.

## Answeres to the following questions

    1) How acts the code run-time and how gig is the storage-usage ?
       The run-time and the storage-usage will be expanding if we had 
       more data in the input.
       If we would change the code for a better manner of this we 
       should test this code with bigger data sets and check the 
       bottlenecks with added go-packages like "benchmarks" or "perf" 
       and then we could see which code parts needs a better coding.

    2) How could we get a better robustness eg. of a bigger data-input ?
       A better way to produce bigger data for the data input is read them 
       from files.
       We can made excel-sheets with interval-datas and export them to 
       .csv-files.
       To read from .csv-files is actually realized in this coding!
       Another way could be to get the datas from databases and store 
       the results in them.
