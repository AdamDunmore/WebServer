{
    description = "WebServer Control Panel Dev Flake";
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";   
    };
    outputs = { self, ... } @inputs: 
    let
        system = "x86_64-linux";
        pkgs = import inputs.nixpkgs {
            inherit system;
            config = {
                allowUnfree = true;
            };
        };
        nixModules = import ./nix { inherit pkgs; config = self.config; };
    in
    {     
        nixosModules.default = nixModules.modules.nixosModule;
        homeManagerModules.default = nixModules.modules.homeModule;
        devShells.${system}.default = nixModules.devShells.default;
        packages.${system}.default = nixModules.package { inherit pkgs; };
        apps.${system}.default = { type = "app"; program = "${self.packages.${system}.default}/bin/webserver"; };
    };
}
