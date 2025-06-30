

# Proyecto 1
### Sistemas Operativos 2

##### Kevin Golwer Enrique Ruiz Barbales
##### 201603009

### Descripcion:

El siguiente proyecto tuvo como finalidad el poner en practica los conocimientos en implementacion de llamadas al sistema y uso de estas para extender la funcionalidad de nuestro kernel de manera personalizada , ademas de eso el poder darle uso a esta informacion para algo practico y util ,en este caso visualizar la informacion de procesos y red , tanto para un proceso individual y para nuestra pc en general 


### Pasos usados para el desarrollo del proyecto

#### 1.) Configuraciones basicas

necesitaremos un archivo de configuracion para poder compilar nuestro kernel , el cual obtendremos del sistema en el que estamos,es fundamental el usar este comando ya que si copiamos un archivo de otro sistema no nos funcionara, con el siguiente comando obtendremos el archivo y lo copiara en la carpeta de nuestro nuevo kernel :

```bash
#Copia el archivo de conf a nuestro kernel 
cp -v /boot/config-$(uname -r) .config
``` 

<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/imagen6.png" width="500px" height="300px" align="center">

posterior a eso , se nos copiara ese archivo .conf ,tendremos que ejecutar el siguiente comando para tener limipio nuestro ambiente de compilacion:

```bash
#Limpia nuestro ambiente de trabajo 
make clean
``` 

luego para dejar en nuestra configuracion de compilacion los drivers y archivos necesarios y quitando todo aquello que no usaremos ,ejecutaremos el siguiente comando :

```bash
#Elimina compilacion innecesaria
make oldconfig
make localmodconfig
``` 

por ulitmo deshabilitamos ciertos certificados que no nos serviran para nuestro caso:

```bash
#Deshabilita certificados
scripts/config --disable SYSTEM_TRUSTED_KEYS
scripts/config --disable SYSTEM_REVOCATION_KEYS
scripts/config --set-str CONFIG_SYSTEM_TRUSTED_KEYS ""
scripts/config --set-str CONFIG_SYSTEM_REVOCATION_KEYS ""
``` 

con esto tendremos ya listo nuestro ambiente para poder compilar nuestro kernel , ahora tendremos que proceder a crear nuestra syscall e implementarla para luego poder compilar nuestro nuevo kernel

 

#### 3.) Creacion de nuestra nueva syscall 
tendremos que crear un archivo llamado "syscalls_usac.h" en esta ruta: "./linux-6.8/kernel/"
en este archivo tendremos que colocar la declaracion de nuestra syscall, que vendria a ser el codigo fuente
para nuestro caso crearemos una syscall sencilla llamada hola_mundo , que no recibe parametros unicamente imprime un mensaje el codigo seria el siguiente:

```bash
#Definicion simple de syscall
SYSCALL_DEFINE0(hola_mundo)
{
    printk(KERN_INFO "Hola mundo desde el kernel!\n");
    return 0;
}
``` 


luego de haber creado nuestro archivo lo unico que tendremos que hacer es agregar esta nueva llamada primero a la tabla de syscalls que se encuentra en esta ruta: "linux-6.8/arch/x86/entry/syscalls/syscall_64.tbl"  y agregaremos una nueva linea con la siguiente forma :
```bash
#Linea ejemplo tabla syscall
548 common hola_mundo sys_hola_mundo
``` 

tambien tendremos que declarar nuestra funcion en los siguientes archivos : "linux-6.8/arch/x86/include/asm/syscalls.h" y 
"linux-6.8/kernel/sys.c" , en ellos tendremos importar nuestro archivo syscalls_usac.c y luego colocar la declaracion de la syscall de la siguiente forma: 
```bash
#Linea ejemplo sys.c o syscalls.h

asmlinkage long sys_hola_mundo(void);
``` 



#### 4.) Compilacion Kernel

Una vez tengamos esto listo y tengamos guardados los cambios , procederemos a compilar el kernel con el siguiente comando:

```bash
#Compila el kernel
fakeroot make 
``` 

si quisieramos ejecutar el proceso usando mas cpu's ejecutamos el comando con la siguiente configuracion:

```bash
#Compila el kernel 3cpu's
fakeroot make -j3
``` 

una vez empiece esto se vera como en la imagen:

<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/imagen1.png" width="500px" height="300px" align="center">


esto tardara como una hora dependiendo los cpu que usemos 


#### 4.) Instalacion Kernel

cuando este termine de compilar ,tendremos que instalarlo y lo haremos con los siguientes comandos:

```bash
#Instala los modulos
make modules_install

#Instala el kernel como tal 
make install

#Instala cabeceras 
make headers_install
``` 

