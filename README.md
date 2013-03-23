# Go Skeleton Project

This is a skeleton of a [Go](http://golang.org/) project which sets up
a Go workspace and places the project into that workspace.

The idea behind this skeleton is to NOT require anything aside from Go
to be set up, and that the skeleton will make a "Go Workspace" for
you, for your project. When you build the project it will create a
workspace, symlink the project into the right place in the workspace,
and fetch dependencies into the workspace for you (via <code>go
get</code>).

The workspace will be in the <code>WORKSPACE</code> directory of the
checkout. This is probably controversial within the Go community, but
it makes life much easier in my opinion. It is set up this way so that
projects using this skeleton remain compatible with <code>go
remote</code> expectations, but not require setting up any global
GOPATH or workspace, or having to deal with checking things out to the
right place in your workspace. Automation, such as make, is supposed
to solve this for us, so let us let it do that.

# Customizing the Skeleton

Folks using this skeleton will almost certainlyneed to customize it.
Beyond making it fancier, the two things to look out for:

## The PACKAGE Makefile Variable 

The first customization will be to make sure the PACKAGE variable is
sane, right now the Makefile assumes it is at the root of a git repo,
and does some munging to guess. If it guesses correctly for you, rock
on. If it guesses incorrectly for you, just set the PACKAGE variable
to the right thing.

## The build Makefile Target
Running

    $ make
    
by itself will run the <code>build</code> target. As exactly what you
want to build will vary by project, you will probably need to
customize this target to build the right thing. Right now it builds a
binary called <code>hello</code>. You'll never guess what it does!
