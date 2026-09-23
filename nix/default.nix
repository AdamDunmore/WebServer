{ pkgs, ... }:
{
    devShells = import ./devshells.nix { inherit pkgs; };
    modules = import ./module.nix;
}
