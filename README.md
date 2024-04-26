<br />

<h1 align="center">waybar-ical-timediff</h1>

  <p align="center">
    An application that displays remaining time of an iCal event
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>


<!-- GETTING STARTED -->
## Getting Started

### Installation

#### With `nix`

1. Download repo
    Clone the repo
    ```sh
    git clone https://github.com/jkulzer/waybar-ical-timediff
    cd waybar-ical-timediff
    ```

2. Build the binary
    ```sh
    nix build
    ```

3. Execute it
    ```
    ./result/bin/waybar-ical-timediff
    ```

#### With `makepkg`
This doesn't build the application, it pulls the latest GitHub release
1. Download repo
    ```sh
    git clone https://github.com/jkulzer/waybar-ical-timediff
    cd waybar-ical-timediff
    ```
2. Create package
    ```sh
    makepkg
    ```
3. Install package
    ```sh
    pacman -U <path to generated package>
    ```

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- ROADMAP -->
## Roadmap

- [ ] Automatic hash generation for `PKGBUILD`
- [ ] Flake for easy installation on NixOS


<p align="right">(<a href="#top">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the GNU GPL v3 License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

