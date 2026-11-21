self: { config, ... }:
let 
    pkgs = self.pkgs;
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
            path = [ pkgs.streamrip ];

            serviceConfig = {
                ExecStart = pkgs.writeShellScript "start-webserver" ''
                    export WEBSERVER_PASSWORD="$(cat ${cfg.passwordFile})" 
                    ${package}/bin/webserver
                '';
                environment.HOME = cfg.user;
                Restart = "on-failure";
                RestartSec = 5;
            };

            environment = cfg.environmentVariables;

            wantedBy = [ "multi-user.target" ];
        };
    };
}
