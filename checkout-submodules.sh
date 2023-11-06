#!/bin/bash

# Define el arreglo de carpetas
folders=("companies" "users" "utils")

# Define la rama a la que hacer checkout
branch=$1

checkout(){
  # Entra en la subcarpeta
  cd "$subfolder"

    # Haz checkout a la rama especificada
    git fetch
    git checkout "$branch"
    git pull

    # Vuelve a la carpeta anterior
    cd ..
}

# Recorre cada carpeta en el arreglo
for folder in "${folders[@]}"; do
  # Entra en la carpeta
  cd "$folder"

  # Encuentra todas las subcarpetas en la carpeta actual, excluyendo el directorio actual
  subfolders=$(find . -mindepth 1 -maxdepth 1 -type d)

  # Recorre cada subcarpeta
  for subfolder in $subfolders; do
    echo "checkout a $subfolder branch $branch"
    checkout $subfolder $branch &
  done

  # Vuelve a la carpeta anterior
  cd ..
done
wait 
echo "Todos los repos en $branch"