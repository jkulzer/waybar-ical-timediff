# Maintainer: jkulzer <git at jkulzer dot dev>
pkgname=waybar-ical-timediff-bin
pkgver=0.1.6
pkgrel=1
pkgdesc="A tool to display current calendar appointment progress in waybar"
arch=(x86_64)
url="https://github.com/jkulzer/waybar-ical-timediff"
license=('GPL')
groups=()
depends=()
makedepends=()
optdepends=()
provides=()
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=(https://github.com/jkulzer/waybar-ical-timediff/releases/download/$pkgver/waybar-ical-timediff)
noextract=()

package() {
	mkdir -p $pkgdir/usr/bin
	cp $srcdir/waybar-ical-timediff $pkgdir/usr/bin/waybar-ical-timediff
}
sha256sums=('9dfcb1f4e30a4671bfb84123f1939a334bbf692a5103e199c352cdc01cad5392')
