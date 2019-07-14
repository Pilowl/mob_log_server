import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { LogComponent } from './log/log.component';
import { SessionsComponent } from './sessions/sessions.component';


const routes: Routes = [
  { path: 'log', component: LogComponent},
  { path: '', component: SessionsComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
