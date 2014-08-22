GP=~/.gospace
echo "Creating $GP symlink to workspace"

rm -f $GP
ln -s $DIR/_workspace $GP