<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/imagen7.png" width="500px" height="300px" align="center">


<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/imagen8.png" width="500px" height="300px" align="center">


#### 5.) Prueba de funcionamiento

cuando ya lo hallamos instalado , tendremos que reiniciar nuestra maquina, y en eso tendremos que presionar la tecla "Shift" ,con ella habilitaremos el grub y tendremos que seleccionar el kernel que nosotros modificamos:

<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/imagen14.png" width="500px" height="300px" align="center">


#### 1.) Creacion de syscall (funcionalidad real)

para esta parte tendremos que crear nuestra syscall en c , esta sera la que se implementara en el kernel para poder extender la funcionalidad del mismo para nuestro caso mostraremos la implementacion de la syscall de matar proceso , el codigo que usa es el siguiente:

```c
SYSCALL_DEFINE1(kill_process_by_pid, pid_t, pid)
{
    struct task_struct *task;
    struct pid *pid_struct;

    // Buscar el proceso por su PID
    pid_struct = find_get_pid(pid);
    if (!pid_struct)
        return -ESRCH; // No existe ese PID

    task = pid_task(pid_struct, PIDTYPE_PID);
    if (!task)
        return -ESRCH; // No existe tarea asociada

    // Enviar la señal SIGKILL
    send_sig(SIGKILL, task, 0);

    return 0; // Exito
}

``` 


como vemos la primera parte nos indica la cantidad  de parametros que recibe nuestra syscall, seguido del nombre de la syscall , luego se declara el tipo de dato que retornara , ya dentro va el codigo fuente de nuestra syscall el que se ejecutara cuando lo llamemos 


#### 2.) Implementacion syscall

para la implementacion de la syscall en nuestro caso , tendremos que agregarla a nuestro archivo syscalls_usac.c donde tenemos creadas todas nuestras syscalls , agregamos las librerias necesarias para que funcione y tendremos que modificar los archivos que se mostraran a continuacion y su codigo correspondiente:

```bash
 #sys.c
 #include "syscalls_usac.c"
 asmlinkage long sys_kill_process_by_pid(pid_t pid);
 
 #syscall_64.tbl
 552 common kill_process_by_pid  sys_kill_process_by_pid
 
 #syscalls.h
 asmlinkage long sys_kill_process_by_pid(pid_t pid);
 
 #syscalls_usac.c
SYSCALL_DEFINE1(kill_process_by_pid, pid_t, pid)
{
    struct task_struct *task;
    struct pid *pid_struct;

    // Buscar el proceso por su PID
    pid_struct = find_get_pid(pid);
    if (!pid_struct)
        return -ESRCH; // No existe ese PID

    task = pid_task(pid_struct, PIDTYPE_PID);
    if (!task)
        return -ESRCH; // No existe tarea asociada

    // Enviar la señal SIGKILL
    send_sig(SIGKILL, task, 0);

    return 0; // Exito
}


``` 


tendremos que asegurarnos que coincida la declaraciones y nombre del codigo de nuestra syscall en los archivos que mencione antes, ya que de lo contrario no funcionara correctamente ,de igual forma con los nombres





#### 3.) Compilacion e Instalacion del nuevo Kernel

para esta parte tendremos que tener listo nuestro kernel para poder hacer uso de nuestras syscall , procederemos a compilar el kernel dependiendo si este ya habia sido compilado con anterioridad o es la primera vez que se compila , una vez terminado lo instalaremos con los siguentes comandos:

```bash
#Instala los modulos
make modules_install

#Instala el kernel como tal 
make install

#Instala cabeceras 
make headers_install
``` 

<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/imagen7.png" width="500px" height="300px" align="center">


<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/imagen8.png" width="500px" height="300px" align="center">



#### 4.) Consumo de syscalls

para esta parte , tendremos que consumir nuestra syscall creada , en mi caso use el lenguaje golang para consumir la syscall , la declaracion para poder acceder al kernel y consumirla seria la siguiente:

```golang
 const SYS_KILL_PROCESS_BY_PID = 552

func KillProcessByPID(pid int) error {
	_, _, errno := syscall.Syscall(uintptr(SYS_KILL_PROCESS_BY_PID), uintptr(pid), 0, 0)
	if errno != 0 {
		return errno
	}
	return nil
}

``` 


nos tendremos que asegurar que use el mismo nombre que le colocamos , y de igual forma tendremos que verificar que posea el mismo numero de syscall que le asignamos , tambien de pasarle los parametros necesarios si los tuviera y de guardar en alguna variable del tipo que retorna


#### 5.) Implementacion de syscalls

