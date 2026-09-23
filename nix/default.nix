{ pkgs, lib, config, ... }:
{
    devShells = import ./devshells.nix { inherit pkgs; };
    modules = import ./module.nix { inherit lib; inherit config; };
}
