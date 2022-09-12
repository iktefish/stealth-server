with import <nixpkgs> {}; let
  mongoGoDriver = buildGoModule {
    src =
      fetchFromGitHub
      {}
      + "github.com/mongodb/mongo-go-driver";
  };
  googleUuid = buildGoModule {
    src =
      fetchFromGitHub
      {}
      + "github.com/google/uuid";
  };
in
  stdenv.mkDerivation {
    name = "go-mongoDb";
    buildInputs = with pkgs; [
      graphviz
    ];
  }
