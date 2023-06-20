# Task 0: Prerequisites
Clone the Lab’s GitHub Repo

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/1.png)


# task 1 : Run some simple Docker containers

## Run a single task in an Alpine Linux container

1.
![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/2.png)

2. melihat list containers

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/3.png)

## Run an interactive Ubuntu container
1. Run a Docker container and access its shell.

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/4.png)

2. Run the following commands in the container.

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/5.png)

3. exit

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/6.png)

4. mengecek versi alpine linux

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/7.png)

## Run a background MySQL container

1. menjalankan mysql containers

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/8.png)

2. melihat daftar containers

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/9.png)

3. mengecek logs

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/10.png)

melihat proses yang berjalan di dalam containers

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/11.png)

4. cek versi mysql

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/12.png)

5. menghubungkan terminal ke sh

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/13.png)

6. cek versi 

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/14.png)


# Task 2: Package and run a custom app using Docker

## Build a simple website image

1. eksport dengan ID Docker, dan mengecek apakah sudah terhubung 

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/15.png)

2. membuat docker image

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/16.png)

3. Menjalankan container

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/17.png)

4. tampilan website

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/18.png)

5. menghentikan containers

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/19.png)


# Task 3: Modify a running website
## Start our web app with a bind mount

1. menjalankan containers

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/20.png)

2. mengecek website

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/21.png)

## Modify the running website

1. mengubah tampilan website

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/22.png)

2. tampilan website yang sudah dirubah

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/23.png)

3. stop dan menghapus containers yang berjalan

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/24.png)

4. menjalankan kembali tanpa bind mount

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/25.png)

5. website telah kembali ke tampilan awal

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/26.png)

6. menghentikan containers

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/27.png)


## Update the image

1. membuat image baru

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/28.png)

2. melihat list image

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/29.png)

## Test The New Version

1. menjalankan containers
2. 
![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/30.png)

3. tampilan website

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/31.png)

3. menjalankan containers baru

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/32.png)


## Push Your Image To Docker Hub

1. list image

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/34.png)

2. login ke dockerhub

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/35.png)

3.Push image 1.0

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/36.png)

4. Push image 2.0

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/37.png)

5. Hasil Push

![image](https://github.com/Dean-182/tekn-cloud-computing/blob/main/minggu-09/PTC/task2/task3/38.png)
