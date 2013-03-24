# Go Skeleton Project

This is a skeleton of a [Go](http://golang.org/) project which sets up
a Go workspace and places the project into that workspace. It comes
with a Makefile to make life easier for you, but it sets up a "go
remote" compatible project (and you should keep this compatibility!).
The Makefile exists to set up a Go workspace and to make life easier
for common cases of using that workspace (such as building the project
in the workspace).

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

# Don't Panic

## I don't know Go, I just want to build this project!

If you just want to build this project, not develop it, run make:

    $ make
    
Output will be in the <code>WORKSPACE/bin</code> directory.

## I *am* a Go programmer, Makefiles are the devil!

It's okay, the Makefile just automates workspace setup, and common
tasks (like building the primary output of the project). You don't
need to use it.

## Some Useful Targets

* <code>make</code> The default target is "build", which will build
  your project.
* <code>make test</code> Run tests.
* <code>make clean</code> Will remove all compiled stuff from the
  workspace, specifically WORKSPACE/pkg and WORKSPACE/bin
* <code>make clean-workspace</code> Will completely wipe out the
  WORKSPACE.
* <code>make docserver</code> Run the godoc server on port 5050
* <code>make fmt</code> Run <code>go fmt</code> on project sources. 

There are also <code>workspace</code> and <code>deps</code> targets,
which set up the basic workspace, and run <code>go get</code>
respectively.

## "Normal" Go Development

The canonical way of working with your project would be to export your
GOPATH and then cd down to the code you are working on, and work on it
there. This works fine. If you want to avoid using make for typical
build and test, this is a good thing. If you don't want to cd down and
muck about, you can just work in the project checkout *outside* the
GOPATH. Folks in #go-nuts will think you are weird, but it's okay. In
this case, do use make to build things for you, as it will do stuff
within the workspace on your behalf.

If you run the <code>activate.sh</code> bash script, it will start a
child bash process designed to work in your WORKSPACE:

    $ ./activate.sh
    
This will export the correct <code>GOPATH</code>, will prefix
<code>PATH</code> with <code>WORKSPACE/bin</code>, will export a
<code>project</code> environment variable which is the path to the
current project in the <code>GOPATH</code>, a <code>root</code>
environment variable pointing at the checkout root, and a
<code>WORKING_ON</code> environment variable with the package name
currently being worked on (useful for inclusion in <code>PS1</code>).

Additionally, if there is a <code>.bash\_local</code> file in the
project directory, it will source that after everything else is set
up. Personally, my <code>.bash\_local</code> prepends $WORKING_ON to
my PS1, and cd's into the project within the GOPATH:

    export PS1="[$WORKING_ON] $PS1"
    cd $project
    
So that I know what workspace I am in.

# Customizing the Skeleton

Folks using this skeleton will almost certainly need to customize it.
Beyond making it fancier, the two things to look out for:

## The PACKAGE Makefile Variable 

The first customization will be to make sure the PACKAGE variable is
sane, right now the Makefile assumes it is at the root of a git repo,
and does some munging to guess. If it guesses correctly for you, rock
on. If it guesses incorrectly for you, just set the PACKAGE variable
to the right thing.

You can find out what it is guessing via the <code>check-sanity</code>
target:

    $ make check-sanity
    
Which will output the various environment stuff being inferred. Check
the <code>PACKAGE</code> one, it is the the one which is fragile. It
should match the base package folks would import, ie
<code>github.com/brianm/variant</code>.

## The build Makefile Target

Running

    $ make
    
by itself will run the <code>build</code> target. As exactly what you
want to build will vary by project, you will probably need to
customize this target to build the right thing. Right now it builds a
binary called <code>hello</code>. You'll never guess what it does!

## The test Makefile Target

The <code>test</code> target tries to run tests in the project root,
and all directories other then WORKSPACE and .git off of the root. You
may need to customize this behavior for your needs. If you do need to
customize this behavior, take a look at the <code>fmt</code> target as
well, as it uses the same heuristic as <code>test</code>.
