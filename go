[ ! -d /usr/local/bin ] && echo "/usr/local/bin is nonexistant" && exit 1
[ ! -O /usr/local/bin ] && SUDO_MAYBE=sudo
$SUDO_MAYBE curl -fso /usr/local/bin/getnjs https://raw.githubusercontent.com/benchOS/getnjs/master/go
$SUDO_MAYBE chmod +x /usr/local/bin/getnjs
