{
    description = "WebServer Control Panel Dev Flake";
    inputs = {
        nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";   
    };
    outputs = { ... } @inputs: 
    let
        system = "x86_64-linux";
        pkgs = import inputs.nixpkgs {
            inherit system;
            config = {
                allowUnfree = true;
            };
        };
    in
    {
        devShells.${system}.default = pkgs.mkShell { 
            packages = with pkgs; [
                go
            ];
            shellHook = ''
                zsh
            '';
        };
    };
}
