## Deskripsi:
Proyek Backend Golang sederhana untuk Project Manager, dikembangkan dengan menggunakan IDE Visual Studio Code, diuji APInya dengan Postman, dan dihubungkan database Postgresql dengan Xata. Pengguna dapat register, login, dan melakukan CRUD untuk project.

## Hasil API Test menggunakan Postman:
# User Register
1. Berhasil mendaftarkan akun
![image](https://github.com/user-attachments/assets/166e5d67-9be8-49f2-8982-295df6c99b7c)

2. Tidak berhasil mendaftarkan akun karena username
![image](https://github.com/user-attachments/assets/df92cf26-d1c4-431e-aa0f-4153652f6f07)

3. Tidak berhasil mendaftarkan akun karena password
![image](https://github.com/user-attachments/assets/7fb8dbe7-5b92-447b-aef1-e1fc1eeb08b8)

4. Membuat user baru dengan username yang registered
![image](https://github.com/user-attachments/assets/0990d42f-f76e-4201-9f97-b5567272ad62)

5. Hasil di database Xata
![image](https://github.com/user-attachments/assets/edbbe23f-ff24-49c0-b75b-793b491e4398)

# User Login
1. Kredensial valid
![image](https://github.com/user-attachments/assets/292a8e36-e37d-4bd6-ba95-25c55a51e281)

2. Kredensial invalid
![image](https://github.com/user-attachments/assets/cefda2ed-3301-46e1-8f9a-9725e10be126)

# Create Project
1. Membuat project
![image](https://github.com/user-attachments/assets/00797cfd-bbc2-4bc0-bef2-aee3e73c9c38)
![image](https://github.com/user-attachments/assets/d188dbab-9e32-4dbd-8da0-e07fb8236d21)

2. Hasil di database Xata
![image](https://github.com/user-attachments/assets/d50eecd1-e2ba-436a-befb-45570adbbd32)

# Get All Project
![image](https://github.com/user-attachments/assets/0e47ecc9-6c24-4d6d-bbd0-00fdf2afdfe5)

# Get Project by ID:
1. ID valid
![image](https://github.com/user-attachments/assets/9e568715-a865-49db-a6d3-2de2ac40c572)

2. ID invalid
![image](https://github.com/user-attachments/assets/74b42074-85fb-4d45-aef6-45034397e169)

# Update Project by ID:
1. Update project by ID
![image](https://github.com/user-attachments/assets/a4b36dc5-d8cf-40f1-964e-db0bde0fab2e)

2. Hasil di database Xata
![image](https://github.com/user-attachments/assets/5ce582da-a90f-40f6-82fc-3319a6e21548)

3. Dicoba kembali Get Project by ID setelah Update
![image](https://github.com/user-attachments/assets/468d88c9-1de7-4bea-b534-9c90b8143185)

4. ID invalid
![image](https://github.com/user-attachments/assets/60f7438a-78ac-48b9-9b31-105751179dec)

# Delete Project by ID:
1. Delete project by ID
![image](https://github.com/user-attachments/assets/b6ca2968-916f-494a-9dd0-f87fc3c7088f)

2. Hasil di database Xata
![image](https://github.com/user-attachments/assets/def19a0d-13ad-4719-a6bd-f35658ba5489)

3. ID invalid
![image](https://github.com/user-attachments/assets/34d02792-4b00-4557-becd-5af61a871ab5)

