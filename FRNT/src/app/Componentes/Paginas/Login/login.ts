import { Component } from '@angular/core';
import { NombreServicioService } from '../../../../servicio/nombre-servicio.service';
import { HomeComponent } from '../../redirect';
import Swal from 'sweetalert2';
 


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


interface Process {

CPUUsage: number
Energy: number;
Name: string;
PID: number;
Priority: number;
RamUsageKB: number;
StartTime: number;
UID: number;

}




@Component({
  selector: 'app-root5',
  templateUrl:'Login.html',
  styleUrls:['login.css']
})




export class LoginC {
    public cpu_chart: ChartOptions;
    public ram_chart: ChartOptions;
  
    public cpu_line: ChartOptions2;
    public ram_line: ChartOptions2;
    public ene_line: ChartOptions2;

    pidG:number=1

    cpuUsageHistory: number[] = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0,0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    timestamps: string[] = [
      '1', '2', '3', '4', '5', '6', '7', '8', '9', '10',
      '11', '12', '13', '14', '15', '16', '17', '18', '19', '20'
    ];

    cpuUsedPercentage: number = 65;
    cpuFreePercentage: number = 100 - this.cpuUsedPercentage;
  
    ramUsedPercentage: number = 65;
    ramFreePercentage: number = 100 - this.cpuUsedPercentage;


  processes: Process[] = [];
  totalRam: string = '5100 MB';  // Reemplaza con el valor real
  usedRam: string = '248 MB';   // Reemplaza con el valor real
  cpuLoad: string = '75%';     // Reemplaza con el valor real
  activeProcesses: string = '120'; // Reemplaza con el valor real
  useEner: string = '0';


  constructor(private servicio0:NombreServicioService,
    private servicio1:NombreServicioService,
    private servicio2:NombreServicioService,
    private redire:HomeComponent  ) {
 

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


      this.ene_line = {
        series: [
          {
            name: "Uso de ENERGIA U",
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
          text: 'Historial de Uso de ENERGIA',
          align: 'left'
        }
      };
    }

 
 



  async ngOnInit(){
    this.loadProcessData();

    setInterval(async () => {

      this.loadProcessData();
      await this.updateData(this.pidG)

      console.log('Array actualizado:');
    }, 1000); // Actualizar cada segundo

  }



  
  async loadProcessData() {

    var resultado0 = await this.servicio0.info_process()

    if(resultado0[0]!=undefined){
      var data = resultado0[0]["Data"]
      console.log(data)

      this.processes=data

      //this.procesosTotal= resultado2[0]["Data"].length
    }
  }


  async showDetails(processId: number) {
    this.pidG=processId
    await this.updateData(processId)
  }




  async killProcess(processId: number) {
    console.log(`KILL proceso con ID: ${processId}`);
    var resultado0 = await this.servicio2.kill_process(processId)
    alert('Se termino el proceso con PID:'+processId);
  }



  async updateData(processId: number){
    console.log(`Mostrar detalles del proceso con ID: ${processId}`);
    var resultado0 = await this.servicio2.info_process_unique(processId)

    if(resultado0!=undefined){
      //carga lo necesario para las graficas falta la de energia
      console.log(resultado0);

      var arr_cpu=[]
      var arr_ram=[]
      var arr_ener=[]

      var cpu =0;
      var ram =0;

      for (let index = 0; index < resultado0.length; index++) {
        const element = resultado0[index];
          //vaores en porcentaje 
          var v1=(element["CPUUsage"]/element["CPUT"]) 
          var v2=(element["RamUsageKB"]/element["RAMT"]) 
          //var v1=(element["CPUUsage"])
          //var v2=(element["RamUsageKB"])
          arr_cpu.push(this.rd(this.rd(v1,5)*100,2))
          arr_ram.push(this.rd(this.rd(v2,5)*100,2))
          arr_ener.push(element["Energy"])
      }


      if(resultado0.length>0){
        cpu=(resultado0[resultado0.length-1]["CPUUsage"]/resultado0[resultado0.length-1]["CPUT"])*100
        ram=(resultado0[resultado0.length-1]["RamUsageKB"]/resultado0[resultado0.length-1]["RAMT"])*100
      }

      this.cpu_chart.series=[cpu, 100-cpu]
      this.ram_chart.series=[ram, 100-ram]

      this.totalRam = ram.toFixed(2)+' %';  
      this.usedRam = resultado0[resultado0.length-1]["RamUsageKB"]+' KB';    
      this.cpuLoad = cpu.toFixed(2)+' %';     
      this.activeProcesses = this.processes.length+'';  
      this.useEner = resultado0[resultado0.length-1]["Energy"]+' U';   

      console.log(arr_ram)

      this.ram_line.series= [{
        name: "Uso de RAM (%)",
        data:  arr_ram
      }]

      this.cpu_line.series= [{
        name: "Uso de CPU (%)",
        data:  arr_cpu
      }]

      this.ene_line.series= [{
        name: "Uso de ENERGIA U",
        data:  arr_ener
      }]


    }



  }


  redirecciona(ruta:string){
    this.redire.refresh(ruta)
  }

  rd(numero:number, decimales:number) {
    const factor = Math.pow(10, decimales);
    return Math.round(numero * factor) / factor;
  }
  
 
}
