``` bash tab="Bash"
for OPT in "$@"
do
  case "$OPT" in
    '-f' )  canonicalize=1 ;;
    '-n' )  switchlf="-n" ;;
  esac
done

# readlink -f
function __readlink_f {
  target="$1"
  while test -n "$target"; do
    filepath="$target"
    cd `dirname "$filepath"`
    target=`readlink "$filepath"`
  done
  /bin/echo $switchlf `pwd -P`/`basename "$filepath"`
}

if [ ! "$canonicalize" ]; then
  readlink $switchlf "$@"
else
  for file in "$@"
  do
    case "$file" in
      -* )  ;;
      *  )  __readlink_f "$file" ;;
    esac
    done
fi

exit $?
```

``` python tab="Python"

for x in y:
  print(x)
```