una vez pudimos consumir nuestra syscall ,procedimos a crear una API , ya que con ella podremos acceder a ella sin tener que estar creando un codigo especifico por cada llamada que necesitemos hacer ,en nuestra API creamos un endpoint que se llama kill_process , y en el la llamamos de la siguiente manera, asegurandonos que se le pasen por parametro el mismo tipo de dato , en otra de nuestras sycall usamos mutex para sincronizar nuestras variables para evitar condiciones de carrera o lecturas malas por culpa de la sincronizacion

```golang


func kill_process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // o "http://tu-dominio.com"
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodGet {
		print(r.Method)
		http.Error(w, "Método no permitido, usa GET", http.StatusMethodNotAllowed)
		return
	}
	pid := r.URL.Query().Get("pid")

	println(pid)

	println("adfads2")
	// Validar campos requeridos
	if pid == "" {
		println("adfads")
		http.Error(w, "El campo PID es requerido", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}

	err = KillProcessByPID(num)
	if err != nil {
		fmt.Printf("Error al matar el proceso: %v\n", err)
	} else {
		fmt.Println("Proceso terminado con éxito.")
	}

	response := Response{
		Status: num,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

}

``` 



#### 6.) Consumiendo nuestra syscall

cuando ya terminamos nuestra api procedemos a consumirla desde nuestra vista , para nuestro caso tenemos 2 dashbord, uno que es global , donde mostrarmos el status general de nuestra maquina y el segundo que nos muestra la informacion especifica de un proceso , ademas de que en esta tenemos implementado la parte de matar proceso ,para matar el proceso unicamente tendremos que pasarle por parametro el PID del proceso:

```typescript
#Implementacion en angular servicio

  async kill_process(PID:number){

    const url = `${this.direccion}/kill_process?pid=${PID}`;
     //const url = this.direccion+'/info_process_unique';
     const {data} = await axios.get(url);
     return data

   }
```


#### 7.) Mostrando Vistas de htop

ahora mostraremos como quedo la interfaz implementando no solo la syscall de kill process sino tambien las otras llamadas que fueron solicitadas 

##### 1) Vista Inicial
<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/IMG1.png" width="600px" align="center">

nos muestra la cantidad de ram total usada en nuestro sistema en tiempo real, ademas de que nos muestra lo mismo pero para el cpu , ademas de eso nos da la cantidad de procesos activos 

##### 2) Vista RAM
<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/IMG2.png" width="600px" align="center">

aca tenemos el grafico de lineas para mostrar el historial de RAM de nuestro sistema donde si fluctua pues esto se vera reflejado en el grafico


##### 3) Vista CPU
<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/IMG3.png" width="600px" align="center">

aca tenemos el grafico de lineas para mostrar el historial de CPU de nuestro sistema donde si fluctua pues esto se vera reflejado en el grafico


##### 4) Vista RX y TX
<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/IMG4.png" width="600px" align="center">

aca tenemos el grafico de lineas para mostrar el historial de RX (KB recibidos) y TX (KB transmitidos) de nuestro sistema donde si fluctua pues esto se vera reflejado en el grafico

##### 5) Dashboard Procesos
<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/IMG5.png" width="600px" align="center">

aca tenemos el dashboard para mostrar la informacion individual de los procesos, para nuestro caso mostramos el % de RAM usado en el momento , tambien % de RAM en el tiempo , % de CPU usado en el momento , % de CPU usado en el tiempo y por ultimo la cantidad de energia usada en el tiempo 


##### 6) Dashboard Procesos Listado
<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/IMG6.png" width="600px" align="center">

tambien tenemos el listado de los procesos donde podremos ver lo que es el PID , Nombre, CPU, RAM , Prioridad ,Inicio del proceso , UID y por ultimo el consumo de energia, para poder visualizar la info de nuestro proceso en los dashboard solo tendremos que dar click en ver para que nos carque los cambios 


##### 7) Terminar Proceso
<img src="https://github.com/gkruiz/syscalls_linux/blob/main/IMAGENES/IMG7.png" width="600px" align="center">

luego para terminar un proceso ,tendremos que dar click en terminar y automaticamente este se terminara en nuestro sistema y nos mostrara el mensaje que tiene la imagen 


#### 8.) Errores Encontrados y soluciones

durante el proceso de compilacion tuve varios errores, entre los cuales tengo los siguientes:

##### a.) Error no mostraba grub
tuve el problema que no me mostraba el grub cuando iniciaba la maquina virtual,presionando las teclas "Shift+Ctrl" o "Shift+Esc"

###### Solucion
presionar unicamente la tecla "Shift"



