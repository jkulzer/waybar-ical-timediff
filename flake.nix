{
  description = "A basic flake with a shell";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
    # gomod2nix = {
    #   url = "github:tweag/gomod2nix";
    #   inputs.nixpkgs.follows = "nixpkgs";
    #   inputs.flake-utils.follows = "flake-utils";
    # };
  };

  outputs = {
    nixpkgs,
    flake-utils,
    # gomod2nix,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      # pkgs = nixpkgs.legacyPackages.${system};
      pkgs = import nixpkgs {inherit system;};
    in {
      packages.default = 
				pkgs.buildGoModule {
						name = "waybar-ical-timediff";
					  pname = "waybar-ical-timediff";
            src = ./.;
						vendorHash = "sha256-u/nSxAOLOflGgZUjW2v3moSfWB1I0Qko7jcxbDxIkXM=";
			};
      devShells.default = pkgs.mkShell {
        packages = with pkgs; [
          go
        ];
      };
    });
}
