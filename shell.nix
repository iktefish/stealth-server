with import <nixpkgs> {}; let
  run = pkgs.writeShellScriptBin ''run'' ''make run'';
  test = pkgs.writeShellScriptBin ''run'' ''make test'';
  build = pkgs.writeShellScriptBin ''run'' ''make build'';
in
  stdenv.mkDerivation {
    name = "stealth-server";
    buildInputs = with pkgs; [
      run
      test
      build
    ];
  }
