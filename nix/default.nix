{ config, pkgs, ... }:
{
    devShells = import ./devshells.nix { inherit pkgs; };
    modules = import ./module.nix { inherit pkgs; inherit config; };
    package = import ./pkg.nix { inherit pkgs; };
}
