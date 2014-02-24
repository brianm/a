# Skeleton Go Project

Base project for go stuffs. Probably not #gonuts compatible.

# Versions, Release, Tags, and Branches 

We make releases, master is not always stable between releases. When
we make a release we'll merge the current release into the
<code>go1</code> branch.

# Working with this Project 

So goskel has (or will have shortly) stuff for buildings debs, man
pages, etc. This doesn't fit neatly into the <code>go get</code>-able
model (though it can). I also have beliefs about workspace-per-project
which are not widely held in the greater go community. In the interest
of accomodating everyone, the project is <code>go get</code> friendly,
but still has its own full workspace, etc. How do we do this you ask?
I am glad you asked!

You can checkout the <code>project</code> branch and run make to build
a workspace. You can then run <code>make activate</code> to enter the
workspace and do you stuff! It will drop you in a local checkout of
this project, inside the workspace.

Finally, as I don't hold with master always being perfectly stable,
there is a <code>go1</code> branch which is always the most current
release. <code>go get</code> prefers <code>go1</code> to
<code>master</code> so folks who want to rely on it as a library, or
install it via <code>go get</code> will get the most recent release,
rather than the most recent checkin to master.

If you hate all of this, ignore it. Pull requests, <code>go get</code>
etc all work as with any other Go project. Have fun and ignore the
<code>project</code> and <code>go1</code> branches for your hackery.