##### b.) Error variable cpu me decia que no existia
tuve el problema que cuando trate de obtener la informacion global del cpu y esta era necesaria para poder obtener la informacion completa del cpu, probe delcarar la variable para que funcionara pero cuando lo compile no hacia nada

###### Solucion
importar la libreria cpufreq.h
```c
 #include <linux/cpufreq.h>


 SYSCALL_DEFINE2(sysinfo_usage, unsigned int __user *, cpu_usage, unsigned int __user *, ram_usage)
{
//cpu , es la variable que debe de obtener de la libreria
     for_each_possible_cpu(cpu) {
        struct kernel_cpustat *kcs = &kcpustat_cpu(cpu);

        work_jiffies += kcs->cpustat[CPUTIME_USER] +
                       kcs->cpustat[CPUTIME_NICE] +
                       kcs->cpustat[CPUTIME_SYSTEM] +
                       kcs->cpustat[CPUTIME_SOFTIRQ] +
                       kcs->cpustat[CPUTIME_IRQ];

        total_jiffies += work_jiffies + kcs->cpustat[CPUTIME_IDLE] +
                        kcs->cpustat[CPUTIME_IOWAIT] +
                        kcs->cpustat[CPUTIME_STEAL] +
                        kcs->cpustat[CPUTIME_GUEST];
    }

```


##### c.) Error al momento de leer una estructura con un array compartido
tuve el problema que tenia un array compartido , y necesitaba escribir y leer este en diferentes hilos en mi aplicacion, el problema que tenia es que cuando leia los valores en la api siempre me retornaba nulo , o vacio probe varias veces pero no me funcionaba

###### Solucion
usar mutex para sincronizar los hilo , mientras uno leia el otro espera para escribir 
```golang
type FixedQueue struct {
	data []StructTraffic
	cap  int
	mu   sync.Mutex
}
```


##### d.) Actualizar mi informacion automaticamente
tuve el problema que necesitaba actualizar mi informacion en la api , pero no podia estar haciendo una peticion manual cada vez que quisiera agregar un registro de historial a mi array , de igual forma sino la informacion iba a ser erronea pues no se habia consultado en un intervalo constante

###### Solucion
usar gorutinas para mantener un hilo siempre consultando la syscall y tener mi array de informacion siempre actualizado igua siempre usando mutex para sincronizar
```golang
 #Ejemplo para la syscall de red
 func getNetworkDataTime(datoRed *FixedQueue) error {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		//Obtiene la informacion de red
		rx, tx, err := GetNetworkStats()
		if err != nil {
			fmt.Printf("Error al obtener estadísticas de red: %v\n", err)
			return nil
		}

		temp := StructTraffic{
			RX: rx,
			TX: tx,
		}

		datoRed.Enqueue(temp)
		//fmt.Printf("Tráfico recibido: %d KB\n", rx)
		//fmt.Printf("Tráfico transmitido: %d KB\n", tx)

	}

}

```


##### e.) Error al recompilar el kernel
tenia el detalle que queria compilar el kenel mas rapido , pues aunque seguia los pasos para hacer el procedimiento normal siempre se tardaba aun mucho , y eso me quitaba mucho tiempo pues empezaba todo de cero 

###### Solucion
usar el comando make y configuracion para hacer la compilacion mas rapida, pues no empieza de cero sino que solo agrega lo que se modifico y lo agrega a lo que ya se habia compilado anteriormente 

```bash
#Solo compila lo necesario no todo el kernel 
make -j$(nproc)
```


##### d.) Error tipo de datos
tuve otro detalle y es que normalmente usamos tipos de datos comunes como int o string o char , pero en este caso tuve que usar algunos tipos de datos especiales pues de lo contrario no iba a poder enviar u obtener la informacion de la syscall ejemplo: uint32 ,uint64,uintptr ,int64,int32 , tomando en cuenta que la syscall esta en c y mi api en go

###### Solucion
me correspondio mapear los tipo de datos de c a golang para que coincidieran y asi pudiera tener el valor real en la variable , de lo contrario se podia truncar o dar error ya que el tipo de dato no coincidia con el colocado en c

```C
#C estructura
struct proc_info2 {
    pid_t pid;
    char name[TASK_COMM_LEN];
    uid_t uid;
    unsigned long ram_usage_kb;
    int priority;
    unsigned long cpu_usage; // no perfecto, solo a modo de ejemplo
    unsigned long start_time;
};
```


```golang
#go estructura
type ProcInfo struct {
	PID        int32
	Name       [TASK_COMM_LEN]byte
	UID        uint32
	RamUsageKB uint64
	Priority   int32
	CPUUsage   uint64
	StartTime  uint64
}
```










