{ config, pkgs, ... }:
let 
    inherit (pkgs.lib) mkIf;
    cfg = config.services.webserver;
    package = import ./pkg.nix { inherit pkgs; };
in
{
    imports = [ ./options.nix ];
    config = mkIf cfg.enable {
        home.packages = [ package ];
        systemd.user.services.webserver = {
            Unit = {
                Description = "Personal Go Web Server";
                After = [ "network-online.target" ];
            };

            Service = {
                ExecStart = "${package}/bin/webserver";
                # WorkingDirectory = "/home/adam/WebServer";
                Restart = "on-failure";
                RestartSec = 5;
                Environment = cfg.environmentVariables;
                EnvironmentFile = cfg.passwordFile; # TODO needs testing
            };

            Install.WantedBy = [ "default.target" ];
        };
    };
}
