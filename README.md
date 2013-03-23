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
projects using thisg skeleton remain compatible with <code>go
remote</code> expectations, but not require setting up any global
GOPATH or workspace, or having to deal with checking things out to the
right place in your workspace. Automation, such as make, is supposed
to solve this for us, so let us let it do that.

## Useful Targets

* <code>make</code> The default target is "build", which will build
  your project.
* <code>make clean</code> Will remove all compiled stuff from the
  workspace, specifically WORKSPACE/pkg and WORKSPACE/bin
* <code>make clean-workspace</code> Will complrtely wipe out the
  WORKSPACE.

There are also <code>workspace</code> and <code>deps</code> targets,
which set up the basic workspace, and run <code>go get</code>
respectively.

## "Normal" Go Development

The canonical way of working with your project would be to export your
GOPATH and then cd down to the code you are working on, and work on it
there. This works fine. If you want to avoid using make for typical
build and test, this is a good thing. If you don't want to cd down and
muck about, you can just work in the project checkout *outside* the
GOPATH. Folks in #go-nuts will think you are weird, but it's okay.
You'll want to make sure to use <code>make</code> to build things
though, so it can keep the GOPATH sane for you.

If you run the <code>activate.sh</code> bash script, it will start a
child bash process designed to work in your WORKSPACE:

    $ ./activate.sh
    
This will export the correct GOPATH, will prefix PATH with
WORKSPACE/bin, will export a <code>$project</code> environment
variable which is the path to the current project in the GOPATH, a
<code>$root</code> environment variable pointing at the checkout root,
and a <code>WORKING_ON</code> environment variable with the package
name currently being worked on (useful for inclusion in PS1).

Additionally, if there is a <code>.bash\_local</code> file in the
project directory, it will source that after everything else is set
up. Personally, my <code>.bash\_local</code> prepeds $WORKING_ON to my
PS1, a la

    export PS1="[$WORKING_ON] $PS1"
    
So that I know what workspace I am in.

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
