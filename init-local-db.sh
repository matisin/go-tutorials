#!/bin/bash

# Define el arreglo de carpetas
folders=("companies" "users" "utils")

# Define la rama a la que hacer checkout
branch=$1
docker compose up db -d
while ! docker exec microservices-db-1 pg_isready -d postgres://localhost:5432/template; do
    echo "Esperando a que la base de datos esté en ejecución..."
    sleep 1
done
echo "La base de datos está en ejecución"
# Recorre cada carpeta en el arreglo
for folder in "${folders[@]}"; do
  # Entra en la carpeta
  cd "$folder"

  # Encuentra todas las subcarpetas en la carpeta actual que contengan la palabra "api" en su nombre, excluyendo el directorio actual
  subfolders=$(find . -mindepth 1 -maxdepth 1 -type d -name "*api*")

  # Recorre cada subcarpeta
  for subfolder in $subfolders; do
    # Entra en la subcarpeta
    cd "$subfolder"

    # Haz checkout a la rama especificada
    pnpm i
    ./get-secrets.sh local
    echo "DATABASE_ADDRESS=localhost" >> .env
    pnpm run-migrations

    # Vuelve a la carpeta anterior
    cd ..
  done

  # Vuelve a la carpeta anterior
  cd ..
done