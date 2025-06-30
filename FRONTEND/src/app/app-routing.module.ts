import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { LoginC  } from './Componentes/Paginas/Login/login';
import { InicioC } from './Componentes/Paginas/Inicio/inicio';
import { TokenC } from './Componentes/Paginas/Token/Token';


const routes: Routes = [
  //{ path: '', component: Principal },
  { path: 'login', component: LoginC },
  { path: 'inicio', component: InicioC },
  { path: 'verificacion', component: TokenC },
  { path: '**', redirectTo: 'inicio' }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {


}
