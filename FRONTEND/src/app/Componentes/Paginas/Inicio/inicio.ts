import { Component, OnInit, ViewChild } from '@angular/core';
import { NombreServicioService } from '../../../../servicio/nombre-servicio.service';
import { CookieService } from 'ngx-cookie-service';
import { HomeComponent } from '../../redirect';
 
import Swal from 'sweetalert2';
import { ChartComponent } from 'ngx-apexcharts';
import {
  ApexNonAxisChartSeries,
  ApexResponsive,
  ApexChart,
  ApexAxisChartSeries,
  ApexXAxis,
  ApexDataLabels,
  ApexTitleSubtitle,
  ApexStroke,
  ApexGrid
} from 'ngx-apexcharts';


export type ChartOptions = {
  series: ApexNonAxisChartSeries;
  chart: ApexChart;
  responsive: ApexResponsive[];
  labels: any;
};

export type ChartOptions2 = {
  series: ApexAxisChartSeries;
  chart: ApexChart;
  xaxis: ApexXAxis;
  dataLabels: ApexDataLabels;
  title: ApexTitleSubtitle;
  stroke: ApexStroke;
  grid: ApexGrid;
};



@Component({
  selector: 'app-root5',
  templateUrl:'Inicio.html',
  styleUrls:['inicio.css']
})




export class InicioC  {
  //@ViewChild("chart") chart: ChartComponent; 
  public cpu_chart: ChartOptions;
  public ram_chart: ChartOptions;

  public cpu_line: ChartOptions2;
  public ram_line: ChartOptions2;

  public rx_line: ChartOptions2;
  public tx_line: ChartOptions2;

  loading: boolean = false;

  CantMod:boolean=true;

  //variables para mostarar publicaciones
  info:any=[];
  almacenes:any=[];
  versiones:any=[];
  versx:any=[];

  cpuUsedPercentage: number = 65;
  cpuFreePercentage: number = 100 - this.cpuUsedPercentage;

  ramUsedPercentage: number = 65;
  ramFreePercentage: number = 100 - this.cpuUsedPercentage;

  //sirve para mostrar el total de procesos pantalla
  procesosTotal: number =0;

  cpuUsageHistory: number[] = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
  timestamps: string[] = [
    '1', '2', '3', '4', '5', '6', '7', '8', '9', '10',
    '11', '12', '13', '14', '15', '16', '17', '18', '19', '20'
  ];

  constructor(
    private cookieService:CookieService,
    private servicio0:NombreServicioService,
    private servicio1:NombreServicioService ,
    private redire:HomeComponent) {

      

      this.cpu_chart = {
        series: [this.cpuUsedPercentage, this.cpuFreePercentage],
        chart: {
          type: 'donut',
          height: 300, // Ajusta la altura según necesites
        },
        labels: ['Usado', 'Libre'],
        responsive: [
          {
            breakpoint: 480,
            options: {
              chart: {
                width: 200
              },
              legend: {
                position: 'bottom'
              }
            }
          }
        ]
      };


      this.ram_chart = {
        series: [this.cpuUsedPercentage, this.cpuFreePercentage],
        chart: {
          type: 'donut',
          height: 300, // Ajusta la altura según necesites
        },
        labels: ['Usado', 'Libre'],
        responsive: [
          {
            breakpoint: 480,
            options: {
              chart: {
                width: 200
              },
              legend: {
                position: 'bottom'
              }
            }
          }
        ]
      };

      /*INICIA GRAFICA DE LINEAS RAM*/

      this.ram_line = {
        series: [
          {
            name: "Uso de RAM (%)",
            data: this.cpuUsageHistory
          }
        ],
        chart: {
          height: 350,
          type: 'line',
          zoom: {
            enabled: false
          }
        },
        dataLabels: {
          enabled: false
        },
        stroke: {
          curve: 'smooth'
        },
        grid: {
          row: {
            colors: ['#f3f3f3', 'transparent'], // toma dos colores que se alternarán
            opacity: 0.5
          },
        },
        xaxis: {
          categories: this.timestamps,
          title: {
            text: 'Tiempo'
          }
        },
        title: {
          text: 'Historial de Uso de RAM',
          align: 'left'
        }
      };


      this.cpu_line = {
        series: [
          {
            name: "Uso de CPU (%)",
            data: this.cpuUsageHistory
          }
        ],
        chart: {
          height: 350,
          type: 'line',
          zoom: {
            enabled: false
          }
        },
        dataLabels: {
          enabled: false
        },
        stroke: {
          curve: 'smooth'
        },
        grid: {
          row: {
            colors: ['#f3f3f3', 'transparent'], // toma dos colores que se alternarán
            opacity: 0.5
          },
        },
        xaxis: {
          categories: this.timestamps,
          title: {
            text: 'Tiempo'
          }
        },
        title: {
          text: 'Historial de Uso de CPU',
          align: 'left'
        }
      };

/*red envia datos*/




this.rx_line = {
  series: [
    {
      name: "KB recibidos",
      data: this.cpuUsageHistory
    }
  ],
  chart: {
    height: 350,
    type: 'line',
    zoom: {
      enabled: false
    }
  },
  dataLabels: {
    enabled: false
  },
  stroke: {
    curve: 'smooth'
  },
  grid: {
    row: {
      colors: ['#f3f3f3', 'transparent'], // toma dos colores que se alternarán
      opacity: 0.5
    },
  },
  xaxis: {
    categories: this.timestamps,
    title: {
      text: 'Tiempo'
    }
  },
  title: {
    text: 'KB recibidos',
    align: 'left'
  }
};



this.tx_line = {
  series: [
    {
      name: "KB enviados",
      data: this.cpuUsageHistory
    }
  ],
  chart: {
    height: 350,
    type: 'line',
    zoom: {
      enabled: false
    }
  },
  dataLabels: {
    enabled: false
  },
  stroke: {
    curve: 'smooth'
  },
  grid: {
    row: {
      colors: ['#f3f3f3', 'transparent'], // toma dos colores que se alternarán
      opacity: 0.5
    },
  },
  xaxis: {
    categories: this.timestamps,
    title: {
      text: 'Tiempo'
    }
  },
  title: {
    text: 'KB enviados',
    align: 'left'
  }
};



      
  }


 
 
