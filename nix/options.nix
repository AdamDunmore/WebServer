{ pkgs, ... }:
let 
    inherit (pkgs.lib) mkEnableOption mkOption types;
in
{
    options.services.webserver = {
        enable = mkEnableOption "Enables the Go Webserver Module";
        environmentVariables = mkOption {
            type = types.attrsOf types.str;
            default = {};
            example = { WEBSERVER_PORT = "9999"; };
            description = "Sets environment variables for the server";
        };
    };
}
