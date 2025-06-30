import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';

import { HttpClientModule, HttpHeaders } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { QrCodeModule } from 'ng-qrcode';



import { Principal } from './Componentes/Principal';
import { LoginC  } from './Componentes/Paginas/Login/login';
import { InicioC } from './Componentes/Paginas/Inicio/inicio';

 
import { TokenC } from './Componentes/Paginas/Token/Token';
import { PdfViewerModule } from 'ng2-pdf-viewer';


import { InfiniteScrollModule } from 'ngx-infinite-scroll';
import { FontAwesomeModule } from '@fortawesome/angular-fontawesome';
import {NgxPaginationModule} from 'ngx-pagination';
import { NgxApexchartsModule } from "ngx-apexcharts";


@NgModule({
  declarations: [
    Principal , LoginC  , InicioC 
    ,TokenC, 
 
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    PdfViewerModule,
    FormsModule,
    ReactiveFormsModule,
    InfiniteScrollModule,

    FontAwesomeModule,
    BrowserModule,
     NgxPaginationModule,
     QrCodeModule,
     NgxApexchartsModule,

  ], 
  providers: [
    
  ],
  bootstrap: [Principal]
})



export class AppModule { 

}
