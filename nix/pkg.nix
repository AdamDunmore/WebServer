{ pkgs, ... }:
pkgs.buildGoModule {
    pname = "webserver";
    version = "0.0.1";
    src = ../.;
    vendorHash = null;

    meta = {
        description = "Go webserver for my personal use";
        homepage = "https://github.com/AdamDunmore/WebServer";
        license = pkgs.lib.licenses.mit;
        maintainers = with pkgs.lib.maintainers; [ AdamDunmore ];
    };
}
