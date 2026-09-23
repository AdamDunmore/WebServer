{ pkgs, ... }:
{
    devShells = import ./devshells.nix { inherit pkgs; };
    modules = import ./module.nix { inherit pkgs; };
    package = import ./pkg.nix { inherit pkgs; };
}
