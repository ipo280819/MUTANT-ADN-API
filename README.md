# <b>MUTANT ADN API</b>
<i><b>[Estadisticas de mutantes encontrados hasta el momento](http://ipo-mutant.tk/stats)</b></i>
## <b>Correr el códido</b>

#### <b>Prerrequisitos</b>

- [Golang](https://golang.org/dl/)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker compose](https://docs.docker.com/compose/install/)

Para levantar de manera local el proyecto es necesario tener el archivo de entorno <b>_.env_</b> configurado, para el cual puedes basarte en el <b>_.env.example_</b>. Luego de eso te bastara con estar en la raíz del proyecto desde la consola y ejecutar los siguientes comandos:

<code>
$ docker-compose build

$ docker-compose up
</code>

## <b>Correr suite de test</b>

Los test se encuentran unicamente enfocados a los servicios que son quienes cargan con la lógica del negocio, por tanto para correr la suite de tests

<code>
$ go test ./services/ -cover
</code>

## <b>Observaciones</b>

<ul>
<li>
    Las secuencias que venían de la forma <b>{A A A A A}</b> se interpretaron como:
        
<ol>
    <li><b>[A A A A] A</b></li>  
    <li> <b>A [A A A A]</b></li>
</ol>
    Es decir, dos subsecuencias diferentes.

</li>

<br>
<li>Para el cálculo del ratio no estaba definida la regla en caso de que el número de humanos fuera 0 por lo cual en su defecto se determinó dejarlo como 1.</li>
</ul>
