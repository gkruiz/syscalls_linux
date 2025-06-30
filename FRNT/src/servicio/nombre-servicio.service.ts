import { Injectable } from '@angular/core';
import { HttpClient ,HttpHeaders  } from '@angular/common/http';
import { AxiosHeaders } from 'axios';
//import { Base64 } from 'js-base64';
//const axios = require('axios');
import axios from 'axios';


@Injectable({
  providedIn: 'root'
})


export class NombreServicioService {

  constructor(private  http:HttpClient) {}


  direccion:string="http://localhost:8080";

  //PRODUCTIVA
  PRB:String=""




  async get_global_info_unique(){
    
    const url = this.direccion+'/get_global_info_unique';
    const {data} = await axios.get(url);

    return data

  }


  async info_process(){
    
    const url = this.direccion+'/info_process';
    const {data} = await axios.get(url);

    return data

  }



  async get_global_info(){
    
    const url = this.direccion+'/get_global_info';
    const {data} = await axios.get(url);

    return data

  }


  async info_process_unique(PID:number){
    
    
   const url = `${this.direccion}/info_process_unique?pid=${PID}`; 
    //const url = this.direccion+'/info_process_unique';
    const {data} = await axios.get(url);

    return data

  }
  


  async kill_process(PID:number){
    
    
    const url = `${this.direccion}/kill_process?pid=${PID}`; 
     //const url = this.direccion+'/info_process_unique';
     const {data} = await axios.get(url);
 
     return data
 
   }
  

   async network_process(){
    
    const url = this.direccion+'/network_process';
    const {data} = await axios.get(url);

    return data

  }


  async sessionl(){
      
    const url = 'https://api.ipify.org/?format=json';
    const {data} = await axios.get(url);
    
    return data

  }




  

  vs(usuario:string,modulo:number){
    let user = { 
      usuario: usuario,
      modulo:modulo
   }

  /* const url = this.direccion+'/vs';
   const {data} = await axios.post(url,user);

   return data*/

    //return this.http.get('https://jsonip.com/');
    return this.http.post(this.direccion+'/vs',user);
  }


 

  //ejercicio consumir servicio
 async upImage2(imagen:any){
    const url = 'http://192.178.10.3:8080/upload';
    const {data} = await axios.post(url,imagen);
 
    return data
    //return this.http.post('http://192.178.10.3:8080/upload',imagen);
  }





 async logout(usuario:string,descripcion:string){
    let user = { 
       usuario: usuario,
       descripcion:descripcion
    }
    const url = this.direccion+'/endSession';
    const {data} = await axios.post(url,user);

    return data

    //return this.http.post(this.direccion+'/endSession',user);
  }






  async Login(usuario:string,contrasena:string,ip:string){
    let user = { 
       usuario: usuario,
       contrasena:contrasena,
       ip:ip
    }

    const url = this.direccion+'/login';
    const {data} = await axios.post(url,user);

    return data
   // return this.http.post(this.direccion+'/login',user);
  }





 async Token(llave:string,codigo:string,ip:string){
    let user = { 
       llave: llave,
       sub:codigo,
       ip:ip
    }

    const url = this.direccion+'/validaToken';
    const {data} = await axios.post(url,user);

    return data
    //return this.http.post(this.direccion+'',user);
  }



 

 


}





