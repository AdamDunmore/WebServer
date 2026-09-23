{ config, pkgs, ... }:
let 
    inherit (pkgs.lib) mkIf;
    cfg = config.services.webserver;
    package = import ./pkg.nix { inherit pkgs; };
in
{
    imports = [ ./options.nix ];
    config = mkIf cfg.enable {
        environment.systemPackages = [ package ];
        systemd.services.webserver = {
            description = "Personal Go Web Server";
            after = [ "network-online.target" ];
            wants = [ "network-online.target" ];

            serviceConfig = {
                ExecStart = "${package}/bin/webserver";
                # WorkingDirectory = "/home/adam/WebServer"; # TODO test
                Restart = "on-failure";
                RestartSec = 5;
                EnvironmentFile = cfg.passwordFile;
            };

            environment = cfg.environmentVariables;

            wantedBy = [ "multi-user.target" ];
        };
    };
}
