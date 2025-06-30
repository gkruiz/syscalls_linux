
import { Injectable, OnInit } from '@angular/core';
import {Router} from '@angular/router';
  
  
@Injectable({
    providedIn: 'root'
})

export class HomeComponent implements OnInit {
  
  constructor(private router:Router) { }
   
 ngOnInit(){}
 onSelect(ruta:string){
   
      this.router.navigate([ruta]);
 }


 refresh(ruta:string){
  this.router.navigateByUrl(ruta);

 }

 
}