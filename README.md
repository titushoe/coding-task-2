# coding-task-2
Exercise: implement a merge function for a list of intervals 

- The exercise will be implemeted in golang
- all needed func's are inplemented in the main-module "merge-intervalls.go" or imported from the go std-librarys
- if you have installed on your os golang of the right way, you should be able to run, build or install this module for testing

# Install "merge-intervals.go" from github
   - perequisite: installed go on your os
   - go to your "$GOPATH/src" and then:
   
   % mkdir coding-task-2  # whatever !!
   % cd coding-task2
   % git init
     Initialized empty Git repository in  
        ... coding-task-2/.git/
   % git pull https://github.com/titushoe/coding-task-2
      remote: Enumerating objects: 11, done.
       .
       .
      Unpacking objects: 100% (11/11), done.
      From https://github.com/titushoe/coding-task-2
      
   - and now you could run:
    % echo " [25,30][2,19][14,23][4,8]" | go run merge-intervalls.go
    
    or
    
    % go run merge-intervalls.go testfiles/test1.csv
   
   
