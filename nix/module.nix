{ lib, config, ... }:
let 
    inherit (lib) mkEnableOption mkOption types mkIf;
    cfg = config.services.webserver;
    options = {
        enable = mkEnableOption "Enables the Go Webserver Module";
        environmentVariables = mkOption {
            type = types.attrsOf types.str;
            default = {};
            example = { WEBSERVER_PORT = "1913"; };
            description = "Sets environment variables for the server";
        };
    };
in
{
    nixosModule = {
        options.services.webserver = options;
        config = mkIf cfg.enable {
            systemd.services.webserver = {
                description = "Personal Go Web Server";
                after = [ "network-online.target" ];
                wants = [ "network-online.target" ];

                serviceConfig = {
                    ExecStart = "/home/adam/bin/webserver"; # TODO move when add package
                    WorkingDirectory = "/home/adam/WebServer"; # TODO move when add package
                    Restart = "on-failure";
                    RestartSec = 5;
                };

                environment = cfg.environmentVariables;

                wantedBy = [ "multi-user.target" ];
            };
        };
    };


    homeModule = {
        options.services.webserver = options;
        systemd.user.services.webserver = mkIf cfg.enable {
            Unit = {
                Description = "Personal Go Web Server";
                After = [ "network-online.target" ];
            };

            Service = {
                ExecStart = "/home/adam/bin/webserver"; # TODO move when add package
                WorkingDirectory = "/home/adam/WebServer"; # TODO move when add package
                Restart = "on-failure";
                RestartSec = 5;
                Environment = cfg.environmentVariables;
            };

            Install.WantedBy = [ "default.target" ];
        };
    };
}
