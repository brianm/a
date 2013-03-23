# [Go](http://golang.org/) Skeleton Project

This is a skeleton of a Go project which sets up a Go workspace and
places the project into that workspace.

# The PACKAGE Makefile Variable
Folks using this skeleton will need to customize it. The first
customization will be to make sure the PACKAGE variable is sane, right
now the Makefile assumes it is at the root of a git repo, and does
some munging to guess. If it guesses correctly for you, rock on. If it
guesses incorrectly for you, just set the PACKAGE variable to the
right thing.

# The build Makefile Target
Running

    $ make
    
by itself will run the <code>build</code> target. As exactly what you
want to build will vary by project, you will probably need to
customize this target to build the right thing. Right now it builds a
binary called <code>hello</code>. You'll never guess what it does!
