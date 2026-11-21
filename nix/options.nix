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
        user = mkOption {
            type = types.attrsOf types.str;
            default = "root";
            example = "adam";
            description = "User to run service as";
        };
        passwordFile = mkOption {
            type = types.str;
            default = "";
            example = "config.sops.secrets.ts_key.path";
            description = "The path to your secret file for reading the password, the file must be globally readable";
        };
    };
}
