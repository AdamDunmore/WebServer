{
    description = "WebServer Control Panel Dev Flake";
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";   
    };
    outputs = { self, ... } @inputs: 
    let
        system = "x86_64-linux";
        pkgs = inputs.nixpkgs.legacyPackages.${system};
    in
    {     
        packages.${system}.default = import ./nix/pkg.nix { inherit pkgs; };
        apps.${system}.default = { type = "app"; program = "${self.packages.${system}.default}/bin/webserver"; };
        devShells.${system}.default = import ./nix/devshells.nix { inherit pkgs; };
        nixosModules.default = ./nix/nixos.nix;
        homeManagerModules.default = ./nix/home.nix;
    };
}
