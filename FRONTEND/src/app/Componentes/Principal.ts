import { Component } from '@angular/core';
import { HomeComponent } from './redirect';
import { Router } from '@angular/router';
//import { rutaAC } from './rutaA';
import { CookieService } from 'ngx-cookie-service';

@Component({
  selector: 'app-root',
  templateUrl: './Principal.html'
})


export class Principal {
  title = 'pro9';




constructor(private redire:HomeComponent,
  private cookieService:CookieService,
  private router: Router){ }

  

  ngOnInit(): void {

  
  }

  
}

