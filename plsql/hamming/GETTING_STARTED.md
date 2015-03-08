# Getting Started

These exercises lean on Test-Driven Development (TDD), but they're not an
exact match.

## Setup

You'll need access to a mounted Oracle DB. If you don't have one already
installed, here are a few options:
* download VirtualBox from Oracle and run one of the freely
available images; at the time of writing, the easiest to get started with
at the time of writing might be _Database App Development VM_. The
image is quite large...
* download and install the a version of the Oracle DB itself. Developer licenses
are free.
* get a free workspace at https://apex.oracle.com

**Note**: if you're using the online version of APEX, compilation errors will
not be indicated very clearly when working in _SQL Commands_ - you will simply
get "Error at line XX: PL/SQL: Statement ignored"... More insight can be
found using the _Object Browser_ and navigating to the object you created
(select either _Packages_ or _Procedures_ in the dropdown menu showing _Tables_,
depending on what you created for the exercise). Also, when you run statements,
"run" each individual `create` statement individually by selecting its text.
APEX does not seem to like doing too much work at once...

To work on individual problems, a nice and free way is to use SQL Developer. If 
you don't want to use yet another IDE, you can simply copy and paste your code
into a terminal / command prompt connected to the database. The files are
prepared in a way that will simply overwrite the previously compiled version.

#Exercise

## Step 1

Compile the test suite. You can easily do that by copy / pasting (let's call
that _installing_ for simplicity) into your terminal connected to a mounted
Oracle database.

This will fail, complaining that there is no package called `HAMMING#`.

To fix the error create the package by installing the prepared solution stub.
Note that you will have to re-install the package body of `UT_HAMMING#`.

A few words about naming: the `#` at the end of the name signifies that this
is a package. `UT` means _unit test_. _PL/SQL_ has a maximum identifier length
of 30 characters, which is why you will find that many words are abbreviated.

If you've worked with PL/SQL before, you might wonder why the template is a
package and not simply a standalone function. That would of course also be a
possibility, but in practice standalone procedures or functions are rarely used.

## Step 2

Try to run the test. You will notice that you are missing the function's
implementation. Create it (see the test package for examples).

Note that functions have to return a value, so for now just `return null;`.

## Step 3

Run the test again. It should now execute, but the test will fail.
That's where you get to actually implement the function!

## Wash, Rinse, Repeat

Only the first test is enabled at first, the others are commented out. To enable
them, simply delete the `--` in front of the procedure call and they'll run
when you next install the test package!

## Submit

When everything is passing, you can submit your code with the following
command:

    $ exercism submit hamming#.plsql

