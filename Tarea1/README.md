## Integrantes:
- **Felipe León Córdova** - 202173598-4
- **Benjamín Lopez Fuenzalida** - 202173533-k

## Ver archivos caravanas registros

Para poder ver los archivos de los registros, despues de ejecutar el contenedor de caravanas, se muestra una terminal en donde se puede hacer el comando "ls" para ver los archivos, luego se puede acceder con "cat "nombre_archivo.txt"" a alguno de los 3 archivos que se crearon de los registros de las 3 caravanas. 

Para salirse de esta terminal solo ingresar "exit".

Esto lo implementamos en otra rama ya que en "main" se creaban los archivos pero se cerraban al terminar el dockerfile, pero esto se implementó después que se cayeran las MVs el domingo entonces preferimos subirlo a otra rama llamada "ver_archivos_MV" para irse a la segura en caso de que fallara algo ya que en local funcionaba pero no lo probamos en las MVs, pero lo unico que cambia es eso mencionado

Por lo que es recomendable probar primero en rama "ver_archivos_MV" y si por alguna razón (ya que no se pudo probar) llegase a fallar, ejecutar en rama "main"

## Parámetros modificables

Los siguientes parametros se pueden modificar en el código:

1. **Tiempo de espera entre caravanas**  
   - Ubicación: `caravanas/caravanas.go`, linea 29.  
   - Variable: `tiempoEspera` (por defecto: 6).

2. **Tiempo de simulación de envíos de caravanas**  
   - Ubicación: `caravanas/caravanas.go`, linea 28.  
   - Variable: `tiempoEntrega` (por defecto: 3).

3. **Tiempo entre los envíos de pedidos desde los clientes**  
   - Ubicación: `clientes/clientes.go`, linea 20.  
   - Variable: `intervaloEnvio` (por defecto: 0).

## Instrucciones

Para ejecutar la tarea correctamente, ejecutar en el siguiente orden (en directorio Tarea1) ademas de que deben ser ejecutadas en las maquinas correctas dado que estan seteadas con la ip de sus respectivas maquinas, esto se detallara a continuacion:

1. **Máquina 017 (Caravanas)**   
   ```bash
   sudo make docker-caravanas

2. **Máquina 018 (Logística)**  
   ```bash
   sudo make docker-logistica

3. **Máquina 019 (Finanzas)**  
   ```bash
   sudo make docker-finanzas

4. **Máquina 020 (Clientes)**   
   ```bash
   sudo make docker-clientes

## Formato solicitudes input: solicitudes.txt

En el directorio principal, debe existir un archivo llamado solicitudes.txt que cumpla con el siguiente formato:

1. **Prímera linea: Número de paquetes a enviar con un ,**
2. **Líneas siguientes: Información de los paquetes con el siguiente formato:**
   ```text
   Tipo,Puesto_Destino,Warframe_Asignado,Cargamento,Valor
3.**Un ejemplo seria:**
   ```text
   11,
   Ostronitas,Puesto A,Excalibur,Antitoxinas,150
   Prioritario,Puesto B,Rhino,Municiones,300
   Prioritario,Puesto A,Rhino,Suministros,150
   Prioritario,Puesto C,Rhino,Suministros,150
   Prioritario,Puesto B,Rhino,Municiones,300
   Prioritario,Puesto B,Rhino,Municiones,300
   Prioritario,Puesto A,Rhino,Antitoxinas,150
   Prioritario,Puesto A,Rhino,Suministros,100
   Normal,Puesto C,Nova,Suministros,100
   Normal,Puesto C,Nova,Suministros,100
   Normal,Puesto C,Nova,Suministros,100
   
El tipo debe ser Ostronitas, Prioritario o Normal (sensible a mayusculas) y deben estar los parametros separados por una ","
