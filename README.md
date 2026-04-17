# Tugas Praktikum 1 - Docker Dockerfile

**Mata Kuliah:** Sistem Paralel dan Terdistribusi - E2526  
**Nama:** Fadhil Awalia Kusuma  
**NIM:** 11231023

## Deskripsi

Praktikum ini mencakup implementasi berbagai instruksi Dockerfile:

| Instruksi | Direktori | Keterangan |
|-----------|-----------|------------|
| FROM | `praktikum/from/` | Menentukan base image |
| RUN | `praktikum/run/` | Menjalankan perintah saat build |
| CMD | `praktikum/command/` | Perintah default saat runtime |
| LABEL | `praktikum/label/` | Menambahkan metadata |
| ADD | `praktikum/add/` | Menambahkan file (+ extract + URL) |
| COPY | `praktikum/copy/` | Menyalin file dari build context |
| EXPOSE | `praktikum/expose/` | Mendeklarasikan port |
| ENV | `praktikum/env/` | Mengatur environment variable |
| VOLUME | `praktikum/volume/` | Membuat mount point |
| WORKDIR | `praktikum/workdir/` | Mengatur working directory |
| USER | `praktikum/user/` | Mengatur user/group |
| ARG | `praktikum/arg/` | Build-time variable |
| HEALTHCHECK | `praktikum/healthcheck/` | Health monitoring |
| ENTRYPOINT | `praktikum/entrypoint/` | Menentukan executable |
| Multi Stage | `praktikum/multi-stage/` | Optimasi ukuran image |
| .dockerignore | `praktikum/ignore/` | Mengecualikan file dari build context |

## Cara Menjalankan

Masuk ke direktori instruksi yang diinginkan, lalu build dan run:

```bash
cd praktikum/<instruksi>
docker build -t praktikum/<instruksi> .
docker run --rm praktikum/<instruksi>
```

Contoh untuk Multi Stage Build:

```bash
cd praktikum/multi-stage
docker build -t praktikum/multi .
docker run --rm praktikum/multi
```

## Prasyarat

- Docker 29.x+
- Go 1.18 (tersedia via Docker image `golang:1.18-alpine`)
