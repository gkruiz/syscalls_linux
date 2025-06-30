import { Component } from '@angular/core';
import { NombreServicioService } from '../../../../servicio/nombre-servicio.service';
import { CookieService } from 'ngx-cookie-service';
import { HomeComponent } from '../../redirect';
 


@Component({
  selector: 'app-root5',
  templateUrl:'Token.html'
})


export class TokenC {
  
  usuario:string='';
  contrasena:string='';
  existe:any=[];
  valores:any=[];
  contador=0;
  ContBaneo=0;
  constructor(private servicio:NombreServicioService,
    private servicio2:NombreServicioService,
    private cookieService:CookieService,
    private redire:HomeComponent  ) { }

    id_usu:string="";
    nombre:string="";
    usuariox:string="";

    
  ngOnInit(): void {

    //let idtempo=localStorage.getItem('id_usu')+"";
    let idtempo=this.cookieService.get('id_usu').toString();  


  }

  



  
  async loginfun(codigo:string) {
      ////console.log(this.usuario);
     // //console.log(this.contrasena);

  }



  
 
}