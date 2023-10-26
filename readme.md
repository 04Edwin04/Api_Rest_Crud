Nombre del Alumno: Edwin Jair Castañeda Montoya.
Matricula: 200506.

Descripsion del funcionamiento del proyecto:

El siguiente proyecto es una API REST CRUD escrita en Go, la cual utiliza un ORM llamado GORM para interactuar con una base de datos MySQL. La API REST permite realizar las operaciones CRUD (Create, Read, Update, Delete) en la tabla users de la base de datos.

Archivo main.go: Este archivo actúa como el punto de entrada del programa, donde se establecen las rutas para la API REST y se configura el servidor web Gin, que sirve como el núcleo de la aplicación.

Archivo User: Esta estructura (struct) define el modelo "User", que se emplea para gestionar y representar los datos asociados a la tabla de usuarios en la base de datos.

Archivo db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}): Esta línea de código establece una conexión con la base de datos MySQL mediante el uso de GORM, una biblioteca de mapeo objeto-relacional. GORM facilita la interacción con la base de datos y la manipulación de datos.

Archivo db.AutoMigrate(&User{}): Este fragmento de código se encarga de ejecutar las migraciones necesarias para crear la tabla "users" en la base de datos. Esto garantiza que la estructura de la base de datos esté sincronizada con la definición del modelo "User", lo que facilita la gestión de datos relacionados con los usuarios en la aplicación.