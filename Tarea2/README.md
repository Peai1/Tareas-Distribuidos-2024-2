Benjamin Lopez 202173533-k 

Felipe Leon 202173598-4

Las entidades en las MVs se distribuyen de la siguiente forma:

• dist017.inf.santiago.usm.cl: Isla File / Data Node 1

• dist018.inf.santiago.usm.cl: Continente Folder / Diaboromon

• dist019.inf.santiago.usm.cl: Continente Server / Data Node 2

• dist020.inf.santiago.usm.cl: Primary Node / Nodo Tai

Ejecución:
Se deben abrir 8 terminales, 2 por cada MV, en donde se deben ejecutar en este orden y esperar el mensaje de que se está en ejecución (ejecutar en grupo05/Tarea2):

1) Maquina 020 -> sudo make docker-primary-node
2) Maquina 020 -> sudo make docker-nodo-tai

3) Maquina 019 -> sudo make docker-datanode2
4) Maquina 019 -> sudo make docker-continente-server

5) Maquina 017 -> sudo make docker-datanode1
6) Maquina 017 -> sudo make docker-isla-file

7) Maquina 018 -> sudo make docker-continente-folder
8) Maquina 018 -> sudo make docker-diaboromon

Si se quieren hacer consultas al primary node desde el nodo tai, se debe ingresar 1 por la terminal. Se puede hacer en cualquier momento que se desee.

Para ver los archivos, despues de terminar la ejecución de los nodos, se puede ver los archivos creados con el comando "ls" y acceder a los .txt con el comando "cat nombre_archivo.txt" al archivo correspondiente ya sea en primary_node, datanode1 y datanode2. Si se quiere salir de esta terminal se debe ingresar "exit", notar que esta terminal estara en root..., por tanto para volver a hacer la ejecucion de los archivos se debe salir de esta.