  async ngOnInit() {
  
    this.actualiza_datos()
    setInterval(() => {
      this.actualiza_datos()
      console.log('Array actualizado:');
    }, 1000); // Actualizar cada segundo

  }
 


  async actualiza_datos(){


    //obtiene valores globales para mostrar dashboard
    var resultado = await this.servicio0.get_global_info_unique()

    if(resultado[0]!=undefined){
      var datos = resultado[0]
      var cpu= datos["CPU"]
      var ram= datos["RAM"]
      //console.log("Asigna cpu");
      //console.log( cpu);
      this.cpu_chart.series=[cpu, 100-cpu]
      this.ram_chart.series=[ram, 100-ram]

      this.cpuUsedPercentage = cpu;
      this.cpuFreePercentage= 100 - this.cpuUsedPercentage;

      this.ramUsedPercentage= ram;
      this.ramFreePercentage = 100 - this.ramUsedPercentage;

    }

    //obtiene valores globales para mostrar dashboard
    var resultado2 = await this.servicio0.info_process()
    if(resultado2[0]!=undefined){
      console.log(resultado2[0]["Data"])
      this.procesosTotal= resultado2[0]["Data"].length
    }

    //obtiene listado historial de cpu
    var resultado3 = await this.servicio0.get_global_info()
    if(resultado3!=undefined){
      console.log(resultado3)
      var valores=resultado3 

      var arr_cpu=[]
      var arr_ram=[]

      for (let index = 0; index < valores.length; index++) {
        const element = valores[index];
          arr_cpu.push(element["CPU"])
          arr_ram.push(element["RAM"])

      }

       

      this.ram_line.series= [{
        name: "Uso de RAM (%)",
        data:  arr_ram
      }]

      this.cpu_line.series= [{
        name: "Uso de CPU (%)",
        data:  arr_cpu
      }]

    }


    //obtiene listado  RED
    var resultado4 = await this.servicio1.network_process()
    if(resultado4!=undefined){
      console.log(resultado4)
      var valores=resultado4 

      var arr_rx=[]
      var arr_tx=[]

      for (let index = 0; index < valores.length; index++) {
        const element = valores[index];
        arr_rx.push(element["rx"])
        arr_tx.push(element["tx"])

      }

       

      this.rx_line.series= [{
        name: "KB Recibidos",
        data:  arr_rx
      }]

      this.tx_line.series= [{
        name: "KB Enviados",
        data:  arr_tx
      }]

    }


  }


  

  redirecciona(ruta:string){
    this.redire.refresh(ruta)
  }


























  msp(){ 

    Swal.fire({
      html:"Margarita Fuentes",
      icon: 'success',
      title: 'Por:',
      showConfirmButton: true
    })

  }









}