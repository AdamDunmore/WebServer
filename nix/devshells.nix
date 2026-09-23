{ pkgs, ... }:
{
    default = pkgs.mkShell { 
        packages = with pkgs; [
            go
        ];
        shellHook = ''
            zsh
        '';
    };
}
