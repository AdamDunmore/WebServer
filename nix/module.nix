{ lib, config, pkgs, ... }:
let 
    inherit (lib) mkEnableOption mkOption types mkIf;
    cfg = config.services.webserver;
    package = import ./pkg.nix { inherit pkgs; };
    options = {
        enable = mkEnableOption "Enables the Go Webserver Module";
        environmentVariables = mkOption {
            type = types.attrsOf types.str;
            default = {};
            example = { WEBSERVER_PORT = "9999"; };
            description = "Sets environment variables for the server";
        };
    };
in
{
    nixosModule = {
        options.services.webserver = options;
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
    };


    homeModule = {
        options.services.webserver = options;
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
    };
}
