self:
{
  pkgs,
  config,
  ...
}:
{
    devShells = import ./devshells.nix { inherit pkgs; };
    package = import ./pkg.nix { inherit pkgs; };
}
