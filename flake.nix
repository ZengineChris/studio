{
  description = "Zengine Studio Developer Environment";
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
  };
  outputs = {
    self,
    nixpkgs,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixpkgs {
        inherit system;
      };
    in {
      devShells.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          nodejs_24
          biome
          go
          air
          golangci-lint
          gopls
          go-task
          templ
          goose
          typescript-language-server
          sqls
          sql-formatter
        ];
      };

      packages = {};
    });
}
