{ pkgs ? import <nixpkgs> {} }:
pkgs.mkShell {
  buildInputs = [
    pkgs.pkg-config
    pkgs.go
    pkgs.gobjectIntrospection
    pkgs.gtk4
    pkgs.gnome3.gtk
  ];
}
