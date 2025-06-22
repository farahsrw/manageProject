## Deskripsi:
Proyek Backend Golang sederhana untuk Project Manager, dikembangkan dengan menggunakan IDE Visual Studio Code, diuji APInya dengan Postman, dan dihubungkan database Postgresql dengan Xata. Pengguna dapat register, login, dan melakukan CRUD untuk project.

## Hasil API Test menggunakan Postman:
# User Register
1. Berhasil mendaftarkan akun
![image](https://github.com/user-attachments/assets/5d98aa7d-6d4f-4820-9804-edd40eaaef1b)

2. Tidak berhasil mendaftarkan akun karena username
![image](https://github.com/user-attachments/assets/4bed3e39-a5aa-43fb-bbfd-5f6dc819d9fb)

3. Tidak berhasil mendaftarkan akun karena password
![image](https://github.com/user-attachments/assets/ea2aac1b-b33a-4714-8608-82638b9a9c9f)

4. Hasil di database Xata
![image](https://github.com/user-attachments/assets/a741df37-7c67-4ed9-84df-fca1c41a58fc)

# User Login
1. Kredensial valid
![image](https://github.com/user-attachments/assets/d895579d-1752-4749-bd3a-cec94eec816e)

2. Kredensial invalid
![image](https://github.com/user-attachments/assets/7e41bc14-0e78-4327-9f5c-26277179a58c)

# Create Project
1. Membuat project
![image](https://github.com/user-attachments/assets/52a819fe-c113-464a-9a94-86829b904ba4)
![image](https://github.com/user-attachments/assets/c9ea8a7f-010a-408e-99de-ed51ca841d13)

2. Hasil di database Xata
![image](https://github.com/user-attachments/assets/7363a529-4b97-4902-8445-6b2864d093f6)

# Get All Project
![image](https://github.com/user-attachments/assets/839be16b-ec38-4b46-bdfd-f4d7e164bd7d)

# Get Project by ID:
1. ID valid
![image](https://github.com/user-attachments/assets/28ac5cb2-e2d8-467a-9042-54923a30191c)

2. ID invalid
![image](https://github.com/user-attachments/assets/1c0440a2-69d3-4e1e-8a0f-b8789caaa97b)

# Update Project by ID:
1. Update project by ID
![image](https://github.com/user-attachments/assets/8bca7051-c626-43ef-af1a-b0b73b4aea53)

2. Hasil di database Xata
![image](https://github.com/user-attachments/assets/b8530b48-2b89-48e2-bc74-06381a8c22ba)

3. Dicoba kembali Get Project by ID setelah Update
![image](https://github.com/user-attachments/assets/62aacc17-1bc4-4017-bfd2-8ea0bf1b8e12)

4. ID invalid
![image](https://github.com/user-attachments/assets/d57ecfd9-c277-4698-871f-8010f0c8ac84)


# Delete Project by ID:
1. Delete project by ID
![image](https://github.com/user-attachments/assets/3615337c-c621-4267-bb64-56d14d1ae263)

2. Hasil di database Xata
![image](https://github.com/user-attachments/assets/69780014-1c5e-4737-aaf8-53c9e5210a21)

3. ID invalid
![image](https://github.com/user-attachments/assets/b15d86c4-aa7e-47a5-aaee-6227db59dd5e)
