package config

import (
	"testing"
)

var TEST_STRING = `
name        : nano
version     : 2.5.3
source      :
    - http://www.nano-editor.org/dist/v2.5/nano-2.5.3.tar.gz ; b2b060129b9feff2d4870d803a441178c96531de9aed144ec0b83bd63ccb12ee
license     : GPL-3.0
description : |
  GNU nano is an easy-to-use text editor originally designed as a replacement for Pico, the ncurses-based editor from the non-free mailer package Pine (itself now available under the Apache License as Alpine).
conflicts:
    - pico
prepend_path:
    PATH:
        - $INSTALL_DIR/bin
setup  :
    - autoreconf -vdi
    - ./configure --enable-utf8 --docdir=$INSTALL_DIR/doc
build  :
    - make --prefix=$INSTALL_DIR
install:
    - make install
    - install -D -m 00644 $BUILD_DIR/nanorc $INSTALL_DIR/nanorc
`

func TestParse(t *testing.T) {
	c := Parse([]byte(TEST_STRING))
	if c.Name == "" {
		t.Error(c.Name)
	}
	t.Error(c)
}
